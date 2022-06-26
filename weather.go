package trafikinfo

import "time"

type WeatherStation1Dot0 struct {
	Active      *bool     `json:"Active,omitempty"`
	County      []County  `json:"CountyNo,omitempty"`
	Deleted     *bool     `json:"Deleted,omitempty"`
	Geometry    *Geometry `json:"Geometry,omitempty"`
	IconID      *string   `json:"IconId,omitempty"`
	ID          *string   `json:"Id,omitempty"`
	Measurement *struct {
		Air *struct {
			RelativeHumidity  *float64 `json:"RelativeHumidity,omitempty"`
			Temperature       *float64 `json:"Temp,omitempty"`
			TemperatureIconID *string  `json:"TempIconId,omitempty"`
		} `json:"Air,omitempty"`
		MeasureTime   *time.Time `json:"MeasureTime,omitempty"`
		Precipitation *struct {
			Amount     *float64             `json:"Amount,omitempty"`
			AmountName *PrecipitationAmount `json:"AmountName,omitempty"`
			Type       *PrecipitationType   `json:"Type,omitempty"`
			TypeIconID *string              `json:"TypeIconId,omitempty"`
		} `json:"Precipitation,omitempty"`
		Road *struct {
			Temperature       *float64 `json:"Temp,omitempty"`
			TemperatureIconID *string  `json:"TempIconId,omitempty"`
		} `json:"Road,omitempty"`
		Wind *struct {
			Direction       *float64       `json:"Direction,omitempty"`
			DirectionIconID *string        `json:"DirectionIconId,omitempty"`
			DirectionText   *WindDirection `json:"DirectionText,omitempty"`
			Force           *float64       `json:"Force,omitempty"`
			ForceMax        *float64       `json:"ForceMax,omitempty"`
		} `json:"Wind,omitempty"`
	} `json:"Measurement,omitempty"`
	MeasurementHistory []struct {
		Air *struct {
			RelativeHumidity  *float64 `json:"RelativeHumidity,omitempty"`
			Temperature       *float64 `json:"Temp,omitempty"`
			TemperatureIconID *string  `json:"TempIconId,omitempty"`
		} `json:"Air,omitempty"`
		MeasureTime   *time.Time `json:"MeasureTime,omitempty"`
		Precipitation *struct {
			Amount     *float64             `json:"Amount,omitempty"`
			AmountName *PrecipitationAmount `json:"AmountName,omitempty"`
			Type       *PrecipitationType   `json:"Type,omitempty"`
			TypeIconID *string              `json:"TypeIconId,omitempty"`
		} `json:"Precipitation,omitempty"`
		Road *struct {
			Temperature       *float64 `json:"Temp,omitempty"`
			TemperatureIconID *string  `json:"TempIconId,omitempty"`
		} `json:"Road,omitempty"`
		Wind *struct {
			Direction       *float64       `json:"Direction,omitempty"`
			DirectionIconID *string        `json:"DirectionIconId,omitempty"`
			DirectionText   *WindDirection `json:"DirectionText,omitempty"`
			Force           *float64       `json:"Force,omitempty"`
			ForceMax        *float64       `json:"ForceMax,omitempty"`
		} `json:"Wind,omitempty"`
	} `json:"MeasurementHistory,omitempty"`
	ModifiedTime *time.Time `json:"ModifiedTime,omitempty"`
	Name         *string    `json:"Name,omitempty"`
	RoadNumber   *int       `json:"RoadNumberNumeric,omitempty"`
}

type PrecipitationAmount string

const (
	PrecipitationDataMissing PrecipitationAmount = "Givare saknas/Fel på givare"
	PrecipitationLightRain   PrecipitationAmount = "Lätt regn"
	PrecipitationMildRain    PrecipitationAmount = "Måttligt regn"
	PrecipitationHeavyRain   PrecipitationAmount = "Kraftigt regn"

	PrecipitationLightSleet PrecipitationAmount = "Lätt snöblandat regn"
	PrecipitationMildSleet  PrecipitationAmount = "Måttligt snöblandat regn"
	PrecipitationHeavySleet PrecipitationAmount = "Kraftigt snöblandat regn"

	PrecipitationLightSnow PrecipitationAmount = "Lätt snöfall"
	PrecipitationMindSnow  PrecipitationAmount = "Måttligt snöfall"
	PrecipitationHeavySnow PrecipitationAmount = "Kraftigt snöfall"

	PrecipitationOther   PrecipitationAmount = "Annan nederbördstyp"
	PrecipitationUnknown PrecipitationAmount = "Okänd nederbördstyp"
	PrecipitationNone    PrecipitationAmount = "Ingen nederbörd"
)

type PrecipitationType string

const (
	PrecipitationTypeDrizzle      PrecipitationType = "Duggregn"
	PrecipitationTypeHail         PrecipitationType = "Hagel"
	PrecipitationTypeRain         PrecipitationType = "Regn"
	PrecipitationTypeSnow         PrecipitationType = "Snö"
	PrecipitationTypeSleet        PrecipitationType = "Snöblandat regn"
	PrecipitationTypeFreezingRain PrecipitationType = "Underkylt regn"
	PrecipitationTypeNone         PrecipitationType = "Ingen nederbörd"
)

type WindDirection string

const (
	WindDirectionN   WindDirection = "Norr"
	WindDirectionNE  WindDirection = "Nordöst"
	WindDirectionNNE WindDirection = "Nordnordöst"
	WindDirectionNNW WindDirection = "Nordnordväst"
	WindDirectionNW  WindDirection = "Nordväst"

	WindDirectionE   WindDirection = "Öst"
	WindDirectionESE WindDirection = "Östsydöst"

	WindDirectionS   WindDirection = "Söder"
	WindDirectionSE  WindDirection = "Sydöst"
	WindDirectionSSW WindDirection = "Sydsydväst"
	WindDirectionSW  WindDirection = "Sydväst"

	WindDirectionW WindDirection = "Väst"
)
