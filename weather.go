package trafikinfo

import "time"

// WeatherStation1Dot0 is the WeatherStation v1.0 model. It returns all kinds
// of environmental data
type WeatherStation1Dot0 struct {
	// Active indicates this station is returning data
	Active *bool `json:"Active,omitempty"`
	// County is the Swedish county in which this station is located
	County []County `json:"CountyNo,omitempty"`
	// Deleted indicates if this is a deleted object. This should only
	// ever be true if you created a Query with IncludeDeletedItems
	Deleted *bool `json:"Deleted,omitempty"`
	// Geometry contains coordinates for where this station is located
	Geometry *Geometry `json:"Geometry,omitempty"`
	IconID   *string   `json:"IconId,omitempty"`
	// ID is the unique ID identifying this station
	ID *string `json:"Id,omitempty"`
	// Measurement represents measurement data from environmental sensors
	Measurement *struct {
		// Air contains the typical things like relative humidity and temperature
		// as measured in the air
		Air *struct {
			// RelativeHumidity is the relative humidity in percent
			RelativeHumidity *float64 `json:"RelativeHumidity,omitempty"`
			// Temperature is measured in degrees Celsius
			Temperature       *float64 `json:"Temp,omitempty"`
			TemperatureIconID *string  `json:"TempIconId,omitempty"`
		} `json:"Air,omitempty"`
		// MeasurTime indicates when these values where recorded
		MeasureTime *time.Time `json:"MeasureTime,omitempty"`
		// Precipitation contains the type and amount of precipitation
		Precipitation *struct {
			// Amount is the mount of rain in mm per hour. It's worth noting
			// that in practice it turns out the absence of a value indicates
			// "no rain", i.e precipitation of 0, not "this station doesn't
			// report precipitation", unless Measurement was Excluded by the
			// Query
			Amount *float64 `json:"Amount,omitempty"`
			// AmountName is a predefined string describing the type and amount
			// of precipitation in Swedish
			AmountName *PrecipitationAmount `json:"AmountName,omitempty"`
			// Type is a predefined string describing the type of precipitation
			// in Swedish
			Type       *PrecipitationType `json:"Type,omitempty"`
			TypeIconID *string            `json:"TypeIconId,omitempty"`
		} `json:"Precipitation,omitempty"`
		// Road contains the temperature as measured at the road deck level.
		// This can be much higher or lower than the Air values
		Road *struct {
			// Temperature is measured in degrees Celsius
			Temperature       *float64 `json:"Temp,omitempty"`
			TemperatureIconID *string  `json:"TempIconId,omitempty"`
		} `json:"Road,omitempty"`
		// Wind contains wind direction and strength
		Wind *struct {
			// Direction is represented in degrees
			Direction       *float64 `json:"Direction,omitempty"`
			DirectionIconID *string  `json:"DirectionIconId,omitempty"`
			// DirectionText is a predefined string representing the wind
			// direction in Swedish
			DirectionText *WindDirection `json:"DirectionText,omitempty"`
			// Force is the 10 minute average in m/sec
			Force *float64 `json:"Force,omitempty"`
			// ForceMax is the 30 minute max in m/sec
			ForceMax *float64 `json:"ForceMax,omitempty"`
		} `json:"Wind,omitempty"`
	} `json:"Measurement,omitempty"`
	// MeasurementHistory contains multiple instances of Measurement
	// over time. See the comments on Measurement for how to interpret
	// the data
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
	// ModifiedTime represents when this object was last modified
	ModifiedTime *time.Time `json:"ModifiedTime,omitempty"`
	// Name is the name of the station
	Name *string `json:"Name,omitempty"`
	// RoadNumber is the Swedish road number this station is located at
	RoadNumber *int `json:"RoadNumberNumeric,omitempty"`
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

type WeatherObservation1Dot0 struct {
	Aggregated10Minutes *struct {
		Precipitation *struct {
			Rain    *bool `json:"Rain,omitempty"`
			RainSum *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"TotalWaterEquivalent,omitempty"`
		} `json:"Precipitation,omitempty"`
		Wind *struct {
			Height       *int `json:"Height,omitempty"`
			SpeedAverage *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedAverage,omitempty"`
			SpeedMax *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedMax,omitempty"`
		} `json:"Wind,omitempty"`
	} `json:"Aggregated10minutes,omitempty"`
	Aggregated30Minutes *struct {
		Precipitation *struct {
			Rain    *bool `json:"Rain,omitempty"`
			RainSum *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"TotalWaterEquivalent,omitempty"`
		} `json:"Precipitation,omitempty"`
		Wind *struct {
			Height       *int `json:"Height,omitempty"`
			SpeedAverage *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedAverage,omitempty"`
			SpeedMax *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedMax,omitempty"`
		} `json:"Wind,omitempty"`
	} `json:"Aggregated30minutes,omitempty"`
	Aggregated5Minutes struct {
		Precipitation *struct {
			Rain    *bool `json:"Rain,omitempty"`
			RainSum *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"TotalWaterEquivalent,omitempty"`
		} `json:"Precipitation,omitempty"`
		Wind *struct {
			Height       *int `json:"Height,omitempty"`
			SpeedAverage *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedAverage,omitempty"`
			SpeedMax *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedMax,omitempty"`
		} `json:"Wind,omitempty"`
	} `json:"Aggregated5minutes,omitempty"`
	Air *struct {
		Dewpoint *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Dewpoint,omitempty"`
		RelativeHumidity *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"RelativeHumidity,omitempty"`
		Temperature *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Temperature,omitempty"`
		VisibleDistance *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"VisibleDistance,omitempty"`
	} `json:"Air,omitempty"`
	DeicingChemical *struct {
		Amount *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Amount,omitempty"`
	} `json:"DeicingChemical,omitempty"`
	Deleted      *bool      `json:"Deleted,omitempty"`
	ID           *string    `json:"Id,omitempty"`
	ModifiedTime *time.Time `json:"ModifiedTime,omitempty"`
	Measurepoint *struct {
		ID       *int      `json:"Id,omitempty"`
		Name     *string   `json:"Name,omitempty"`
		Geometry *Geometry `json:"Geometry,omitempty"`
	} `json:"Measurepoint,omitempty"`
	Sample     *time.Time `json:"Sample,omitempty"`
	Subsurface *struct {
		Ground []struct {
			Depth *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Depth,omitempty"`
			Temperature *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Temperature,omitempty"`
		} `json:"Ground,omitempty"`
	} `json:"Subsurface,omitempty"`
	Surface *struct {
		Grip *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Grip,omitempty"`
		Ice      *bool `json:"Ice,omitempty"`
		IceDepth *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"IceDepth,omitempty"`
		Snow      *bool `json:"Snow,omitempty"`
		SnowDepth *struct {
			Solid *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Solid,omitempty"`
			WaterEquivalent *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"WaterEquivalent,omitempty"`
		} `json:"SnowDepth,omitempty"`
		Temperature *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Temperature,omitempty"`
		Water      *bool `json:"Water,omitempty"`
		WaterDepth *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"WaterDepth,omitempty"`
	} `json:"Surface,omitempty"`
	Wind [2]struct {
		Direction *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Direction,omitempty"`
		Height *int `json:"Height,omitempty"`
		Speed  *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Speed,omitempty"`
	} `json:"Wind,omitempty"`
}

type PrecipitationTypeMeasure string

const (
	PrecipitationTypeMeasureYes PrecipitationTypeMeasure = "yes"
	PrecipitationTypeMeasureNo  PrecipitationTypeMeasure = "no"

	PrecipitationTypeMeasureRain         PrecipitationTypeMeasure = "rain"
	PrecipitationTypeMeasureFreezingRain PrecipitationTypeMeasure = "freezing_rain"
	PrecipitationTypeMeasureSnow         PrecipitationTypeMeasure = "snow"
	PrecipitationTypeMeasureSleet        PrecipitationTypeMeasure = "sleet"
)

type Origin string

const (
	OriginCalculated Origin = "calculated"
	OriginEstimated  Origin = "estimated"
	OriginMeasured   Origin = "measured"
)

type WeatherObservation2Dot0 struct {
	Aggregated10Minutes *struct {
		Precipitation *struct {
			Rain    *bool `json:"Rain,omitempty"`
			RainSum *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"TotalWaterEquivalent,omitempty"`
		} `json:"Precipitation,omitempty"`
		Wind *struct {
			Height       *int `json:"Height,omitempty"`
			SpeedAverage *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedAverage,omitempty"`
			SpeedMax *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedMax,omitempty"`
		} `json:"Wind,omitempty"`
	} `json:"Aggregated10minutes,omitempty"`
	Aggregated30Minutes *struct {
		Precipitation *struct {
			Rain    *bool `json:"Rain,omitempty"`
			RainSum *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"TotalWaterEquivalent,omitempty"`
		} `json:"Precipitation,omitempty"`
		Wind *struct {
			Height       *int `json:"Height,omitempty"`
			SpeedAverage *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedAverage,omitempty"`
			SpeedMax *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedMax,omitempty"`
		} `json:"Wind,omitempty"`
	} `json:"Aggregated30minutes,omitempty"`
	Aggregated5Minutes struct {
		Precipitation *struct {
			Rain    *bool `json:"Rain,omitempty"`
			RainSum *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin `json:"Origin,omitempty"`
				SensorNames *string `json:"SensorNames,omitempty"`
				Value       *int    `json:"Value,omitempty"`
			} `json:"TotalWaterEquivalent,omitempty"`
		} `json:"Precipitation,omitempty"`
		Wind *struct {
			Height       *int `json:"Height,omitempty"`
			SpeedAverage *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedAverage,omitempty"`
			SpeedMax *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"SpeedMax,omitempty"`
		} `json:"Wind,omitempty"`
	} `json:"Aggregated5minutes,omitempty"`
	Air *struct {
		Dewpoint *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Dewpoint,omitempty"`
		RelativeHumidity *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"RelativeHumidity,omitempty"`
		Temperature *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Temperature,omitempty"`
		VisibleDistance *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"VisibleDistance,omitempty"`
	} `json:"Air,omitempty"`
	DeicingChemical *struct {
		Amount *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Amount,omitempty"`
	} `json:"DeicingChemical,omitempty"`
	Deleted      *bool      `json:"Deleted,omitempty"`
	ID           *string    `json:"Id,omitempty"`
	ModifiedTime *time.Time `json:"ModifiedTime,omitempty"`
	Measurepoint *struct {
		ID       *int      `json:"Id,omitempty"`
		Name     *string   `json:"Name,omitempty"`
		Geometry *Geometry `json:"Geometry,omitempty"`
	} `json:"Measurepoint,omitempty"`
	Sample     *time.Time `json:"Sample,omitempty"`
	Subsurface *struct {
		Ground []struct {
			Depth *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Depth,omitempty"`
			Temperature *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Temperature,omitempty"`
		} `json:"Ground,omitempty"`
	} `json:"Subsurface,omitempty"`
	Surface *struct {
		Grip *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Grip,omitempty"`
		Ice      *bool `json:"Ice,omitempty"`
		IceDepth *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"IceDepth,omitempty"`
		Snow      *bool `json:"Snow,omitempty"`
		SnowDepth *struct {
			Solid *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Solid,omitempty"`
			WaterEquivalent *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"WaterEquivalent,omitempty"`
		} `json:"SnowDepth,omitempty"`
		Temperature *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Temperature,omitempty"`
		Water      *bool `json:"Water,omitempty"`
		WaterDepth *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"WaterDepth,omitempty"`
	} `json:"Surface,omitempty"`
	Weather *struct {
		Precipitation *PrecipitationTypeMeasure `json:"Precipitation,omitempty"`
	} `json:"Weather,omitempty"`
	Wind [2]struct {
		Direction *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Direction,omitempty"`
		Height *int `json:"Height,omitempty"`
		Speed  *struct {
			Origin      *Origin  `json:"Origin,omitempty"`
			SensorNames *string  `json:"SensorNames,omitempty"`
			Value       *float64 `json:"Value,omitempty"`
		} `json:"Speed,omitempty"`
	} `json:"Wind,omitempty"`
}

type WeatherMeasurepoint1Dot0 struct {
	Deleted      *bool      `json:"Deleted,omitempty"`
	Geometry     *Geometry  `json:"Geometry,omitempty"`
	ID           *string    `json:"Id,omitempty"`
	ModifiedTime *time.Time `json:"ModifiedTime,omitempty"`
	Name         *string    `json:"Name,omitempty"`
	Observation  *struct {
		Aggregated10Minutes *struct {
			Precipitation *struct {
				Rain    *bool `json:"Rain,omitempty"`
				RainSum *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin `json:"Origin,omitempty"`
						SensorNames *string `json:"SensorNames,omitempty"`
						Value       *int    `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"TotalWaterEquivalent,omitempty"`
			} `json:"Precipitation,omitempty"`
			Wind *struct {
				Height       *int `json:"Height,omitempty"`
				SpeedAverage *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedAverage,omitempty"`
				SpeedMax *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedMax,omitempty"`
			} `json:"Wind,omitempty"`
		} `json:"Aggregated10minutes,omitempty"`
		Aggregated30Minutes *struct {
			Precipitation *struct {
				Rain    *bool `json:"Rain,omitempty"`
				RainSum *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin `json:"Origin,omitempty"`
						SensorNames *string `json:"SensorNames,omitempty"`
						Value       *int    `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"TotalWaterEquivalent,omitempty"`
			} `json:"Precipitation,omitempty"`
			Wind *struct {
				Height       *int `json:"Height,omitempty"`
				SpeedAverage *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedAverage,omitempty"`
				SpeedMax *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedMax,omitempty"`
			} `json:"Wind,omitempty"`
		} `json:"Aggregated30minutes,omitempty"`
		Aggregated5Minutes *struct {
			Precipitation *struct {
				Rain    *bool `json:"Rain,omitempty"`
				RainSum *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin `json:"Origin,omitempty"`
						SensorNames *string `json:"SensorNames,omitempty"`
						Value       *int    `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"TotalWaterEquivalent,omitempty"`
			} `json:"Precipitation,omitempty"`
			Wind *struct {
				Height       *int `json:"Height,omitempty"`
				SpeedAverage *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedAverage,omitempty"`
				SpeedMax *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedMax,omitempty"`
			} `json:"Wind,omitempty"`
		} `json:"Aggregated5minutes,omitempty"`
		Air *struct {
			Dewpoint *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Dewpoint,omitempty"`
			RelativeHumidity *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"RelativeHumidity,omitempty"`
			Temperature *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Temperature,omitempty"`
			VisibleDistance *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"VisibleDistance,omitempty"`
		} `json:"Air,omitempty"`
		ID         *string    `json:"Id,omitempty"`
		Sample     *time.Time `json:"Sample,omitempty"`
		Subsurface *struct {
			Ground []struct {
				Depth *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Depth,omitempty"`
				Temperature *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Temperature,omitempty"`
			} `json:"Ground,omitempty"`
		} `json:"Subsurface,omitempty"`
		Surface *struct {
			Grip *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Grip,omitempty"`
			Ice      *bool `json:"Ice,omitempty"`
			IceDepth *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"IceDepth,omitempty"`
			Snow      *bool `json:"Snow,omitempty"`
			SnowDepth *struct {
				Solid *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
				WaterEquivalent *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"WaterEquivalent,omitempty"`
			} `json:"SnowDepth,omitempty"`
			Temperature *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Temperature,omitempty"`
			Water      *bool `json:"Water,omitempty"`
			WaterDepth *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"WaterDepth,omitempty"`
		} `json:"Surface,omitempty"`
		Wind [2]struct {
			Direction *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Direction,omitempty"`
			Height *int `json:"Height,omitempty"`
			Speed  *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Speed,omitempty"`
		} `json:"Wind,omitempty"`
	} `json:"Observation,omitempty"`
}

type WeatherMeasurepoint2Dot0 struct {
	Deleted      *bool      `json:"Deleted,omitempty"`
	Geometry     *Geometry  `json:"Geometry,omitempty"`
	ID           *string    `json:"Id,omitempty"`
	ModifiedTime *time.Time `json:"ModifiedTime,omitempty"`
	Name         *string    `json:"Name,omitempty"`
	Observation  *struct {
		Aggregated10Minutes *struct {
			Precipitation *struct {
				Rain    *bool `json:"Rain,omitempty"`
				RainSum *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin `json:"Origin,omitempty"`
						SensorNames *string `json:"SensorNames,omitempty"`
						Value       *int    `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"TotalWaterEquivalent,omitempty"`
			} `json:"Precipitation,omitempty"`
			Wind *struct {
				Height       *int `json:"Height,omitempty"`
				SpeedAverage *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedAverage,omitempty"`
				SpeedMax *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedMax,omitempty"`
			} `json:"Wind,omitempty"`
		} `json:"Aggregated10minutes,omitempty"`
		Aggregated30Minutes *struct {
			Precipitation *struct {
				Rain    *bool `json:"Rain,omitempty"`
				RainSum *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin `json:"Origin,omitempty"`
						SensorNames *string `json:"SensorNames,omitempty"`
						Value       *int    `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"TotalWaterEquivalent,omitempty"`
			} `json:"Precipitation,omitempty"`
			Wind *struct {
				Height       *int `json:"Height,omitempty"`
				SpeedAverage *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedAverage,omitempty"`
				SpeedMax *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedMax,omitempty"`
			} `json:"Wind,omitempty"`
		} `json:"Aggregated30minutes,omitempty"`
		Aggregated5Minutes *struct {
			Precipitation *struct {
				Rain    *bool `json:"Rain,omitempty"`
				RainSum *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin `json:"Origin,omitempty"`
						SensorNames *string `json:"SensorNames,omitempty"`
						Value       *int    `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin `json:"Origin,omitempty"`
					SensorNames *string `json:"SensorNames,omitempty"`
					Value       *int    `json:"Value,omitempty"`
				} `json:"TotalWaterEquivalent,omitempty"`
			} `json:"Precipitation,omitempty"`
			Wind *struct {
				Height       *int `json:"Height,omitempty"`
				SpeedAverage *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedAverage,omitempty"`
				SpeedMax *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"SpeedMax,omitempty"`
			} `json:"Wind,omitempty"`
		} `json:"Aggregated5minutes,omitempty"`
		Air *struct {
			Dewpoint *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Dewpoint,omitempty"`
			RelativeHumidity *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"RelativeHumidity,omitempty"`
			Temperature *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Temperature,omitempty"`
			VisibleDistance *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"VisibleDistance,omitempty"`
		} `json:"Air,omitempty"`
		ID         *string    `json:"Id,omitempty"`
		Sample     *time.Time `json:"Sample,omitempty"`
		Subsurface *struct {
			Ground []struct {
				Depth *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Depth,omitempty"`
				Temperature *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Temperature,omitempty"`
			} `json:"Ground,omitempty"`
		} `json:"Subsurface,omitempty"`
		Surface *struct {
			Grip *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Grip,omitempty"`
			Ice      *bool `json:"Ice,omitempty"`
			IceDepth *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"IceDepth,omitempty"`
			Snow      *bool `json:"Snow,omitempty"`
			SnowDepth *struct {
				Solid *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
				WaterEquivalent *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"WaterEquivalent,omitempty"`
			} `json:"SnowDepth,omitempty"`
			Temperature *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Temperature,omitempty"`
			Water      *bool `json:"Water,omitempty"`
			WaterDepth *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"WaterDepth,omitempty"`
		} `json:"Surface,omitempty"`
		Weather *struct {
			Precipitation *PrecipitationTypeMeasure `json:"Precipitation,omitempty"`
		} `json:"Weather,omitempty"`
		Wind [2]struct {
			Direction *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Direction,omitempty"`
			Height *int `json:"Height,omitempty"`
			Speed  *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"Speed,omitempty"`
		} `json:"Wind,omitempty"`
	} `json:"Observation,omitempty"`
}
