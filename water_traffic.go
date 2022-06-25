package trafikinfo

import (
	"time"
)

type FerryAnnouncement1Dot2 struct {
	Deleted       *bool      `json:"Deleted,omitempty"`
	DepartureTime *time.Time `json:"DepartureTime,omitempty"`
	DeviationID   *string    `json:"DeviationId,omitempty"`
	FromHarbor    *struct {
		ID   *int    `json:"Id,omitempty"`
		Name *string `json:"Name,omitempty"`
	} `json:"FromHarbor,omitempty"`
	ID           *int       `json:"Id,omitempty"`
	Info         []string   `json:"Info,omitempty"`
	ModifiedTime *time.Time `json:"ModifiedTime,omitempty"`
	Route        *struct {
		ID        *int    `json:"Id,omitempty"`
		Name      *string `json:"Name,omitempty"`
		Shortname *string `json:"Shortname,omitempty"`
		Type      *struct {
			ID   *int    `json:"Id,omitempty"`
			Name *string `json:"Name,omitempty"`
		} `json:"Type,omitempty"`
	} `json:"Route,omitempty"`
	ToHarbor *struct {
		ID   *int    `json:"Id,omitempty"`
		Name *string `json:"Name,omitempty"`
	} `json:"ToHarbor,omitempty"`
}

type FerryRoute1Dot2 struct {
	Deleted     *bool     `json:"Deleted,omitempty"`
	DeviationID *string   `json:"DeviationId,omitempty"`
	Geometry    *Geometry `json:"Geometry,omitempty"`
	Harbor      []struct {
		ID        *int    `json:"Id,omitempty"`
		Name      *string `json:"Name,omitempty"`
		SortOrder *int    `json:"SortOrder,omitempty"`
		StopType  *struct {
			ID      *int      `json:"Id,omitempty"`
			Name    *StopType `json:"Name,omitempty"`
			Visible *bool     `json:"Visible,omitempty"`
		} `json:"StopType,omitempty"`
	} `json:"Harbor,omitempty"`
	ID           *int       `json:"Id,omitempty"`
	ModifiedTime *time.Time `json:"ModifiedTime,omitempty"`
	Name         *string    `json:"Name,omitempty"`
	Shortname    *string    `json:"Shortname,omitempty"`
	Timetable    []struct {
		Description *string `json:"Description,omitempty"`
		Period      []struct {
			Name     *string `json:"Name,omitempty"`
			Schedule []struct {
				Deviation []struct {
					Description *string `json:"Description,omitempty"`
					FromDate    *string `json:"FromDate,omitempty"`
					ID          *string `json:"Id,omitempty"`
					SpecDate    *string `json:"SpecDate,omitempty"`
					ToDate      *string `json:"ToDate,omitempty"`
					Type        *struct {
						ID   *string        `json:"Id,omitempty"`
						Name *DeviationType `json:"Name,omitempty"`
					} `json:"Type,omitempty"`
				} `json:"Deviation,omitempty"`
				Harbor *struct {
					ID        *int    `json:"Id,omitempty"`
					Name      *string `json:"Name,omitempty"`
					SortOrder *int    `json:"SortOrder,omitempty"`
					StopType  *struct {
						ID      *int      `json:"Id,omitempty"`
						Name    *StopType `json:"Name,omitempty"`
						Visible *bool     `json:"Visible,omitempty"`
					} `json:"StopType,omitempty"`
				} `json:"Harbor,omitempty"`
				SortOrder *int `json:"SortOrder,omitempty"`
				StopType  *struct {
					ID      *int      `json:"Id,omitempty"`
					Name    *StopType `json:"Name,omitempty"`
					Visible *bool     `json:"Visible,omitempty"`
				} `json:"StopType,omitempty"`
				Time *string `json:"Time,omitempty"`
			} `json:"Schedule,omitempty"`
			SortOrder *int `json:"SortOrder,omitempty"`
			Weekday   []struct {
				ID  *int    `json:"Id,omitempty"`
				Day *string `json:"Day,omitempty"`
			} `json:"Weekday,omitempty"`
		} `json:"Period,omitempty"`
		Priority *int `json:"Priority,omitempty"`
		Valid    []struct {
			From *time.Time `json:"From,omitempty"`
			To   *time.Time `json:"To,omitempty"`
		} `json:"Valid,omitempty"`
	} `json:"Timetable,omitempty"`
	Type *struct {
		ID   *int    `json:"Id,omitempty"`
		Name *string `json:"Name,omitempty"`
	} `json:"Type,omitempty"`
}

type DeviationType string

const (
	DeviationTypeMessage   DeviationType = "Meddelande"
	DeviationTypeSummon    DeviationType = "Kallelse"
	DeviationTypeGoing     DeviationType = "Går"
	DeviationTypeCancelled DeviationType = "Går ej"
)

type StopType string

const (
	StopTypeArrival          StopType = "Ank"
	StopTypeDeparture        StopType = "Avg"
	StopTypeArrivalDeparture StopType = "Ank/Avg"
)
