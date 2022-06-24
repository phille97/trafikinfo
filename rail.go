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

type CountryCode string

const (
	CountryCodeGermany CountryCode = "DE"
	CountryCodeDenmark CountryCode = "DK"
	CountryCodeNorway  CountryCode = "NO"
	CountryCodeSweden  CountryCode = "SE"
)

type TrainStation1Dot0 struct {
	Advertised                  *bool          `json:"Advertised,omitempty"`
	AdvertisedLocationName      *string        `json:"AdvertisedLocationName,omitempty"`
	AdvertisedShortLocationName *string        `json:"AdvertisedShortLocationName,omitempty"`
	CountryCode                 *CountryCode   `json:"CountryCode,omitempty"`
	CountyNo                    []CountyNumber `json:"CountyNo,omitempty"`
	Deleted                     *bool          `json:"Deleted,omitempty"`
	Geometry                    *Geometry      `json:"Geometry,omitempty"`
	LocationInformationText     *string        `json:"LocationInformationText,omitempty"`
	LocationSignature           *string        `json:"LocationSignature,omitempty"`
	ModifiedTime                *time.Time     `json:"ModifiedTime,omitempty"`
	PlatformLine                []string       `json:"PlatformLine,omitempty"`
	Prognosticated              *bool          `json:"Prognosticated,omitempty"`
}

type TrainStation1Dot4 struct {
	TrainStation1Dot0
	OfficialLocationName *string `json:"OfficialLocationName,omitempty"`
}

type trainMessageCommon struct {
	CountyNo            []CountyNumber `json:"CountyNo,omitempty"`
	Deleted             *bool          `json:"Deleted,omitempty"`
	EventID             *string        `json:"EventId,omitempty"`
	ExternalDescription *string        `json:"ExternalDescription,omitempty"`
	Geometry            *Geometry      `json:"Geometry,omitempty"`
	LastUpdateDateTime  *time.Time     `json:"LastUpdateDateTime,omitempty"`
	ModifiedTime        *time.Time     `json:"ModifiedTime,omitempty"`
	StartDateTime       *time.Time     `json:"StartDateTime,omitempty"`
}

type TrainMessage1Dot0 struct {
	trainMessageCommon
	AffectedLocation []string `json:"AffectedLocation,omitempty"`
	ReasonCodeText   *string  `json:"ReasonCodeText,omitempty"`
}

type TrainMessage1Dot3 struct {
	TrainMessage1Dot0
	Header        *string `json:"Header,omitempty"`
	TrafficImpact *struct {
		AffectedLocation []string `json:"AffectedLocation,omitempty"`
		FromLocation     []string `json:"FromLocation,omitempty"`
		ToLocation       []string `json:"ToLocation,omitempty"`
	} `json:"TrafficImpact,omitempty"`
	EndDateTime *time.Time `json:"EndDateTime,omitempty"`
}

type TrainMessage1Dot4 struct {
	TrainMessage1Dot3
	PrognosticatedEndDateTimeTrafficImpact *time.Time `json:"PrognosticatedEndDateTimeTrafficImpact,omitempty"`
	ExpectTrafficImpact                    *bool      `jsno:"ExpectTrafficImpact,omitempty"`
}

type TrainMessage1Dot5 struct {
	trainMessageCommon
	AffectedLocation                       []string   `json:"AffectedLocation,omitempty"`
	EndDateTime                            *time.Time `json:"EndDateTime,omitempty"`
	ExpectTrafficImpact                    *bool      `jsno:"ExpectTrafficImpact,omitempty"`
	Header                                 *string    `json:"Header,omitempty"`
	PrognosticatedEndDateTimeTrafficImpact *time.Time `json:"PrognosticatedEndDateTimeTrafficImpact,omitempty"`
	ReasonCode                             *struct {
		Code        *string `json:"Code,omitempty"`
		Description *string `json:"Description,omitempty"`
	} `json:"ReasonCodeText,omitempty"`
	TrafficImpact *struct {
		AffectedLocation []string `json:"AffectedLocation,omitempty"`
		FromLocation     []string `json:"FromLocation,omitempty"`
		ToLocation       []string `json:"ToLocation,omitempty"`
	} `json:"TrafficImpact,omitempty"`
}

type TrainMessage1Dot6 struct {
	trainMessageCommon
	EndDateTime                            *time.Time `json:"EndDateTime,omitempty"`
	Header                                 *string    `json:"Header,omitempty"`
	PrognosticatedEndDateTimeTrafficImpact *time.Time `json:"PrognosticatedEndDateTimeTrafficImpact,omitempty"`
	ReasonCode                             *struct {
		Code        *string `json:"Code,omitempty"`
		Description *string `json:"Description,omitempty"`
	} `json:"ReasonCodeText,omitempty"`
	TrafficImpact *struct {
		AffectedLocation []string `json:"AffectedLocation,omitempty"`
		FromLocation     []string `json:"FromLocation,omitempty"`
		IsConfirmed      *bool    `json:"IsConfirmed,omitempty"`
		ToLocation       []string `json:"ToLocation,omitempty"`
	} `json:"TrafficImpact,omitempty"`
}

type TrainMessage1Dot7 struct {
	trainMessageCommon
	EndDateTime                            *time.Time `json:"EndDateTime,omitempty"`
	Header                                 *string    `json:"Header,omitempty"`
	PrognosticatedEndDateTimeTrafficImpact *time.Time `json:"PrognosticatedEndDateTimeTrafficImpact,omitempty"`
	ReasonCode                             *struct {
		Code        *string `json:"Code,omitempty"`
		Description *string `json:"Description,omitempty"`
	} `json:"ReasonCodeText,omitempty"`
	TrafficImpact *struct {
		AffectedLocation []struct {
			LocationSignature       *string `json:"LocationSignature,omitempty"`
			ShouldBeTrafficInformed *bool   `json:"ShouldBeTrafficInformed,omitempty"`
		} `json:"AffectedLocation,omitempty"`
		FromLocation []string `json:"FromLocation,omitempty"`
		IsConfirmed  *bool    `json:"IsConfirmed,omitempty"`
		ToLocation   []string `json:"ToLocation,omitempty"`
	} `json:"TrafficImpact,omitempty"`
}

type ActivityType string

const (
	ActivityTypeArrival   ActivityType = "Ankomst"
	ActivityTypeDeparture ActivityType = "Avgang"
)

type TrainAnnouncement1Dot0 struct {
	ActivityID                 *string       `json:"ActivityId,omitempty"`
	ActivityType               *ActivityType `json:"ActivityType,omitempty"`
	Advertised                 *bool         `json:"Advertised,omitempty"`
	AdvertisedTimeAtLocation   *time.Time    `json:"AdvertisedTimeAtLocation,omitempty"`
	AdvertisedTrainIdent       *string       `json:"AdvertisedTrainIdent,omitempty"`
	Booking                    []string      `json:"Booking,omitempty"`
	Canceled                   *bool         `json:"Canceled,omitempty"`
	Deleted                    *bool         `json:"Deleted,omitempty"`
	Deviation                  []string      `json:"Deviation,omitempty"`
	EstimatedTimeAtLocation    *time.Time    `json:"EstimatedTimeAtLocation,omitempty"`
	EstimatedTimeIsPreliminary *bool         `json:"EstimatedTimeIsPreliminary,omitempty"`
	FromLocation               []string      `json:"FromLocation,omitempty"`
	InformationOwner           *string       `json:"InformationOwner,omitempty"`
	LocationSignature          *string       `json:"LocationSignature,omitempty"`
	MobileWebLink              *string       `json:"MobileWebLink,omitempty"`
	ModifiedTime               *time.Time    `json:"ModifiedTime,omitempty"`
	OtherInformation           []string      `json:"OtherInformation,omitempty"`
	ProductInformation         []string      `json:"ProductInformation,omitempty"`
	ScheduledDepartureDateTime *time.Time    `json:"ScheduledDepartureDateTime,omitempty"`
	Service                    []string      `json:"Service,omitempty"`
	TimeAtLocation             *time.Time    `json:"TimeAtLocation,omitempty"`
	ToLocation                 []string      `json:"ToLocation,omitempty"`
	TrackAtLocation            *string       `json:"TrackAtLocation,omitempty"`
	TrainComposition           []string      `json:"TrainComposition,omitempty"`
	TypeOfTraffic              *string       `json:"TypeOfTraffic,omitempty"`
	WebLink                    *string       `json:"WebLink,omitempty"`
}

type trainAnnouncementLocation struct {
	LocationName *string `json:"LocationName,omitempty"`
	Priority     *int    `json:"Priority,omitempty"`
	Order        *int    `json:"Order,omitempty"`
}

type TrainAnnouncement1Dot3 struct {
	ActivityID                            *string                     `json:"ActivityId,omitempty"`
	ActivityType                          *ActivityType               `json:"ActivityType,omitempty"`
	Advertised                            *bool                       `json:"Advertised,omitempty"`
	AdvertisedTimeAtLocation              *time.Time                  `json:"AdvertisedTimeAtLocation,omitempty"`
	AdvertisedTrainIdent                  *string                     `json:"AdvertisedTrainIdent,omitempty"`
	Booking                               []string                    `json:"Booking,omitempty"`
	Canceled                              *bool                       `json:"Canceled,omitempty"`
	Deleted                               *bool                       `json:"Deleted,omitempty"`
	Deviation                             []string                    `json:"Deviation,omitempty"`
	EstimatedTimeAtLocation               *time.Time                  `json:"EstimatedTimeAtLocation,omitempty"`
	EstimatedTimeIsPreliminary            *bool                       `json:"EstimatedTimeIsPreliminary,omitempty"`
	FromLocation                          []trainAnnouncementLocation `json:"FromLocation,omitempty"`
	InformationOwner                      *string                     `json:"InformationOwner,omitempty"`
	LocationSignature                     *string                     `json:"LocationSignature,omitempty"`
	MobileWebLink                         *string                     `json:"MobileWebLink,omitempty"`
	ModifiedTime                          *time.Time                  `json:"ModifiedTime,omitempty"`
	NewEquipment                          *int                        `json:"NewEquipment,omitempty"`
	OtherInformation                      []string                    `json:"OtherInformation,omitempty"`
	PlannedEstimatedTimeAtLocation        *time.Time                  `json:"PlannedEstimatedTimeAtLocation,omitempty"`
	PlannedEstimatedTimeAtLocationIsValid *bool                       `json:"PlannedEstimatedTimeAtLocationIsValid,omitempty"`
	ProductInformation                    []string                    `json:"ProductInformation,omitempty"`
	ScheduledDepartureDateTime            *time.Time                  `json:"ScheduledDepartureDateTime,omitempty"`
	Service                               []string                    `json:"Service,omitempty"`
	TechnicalTrainIdent                   *string                     `json:"TechnicalTrainIdent,omitempty"`
	TimeAtLocation                        *time.Time                  `json:"TimeAtLocation,omitempty"`
	ToLocation                            []trainAnnouncementLocation `json:"ToLocation,omitempty"`
	TrackAtLocation                       *string                     `json:"TrackAtLocation,omitempty"`
	TrainComposition                      []string                    `json:"TrainComposition,omitempty"`
	TypeOfTraffic                         *string                     `json:"TypeOfTraffic,omitempty"`
	ViaFromLocation                       []trainAnnouncementLocation `json:"ViaFromLocation,omitempty"`
	ViaToLocation                         []trainAnnouncementLocation `json:"ViaToLocation,omitempty"`
	WebLink                               *string                     `json:"WebLink,omitempty"`
	WebLinkName                           *string                     `json:"WebLinkName,omitempty"`
}

type TrainAnnouncement1Dot4 struct {
	TrainAnnouncement1Dot3
	Operator                  *string    `json:"Operator,omitempty"`
	TechnicalDateTime         *time.Time `json:"TechnicalDateTime,omitempty"`
	TimeAtLocationWithSeconds *time.Time `json:"TimeAtLocationWithSeconds,omitempty"`
	TrainOwner                *string    `json:"TrainOwner,omitempty"`
}

type trainAnnouncementCodeDescr struct {
	Code        *string `json:"Code,omitempty"`
	Description *string `json:"Description,omitempty"`
}

type TrainAnnouncement1Dot5 struct {
	ActivityID                            *string                      `json:"ActivityId,omitempty"`
	ActivityType                          *ActivityType                `json:"ActivityType,omitempty"`
	Advertised                            *bool                        `json:"Advertised,omitempty"`
	AdvertisedTimeAtLocation              *time.Time                   `json:"AdvertisedTimeAtLocation,omitempty"`
	AdvertisedTrainIdent                  *string                      `json:"AdvertisedTrainIdent,omitempty"`
	Booking                               []trainAnnouncementCodeDescr `json:"Booking,omitempty"`
	Canceled                              *bool                        `json:"Canceled,omitempty"`
	Deleted                               *bool                        `json:"Deleted,omitempty"`
	Deviation                             []trainAnnouncementCodeDescr `json:"Deviation,omitempty"`
	EstimatedTimeAtLocation               *time.Time                   `json:"EstimatedTimeAtLocation,omitempty"`
	EstimatedTimeIsPreliminary            *bool                        `json:"EstimatedTimeIsPreliminary,omitempty"`
	FromLocation                          []trainAnnouncementLocation  `json:"FromLocation,omitempty"`
	InformationOwner                      *string                      `json:"InformationOwner,omitempty"`
	LocationSignature                     *string                      `json:"LocationSignature,omitempty"`
	MobileWebLink                         *string                      `json:"MobileWebLink,omitempty"`
	ModifiedTime                          *time.Time                   `json:"ModifiedTime,omitempty"`
	NewEquipment                          *int                         `json:"NewEquipment,omitempty"`
	Operator                              *string                      `json:"Operator,omitempty"`
	OtherInformation                      []trainAnnouncementCodeDescr `json:"OtherInformation,omitempty"`
	PlannedEstimatedTimeAtLocation        *time.Time                   `json:"PlannedEstimatedTimeAtLocation,omitempty"`
	PlannedEstimatedTimeAtLocationIsValid *bool                        `json:"PlannedEstimatedTimeAtLocationIsValid,omitempty"`
	ProductInformation                    []trainAnnouncementCodeDescr `json:"ProductInformation,omitempty"`
	ScheduledDepartureDateTime            *time.Time                   `json:"ScheduledDepartureDateTime,omitempty"`
	Service                               []trainAnnouncementCodeDescr `json:"Service,omitempty"`
	TechnicalDateTime                     *time.Time                   `json:"TechnicalDateTime,omitempty"`
	TechnicalTrainIdent                   *string                      `json:"TechnicalTrainIdent,omitempty"`
	TimeAtLocation                        *time.Time                   `json:"TimeAtLocation,omitempty"`
	TimeAtLocationWithSeconds             *time.Time                   `json:"TimeAtLocationWithSeconds,omitempty"`
	ToLocation                            []trainAnnouncementLocation  `json:"ToLocation,omitempty"`
	TrackAtLocation                       *string                      `json:"TrackAtLocation,omitempty"`
	TrainComposition                      []trainAnnouncementCodeDescr `json:"TrainComposition,omitempty"`
	TrainOwner                            *string                      `json:"TrainOwner,omitempty"`
	TypeOfTraffic                         *string                      `json:"TypeOfTraffic,omitempty"`
	ViaFromLocation                       []trainAnnouncementLocation  `json:"ViaFromLocation,omitempty"`
	ViaToLocation                         []trainAnnouncementLocation  `json:"ViaToLocation,omitempty"`
	WebLink                               *string                      `json:"WebLink,omitempty"`
	WebLinkName                           *string                      `json:"WebLinkName,omitempty"`
}

type TrainAnnouncement1Dot6 struct {
	ActivityID                            *string                      `json:"ActivityId,omitempty"`
	ActivityType                          *ActivityType                `json:"ActivityType,omitempty"`
	Advertised                            *bool                        `json:"Advertised,omitempty"`
	AdvertisedTimeAtLocation              *time.Time                   `json:"AdvertisedTimeAtLocation,omitempty"`
	AdvertisedTrainIdent                  *string                      `json:"AdvertisedTrainIdent,omitempty"`
	Booking                               []trainAnnouncementCodeDescr `json:"Booking,omitempty"`
	Canceled                              *bool                        `json:"Canceled,omitempty"`
	Deleted                               *bool                        `json:"Deleted,omitempty"`
	Deviation                             []trainAnnouncementCodeDescr `json:"Deviation,omitempty"`
	EstimatedTimeAtLocation               *time.Time                   `json:"EstimatedTimeAtLocation,omitempty"`
	EstimatedTimeIsPreliminary            *bool                        `json:"EstimatedTimeIsPreliminary,omitempty"`
	FromLocation                          []trainAnnouncementLocation  `json:"FromLocation,omitempty"`
	InformationOwner                      *string                      `json:"InformationOwner,omitempty"`
	LocationSignature                     *string                      `json:"LocationSignature,omitempty"`
	MobileWebLink                         *string                      `json:"MobileWebLink,omitempty"`
	ModifiedTime                          *time.Time                   `json:"ModifiedTime,omitempty"`
	NewEquipment                          *int                         `json:"NewEquipment,omitempty"`
	Operator                              *string                      `json:"Operator,omitempty"`
	OtherInformation                      []trainAnnouncementCodeDescr `json:"OtherInformation,omitempty"`
	PlannedEstimatedTimeAtLocation        *time.Time                   `json:"PlannedEstimatedTimeAtLocation,omitempty"`
	PlannedEstimatedTimeAtLocationIsValid *bool                        `json:"PlannedEstimatedTimeAtLocationIsValid,omitempty"`
	ProductInformation                    []trainAnnouncementCodeDescr `json:"ProductInformation,omitempty"`
	ScheduledDepartureDateTime            *time.Time                   `json:"ScheduledDepartureDateTime,omitempty"`
	Service                               []trainAnnouncementCodeDescr `json:"Service,omitempty"`
	TechnicalDateTime                     *time.Time                   `json:"TechnicalDateTime,omitempty"`
	TechnicalTrainIdent                   *string                      `json:"TechnicalTrainIdent,omitempty"`
	TimeAtLocation                        *time.Time                   `json:"TimeAtLocation,omitempty"`
	TimeAtLocationWithSeconds             *time.Time                   `json:"TimeAtLocationWithSeconds,omitempty"`
	ToLocation                            []trainAnnouncementLocation  `json:"ToLocation,omitempty"`
	TrackAtLocation                       *string                      `json:"TrackAtLocation,omitempty"`
	TrainComposition                      []trainAnnouncementCodeDescr `json:"TrainComposition,omitempty"`
	TrainOwner                            *string                      `json:"TrainOwner,omitempty"`
	TypeOfTraffic                         []trainAnnouncementCodeDescr `json:"TypeOfTraffic,omitempty"`
	ViaFromLocation                       []trainAnnouncementLocation  `json:"ViaFromLocation,omitempty"`
	ViaToLocation                         []trainAnnouncementLocation  `json:"ViaToLocation,omitempty"`
	WebLink                               *string                      `json:"WebLink,omitempty"`
	WebLinkName                           *string                      `json:"WebLinkName,omitempty"`
}
