// Package tree contains a bunch of awful code that transforms the XSD into
// an internal representation that we can then template into Go code.
package tree

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/phille97/trafikinfo/internal/meta"
	"github.com/phille97/trafikinfo/internal/xsd"
)

type Root struct {
	Meta meta.Data

	Nodes   []*Node
	Imports []string
}

type Node struct {
	Name          string
	Type          Type
	Nodes         []*Node
	Multiple      bool
	Attr          bool
	Documentation []string
	IsBase        bool
}

type TypeMap struct {
	Types    map[string]*Node
	Bases    map[string]struct{}
	Aliasses map[string]string
}

func WalkXSD(schema *xsd.Schema) *Root {
	// The code you're about to witness is horrendous. But it works, achieves
	// its goal and performance doesn't matter as there's rarely a need to
	// generate the code

	r := &Root{
		Meta: schema.Meta,
	}

	cts := map[string]xsd.ComplexType{}
	ags := map[string]xsd.AttributeGroup{}
	allNodes := []*Node{}

	tm := TypeMap{
		Types:    map[string]*Node{},
		Bases:    map[string]struct{}{},
		Aliasses: map[string]string{},
	}

	// Create a mapping of all complexTypes. These are the equivalent of
	// the package-level structs we need to generate
	for _, ct := range schema.ComplexTypes {
		cts[goName(ct.Name)] = ct
	}

	// Create nodes out of all simpleTypes. Simple types tend to be an end
	// type, like a string, but with some additional restriction imposed.
	// However, for our purpose we don't care about the restrictions as we
	// don't implement validation routines
	for _, st := range schema.SimpleTypes {
		if len(st.Restriction.Enumerations) == 0 {
			n := Node{Name: goName(st.Name), Type: typeMap(st.Restriction.Base)}
			tm.Types[goName(n.Name)] = &n
			continue
		}

		typ := Type{
			Kind:      st.Name,
			Choices:   make([]string, 0, len(st.Restriction.Enumerations)),
			Undertype: typeMap(st.Restriction.Base).Kind,
		}
		for _, enum := range st.Restriction.Enumerations {
			typ.Choices = append(typ.Choices, enum.Value)
		}
		n := Node{Name: goName(st.Name), Type: typ}
		if anot := st.Annotation; anot != nil {
			for _, doc := range anot.Documentations {
				n.Documentation = append(n.Documentation, DocString(doc))
			}
		}
		tm.Types[goName(n.Name)] = &n
		allNodes = append(allNodes, &n)
	}

	// Create a mapping of attributeGroups, so that when we run into a
	// reference to an attribute group we can inline the attributes on
	// the type
	for _, ag := range schema.AttributeGroups {
		ags[goName(ag.Name)] = ag
	}

	refs := map[string]xsd.Element{}
	for _, x := range schema.Elements {
		if len(x.Name) == 0 {
			continue
		}
		refs[goName(x.Name)] = x
	}

	// This is where it starts to get funky. Start by looping over each
	// complexType
	for _, ct := range schema.ComplexTypes {
		// Create a node for each type we encounter
		n := &Node{
			Name: goName(ct.Name),
		}

		// If we have a simpleContent inside a complexType, we're
		// restricting an existing type. This won't be emitted directly
		// but will instead be a type an element uses. This creates a
		// type out of it, so we can do type resolution later
		if sc := ct.SimpleContent; sc != nil {
			n.Type = typeMap(sc.Extension.Base)
			tm.Types[goName(ct.Name)] = n
			// This continue is important, as now we're done processing
			// this simple type and we don't want to add it to the generic
			// nodes, since we'll never emit this type as a struct
			continue
		}

		n.Nodes = []*Node{}
		n.Type = Type{Kind: "struct"}

		// Elements on a complexType are fields in a struct
		for _, elem := range AllElems(ct) {
			node, isType := ElemToNode(elem, refs)
			n.Nodes = append(n.Nodes, node)
			if isType {
				nt := Node{Name: node.Name, Type: node.Type, Documentation: node.Documentation}
				tm.Types[goName(nt.Name)] = &nt
				allNodes = append(allNodes, &nt)
			}
		}

		// Resolve any attribute groups, and create fields on the struct
		for _, ag := range ct.AttributeGroups {
			attrg, ok := ags[goName(ag.Ref)]
			if !ok {
				panic(fmt.Sprintf("did not find attribute group: %s", ag.Name))
			}
			for _, attr := range attrg.Attributes {
				n.Nodes = append(n.Nodes, AttrToNode(attr))
			}
		}

		// If we have complexContent, then we're probably extending an existing
		// type. This usually means taking an existing type and adding one or
		// multiple fields to it
		if cc := ct.ComplexContent; cc != nil {
			extCt, ok := cts[goName(cc.Extension.Base)]

			// If we're not extending another type, it's effectively just an
			// alias. So instead, keep a mapping that we can consult during
			// type resolution
			if cc.Extension.Sequence == nil {
				tm.Aliasses[goName(ct.Name)] = goName(extCt.Name)
				continue
			}

			// Add the base we've used to our typemap, so we can use it
			// for type resolution later. Bases themselves we don't want
			// to emit as structs as we inline them in the type definition
			// of whoever is using it
			tm.Bases[goName(extCt.Name)] = struct{}{}
			if !ok {
				panic(fmt.Sprintf("base: %s not found in ComplexTypes", cc.Extension.Base))
			}
			// In order to respect the intention of the XSD, create field
			// definitions for the elements in the base type first
			for _, elem := range AllElems(extCt) {
				node, _ := ElemToNode(elem, refs)
				n.Nodes = append(n.Nodes, node)
			}
			// Now add the additional fields of our type to it
			for _, elem := range cc.Extension.Sequence.Elements {
				node, _ := ElemToNode(elem, refs)
				n.Nodes = append(n.Nodes, node)
			}
		}
		allNodes = append(allNodes, n)
	}

	// Resolve any indirect types we may have, like A -> B -> string
	// so that looking up A results in string
	for _, val := range tm.Types {
		val.ResolveType(tm)
	}

	// Create the actual result nodes for our tree
	resNodes := []*Node{}

	// Go over the existing nodes and only copy those over
	// that we want to keep
	for _, n := range allNodes {
		// Check if the node is a base and skip those as their
		// elements have been inlined
		if _, ok := tm.Bases[n.Name]; ok {
			continue
		}
		// Do the type resolution for the node and any child
		// nodes
		n.ResolveType(tm)
		resNodes = append(resNodes, n)
	}

	r.Nodes = resNodes
	return r
}

func (n *Node) ResolveType(tm TypeMap) {
	if n.Type.Final {
		return
	}

	if val, ok := tm.Aliasses[n.Type.Kind]; ok {
		n.Type.Kind = val
	}

	if val, ok := tm.Types[n.Type.Kind]; ok {
		n.Type = val.Type
	}

	for _, nd := range n.Nodes {
		nd.ResolveType(tm)
	}
}

type Type struct {
	Kind      string
	Final     bool
	Choices   []string
	Undertype string
}

func AllElems(cts xsd.ComplexType) []xsd.Element {
	if cts.All != nil {
		return cts.All.Elements
	}
	if cts.Sequence != nil {
		return cts.Sequence.Elements
	}
	return nil
}

func AttrToNode(attr xsd.Attribute) *Node {
	n := Node{
		Name: attr.Name,
		Type: typeMap(attr.Type),
		Attr: true,
	}
	return &n
}

func ElemToNode(elem xsd.Element, refs map[string]xsd.Element) (*Node, bool) {
	if elem.Reference != "" {
		refElem, ok := refs[goName(elem.Reference)]
		if !ok {
			panic(fmt.Sprintf("cannot handle elem with ref: %s", elem.Reference))
		}
		elem = refElem
	}
	n := Node{
		Name:     goName(elem.Name),
		Multiple: elem.Multiple,
	}
	isType := false
	if st := elem.SimpleType; st != nil && len(st.Restriction.Enumerations) != 0 {
		isType = true
		n.Type = Type{
			Kind:      elem.Name,
			Undertype: typeMap(st.Restriction.Base).Kind,
		}
		choices := []string{}
		for _, c := range st.Restriction.Enumerations {
			choices = append(choices, c.Value)
		}
		n.Type.Choices = choices
	} else {
		n.Type = typeMap(elem.Type)
	}
	if anot := elem.Annotation; anot != nil {
		for _, doc := range anot.Documentations {
			n.Documentation = append(n.Documentation, DocString(doc))
		}
	}

	return &n, isType
}

func DocString(doc xsd.Documentation) string {
	res := strings.Join(strings.Fields(doc.Content), " ")
	res = strings.TrimSpace(res)
	return fmt.Sprintf("// %s: %s", strings.ToUpper(doc.Language), res)
}

func tag(n *Node) string {
	s := "`xml:\"" + n.Name
	if n.Attr {
		s = s + ",attr"
	}
	s = s + ",omitempty\"`"
	return s
}

func typeMap(s string) Type {
	if s == "" {
		return Type{
			Kind:  "string",
			Final: true,
		}
	}

	if !strings.HasPrefix(s, "xs:") {
		return Type{Kind: goName(s)}
	}

	types := map[string]string{
		"xs:boolean":      "bool",
		"xs:string":       "string",
		"xs:long":         "int64",
		"xs:int":          "int",
		"xs:integer":      "int64",
		"xs:float":        "float64",
		"xs:double":       "float64",
		"xs:dateTime":     "time.Time",
		"xs:date":         "time.Time",
		"xs:unsignedByte": "uint8",
		"xs:decimal":      "float64",
	}
	val, ok := types[s]
	if !ok {
		panic(fmt.Sprintf("no Go type for XML type: %s", s))
	}

	return Type{Kind: val, Final: true}
}

func upperFirst(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func goName(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	sl := strings.Split(s, "_")
	for i, elem := range sl[:] {
		sl[i] = upperFirst(elem)
	}
	s = strings.Join(sl, "")
	s = strings.TrimSuffix(s, "Measure")
	s = strings.TrimSuffix(s, "Config")
	if s == "Amount" {
		return s
	}
	s = strings.TrimSuffix(s, "Amount")
	return s
}

func goIdentifier(s string, multiple bool) string {
	s = goName(s)

	switch strings.ToLower(s) {
	case "countyno":
		s = "County"
	case "checkoutid":
		s = "CheckoutID"
	case "leaseduntil":
		s = "LeasedUntil"
	case "statuscode":
		s = "StatusCode"
	case "itemsacknowledged":
		s = "ItemsAcknowledged"
	case "itemsleft":
		s = "ItemsLeft"
	}

	if strings.HasSuffix(s, "Id") {
		s = strings.TrimSuffix(s, "Id")
		s = s + "ID"
	}

	if strings.HasSuffix(s, "Url") {
		s = strings.TrimSuffix(s, "Url")
		s = s + "URL"
	}

	if multiple {
		l := strings.ToLower(s)
		if !strings.HasSuffix(l, "info") &&
			!strings.HasSuffix(l, "s") &&
			!strings.HasSuffix(l, "information") {
			switch l {
			case "valid", "equipment", "trafficimpact", "wind", "ground":
			case "facility":
				s = "Facilities"
			case "county":
				s = "Counties"
			case "typeoftraffic":
				s = "TypesOfTraffic"
			default:
				s = s + "s"
			}
		}
	}

	return s
}
