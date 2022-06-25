package trafikinfo

import "time"

type Camera1Dot0 struct {
	Active           *bool       `json:"Active,omitempty"`
	CameraGroup      *string     `json:"CameraGroup,omitempty"`
	ContentType      *string     `json:"ContentType,omitempty"`
	County           []County    `json:"CountyNo,omitempty"`
	Deleted          *bool       `json:"Deleted,omitempty"`
	Description      *string     `json:"Description,omitempty"`
	Direction        *int        `json:"Direction,omitempty"`
	Geometry         *Geometry   `json:"Geometry,omitempty"`
	HasFullSizePhoto *bool       `json:"HasFullSizePhoto,omitempty"`
	HasSketchImage   *bool       `json:"HasSketchImage,omitempty"`
	IconID           *string     `json:"IconId,omitempty"`
	ID               *string     `json:"Id,omitempty"`
	Location         *string     `json:"Location,omitempty"`
	ModifiedTime     *time.Time  `json:"ModifiedTime,omitempty"`
	Name             *string     `json:"Name,omitempty"`
	PhotoTime        *time.Time  `json:"PhotoTime,omitempty"`
	PhotoURL         *string     `json:"PhotoUrl,omitempty"`
	Status           *string     `json:"Status,omitempty"`
	Type             *CameraType `json:"Type,omitempty"`
}

type CameraType string

const (
	CameraTypeSpeed CameraType = "Väglagskamera"
	CameraTypeFlow  CameraType = "Trafikflödeskamera"
)

type TrafficSafetyCamera1Dot0 struct {
	Bearing      *int       `json:"Bearing,omitempty"`
	County       []County   `json:"CountyNo,omitempty"`
	Deleted      *bool      `json:"Deleted,omitempty"`
	Geometry     *Geometry  `json:"Geometry,omitempty"`
	IconID       *string    `json:"IconId,omitempty"`
	ID           *string    `json:"Id,omitempty"`
	ModifiedTime *time.Time `json:"ModifiedTime,omitempty"`
	Name         *string    `json:"Name,omitempty"`
	RoadNumber   *string    `json:"RoadNumber,omitempty"`
}

type RoadConditionOverview1Dot0 struct {
	County                  []County   `json:"CountyNo,omitempty"`
	Deleted                 *bool      `json:"Deleted,omitempty"`
	EndTime                 *time.Time `json:"EndTime,omitempty"`
	Geometry                *Geometry  `json:"Geometry,omitempty"`
	ID                      *string    `json:"Id,omitempty"`
	LocationText            *string    `json:"LocationText,omitempty"`
	ModifiedTime            *time.Time `json:"ModifiedTime,omitempty"`
	StartTime               *time.Time `json:"StartTime,omitempty"`
	Text                    *string    `json:"Text,omitempty"`
	ValidUntilFurtherNotice *bool      `json:"ValidUntilFurtherNotice,omitempty"`
}

type TrafficFlow1Dot0 struct {
	AverageVehicleSpeed            *int       `json:"AverageVehicleSpeed,omitempty"`
	County                         []County   `json:"CountyNo,omitempty"`
	Deleted                        *bool      `json:"Deleted,omitempty"`
	Geometry                       *Geometry  `json:"Geometry,omitempty"`
	MeasurementOrCalculationPeriod *int       `json:"MeasurementOrCalculationPeriod,omitempty"`
	MeasurementSide                *string    `json:"MeasurementSide,omitempty"`
	MeasurementTime                *time.Time `json:"MeasurementTime,omitempty"`
	ModifiedTime                   *time.Time `json:"ModifiedTime,omitempty"`
	RegionID                       *Region    `json:"RegionId,omitempty"`
	SiteID                         *int       `json:"SiteId,omitempty"`
	SpecificLane                   *string    `json:"SpecificLane,omitempty"`
	VehicleFlowRate                *int       `json:"VehicleFlowRate,omitempty"`
	VehicleType                    *Vehicle   `json:"VehicleType,omitempty"`
}

type TrafficFlow1Dot4 struct {
	AverageVehicleSpeed            *float64   `json:"AverageVehicleSpeed,omitempty"`
	County                         []County   `json:"CountyNo,omitempty"`
	Deleted                        *bool      `json:"Deleted,omitempty"`
	Geometry                       *Geometry  `json:"Geometry,omitempty"`
	MeasurementOrCalculationPeriod *int       `json:"MeasurementOrCalculationPeriod,omitempty"`
	MeasurementSide                *string    `json:"MeasurementSide,omitempty"`
	MeasurementTime                *time.Time `json:"MeasurementTime,omitempty"`
	ModifiedTime                   *time.Time `json:"ModifiedTime,omitempty"`
	RegionID                       *Region    `json:"RegionId,omitempty"`
	SiteID                         *int       `json:"SiteId,omitempty"`
	SpecificLane                   *string    `json:"SpecificLane,omitempty"`
	VehicleFlowRate                *int       `json:"VehicleFlowRate,omitempty"`
	VehicleType                    *Vehicle   `json:"VehicleType,omitempty"`
}

type Parking1Dot0 struct {
	County                []County `json:"CountyNo,omitempty"`
	Deleted               *bool    `json:"Deleted,omitempty"`
	Description           *string  `json:"Description,omitempty"`
	DistanceToNearestCity *string  `json:"DistanceToNearestCity,omitempty"`
	Equipment             []struct {
		Accessibility *Accessibility `json:"Accessibility,omitempty"`
		Type          *Equipment     `json:"Type,omitempty"`
	} `json:"Equipment,omitempty"`
	Facility []struct {
		Accessibility *Accessibility `json:"Accessibility,omitempty"`
		Type          *Facility      `json:"Type,omitempty"`
	} `json:"Facility,omitempty"`
	Geometry            *Geometry  `json:"Geometry,omitempty"`
	IconID              *string    `json:"IconId,omitempty"`
	ID                  *string    `json:"Id,omitempty"`
	LocationDescription *string    `json:"LocationDescription,omitempty"`
	ModifiedTime        *time.Time `json:"ModifiedTime,omitempty"`
	Name                *string    `json:"Name,omitempty"`
	OpenStatus          *string    `json:"OpenStatus,omitempty"`
	OperationStatus     *string    `json:"OperationStatus,omitempty"`
	Operator            *struct {
		Contact                *string `json:"Contact,omitempty"`
		ContactEmail           *string `json:"ContactEmail,omitempty"`
		ContactTelephoneNumber *string `json:"ContactTelephoneNumber,omitempty"`
		Name                   *string `json:"Name,omitempty"`
	} `json:"Operator,omitempty"`
	ParkingAccess []Geometry `json:"ParkingAccess,omitempty"`
	Photo         []struct {
		Title *string `json:"Title,omitempty"`
		URL   *string `json:"Url,omitempty"`
	} `json:"Photo,omitempty"`
	TariffsAndPayment *struct {
		FreeOfCharge *bool   `json:"FreeOfCharge,omitempty"`
		Tariff       *string `json:"Tariff,omitempty"`
	} `json:"TariffsAndPayment,omitempty"`
	UsageSenario           []string `json:"UsageSenario,omitempty"`
	VehicleCharacteristics []struct {
		LoadType       *string  `json:"LoadType,omitempty"`
		NumberOfSpaces *int     `json:"NumberOfSpaces,omitempty"`
		VehicleType    *Vehicle `json:"VeichleType,omitempty"`
	} `json:"VehicleCharacteristics,omitempty"`
}

type Parking1Dot4 struct {
	County                []County `json:"CountyNo,omitempty"`
	Deleted               *bool    `json:"Deleted,omitempty"`
	Description           *string  `json:"Description,omitempty"`
	DistanceToNearestCity *string  `json:"DistanceToNearestCity,omitempty"`
	Equipment             []struct {
		Accessibility *Accessibility `json:"Accessibility,omitempty"`
		Type          *Equipment     `json:"Type,omitempty"`
	} `json:"Equipment,omitempty"`
	Facility []struct {
		Accessibility *Accessibility `json:"Accessibility,omitempty"`
		Type          *Facility      `json:"Type,omitempty"`
	} `json:"Facility,omitempty"`
	Geometry            *Geometry  `json:"Geometry,omitempty"`
	IconID              *string    `json:"IconId,omitempty"`
	ID                  *string    `json:"Id,omitempty"`
	LocationDescription *string    `json:"LocationDescription,omitempty"`
	ModifiedTime        *time.Time `json:"ModifiedTime,omitempty"`
	Name                *string    `json:"Name,omitempty"`
	OpenStatus          *string    `json:"OpenStatus,omitempty"`
	OperationStatus     *string    `json:"OperationStatus,omitempty"`
	Operator            *struct {
		Contact                *string `json:"Contact,omitempty"`
		ContactEmail           *string `json:"ContactEmail,omitempty"`
		ContactTelephoneNumber *string `json:"ContactTelephoneNumber,omitempty"`
		Name                   *string `json:"Name,omitempty"`
	} `json:"Operator,omitempty"`
	ParkingAccess []Geometry `json:"ParkingAccess,omitempty"`
	Photo         []struct {
		Title *string `json:"Title,omitempty"`
		URL   *string `json:"Url,omitempty"`
	} `json:"Photo,omitempty"`
	TariffsAndPayment *struct {
		FreeOfCharge *bool   `json:"FreeOfCharge,omitempty"`
		Tariff       *string `json:"Tariff,omitempty"`
	} `json:"TariffsAndPayment,omitempty"`
	UsageSenario           []string `json:"UsageSenario,omitempty"`
	VehicleCharacteristics []struct {
		LoadType       *string  `json:"LoadType,omitempty"`
		NumberOfSpaces *int     `json:"NumberOfSpaces,omitempty"`
		VehicleType    *Vehicle `json:"VehicleType,omitempty"`
	} `json:"VehicleCharacteristics,omitempty"`
}

type RoadCondition1Dot0 struct {
	Cause             []Cause        `json:"Cause,omitempty"`
	ConditionCode     *ConditionCode `json:"ConditionCode,omitempty"`
	ConditionInfo     []string       `json:"ConditionInfo,omitempty"`
	ConditionText     *string        `json:"ConditionText,omitempty"`
	County            []County       `json:"CountyNo,omitempty"`
	Creator           *string        `json:"Creator,omitempty"`
	Deleted           *bool          `json:"Deleted,omitempty"`
	EndTime           *time.Time     `json:"EndTime,omitempty"`
	Geometry          *Geometry      `json:"Geometry,omitempty"`
	IconID            *string        `json:"IconId,omitempty"`
	ID                *string        `json:"Id,omitempty"`
	LocationText      *string        `json:"LocationText,omitempty"`
	Measurement       []Measurement  `json:"Measurement,omitempty"`
	ModifiedTime      *time.Time     `json:"ModifiedTime,omitempty"`
	RoadNumber        *string        `json:"RoadNumber,omitempty"`
	RoadNumberNumeric *int           `json:"RoadNumberNumeric,omitempty"`
	StartTime         *time.Time     `json:"StartTime,omitempty"`
	Warning           []Warning      `json:"Warning,omitempty"`
}

type RoadCondition1Dot1 struct {
	Cause                []Cause        `json:"Cause,omitempty"`
	ConditionCode        *ConditionCode `json:"ConditionCode,omitempty"`
	ConditionInfo        []string       `json:"ConditionInfo,omitempty"`
	ConditionText        *string        `json:"ConditionText,omitempty"`
	County               []County       `json:"CountyNo,omitempty"`
	Creator              *string        `json:"Creator,omitempty"`
	Deleted              *bool          `json:"Deleted,omitempty"`
	EndTime              *time.Time     `json:"EndTime,omitempty"`
	Geometry             *Geometry      `json:"Geometry,omitempty"`
	IconID               *string        `json:"IconId,omitempty"`
	ID                   *string        `json:"Id,omitempty"`
	LocationText         *string        `json:"LocationText,omitempty"`
	Measurement          []Measurement  `json:"Measurement,omitempty"`
	ModifiedTime         *time.Time     `json:"ModifiedTime,omitempty"`
	RoadNumber           *string        `json:"RoadNumber,omitempty"`
	RoadNumberNumeric    *int           `json:"RoadNumberNumeric,omitempty"`
	SafetyRelatedMessage *bool          `json:"SafetyRelatedMessage,omitempty"`
	StartTime            *time.Time     `json:"StartTime,omitempty"`
	Warning              []Warning      `json:"Warning,omitempty"`
}

type RoadCondition1Dot2 struct {
	Cause                []Cause          `json:"Cause,omitempty"`
	ConditionCode        *ConditionCode   `json:"ConditionCode,omitempty"`
	ConditionInfo        []string         `json:"ConditionInfo,omitempty"`
	ConditionText        *string          `json:"ConditionText,omitempty"`
	County               []County         `json:"CountyNo,omitempty"`
	Creator              *string          `json:"Creator,omitempty"`
	Deleted              *bool            `json:"Deleted,omitempty"`
	EndTime              *time.Time       `json:"EndTime,omitempty"`
	Geometry             *GeometryWithMod `json:"Geometry,omitempty"`
	IconID               *string          `json:"IconId,omitempty"`
	ID                   *string          `json:"Id,omitempty"`
	LocationText         *string          `json:"LocationText,omitempty"`
	Measurement          []Measurement    `json:"Measurement,omitempty"`
	ModifiedTime         *time.Time       `json:"ModifiedTime,omitempty"`
	RoadNumber           *string          `json:"RoadNumber,omitempty"`
	RoadNumberNumeric    *int             `json:"RoadNumberNumeric,omitempty"`
	SafetyRelatedMessage *bool            `json:"SafetyRelatedMessage,omitempty"`
	StartTime            *time.Time       `json:"StartTime,omitempty"`
	Warning              []Warning        `json:"Warning,omitempty"`
}

type TravelTimeRoute1Dot0 struct {
	AverageFunctionalRoadClass *int       `json:"AverageFunctionalRoadClass,omitempty"`
	Country                    *Country   `json:"CountryCode,omitempty"`
	County                     []County   `json:"CountyNo,omitempty"`
	Deleted                    *bool      `json:"Deleted,omitempty"`
	ExpectedFreeFlowTravelTime *int       `json:"ExpectedFreeFlowTravelTime,omitempty"`
	FreeFlowTravelTime         *int       `json:"FreeFlowTravelTime,omitempty"`
	Geometry                   *Geometry  `json:"Geometry,omitempty"`
	ID                         *string    `json:"Id,omitempty"`
	Length                     *int       `json:"Length,omitempty"`
	MeasureTime                *time.Time `json:"MeasureTime,omitempty"`
	ModifiedTime               *time.Time `json:"ModifiedTime,omitempty"`
	Name                       *string    `json:"Name,omitempty"`
	Speed                      *float64   `json:"Speed,omitempty"`
	Status                     *Status    `json:"Status,omitempty"`
	TravelTime                 *int       `json:"TravelTime,omitempty"`
}

type TravelTimeRoute1Dot3 struct {
	AverageFunctionalRoadClass *int             `json:"AverageFunctionalRoadClass,omitempty"`
	Country                    *Country         `json:"CountryCode,omitempty"`
	County                     []County         `json:"CountyNo,omitempty"`
	Deleted                    *bool            `json:"Deleted,omitempty"`
	ExpectedFreeFlowTravelTime *int             `json:"ExpectedFreeFlowTravelTime,omitempty"`
	FreeFlowTravelTime         *int             `json:"FreeFlowTravelTime,omitempty"`
	Geometry                   *GeometryWithMod `json:"Geometry,omitempty"`
	ID                         *string          `json:"Id,omitempty"`
	Length                     *int             `json:"Length,omitempty"`
	MeasureTime                *time.Time       `json:"MeasureTime,omitempty"`
	ModifiedTime               *time.Time       `json:"ModifiedTime,omitempty"`
	Name                       *string          `json:"Name,omitempty"`
	RouteOwner                 *int             `json:"RouteOwner,omitempty"`
	Speed                      *float64         `json:"Speed,omitempty"`
	Status                     *Status          `json:"Status,omitempty"`
	TravelTime                 *int             `json:"TravelTime,omitempty"`
}

type TravelTimeRoute1Dot4 struct {
	AverageFunctionalRoadClass *int             `json:"AverageFunctionalRoadClass,omitempty"`
	Country                    *Country         `json:"CountryCode,omitempty"`
	County                     []County         `json:"CountyNo,omitempty"`
	Deleted                    *bool            `json:"Deleted,omitempty"`
	ExpectedFreeFlowTravelTime *int             `json:"ExpectedFreeFlowTravelTime,omitempty"`
	FreeFlowTravelTime         *int             `json:"FreeFlowTravelTime,omitempty"`
	Geometry                   *GeometryWithMod `json:"Geometry,omitempty"`
	ID                         *string          `json:"Id,omitempty"`
	Length                     *int             `json:"Length,omitempty"`
	MeasureTime                *time.Time       `json:"MeasureTime,omitempty"`
	ModifiedTime               *time.Time       `json:"ModifiedTime,omitempty"`
	Name                       *string          `json:"Name,omitempty"`
	RouteOwner                 *int             `json:"RouteOwner,omitempty"`
	Speed                      *float64         `json:"Speed,omitempty"`
	TrafficStatus              *TrafficStatus   `json:"TrafficStatus,omitempty"`
	TravelTime                 *int             `json:"TravelTime,omitempty"`
}

type TravelTimeRoute1Dot5 struct {
	AverageFunctionalRoadClass *int             `json:"AverageFunctionalRoadClass,omitempty"`
	Country                    *Country         `json:"CountryCode,omitempty"`
	County                     []County         `json:"CountyNo,omitempty"`
	Deleted                    *bool            `json:"Deleted,omitempty"`
	ExpectedFreeFlowTravelTime *float64         `json:"ExpectedFreeFlowTravelTime,omitempty"`
	FreeFlowTravelTime         *float64         `json:"FreeFlowTravelTime,omitempty"`
	Geometry                   *GeometryWithMod `json:"Geometry,omitempty"`
	ID                         *string          `json:"Id,omitempty"`
	Length                     *float64         `json:"Length,omitempty"`
	MeasureTime                *time.Time       `json:"MeasureTime,omitempty"`
	ModifiedTime               *time.Time       `json:"ModifiedTime,omitempty"`
	Name                       *string          `json:"Name,omitempty"`
	RouteOwner                 *int             `json:"RouteOwner,omitempty"`
	Speed                      *float64         `json:"Speed,omitempty"`
	TrafficStatus              *TrafficStatus   `json:"TrafficStatus,omitempty"`
	TravelTime                 *float64         `json:"TravelTime,omitempty"`
}
