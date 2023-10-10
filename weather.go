package trafikinfo

import "time"

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
	PrecipitationMildSnow  PrecipitationAmount = "Måttligt snöfall"
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
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
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
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
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
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
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
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
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
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
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
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
			} `json:"RainSum"`
			Snow    *bool `json:"Snow,omitempty"`
			SnowSum *struct {
				Solid *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"Solid,omitempty"`
			} `json:"SnowSum,omitempty"`
			TotalWaterEquivalent *struct {
				Origin      *Origin  `json:"Origin,omitempty"`
				SensorNames *string  `json:"SensorNames,omitempty"`
				Value       *float64 `json:"Value,omitempty"`
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
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin  `json:"Origin,omitempty"`
						SensorNames *string  `json:"SensorNames,omitempty"`
						Value       *float64 `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
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
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin  `json:"Origin,omitempty"`
						SensorNames *string  `json:"SensorNames,omitempty"`
						Value       *float64 `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
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
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin  `json:"Origin,omitempty"`
						SensorNames *string  `json:"SensorNames,omitempty"`
						Value       *float64 `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
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
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin  `json:"Origin,omitempty"`
						SensorNames *string  `json:"SensorNames,omitempty"`
						Value       *float64 `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
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
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin  `json:"Origin,omitempty"`
						SensorNames *string  `json:"SensorNames,omitempty"`
						Value       *float64 `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
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
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
				} `json:"RainSum"`
				Snow    *bool `json:"Snow,omitempty"`
				SnowSum *struct {
					Solid *struct {
						Origin      *Origin  `json:"Origin,omitempty"`
						SensorNames *string  `json:"SensorNames,omitempty"`
						Value       *float64 `json:"Value,omitempty"`
					} `json:"Solid,omitempty"`
				} `json:"SnowSum,omitempty"`
				TotalWaterEquivalent *struct {
					Origin      *Origin  `json:"Origin,omitempty"`
					SensorNames *string  `json:"SensorNames,omitempty"`
					Value       *float64 `json:"Value,omitempty"`
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
