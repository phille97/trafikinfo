package trafikinfo

import "time"

type MeasurementData1001Dot0 struct {
	County    *CountyNumber `json:"County,omitempty"`
	Deleted   *bool         `json:"Deleted,omitempty"`
	Direction *struct {
		Code  *int    `json:"Code,omitempty"`
		Value *string `json:"Value,omitempty"`
	} `json:"Direction,omitempty"`
	EdgeDepthAverageValue       *float64 `json:"EdgeDepthAverageValue,omitempty"`
	EndContinuousLength         *int     `json:"EndContinuousLength,omitempty"`
	IRIRightAverageValue        *float64 `json:"IRIRightAverageValue,omitempty"`
	Lane                        *int     `json:"Lane,omitempty"`
	Length                      *int     `json:"Length,omitempty"`
	MPDMacrotextureAverageValue *float64 `json:"MPDMacrotextureAverageValue,omitempty"`
	MeasurementDataType         *struct {
		Code  *int    `json:"Code,omitempty"`
		Value *string `json:"Value,omitempty"`
	} `json:"MeasurementDataType,omitempty"`
	MeasurementDate           *string    `json:"MeasurementDate,omitempty"`
	MeasurementDateSpecific   *string    `json:"MeasurementDateSpecific,omitempty"`
	ModifiedTime              *time.Time `json:"ModifiedTime,omitempty"`
	RoadMainNumber            *int       `json:"RoadMainNumber,omitempty"`
	RoadSubNumber             *int       `json:"RoadSubNumber,omitempty"`
	RutDepthMax15AverageValue *float64   `json:"RutDepthMax15AverageValue,omitempty"`
	RutDepthMax17AverageValue *float64   `json:"RutDepthMax17AverageValue,omitempty"`
	StartContinuousLength     *int       `json:"StartContinuousLength,omitempty"`
	Timestamp                 *time.Time `json:"TimeStamp,omitempty"`
}

type MeasurementData201Dot0 struct {
	County             *CountyNumber `json:"County,omitempty"`
	CrossfallRutBottom *float64      `json:"CrossfallRutBottom,omitempty"`
	Curvature          *float64      `json:"Curvature,omitempty"`
	Deleted            *bool         `json:"Deleted,omitempty"`
	Direction          *struct {
		Code  *int    `json:"Code,omitempty"`
		Value *string `json:"Value,omitempty"`
	} `json:"Direction,omitempty"`
	EdgeDepth             *float64 `json:"EdgeDepth,omitempty"`
	EndContinuousLength   *int     `json:"EndContinuousLength,omitempty"`
	Hilliness             *float64 `json:"Hilliness,omitempty"`
	IRILeft               *float64 `json:"IRILeft,omitempty"`
	IRIRight              *float64 `json:"IRIRight,omitempty"`
	Lane                  *int     `json:"Lane,omitempty"`
	Length                *int     `json:"Length,omitempty"`
	MPDMacrotextureLeft   *float64 `json:"MPDMacrotextureLeft,omitempty"`
	MPDMacrotextureMiddle *float64 `json:"MPDMacrotextureMiddle,omitempty"`
	MPDMacrotextureRight  *float64 `json:"MPDMacrotextureRight,omitempty"`
	MeasurementDataType   *struct {
		Code  *int    `json:"Code,omitempty"`
		Value *string `json:"Value,omitempty"`
	} `json:"MeasurementDataType,omitempty"`
	MeasurementDate         *string    `json:"MeasurementDate,omitempty"`
	MeasurementDateSpecific *string    `json:"MeasurementDateSpecific,omitempty"`
	MeasurementVehicleSpeed *float64   `json:"MeasurementVehicleSpeed,omitempty"`
	MegatextureLeft         *float64   `json:"MegatextureLeft,omitempty"`
	MegatextureRight        *float64   `json:"MegatextureRight,omitempty"`
	ModifiedTime            *time.Time `json:"ModifiedTime,omitempty"`
	RoadMainNumber          *int       `json:"RoadMainNumber,omitempty"`
	RoadSubNumber           *int       `json:"RoadSubNumber,omitempty"`
	RutArea                 *float64   `json:"RutArea,omitempty"`
	RutBottomDistance       *float64   `json:"RutBottomDistance,omitempty"`
	RutDepthLeft17          *float64   `json:"RutDepthLeft17,omitempty"`
	RutDepthMax15           *float64   `json:"RutDepthMax15,omitempty"`
	RutDepthMax17           *float64   `json:"RutDepthMax17,omitempty"`
	RutDepthRight15         *float64   `json:"RutDepthRight15,omitempty"`
	RutDepthRight17         *float64   `json:"RutDepthRight17,omitempty"`
	RutWidthLeft            *float64   `json:"RutWidthLeft,omitempty"`
	RutWidthRight           *float64   `json:"RutWidthRight,omitempty"`
	StartContinuousLength   *int       `json:"StartContinuousLength,omitempty"`
	Timestamp               *time.Time `json:"TimeStamp,omitempty"`
	WaterArea               *float64   `json:"WaterArea,omitempty"`
}

type PavementData1Dot0 struct {
	BallMillValue *float64      `json:"BallMillValue,omitempty"`
	Binder        *string       `json:"Binder,omitempty"`
	Contractor    *string       `json:"Contractor,omitempty"`
	County        *CountyNumber `json:"County,omitempty"`
	Coverage      *string       `json:"Coverage,omitempty"`
	Deleted       *bool         `json:"Deleted,omitempty"`
	Direction     *struct {
		Code  *int    `json:"Code,omitempty"`
		Value *string `json:"Value,omitempty"`
	} `json:"Direction,omitempty"`
	EndContinuousLength   *int       `json:"EndContinuousLength,omitempty"`
	FinalInspectionDate   *time.Time `json:"FinalInspectionDate,omitempty"`
	Lane                  *int       `json:"Lane,omitempty"`
	Length                *int       `json:"Length,omitempty"`
	ManufacturingMethod   *string    `json:"ManufacturingMethod,omitempty"`
	MaxStoneSize          *int       `json:"MaxStoneSize,omitempty"`
	ModifiedTime          *time.Time `json:"ModifiedTime,omitempty"`
	PavementDate          *time.Time `json:"PavementDate,omitempty"`
	PavementType          *string    `json:"PavementType,omitempty"`
	PavingMethod          *string    `json:"PavingMethod,omitempty"`
	RoadMainNumber        *int       `json:"RoadMainNumber,omitempty"`
	RoadSubNumber         *int       `json:"RoadSubNumber,omitempty"`
	StartContinuousLength *int       `json:"StartContinuousLength,omitempty"`
	Thickness             *float64   `json:"Thickness,omitempty"`
	Timestamp             *time.Time `json:"TimeStamp,omitempty"`
	TreatmentCategory     *string    `json:"TreatmentCategory,omitempty"`
	Warranty              *int       `json:"Warranty,omitempty"`
	WarrantyIsDue         *time.Time `json:"WarrantyIsDue,omitempty"`
}
