package trafikinfo

import "time"

type WeatherStation1Dot0 struct {
	Active             *bool                `json:"Active,omitempty"`
	County             []County             `json:"CountyNo,omitempty"`
	Deleted            *bool                `json:"Deleted,omitempty"`
	Geometry           *Geometry            `json:"Geometry,omitempty"`
	IconID             *string              `json:"IconId,omitempty"`
	ID                 *string              `json:"Id,omitempty"`
	Measurement        *WeatherMeasurement  `json:"Measurement,omitempty"`
	MeasurementHistory []WeatherMeasurement `json:"MeasurementHistory,omitempty"`
	ModifiedTime       *time.Time           `json:"ModifiedTime,omitempty"`
	Name               *string              `json:"Name,omitempty"`
	RoadNumber         *int                 `json:"RoadNumberNumeric,omitempty"`
}

type WeatherMeasurement struct {
	Air *struct {
		RelativeHumidity  *float64 `json:"RelativeHumidity,omitempty"`
		Temperature       *float64 `json:"Temp,omitempty"`
		TemperatureIconID *string  `json:"TempIconId,omitempty"`
	} `json:"Air,omitempty"`
	MeasureTime   *time.Time `json:"MeasureTime,omitempty"`
	Precipitation *struct {
		Amount     *float64                 `json:"Amount,omitempty"`
		AmountName *PrecipitationAmountName `json:"AmountName,omitempty"`
		Type       *PrecipitationType       `json:"Type,omitempty"`
		TypeIconID *string                  `json:"TypeIconId,omitempty"`
	} `json:"Precipitation,omitempty"`
	Road *Road `json:"Road,omitempty"`
	Wind *struct {
		Direction       *float64       `json:"Direction,omitempty"`
		DirectionIconID *string        `json:"DirectionIconId,omitempty"`
		DirectionText   *WindDirection `json:"DirectionText,omitempty"`
		Force           *float64       `json:"Force,omitempty"`
		ForceMax        *float64       `json:"ForceMax,omitempty"`
	} `json:"Wind,omitempty"`
}

type PrecipitationAmountName string

const (
	PrecipitationDataMissing PrecipitationAmountName = "Givare saknas/Fel på givare"
	PrecipitationLightRain   PrecipitationAmountName = "Lätt regn"
	PrecipitationMildRain    PrecipitationAmountName = "Måttligt regn"
	PrecipitationHeavyRain   PrecipitationAmountName = "Kraftigt regn"

	PrecipitationLightSleet PrecipitationAmountName = "Lätt snöblandat regn"
	PrecipitationMildSleet  PrecipitationAmountName = "Måttligt snöblandat regn"
	PrecipitationHeavySleet PrecipitationAmountName = "Kraftigt snöblandat regn"

	PrecipitationLightSnow PrecipitationAmountName = "Lätt snöfall"
	PrecipitationMindSnow  PrecipitationAmountName = "Måttligt snöfall"
	PrecipitationHeavySnow PrecipitationAmountName = "Kraftigt snöfall"

	PrecipitationOther   PrecipitationAmountName = "Annan nederbördstyp"
	PrecipitationUnknown PrecipitationAmountName = "Okänd nederbördstyp"
	PrecipitationNone    PrecipitationAmountName = "Ingen nederbörd"
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
