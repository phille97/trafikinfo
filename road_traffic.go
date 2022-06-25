package trafikinfo

import "time"

type Camera1Dot0 struct {
	Active           *bool          `json:"Active,omitempty"`
	CameraGroup      *string        `json:"CameraGroup,omitempty"`
	ContentType      *string        `json:"ContentType,omitempty"`
	CountyNo         []CountyNumber `json:"CountyNo,omitempty"`
	Deleted          *bool          `json:"Deleted,omitempty"`
	Description      *string        `json:"Description,omitempty"`
	Direction        *int           `json:"Direction,omitempty"`
	Geometry         *Geometry      `json:"Geometry,omitempty"`
	HasFullSizePhoto *bool          `json:"HasFullSizePhoto,omitempty"`
	HasSketchImage   *bool          `json:"HasSketchImage,omitempty"`
	IconID           *string        `json:"IconId,omitempty"`
	ID               *string        `json:"Id,omitempty"`
	Location         *string        `json:"Location,omitempty"`
	ModifiedTime     *time.Time     `json:"ModifiedTime,omitempty"`
	Name             *string        `json:"Name,omitempty"`
	PhotoTime        *time.Time     `json:"PhotoTime,omitempty"`
	PhotoURL         *string        `json:"PhotoUrl,omitempty"`
	Status           *string        `json:"Status,omitempty"`
	Type             *CameraType    `json:"Type,omitempty"`
}

type CameraType string

const (
	CameraTypeSpeed CameraType = "Väglagskamera"
	CameraTypeFlow  CameraType = "Trafikflödeskamera"
)

type TrafficSafetyCamera1Dot0 struct {
	Bearing      *int           `json:"Bearing,omitempty"`
	CountyNo     []CountyNumber `json:"CountyNo,omitempty"`
	Deleted      *bool          `json:"Deleted,omitempty"`
	Geometry     *Geometry      `json:"Geometry,omitempty"`
	IconID       *string        `json:"IconId,omitempty"`
	ID           *string        `json:"Id,omitempty"`
	ModifiedTime *time.Time     `json:"ModifiedTime,omitempty"`
	Name         *string        `json:"Name,omitempty"`
	RoadNumber   *string        `json:"RoadNumber,omitempty"`
}
