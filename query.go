package trafikinfo

import (
	"encoding/xml"
	"strconv"

	"github.com/phille97/trafikinfo/trv"
)

// Query is used to request information from the Trafikinfo API
type Query struct {
	query *query
}

func (q Query) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(q.query, start)
}

type query struct {
	ObjectType            string       `xml:"objecttype,attr"`
	SchemaVersion         string       `xml:"schemaversion,attr"`
	Namespace             string       `xml:"namespace,attr,omitempty"`
	ID                    string       `xml:"id,attr,omitempty"`
	IncludeDeletedObjects bool         `xml:"includedeletedobjects,attr,omitempty"`
	Limit                 int          `xml:"limit,attr,omitempty"`
	OrderBy               string       `xml:"orderby,attr,omitempty"`
	Skip                  int          `xml:"skip,attr,omitempty"`
	LastModified          bool         `xml:"lastmodified,attr,omitempty"`
	ChangeID              string       `xml:"changeid,attr"`
	SSEURL                bool         `xml:"sseurl,attr,omitempty"`
	Filter                *Filter      `xml:"FILTER"`
	Include               []string     `xml:"INCLUDE"`
	Exclude               []string     `xml:"EXCLUDE"`
	Distinct              string       `xml:"DISTINCT,omitempty"`
	Eval                  []Evaluation `xml:"EVAL"`
}

// NewQuery returns a query that other methods can be chained on to further
// customise the request.
func NewQuery(obj trv.ObjectType) Query {
	return Query{
		query: &query{
			ObjectType:    obj.Kind,
			Namespace:     obj.Namespace,
			SchemaVersion: obj.Version,
			ChangeID:      "0",
			LastModified:  true,
		},
	}
}

// ID is an arbitrary value which will be echoed in the response. It can
// be used to associate queries with responses, especially when a request
// includes multiple queries.
func (q Query) ID(opt string) Query {
	q.query.ID = opt
	return q
}

// IncludeDeletedObjects requests that deleted objects also be returned.
// This defaults to false.
func (q Query) IncludeDeletedObjects(opt bool) Query {
	q.query.IncludeDeletedObjects = opt
	return q
}

// Limit sets the limit for the amount of items to be returned. This can be
// used together with Skip to implement pagination. A Limit of 0 means no
// limit at all, i.e return everything.
func (q Query) Limit(opt int) Query {
	q.query.Limit = opt
	return q
}

// OrderBy takes a sorting expression the items in the response will be
// sorted by.
//
// For example:
//
//	OrderBy("SomeData.Name desc, SomeData.Description asc")
func (q Query) OrderBy(opt string) Query {
	q.query.OrderBy = opt
	return q
}

// Skip sets how many items to skip/exclude from the response. This can be
// used together with Limit to implement pagination.
func (q Query) Skip(opt int) Query {
	q.query.Skip = opt
	return q
}

// ChangeID sets the change ID for the request. This should initially be
// 0 to request all data, and then be set to the value of the change ID in
// the response to only get updated/deleted objects since the previous change
// ID.
func (q Query) ChangeID(opt string) Query {
	if opt == "" {
		q.query.ChangeID = "0"
	} else {
		q.query.ChangeID = opt
	}
	return q
}

func (q Query) SSEURL(opt bool) Query {
	q.query.SSEURL = opt
	return q
}

func (q Query) Filter(filters ...Filter) Query {
	f := merge(filters...)
	q.query.Filter = &Filter{filter: f}
	return q
}

// Distinct returns an array of unique values of this field
func (q Query) Distinct(field string) Query {
	q.query.Distinct = field
	return q
}

// Include ensures only elements matching the specified fields are returned
func (q Query) Include(fields ...string) Query {
	q.query.Include = append(q.query.Include, fields...)
	return q
}

// Exclude ensures element matching the specifields fields are omitted
func (q Query) Exclude(fields ...string) Query {
	q.query.Exclude = append(q.query.Exclude, fields...)
	return q
}

func (q Query) Eval(evaluations ...Evaluation) Query {
	q.query.Eval = append(q.query.Eval, evaluations...)
	return q
}

// And builds an AND filter
//
// This filter should be called with at least 2 filters. With only
// one filter the AND doesn't make much sense and is equivalent to
// adding the child filter on the Query directly. Despite this, an
// AND filter with only one child is still considered valid by the
// API.
func And(filters ...Filter) Filter {
	return Filter{
		filter: filter{And: []filter{merge(filters...)}},
	}
}

// Or builds an OR filter
//
// This filter should be called with at least 2 filters. With only
// one filter the OR doesn't make much sense and is equivalent to
// adding the child filter on the Query directly. Despite this, an
// OR filter with only one child is still considered valid by the
// API.
func Or(filters ...Filter) Filter {
	return Filter{
		filter: filter{Or: []filter{merge(filters...)}},
	}
}

// Not builds an NOT filter
//
// This filter should be called with at least 1 filter.
func Not(filters ...Filter) Filter {
	return Filter{
		filter: filter{Not: []filter{merge(filters...)}},
	}
}

// ElementMatch builds an ELEMENTMATCH filter
//
// This filter should be called with at least 2 filters. Using any less
// than 2 filters will result in the API returning an error.
func ElementMatch(filters ...Filter) Filter {
	return Filter{
		filter: filter{ElementMatch: []filter{merge(filters...)}},
	}
}

type Evaluation struct {
	evaluation evaluation
}

func (v Evaluation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(v.evaluation, start)
}

type evaluation struct {
	Alias    string `xml:"alias,attr,omitempty"`
	Function string `xml:"function,attr,omitempty"`
}

func Eval(alias, function string) Evaluation {
	return Evaluation{
		evaluation: evaluation{
			Alias:    alias,
			Function: function,
		},
	}
}

// Filter represents a filter element in the query.
type Filter struct {
	filter filter
}

func (f Filter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
