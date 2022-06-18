package trafikinfo

import (
	"encoding/xml"
	"strconv"
)

// Query is used to request information from the Trafikinfo API
type Query struct {
	query query
}

func (q *Query) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(q.query, start)
}

type query struct {
	ObjectType    ObjectType `xml:"objecttype,attr"`
	SchemaVersion float64    `xml:"schemaversion,attr"`
	Filter        *Filter    `xml:"FILTER"`
	Include       []string   `xml:"INCLUDE"`
	Exclude       []string   `xml:"EXCLUDE"`
	Distinct      string     `xml:"DISTINCT,omitempty"`
}

// NewQuery returns a query with the provided filters
func NewQuery(objectType ObjectType, schemaVersion float64) *Query {
	return &Query{
		query: query{
			ObjectType:    objectType,
			SchemaVersion: schemaVersion,
		},
	}
}

func (q *Query) Filter(filters ...Filter) *Query {
	f := merge(filters...)
	q.query.Filter = &Filter{filter: f}
	return q
}

// Distinct returns an array of unique values of this field
func (q *Query) Distinct(field string) *Query {
	q.query.Distinct = field
	return q
}

// Include ensures only elements matching the specified fields are returned
func (q *Query) Include(fields ...string) *Query {
	q.query.Include = append(q.query.Include, fields...)
	return q
}

// Exclude ensures element matching the specifields fields are ommitted
func (q *Query) Exclude(fields ...string) *Query {
	q.query.Exclude = append(q.query.Exclude, fields...)
	return q
}

// And builds an AND filter
func And(f1, f2 Filter, rest ...Filter) Filter {
	f := []Filter{f1, f2}
	f = append(f, rest...)
	return Filter{
		filter: filter{And: []filter{merge(f...)}},
	}
}

// Or builds an OR filter
func Or(f1, f2 Filter, rest ...Filter) Filter {
	f := []Filter{f1, f2}
	f = append(f, rest...)
	return Filter{
		filter: filter{Or: []filter{merge(f...)}},
	}
}

// Not builds an NOT filter
func Not(f1, f2 Filter, rest ...Filter) Filter {
	f := []Filter{f1, f2}
	f = append(f, rest...)
	return Filter{
		filter: filter{Not: []filter{merge(f...)}},
	}
}

// ElementMatch builds an ELEMENTMATCH filter
func ElementMatch(f1, f2 Filter, rest ...Filter) Filter {
	f := []Filter{f1, f2}
	f = append(f, rest...)
	return Filter{
		filter: filter{ElementMatch: []filter{merge(f...)}},
	}
}

// Filter represents a filter element in the query.
type Filter struct {
	filter filter
}

func (f *Filter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(f.filter, start)
}

type filter struct {
	And          []filter `xml:"AND"`
	Or           []filter `xml:"OR"`
	Not          []filter `xml:"NOT"`
	ElementMatch []filter `xml:"ELEMENTMATCH"`

	Equal              []comparison `xml:"EQ"`
	Exists             []comparison `xml:"EXISTS"`
	GreaterThan        []comparison `xml:"GT"`
	GreaterThanOrEqual []comparison `xml:"GTE"`
	LessThan           []comparison `xml:"LT"`
	LessThanOrEqual    []comparison `xml:"LTE"`
	NotEqual           []comparison `xml:"NE"`
	Like               []comparison `xml:"LIKE"`
	NotLike            []comparison `xml:"NOTLIKE"`
	In                 []comparison `xml:"IN"`
	NotIn              []comparison `xml:"NOTIN"`
	Within             []comparison `xml:"WITHIN"`
	Intersects         []comparison `xml:"INTERSECTS"`
	Near               []comparison `xml:"NEAR"`
}

// merge combines a set of filters into their final form
func merge(filters ...Filter) filter {
	f := filter{}
	for _, elem := range filters {
		f.And = append(f.And, elem.filter.And...)
		f.Or = append(f.Or, elem.filter.Or...)
		f.Not = append(f.Not, elem.filter.Not...)
		f.ElementMatch = append(f.ElementMatch, elem.filter.ElementMatch...)
		f.Equal = append(f.Equal, elem.filter.Equal...)
		f.Exists = append(f.Exists, elem.filter.Exists...)
		f.GreaterThan = append(f.GreaterThan, elem.filter.GreaterThan...)
		f.GreaterThanOrEqual = append(f.GreaterThanOrEqual, elem.filter.GreaterThanOrEqual...)
		f.LessThan = append(f.LessThan, elem.filter.LessThan...)
		f.LessThanOrEqual = append(f.LessThanOrEqual, elem.filter.LessThanOrEqual...)
		f.NotEqual = append(f.NotEqual, elem.filter.NotEqual...)
		f.Like = append(f.Like, elem.filter.Like...)
		f.NotLike = append(f.NotLike, elem.filter.NotLike...)
		f.In = append(f.In, elem.filter.In...)
		f.NotIn = append(f.NotIn, elem.filter.NotIn...)
		f.Within = append(f.Within, elem.filter.Within...)
		f.Intersects = append(f.Intersects, elem.filter.Intersects...)
		f.Near = append(f.Near, elem.filter.Near...)
	}

	return f
}

// comparison represents a filter operation that will be applied to each
// element to decide on its inclusion in the response
type comparison struct {
	Name        string  `xml:"name,attr"`
	Value       string  `xml:"value,attr"`
	Shape       string  `xml:"shape,attr,omitempty"`
	Radius      float64 `xml:"radius,attr,omitempty"`
	MinDistance *int    `xml:"mindistance,attr"`
	MaxDistance *int    `xml:"maxdistance,attr"`
}

// Equal filters by if the field provided by name matches the specified value
func Equal(name, value string) Filter {
	return Filter{filter: filter{Equal: []comparison{{Name: name, Value: value}}}}
}

// Exists filters by whether the field provided by name exists or not
func Exists(name string, exists bool) Filter {
	return Filter{filter: filter{Exists: []comparison{{Name: name, Value: strconv.FormatBool(exists)}}}}
}

// GreaterThan filters by whether the field specified by name is greater than the specified value
func GreaterThan(name, value string) Filter {
	return Filter{filter: filter{GreaterThan: []comparison{{Name: name, Value: value}}}}
}

// GreaterThanOrEqual filters by whether the field specified by name is greater than or equal to the specified value
func GreaterThanOrEqual(name, value string) Filter {
	return Filter{filter: filter{GreaterThanOrEqual: []comparison{{Name: name, Value: value}}}}
}

// LessThan filters by whether the field specified by name is less than the specified value
func LessThan(name, value string) Filter {
	return Filter{filter: filter{LessThan: []comparison{{Name: name, Value: value}}}}
}

// LessThanOrEqual filters by whether the field specified by name is less than or equal to the specified value
func LessThanOrEqual(name, value string) Filter {
	return Filter{filter: filter{LessThanOrEqual: []comparison{{Name: name, Value: value}}}}
}

// NotEqual is the inverse of Equal
func NotEqual(name, value string) Filter {
	return Filter{filter: filter{NotEqual: []comparison{{Name: name, Value: value}}}}
}

// Like filters by if the field provided by name matches the regex provided by value
func Like(name, value string) Filter {
	return Filter{filter: filter{Like: []comparison{{Name: name, Value: value}}}}
}

// NotLIke is the inverse of Like
func NotLike(name, value string) Filter {
	return Filter{filter: filter{NotLike: []comparison{{Name: name, Value: value}}}}
}

// In filters by if the field specified by name matches any of the provided values
func In(name, value string) Filter {
	return Filter{filter: filter{In: []comparison{{Name: name, Value: value}}}}
}

// NotIn is the inverse of In
func NotIn(name, value string) Filter {
	return Filter{filter: filter{NotIn: []comparison{{Name: name, Value: value}}}}
}

// Within filters by if the field specified in name falls within the
// specified shape, radius and the coordinates in value
func Within(name, value, shape string, radius float64) Filter {
	return Filter{filter: filter{Within: []comparison{{Name: name, Value: value, Shape: shape, Radius: radius}}}}
}

// Intersects filters by if te field specified in name intersects with
// the coordinates provided in value
func Intersects(name, value string) Filter {
	return Filter{filter: filter{Intersects: []comparison{{Name: name, Value: value}}}}
}

// Near filters by if the field specified in name is within the specified
// min/max dinstance from the point coordinates sepcified in value
func Near(name, value string, minDistance, maxDistance int) Filter {
	return Filter{filter: filter{Near: []comparison{{Name: name, Value: value, MinDistance: &minDistance, MaxDistance: &maxDistance}}}}
}
