<h1 align="center">
🚦 Trafikinfo 🦺
</h1>
<h4 align="center">A Go library for the <a href="https://api.trafikinfo.trafikverket.se/">Trafikinfo API</a> from <a href="https://www.trafikverket.se/">Trafikverket</a></h4>
<p align="center">
    <a href="https://github.com/daenney/trafikinfo/actions/workflows/ci.yml"><img src="https://github.com/daenney/trafikinfo/actions/workflows/ci.yml/badge.svg" alt="Build Status"></a>
	<a href="https://github.com/daenney/trafikinfo/releases"><img src="https://img.shields.io/github/release/daenney/trafikinfo.svg" alt="Release"></a>
    <a href="https://codecov.io/gh/daenney/trafikinfo"><img src="https://codecov.io/gh/daenney/trafikinfo/branch/main/graph/badge.svg" alt="Coverage Status"></a>
    <a href="https://goreportcard.com/report/code.dny.dev/trafikinfo"><img src="https://goreportcard.com/badge/code.dny.dev/trafikinfo" alt="Go report card"></a>
    <a href="https://pkg.go.dev/code.dny.dev/trafikinfo"><img src="https://pkg.go.dev/badge/code.dny.dev/trafikinfo.svg" alt="GoDoc"></a>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/daenney/trafikinfo" alt="License: MIT"></a>
</p>

This library provides the necessary primitives to interact with the
Trafikinfo API. It contains a query builder that can be used to build up a
`Request` object. You can then `xml.Marshal` it and pass it on to your
favourite HTTP client to retrieve it.

The API endpoint is available through the `Endpoint` constant. The only
supported endpoint is `v2`, `v1.x` is being deprecated by Trafikverket.

This library is under construction and currently lacks the struct types to
decode the response into.

## Usage

```go
req, err := trafikinfo.NewRequest().
	APIKey("YOUR_API_KEY").
	Query(
		trafikinfo.NewQuery(
			trafikinfo.WeatherStation,
			1.0,
		).Filter(
			trafikinfo.Equal("Id", "YOUR_STATION_ID"),
		),
	).Build()

if err != nil {
	panic(err)
}

// Rest of the code here to do the request, handle non-200 error
// responses etc.

type respMsg struct {
	Resonpse struct {
		Result []struct {
			WeatherStation []trafikinfo.WeatherStation1Dot0 `json:"WeatherStation"`
			Info           trafikinfo.Info                  `json:"INFO"`
		}
	} `json:"RESPONSE"`
}

var c respMsg
d := json.NewDecoder(resp.Body)
err = d.Decode(&c)

if err != nil {
	panic(err)
}

// Enjoy your struct!
```

More complete code can be found in the `examples/` directory.

Multiple queries can be passed by either passing multiple `NewQuery()` into a
single `Query()` call, or chaining `.Query()` multiple times on the result of
`NewRequest()`.

Calling `.Filter()` multiple times on a `Query` will replace the whole filter,
as a query can only have one filter block.

## Supported object types

This library provides structs for response decoding for the following object
types and versions.

| Object | Versions |
:-- | :-----------
`Camera` | 1.0: ❌
`FerryAnnouncement` | 1.2: ❌
`FerryRoute` | 1.2: ❌
`Icon` | 1.0: ❌
`MeasurementData100` | 1.0: ❌
`MeasurementData20` | 1.0: ❌
`Parking` | 1.0: ❌ 1.4: ❌
`PavementData` | 1.0: ❌
`RailCrossing` | 1.4: ❌ 1.5: ❌
`ReasonCode` | 1.0: ❌
`RoadCondition` | 1.0: ❌ 1.1: ❌ 1.2: ❌
`RoadConditionOverview` | 1.0: ❌
`RoadData` | 1.0: ❌
`RoadGeometry` | 1.0: ❌
`TrafficFlow` | 1.0: ❌ 1.4: ❌
`TrafficSafetyCamera` | 1.0: ❌
`TrainAnnouncement` | 1.0: ❌ 1.3: ❌ 1.4: ❌ 1.5: ❌ 1.6: ❌
`TrainMessage` | 1.0: ❌ 1.3: ❌ 1.4: ❌ 1.5: ❌ 1.6: ❌ 1.7: ❌
`TrainStation` | 1.0: ❌ 1.4: ❌
`TravelTimeRoute` | 1.0: ❌ 1.3: ❌ 1.4: ❌ 1.5
`WeatherMeasurepoint` | 1.0: ❌ 2.0: ❌
`WeatherObservation` | 1.0: ❌ 2.0: ❌
`WeatherStation` | 1.0: ✅
