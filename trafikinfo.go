// Package trafikinfo provides the necessary primitives to interact with the
// Trafikinfo API. It contains a query builder that can be used to build up a
// [Request] object. You can then [encoding/xml.Marshal] it yourself or call
// [Request.Build]. It should be passed as the body to an HTTP request.
// You can [encoding/xml.Unmarshal] the response into a struct. Each
// package in [trv] provides a Response object that you can use, instead of having
// to specify your own. See for example
// [code.dny.dev/trafikinfo/trv/icon/v1dot1.Response].
//
// The data returned by the Trafikinfo API is licensed under Creative Commons
// CC0:
//   - https://data.trafikverket.se/documentation/datacache/intro
//   - https://creativecommons.org/publicdomain/zero/1.0/.
//
// The API endpoint is available through the [Endpoint] constant. The only
// supported endpoint is v2. All requests should be done with [net/http.MethodPost].
//
// Much of the code in this library is automatically generated using the XSD
// schemas provided by Trafikverket. This results in some pointer-heavy structs
// in [code.dny.dev/trafikinfo/internal/trv] as all fields are effectively
// optional. In order to make it easier to drill down into nested structs, the
// public structs in [trv] expose all fields as methods instead which you can
// chain, even if they're nil, making it easy to get to a deeply nested value.
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
//     API in practice)
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
package trafikinfo

const (
	Endpoint = "https://api.trafikinfo.trafikverket.se/v2/data.xml"
)
