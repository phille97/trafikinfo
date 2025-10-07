<h1 align="center">
🚦 Trafikinfo 🦺
</h1>
<h4 align="center">A Go library for the <a href="https://api.trafikinfo.trafikverket.se/">Trafikinfo API</a> from <a href="https://www.trafikverket.se/">Trafikverket</a></h4>
<p align="center">
    <a href="https://github.com/phille97/trafikinfo/actions/workflows/ci.yml"><img src="https://github.com/phille97/trafikinfo/actions/workflows/ci.yml/badge.svg" alt="Build Status"></a>
	<a href="https://github.com/phille97/trafikinfo/releases"><img src="https://img.shields.io/github/release/phille97/trafikinfo.svg" alt="Release"></a>
    <a href="https://pkg.go.dev/github.com/phille97/trafikinfo"><img src="https://pkg.go.dev/badge/github.com/phille97/trafikinfo.svg" alt="GoDoc"></a>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/phille97/trafikinfo" alt="License: MIT"></a>
</p>

This library provides the necessary primitives to interact with the
Trafikinfo API. It contains a query builder that can be used to build up a
`Request` object. You can then `xml.Marshal` it and pass it on to your
favourite HTTP client to retrieve it. The response can be decoded using `xml.Unmarshal`.

The data returned by the [Trafikinfo API is licensed][tl] under [Creative Commons
CC0][cc0].

[tl]: https://data.trafikverket.se/documentation/datacache/intro
[cc0]: https://creativecommons.org/publicdomain/zero/1.0/

## Usage

See the example client in [`examples/client`](examples/client/). It includes
preparing the query, decoding the response etc.

Multiple queries can be passed by either passing multiple `NewQuery()` into a
single `Query()` call, or chaining `.Query()` multiple times on the result of
`NewRequest()`.

Calling `.Filter()` multiple times on a `Query` will replace the whole filter,
as a query can only have one filter block.

## Supported object types

This library provides facilities for response decoding for the following object
types and versions.

| Object | Version(s) |
:-- | :-----------
`Camera` | 1
`FerryAnnouncement` | 1.2
`FerryRoute` | 1.2
`Icon` | 1.1
`MeasurementData20` | 1.0
`MeasurementData100` | 1.0
`Parking` | 1.4
`PavementData` | 1
`RailCrossing` | 1.5
`ReasonCode` | 1
`RoadCondition` | 1.2
`RoadData` | 1
`RoadGeometry` | 1
`Situation` | 1.5
`TrafficFlow` | 1.4
`TrafficSafetyCamera` | 1
`TrainAnnouncement` | 1.9
`TrainMessage` | 1.7
`TrainPosition` | 1.1
`TrainStation` | 1.4
`TrainStationMessage` | 1
`TravelTimeRoute` | 1.5
`WeatherMeasurepoint` | 2.1
`WeatherObservation` | 2.1
