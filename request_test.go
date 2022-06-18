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
				NewQuery(WeatherMeasurePoint, 1.0),
				NewQuery(WeatherObservation, 1.0),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			var r *Request
			if len(tt.Input) == 1 {
				r = NewRequest().Query(tt.Input[0])
			} else {
				r = NewRequest().Query(tt.Input[0], tt.Input[1:]...)
			}
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
