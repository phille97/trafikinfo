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
