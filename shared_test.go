package trafikinfo

import "testing"

func TestCountyNumberString(t *testing.T) {
	tests := []struct {
		In  County
		Out string
	}{
		{In: 1000, Out: "Okänt län (1000)"},
		{In: CountyStockholm, Out: "Stockholms län"},
		{In: CountyStockholmDeprecated, Out: "Stockholms län"},
		{In: CountyUppsala, Out: "Uppsala län"},
		{In: CountySodermanland, Out: "Södermanlands län"},
		{In: CountyOstergotland, Out: "Östergötlands län"},
		{In: CountyJonkoping, Out: "Jönköpings län"},
		{In: CountyKronobergs, Out: "Kronobergs län"},
		{In: CountyKalmar, Out: "Kalmar län"},
		{In: CountyGotland, Out: "Gotlands län"},
		{In: CountyBlekinge, Out: "Blekinge län"},
		{In: CountySkane, Out: "Skåne län"},
		{In: CountyHallands, Out: "Hallands län"},
		{In: CountyVastraGotaland, Out: "Västra Götalands län"},
		{In: CountyVarmland, Out: "Värmlands län"},
		{In: CountyOrebro, Out: "Örebro län"},
		{In: CountyVastmanland, Out: "Västmanlands län"},
		{In: CountyDalarna, Out: "Dalarnas län"},
		{In: CountyGavleborg, Out: "Gävleborgs län"},
		{In: CountyVasternorrland, Out: "Västernorrlands län"},
		{In: CountyJamtland, Out: "Jämtlands län"},
		{In: CountyVasterbotten, Out: "Västerbottens län"},
		{In: CountyNorrbotten, Out: "Norrbottens län"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.In.String(), func(t *testing.T) {
			t.Parallel()
			if tt.In.String() != tt.Out {
				t.Fatalf("Expected: %s, got: %s", tt.Out, tt.In.String())
			}
		})
	}
}

func TestRegionString(t *testing.T) {
	tests := []struct {
		In  Region
		Out string
	}{
		{In: 1000, Out: "Okänd region (1000)"},
		{In: RegionNorth, Out: "Region Norr"},
		{In: RegionMiddle, Out: "Region Mitt"},
		{In: RegionEast, Out: "Region Öst"},
		{In: RegionStockholm, Out: "Region Stockholm"},
		{In: RegionWest, Out: "Region Väst"},
		{In: RegionSouth, Out: "Region Syd"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.In.String(), func(t *testing.T) {
			t.Parallel()
			if tt.In.String() != tt.Out {
				t.Fatalf("Expected: %s, got: %s", tt.Out, tt.In.String())
			}
		})
	}
}

func TestConditionCodeString(t *testing.T) {
	tests := []struct {
		In  ConditionCode
		Out string
	}{
		{In: 1000, Out: "Okänd väglagskod (1000)"},
		{In: ConditionCodeDifficult, Out: "besvärligt (risk för)"},
		{In: ConditionCodeVeryDifficult, Out: "mycket besvärligt"},
		{In: ConditionCodeNormal, Out: "normalt"},
		{In: ConditionCodeIceAndSnow, Out: "is- och snövägbana"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.In.String(), func(t *testing.T) {
			t.Parallel()
			if tt.In.String() != tt.Out {
				t.Fatalf("Expected: %s, got: %s", tt.Out, tt.In.String())
			}
		})
	}
}

func TestStatusString(t *testing.T) {
	tests := []struct {
		In  Status
		Out string
	}{
		{In: 1000, Out: "Okänd status (1000)"},
		{In: StatusFreeAccess, Out: "fri framkomlighet"},
		{In: StatusPassabilityDifficult, Out: "svår framkomlighet"},
		{In: StatusPassabilityImpossile, Out: "framkomlighet omöjlig"},
		{In: StatusPassabilityUnknown, Out: "framkomlighet okänd"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.In.String(), func(t *testing.T) {
			t.Parallel()
			if tt.In.String() != tt.Out {
				t.Fatalf("Expected: %s, got: %s", tt.Out, tt.In.String())
			}
		})
	}
}
