package trafikinfo

import (
	"testing"

	"code.dny.dev/trafikinfo/trv"
	wmp "code.dny.dev/trafikinfo/trv/weathermeasurepoint/v2"
	wo "code.dny.dev/trafikinfo/trv/weatherobservation/v2"
)

func TestRequestQuery(t *testing.T) {
	tests := []struct {
		Name  string
		Input []Query
	}{
		{
			Name: "Single",
			Input: []Query{
				NewQuery(wmp.ObjectType()),
			},
		},
		{
			Name: "Double",
			Input: []Query{
				NewQuery(wmp.ObjectType()),
				NewQuery(wo.ObjectType()),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			r := NewRequest().Query(tt.Input...)
			if qlen := len(r.Queries); qlen != len(tt.Input) {
				t.Fatalf("Expected %d queries, got: %d", len(tt.Input), qlen)
			}
		})
	}
}

func TestRequestLogin(t *testing.T) {
	r := NewRequest().APIKey("test")
	if key := r.Login.AuthenticationKey; key != "test" {
		t.Fatalf("Failed to set API credential, expected: test, got: %s", key)
	}
}

func TestRequestBuild(t *testing.T) {
	tests := []struct {
		name string
		in   *Request
		out  string
		err  bool
	}{
		{name: "without API key", in: NewRequest(), out: "", err: true},
		{name: "with API key but no query", in: NewRequest().APIKey("test"), out: "", err: true},
		{name: "with API key and query", in: NewRequest().APIKey("test").Query(NewQuery(trv.ObjectType{})), out: `<REQUEST><LOGIN authenticationkey="test"></LOGIN><QUERY objecttype="" schemaversion="" lastmodified="true" changeid="0"></QUERY></REQUEST>`},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.in.Build()
			if tt.err && err == nil {
				t.Fatalf("Expected error but didn't get one")
			}
			if !tt.err && err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			}
			if out := string(res); out != tt.out {
				t.Fatalf("Expected: %s, got: %s", tt.out, out)
			}
		})
	}
}
