// Package trafikinfo provides the necessary primitives to interact with the
// Trafikinfo API. It contains a query builder that can be used to build up a
// [Request] object. You can then [encoding/xml.Marshal] it yourself or call
// [Request.Build]. It should be passed as the body to an HTTP request.
// You can [encoding/xml.Unmarshal] the response into a struct.
//
// Each package in [trv] provides a [trv.ObjectType] that you can pass to [NewQuery]
// to create a [Request] for that object type. It also provides a Response object
// that you can use to decode the response, instead of having to specify your own.
// A package's Response type can only be used if all [Query] in a [Request] are
// for the same [trv.ObjectType]. For other cases you'll need to craft a response
// struct yourself.
//
// Requests should be done against the v2 [Endpoint]. Earlier versions of the API
// are deprecated by Trafikverket and not supported in this library. All requests
// must be POST requests, for example using [net/http.NewRequestWithContext]
// together with [net/http.MethodPost].
//
// A number of endpoints return IconID elements, either at the root of the
// object, or in one of its embedded objects. Each of these IconID strings can
// be plugged into a query for the [code.dny.dev/trafikinfo/trv/icon/v1dot1.Icon]
// resource. It in turn will give you access to a base64 encoded PNG in the response,
// as well as a URL at which the icon can be retrieved.
//
// When using certain filters, for example [Equal], querying for an attribute
// with a value that doesn't exist will not result in an error by the API.
// Instead you'll get an empty response.
//
// When creating a query, a Change ID of 0 is automatically included, unless
// explicitly set to something else with [Query.ChangeID]. The query also always
// requests for a [trv.LastModified] to be returned in the response.
//
// # Package design
//
// Much of the code in this library is automatically generated using the XSD
// schemas provided by Trafikverket. This results in some pointer-heavy structs
// in [code.dny.dev/trafikinfo/internal/trv] as all fields are effectively
// optional.
//
// In order to make it easier to drill down into nested structs, the
// public structs in [trv] expose all fields as methods instead which perform
// nil-checks themselves when returning nested values. In the case of nested
// structs, they'll instead return a pointer to an empty struct. This results
// in a fluent-style API, letting you write something like:
//
//	Observation().Aggregated5minutes().Precipitation().TotalWaterEquivalent().Value()
//
// You will get a value back which may be nil, no matter if Observation,
// Aggregated5minutes, Precipitation and TotalWaterEquivalent elements where
// present in the response or not.
//
// When a method returns a slice, that slice may be nil or empty. You can always
// safely iterate over those using a for loop, or check the length of the returned
// slice. A nil slice always has a length of 0.
//
// # XML to Go type mapping
//
// The data types are mapped to their closest representation in Go:
//   - xs:long and :integer are mapped to int64 (to ensure they still decode correctly
//     on 32-bit systems)
//   - xs:int to int (as int is always at least 32 bits)
//   - xs:unsignedByte to uint8
//   - xs:dateTime to time.Time
//   - xs:float, :double and :decimal to float64 (for xs:float float32 is in
//     theory enough but this avoids needing to deal with the extra type, whereas
//     xs:decimal could in theory exceed a float64 but this never happens with the
//     API in practice as it's only used for sensor values)
//
// # Data License
//
// The data returned by the Trafikinfo API is licensed under Creative Commons
// CC0:
//   - https://data.trafikverket.se/documentation/datacache/intro
//   - https://creativecommons.org/publicdomain/zero/1.0/.
package trafikinfo

const (
	Endpoint = "https://api.trafikinfo.trafikverket.se/v2/data.xml"
)
