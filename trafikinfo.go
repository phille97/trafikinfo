package trafikinfo

// Endpoint is the current recommended endpoint for the Trafikinfo API
const Endpoint = "https://api.trafikinfo.trafikverket.se/v2/data.json"

// ObjectType is a type of object you can retrieve from the API
type ObjectType string

const (
	RailCrossing        ObjectType = "RailCrossing"
	ReasonCode          ObjectType = "ReasonCode"
	TrainAnnouncement   ObjectType = "TrainAnnouncement"
	TrainMessage        ObjectType = "TrainMessage"
	TrainStation        ObjectType = "TrainStation"
	TrainStationMessage ObjectType = "TrainStationMessage"
	TrainPosition       ObjectType = "TrainPosition"

	Camera                ObjectType = "Camera"
	FerryAnnouncement     ObjectType = "FerryAnnouncement"
	FerryRoute            ObjectType = "FerryRoute"
	Icon                  ObjectType = "Icon"
	Parking               ObjectType = "Parking"
	RoadCondition         ObjectType = "RoadCondition"
	RoadConditionOverview ObjectType = "RoadConditionOverview"
	Situation             ObjectType = "Situation"
	TrafficFlow           ObjectType = "TrafficFlow"
	TrafficSafetyCamera   ObjectType = "TrafficSafetyCamera"
	TravelTimeRoute       ObjectType = "TravelTimeRoute"
	WeatherMeasurePoint   ObjectType = "WeatherMeasurePoint"
	WeatherObservation    ObjectType = "WeatherObservation"
	WeatherStation        ObjectType = "WeatherStation"

	MeasurementData100 ObjectType = "MeasurementData100"
	MeasurementData20  ObjectType = "MeasurementData20"
	PavementData       ObjectType = "PavementData"
	RoadData           ObjectType = "RoadData"
	RoadGeometry       ObjectType = "RoadGeometry"
)
