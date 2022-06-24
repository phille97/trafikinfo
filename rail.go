package trafikinfo

import "time"

// ReasonCode1Dot0 represents a reason for a train related issue.
//
// The "Code" and "Level3Description" fields correspond to the "Code" and
// "Description" fields for the TrainAnnouncement and TrainMessage object
// types.
type ReasonCode1Dot0 struct {
	Code              *string    `json:"Code,omitempty"`
	Deleted           *bool      `json:"Deleted,omitempty"`
	GroupDescription  *string    `json:"GroupDescription,omitempty"`
	Level1Description *string    `json:"Level1Description,omitempty"`
	Level2Description *string    `json:"Level2Description,omitempty"`
	Level3Description *string    `json:"Level3Description,omitempty"`
	ModifiedTime      *time.Time `json:"ModifiedTime,omitempty"`
}

type MediaType string

const (
	MediaTypeMonitor      MediaType = "Monitor"
	MediaTypePlatformSign MediaType = "Plattformsskylt"
	MediaTypeAnnouncement MediaType = "Utrop"
)

type MessageStatus string

const (
	MessageStatusLow        MessageStatus = "Lag"
	MessageStatusNormal     MessageStatus = "Normal"
	MessageStatusHigh       MessageStatus = "Hog"
	MessageStatusDisruption MessageStatus = "StortLage"
)

type TrainStationMessage1Dot0 struct {
	ActiveDays        *string    `json:"ActiveDays,omitempty"`
	Deleted           *bool      `json:"Deleted,omitempty"`
	EndDateTime       *time.Time `json:"EndDateTime,omitempty"`
	EventID           *string    `json:"EventId,omitempty" yaml:"EventId,omitempty"`
	FreeText          *string    `json:"FreeText,omitempty"`
	ID                *string    `json:"Id,omitempty"`
	LocationCode      *string    `json:"LocationCode,omitempty"`
	MediaType         *MediaType `json:"MediaType,omitempty"`
	ModifiedTime      *time.Time `json:"ModifiedTime,omitempty"`
	MonitorAttributes *struct {
		BigScreenMonitorActivated *bool `json:"BigScreenMonitorActivated,omitempty"`
		CommuterMonitor           *bool `json:"CommuterMonitor,omitempty"`
	} `json:"MonitorAttributes,omitempty"`
	PlatformSignAttributes *struct {
		CommuterPlatformSign *bool `json:"CommuterPlatformSign,omitempty"`
		TrackList            *struct {
			Track []string `json:"Track,omitempty"`
		} `json:"TrackList,omitempty"`
	} `json:"PlatformSignAttributes,omitempty"`
	PublicAnnouncementAttributes *struct {
		EnglishPublicAnnouncementActivated *bool   `json:"EnglishPublicAnnouncementActivated,omitempty"`
		EnglishText                        *string `json:"EnglishText,omitempty"`
		PublicAnnouncementPlanList         *struct {
			PublicAnnouncementPlan []struct {
				ActiveDays                     *string `json:"ActiveDays,omitempty"`
				PublicAnnouncementOccasionList *struct {
					PublicAnnouncementOccasion []int `json:"PublicAnnouncementOccasion,omitempty"`
				} `json:"PublicAnnouncementOccasionList,omitempty"`
				ValidFrom *time.Time `json:"ValidFrom,omitempty"`
				ValidTo   *time.Time `json:"ValidTo,omitempty"`
			} `json:"PublicAnnouncementPlan,omitempty"`
		} `json:"PublicAnnouncementPlanList,omitempty"`
		PublicAnnouncementZoneList *struct {
			PublicAnnouncementZone []string `json:"PublicAnnouncementZone,omitempty"`
		} `json:"PublicAnnouncementZoneList,omitempty"`
	} `json:"PublicAnnouncementAttributes,omitempty"`
	SplitActivationTime *bool          `json:"SplitActivationTime,omitempty"`
	StartDateTime       *time.Time     `json:"StartDateTime,omitempty"`
	Status              *MessageStatus `json:"Status,omitempty"`
	VersionNumber       *int           `json:"VersionNumber,omitempty"`
}
