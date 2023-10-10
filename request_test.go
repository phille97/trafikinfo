package trafikinfo

import (
	"testing"
)

func TestRequestQuery(t *testing.T) {
	tests := []struct {
		Name  string
		Input []*Query
	}{
		{
			Name: "Single",
			Input: []*Query{
				NewQuery(WeatherStation, 1.0),
			},
		},
		{
			Name: "Triple",
			Input: []*Query{
				NewQuery(WeatherStation, 1.0),
				NewQuery(WeatherMeasurepoint, 1.0),
				NewQuery(WeatherObservation, 1.0),
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
	r, err := NewRequest().Build()
	if err != nil {
		t.Fatalf("Expected success, got: %v", err)
	}

	exp := `<REQUEST></REQUEST>`
	if res := string(r); res != exp {
		t.Fatalf("Expected: %s, got: %s", exp, res)
	}
}
