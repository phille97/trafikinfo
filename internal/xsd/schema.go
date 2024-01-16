// Package xsd parses an XML Schema Definition. This package does a number of questionable
// things to handle the Trafikverket XSD, as their XSDs aren't strictly speaking valid.
//
// The package only implements features necessary to deal with the Trafikverket XSDs and
// can't be used as a generic XSD decoder.
package xsd

import (
	"encoding/xml"
	"slices"
	"strings"

	"code.dny.dev/trafikinfo/internal/meta"
)

var complexSkip = []string{
	"response", "result", "error", "lastmodified", "evalresult", "info",
}

type Schema struct {
	XMLName         xml.Name         `xml:"schema"`
	ComplexTypes    []ComplexType    `xml:"complexType"`
	SimpleTypes     []SimpleType     `xml:"simpleType"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
	Annotation      []Annotation     `xml:"annotation"`

	Meta meta.Data `xml:"-"`
}

func (s *Schema) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type schema Schema
	elem := schema{
		Meta: s.Meta,
	}

	if err := d.DecodeElement(&elem, &start); err != nil {
		return err
	}

	cts := []ComplexType{}
	for i, ct := range elem.ComplexTypes {
		if slices.Contains(complexSkip, strings.ToLower(elem.ComplexTypes[i].Name)) {
			continue
		}

		if ct.Name == "DynamicObjectType" {
			elem.ComplexTypes[i].Name = elem.Meta.Name
		}

		cts = append(cts, elem.ComplexTypes[i])
	}

	elem.ComplexTypes = cts
	*s = (Schema)(elem)
	return nil
}

type Element struct {
	XMLName      xml.Name    `xml:"element"`
	Name         string      `xml:"name,attr"`
	Type         string      `xml:"type,attr"`
	Reference    string      `xml:"ref,attr"`
	Multiple     bool        `xml:"-"`
	StrMaxOccurs string      `xml:"maxOccurs,attr"`
	Nillable     bool        `xml:"nillable,attr"`
	Annotation   *Annotation `xml:"annotation"`
	SimpleType   *SimpleType `xml:"simpleType"`

	// This attr is specific to Trafikverket
	Key bool `xml:"key,attr"`
}

func (e *Element) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type element Element
	elem := element{}
	if err := d.DecodeElement(&elem, &start); err != nil {
		return err
	}
	switch elem.StrMaxOccurs {
	case "1", "":
	default:
		elem.Multiple = true
	}
	*e = (Element)(elem)
	return nil
}

type Attribute struct {
	XMLName xml.Name `xml:"attribute"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Use     string   `xml:"use,attr"`
}

type AttributeGroup struct {
	XMLName    xml.Name    `xml:"attributeGroup"`
	Name       string      `xml:"name,attr"`
	Attributes []Attribute `xml:"attribute"`
	Ref        string      `xml:"ref,attr"`
}

type ComplexType struct {
	XMLName         xml.Name         `xml:"complexType"`
	Name            string           `xml:"name,attr"`
	Mixed           bool             `xml:"mixed,attr"`
	All             *All             `xml:"all"`
	Sequence        *Sequence        `xml:"sequence"`
	SimpleContent   *SimpleContent   `xml:"simpleContent"`
	ComplexContent  *ComplexContent  `xml:"complexContent"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
}

type All struct {
	XMLName    xml.Name    `xml:"all"`
	Annotation *Annotation `xml:"annotation"`
	Elements   []Element   `xml:"element"`
}

type SimpleType struct {
	XMLName     xml.Name    `xml:"simpleType"`
	Name        string      `xml:"name,attr"`
	Restriction Restriction `xml:"restriction"`
	Annotation  *Annotation `xml:"annotation"`
}

type Restriction struct {
	XMLName      xml.Name    `xml:"restriction"`
	Base         string      `xml:"base,attr"`
	Attributes   []Attribute `xml:"attribute"`
	Enumerations []struct {
		Value string `xml:"value,attr"`
	} `xml:"enumeration"`
	SimpleContent *struct {
		Value string `xml:"value,attr"`
	} `xml:"simpleContent"`
	MinInclusive *struct {
		Value string `xml:"value,attr"`
	} `xml:"minInclusive"`
	MaxInclusive *struct {
		Value string `xml:"value,attr"`
	} `xml:"maxInclusive"`
	Pattern *struct {
		Value string `xml:"value,attr"`
	} `xml:"pattern"`
}

type Extension struct {
	XMLName    xml.Name    `xml:"extension"`
	Base       string      `xml:"base,attr"`
	Attributes []Attribute `xml:"attribute"`
	Sequence   *Sequence   `xml:"sequence"`
}

type Sequence struct {
	XMLName      xml.Name    `xml:"sequence"`
	MinOccurs    int         `xml:"-"`
	MaxOccurs    int         `xml:"-"`
	Multiple     bool        `xml:"-"`
	StrMaxOccurs string      `xml:"maxOccurs,attr"`
	Elements     []Element   `xml:"element"`
	Annotation   *Annotation `xml:"annotation"`
}

func (s *Sequence) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type sequence Sequence
	elem := sequence{}
	if err := d.DecodeElement(&elem, &start); err != nil {
		return err
	}
	switch elem.StrMaxOccurs {
	case "1", "":
	default:
		elem.Multiple = true
	}
	*s = (Sequence)(elem)
	return nil
}

type SimpleContent struct {
	XMLName    xml.Name    `xml:"simpleContent"`
	Annotation *Annotation `xml:"annotation"`
	Extension  *Extension  `xml:"extension"`
}

type ComplexContent struct {
	XMLName    xml.Name    `xml:"complexContent"`
	Mixed      bool        `xml:"mixed,attr"`
	Annotation *Annotation `xml:"annotation"`
	Extension  Extension   `xml:"extension"`
}

type Annotation struct {
	XMLName        xml.Name        `xml:"annotation"`
	Documentations []Documentation `xml:"documentation"`
}

type Documentation struct {
	XMLName  xml.Name `xml:"documentation"`
	Language string   `xml:"lang,attr"`
	Content  string   `xml:",innerxml"`
}
