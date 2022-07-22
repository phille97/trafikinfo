package trafikinfo

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestQueryID(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).ID("test")
	if res := q.query.ID; res != "test" {
		t.Fatalf("Expected query with id=test, got: %s", res)
	}

	res, err := xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if !strings.Contains(string(res), `id="test"`) {
		t.Fatalf("Expected id attribute is missing in serialised query: %s", string(res))
	}

	q.ID("")
	res, err = xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if strings.Contains(string(res), `id=`) {
		t.Fatalf("Expected id attribute to be missing in serialised query: %s", string(res))
	}
}

func TestQueryIncludeDeletedObjects(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).IncludeDeletedObjects(true)
	if res := q.query.IncludeDeletedObjects; !res {
		t.Fatalf("Expected query with includedeletedobjects=true, got: %t", res)
	}

	res, err := xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if !strings.Contains(string(res), `includedeletedobjects="true"`) {
		t.Fatalf("Expected includedeletedobjects attribute is missing in serialised query: %s", string(res))
	}

	q.IncludeDeletedObjects(false)
	res, err = xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if strings.Contains(string(res), `includedeletedobjects=`) {
		t.Fatalf("Expected includedeletedobjects attribute to be missing in serialised query: %s", string(res))
	}
}

func TestQueryLimit(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).Limit(1)
	if res := q.query.Limit; res != 1 {
		t.Fatalf("Expected query with limit=1, got: %d", res)
	}

	res, err := xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if !strings.Contains(string(res), `limit="1"`) {
		t.Fatalf("Expected limit attribute is missing in serialised query: %s", string(res))
	}

	q.Limit(0)
	res, err = xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if strings.Contains(string(res), `limit=`) {
		t.Fatalf("Expected limit attribute to be missing in serialised query: %s", string(res))
	}
}

func TestQueryOrderBy(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).OrderBy("attr asc, attr desc")
	if res := q.query.OrderBy; res != "attr asc, attr desc" {
		t.Fatalf("Expected query with orderby=attr asc, attr desc, got: %s", res)
	}

	res, err := xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if !strings.Contains(string(res), `orderby="attr asc, attr desc"`) {
		t.Fatalf("Expected orderby attribute is missing in serialised query: %s", string(res))
	}

	q.OrderBy("")
	res, err = xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if strings.Contains(string(res), `orderby=`) {
		t.Fatalf("Expected orderby attribute to be missing in serialised query: %s", string(res))
	}
}

func TestQuerySkip(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).Skip(1)
	if res := q.query.Skip; res != 1 {
		t.Fatalf("Expected query with skip=1, got: %d", res)
	}

	res, err := xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if !strings.Contains(string(res), `skip="1"`) {
		t.Fatalf("Expected skip attribute is missing in serialised query: %s", string(res))
	}

	q.Skip(0)
	res, err = xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if strings.Contains(string(res), `skip=`) {
		t.Fatalf("Expected skip attribute to be missing in serialised query: %s", string(res))
	}
}

func TestQueryLastModified(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).LastModified(true)
	if res := q.query.LastModified; !res {
		t.Fatalf("Expected query with lastmodified=true, got: %t", res)
	}

	res, err := xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if !strings.Contains(string(res), `lastmodified="true"`) {
		t.Fatalf("Expected lastmodified attribute is missing in serialised query: %s", string(res))
	}

	q.LastModified(false)
	res, err = xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if strings.Contains(string(res), `lastmodified=`) {
		t.Fatalf("Expected lastmodified attribute to be missing in serialised query: %s", string(res))
	}
}

func TestQueryChangeID(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).ChangeID("1")
	if res := q.query.ChangeID; res != nil && *res != "1" {
		t.Fatalf("Expected query with changeid=1, got: %s", *res)
	}

	res, err := xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if !strings.Contains(string(res), `changeid="1"`) {
		t.Fatalf("Expected changeid attribute is missing in serialised query: %s", string(res))
	}

	q.ChangeID("")
	if res := q.query.ChangeID; res != nil {
		t.Fatalf("Expected query without changeid, got: %s", *res)
	}
	res, err = xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if strings.Contains(string(res), `changeid=`) {
		t.Fatalf("Expected changeid attribute to be missing in serialised query: %s", string(res))
	}
}

func TestQuerySSEURL(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).SSEURL(true)
	if res := q.query.SSEURL; !res {
		t.Fatalf("Expected query with sseurl=true, got: %t", res)
	}

	res, err := xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if !strings.Contains(string(res), `sseurl="true"`) {
		t.Fatalf("Expected sseurl attribute is missing in serialised query: %s", string(res))
	}

	q.SSEURL(false)
	res, err = xml.Marshal(&q)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if strings.Contains(string(res), `sseurl=`) {
		t.Fatalf("Expected sseurl attribute to be missing in serialised query: %s", string(res))
	}
}

func TestQueryDistinct(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).Distinct("test")
	if q.query.Distinct != "test" {
		t.Fatalf("Expected distinct value of: test, got: %s", q.query.Distinct)
	}

	q2 := NewQuery(WeatherStation, 1.0).Distinct("test").Distinct("foo")
	if q2.query.Distinct != "foo" {
		t.Fatalf("Expected distinct value of: foo, got: %s", q.query.Distinct)
	}
}

func TestQueryInclude(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []string
		Output []string
	}{
		{Name: "One", Input: []string{"one"}, Output: []string{"one"}},
		{Name: "Two", Input: []string{"one", "two"}, Output: []string{"one", "two"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			q := NewQuery(WeatherStation, 1.0).Include(tt.Input...)
			if len(q.query.Include) != len(tt.Output) {
				t.Fatalf("Expected: %d elements, got: %d", len(tt.Output), len(q.query.Include))
			}
			if !reflect.DeepEqual(q.query.Include, tt.Output) {
				t.Fatalf("Expected: %#v, got: %#v", tt.Output, q.query.Include)
			}
		})
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Fluent%s", tt.Name), func(t *testing.T) {
			t.Parallel()
			q := NewQuery(WeatherStation, 1.0)
			for _, in := range tt.Input {
				q.Include(in)
			}
			t.Log(q.query.Include)
			if len(q.query.Include) != len(tt.Output) {
				t.Fatalf("Expected: %d elements, got: %d", len(tt.Output), len(q.query.Include))
			}
			if !reflect.DeepEqual(q.query.Include, tt.Output) {
				t.Fatalf("Expected: %#v, got: %#v", tt.Output, q.query.Include)
			}
		})
	}
}

func TestQueryExclude(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []string
		Output []string
	}{
		{Name: "One", Input: []string{"one"}, Output: []string{"one"}},
		{Name: "Two", Input: []string{"one", "two"}, Output: []string{"one", "two"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			q := NewQuery(WeatherStation, 1.0).Exclude(tt.Input...)
			if len(q.query.Exclude) != len(tt.Output) {
				t.Fatalf("Expected: %d elements, got: %d", len(tt.Output), len(q.query.Exclude))
			}
			if !reflect.DeepEqual(q.query.Exclude, tt.Output) {
				t.Fatalf("Expected: %#v, got: %#v", tt.Output, q.query.Exclude)
			}
		})
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Fluent%s", tt.Name), func(t *testing.T) {
			t.Parallel()
			q := NewQuery(WeatherStation, 1.0)
			for _, in := range tt.Input {
				q.Exclude(in)
			}
			if len(q.query.Exclude) != len(tt.Output) {
				t.Fatalf("Expected: %d elements, got: %d", len(tt.Output), len(q.query.Exclude))
			}
			if !reflect.DeepEqual(q.query.Exclude, tt.Output) {
				t.Fatalf("Expected: %#v, got: %#v", tt.Output, q.query.Exclude)
			}
		})
	}
}

func TestQueryFilter(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).Filter(Equal("a", "b"))
	if len(q.query.Filter.filter.Equal) != 1 {
		t.Fatalf("Expected 1 filter, got: %d", len(q.query.Filter.filter.Equal))
	}
	if q.query.Filter.filter.Equal[0].Name != "a" {
		t.Fatalf("Expected Equal filter with name=a, got name=%s", q.query.Filter.filter.Equal[0].Name)
	}
	if q.query.Filter.filter.Equal[0].Value != "b" {
		t.Fatalf("Expected Equal filter with value=b, got value=%s", q.query.Filter.filter.Equal[0].Value)
	}
}

func TestQueryEval(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).Eval(Eval("a", "b"))
	if len(q.query.Eval) != 1 {
		t.Fatalf("Expected 1 evaluation, got: %d", len(q.query.Eval))
	}
	if res := q.query.Eval[0].evaluation.Alias; res != "a" {
		t.Fatalf("Expected Eval with alias=a, got alias=%s", res)
	}
	if res := q.query.Eval[0].evaluation.Function; res != "b" {
		t.Fatalf("Expected Equal filter with function=b, got function=%s", res)
	}

	q.Eval(Eval("c", "d"))
	if len(q.query.Eval) != 2 {
		t.Fatalf("Expected 2 evaluations, got: %d", len(q.query.Eval))
	}
}

func TestQueryMarshalXML(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0)

	res, err := xml.Marshal(q)
	if err != nil {
		t.Fatal(err)
	}

	exp := `<Query objecttype="WeatherStation" schemaversion="1"></Query>`
	if string(res) != exp {
		t.Fatalf("Expected: %s, got: %s", exp, string(res))
	}
}

func TestEvaluation(t *testing.T) {
	e := Eval("a", "b")
	exp := Evaluation{
		evaluation: evaluation{
			Alias:    "a",
			Function: "b",
		},
	}

	if !reflect.DeepEqual(e, exp) {
		t.Fatalf("Expected: %#v, got: %#v", exp, e)
	}
}

func TestEvaluationMarshalXML(t *testing.T) {
	e := Eval("a", "b")
	res, err := xml.Marshal(&e)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}

	exp := `<Evaluation alias="a" function="b"></Evaluation>`
	if string(res) != exp {
		t.Fatalf("Expected: %s, got %s", exp, string(res))
	}
}

func TestCompositeFilter(t *testing.T) {
	tests := []struct {
		Name    string
		Func    func(filters ...Filter) Filter
		Filters []Filter
		Result  Filter
	}{
		{
			Name:    "And",
			Func:    And,
			Filters: []Filter{Equal("a", "b"), Equal("c", "d")},
			Result: Filter{filter: filter{
				And: []filter{{
					Equal: []comparison{
						{Name: "a", Value: "b"},
						{Name: "c", Value: "d"},
					},
				}},
			}},
		},
		{
			Name:    "Or",
			Func:    Or,
			Filters: []Filter{Equal("a", "b"), Equal("c", "d")},
			Result: Filter{filter: filter{
				Or: []filter{{
					Equal: []comparison{
						{Name: "a", Value: "b"},
						{Name: "c", Value: "d"},
					},
				}},
			}},
		},
		{
			Name:    "Not",
			Func:    Not,
			Filters: []Filter{Equal("a", "b"), Equal("c", "d")},
			Result: Filter{filter: filter{
				Not: []filter{{
					Equal: []comparison{
						{Name: "a", Value: "b"},
						{Name: "c", Value: "d"},
					},
				}},
			}},
		},
		{
			Name:    "ElementMatch",
			Func:    ElementMatch,
			Filters: []Filter{Equal("a", "b"), Equal("c", "d")},
			Result: Filter{filter: filter{
				ElementMatch: []filter{{
					Equal: []comparison{
						{Name: "a", Value: "b"},
						{Name: "c", Value: "d"},
					},
				}},
			}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			f := tt.Func(tt.Filters...)

			if !reflect.DeepEqual(tt.Result, f) {
				t.Fatalf("Did not get expected filter set: %+v", f)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	f := Equal("field", "something")
	if len(f.filter.Equal) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Equal))
	}
	comp := f.filter.Equal[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestExists(t *testing.T) {
	f := Exists("field", true)
	if len(f.filter.Exists) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Exists))
	}
	comp := f.filter.Exists[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "true" {
		t.Fatalf("Expected filter with value=true, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestGreaterThan(t *testing.T) {
	f := GreaterThan("field", "something")
	if len(f.filter.GreaterThan) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.GreaterThan))
	}
	comp := f.filter.GreaterThan[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=true, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	f := GreaterThanOrEqual("field", "something")
	if len(f.filter.GreaterThanOrEqual) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.GreaterThanOrEqual))
	}
	comp := f.filter.GreaterThanOrEqual[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestLessThan(t *testing.T) {
	f := LessThan("field", "something")
	if len(f.filter.LessThan) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.LessThan))
	}
	comp := f.filter.LessThan[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestLeesThanOrEqual(t *testing.T) {
	f := LessThanOrEqual("field", "something")
	if len(f.filter.LessThanOrEqual) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.LessThanOrEqual))
	}
	comp := f.filter.LessThanOrEqual[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestNotEqual(t *testing.T) {
	f := NotEqual("field", "something")
	if len(f.filter.NotEqual) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.NotEqual))
	}
	comp := f.filter.NotEqual[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestLike(t *testing.T) {
	f := Like("field", "something")
	if len(f.filter.Like) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Like))
	}
	comp := f.filter.Like[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestNotLike(t *testing.T) {
	f := NotLike("field", "something")
	if len(f.filter.NotLike) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.NotLike))
	}
	comp := f.filter.NotLike[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestIn(t *testing.T) {
	f := In("field", "something")
	if len(f.filter.In) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.In))
	}
	comp := f.filter.In[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestNotIn(t *testing.T) {
	f := NotIn("field", "something")
	if len(f.filter.NotIn) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.NotIn))
	}
	comp := f.filter.NotIn[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestWithin(t *testing.T) {
	f := Within("field", "something", "box", 2.5)
	if len(f.filter.Within) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Within))
	}
	comp := f.filter.Within[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 2.5 {
		t.Fatalf("Expected filter with radius=2.5, but got radius=%f", radius)
	}
	if shape := comp.Shape; shape != "box" {
		t.Fatalf("Expected filter with shape=box, but got shape=%s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestIntersects(t *testing.T) {
	f := Intersects("field", "something")
	if len(f.filter.Intersects) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Intersects))
	}
	comp := f.filter.Intersects[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestNear(t *testing.T) {
	f := Near("field", "something", 1, 10)
	if len(f.filter.Near) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Near))
	}
	comp := f.filter.Near[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := *comp.MinDistance; dst != 1 {
		t.Fatalf("Expected filter with mindistance=1, but got mindistance=%d", dst)
	}
	if dst := *comp.MaxDistance; dst != 10 {
		t.Fatalf("Expected filter with maxdistance=10, but got maxdistance=%d", dst)
	}
}

func TestFilterMarshalXML(t *testing.T) {
	f := Equal("a", "b")

	res, err := xml.Marshal(&f)
	if err != nil {
		t.Fatal(err)
	}

	exp := `<Filter><EQ name="a" value="b"></EQ></Filter>`
	if string(res) != exp {
		t.Fatalf("Expected: %s, got: %s", exp, string(res))
	}
}
