<h1 align="center">
🚦 Trafikinfo 🦺
</h1>
<h4 align="center">A Go library for the <a href="https://api.trafikinfo.trafikverket.se/">Trafikinfo API</a> from <a href="https://www.trafikverket.se/">Trafikverket</a></h4>
<p align="center">
    <a href="https://github.com/daenney/trafikinfo/actions/workflows/ci.yml"><img src="https://github.com/daenney/trafikinfo/actions/workflows/ci.yml/badge.svg" alt="Build Status"></a>
    <a href="https://codecov.io/gh/daenney/trafikinfo"><img src="https://codecov.io/gh/daenney/trafikinfo/branch/main/graph/badge.svg" alt="Coverage Status"></a>
    <a href="https://goreportcard.com/report/code.dny.dev/trafikinfo"><img src="https://goreportcard.com/badge/code.dny.dev/trafikinfo" alt="Go report card"></a>
    <a href="https://pkg.go.dev/code.dny.dev/trafikinfo"><img src="https://pkg.go.dev/badge/code.dny.dev/trafikinfo.svg" alt="GoDoc"></a>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/daenney/trafikinfo" alt="License: MIT"></a>
</p>

This library provides the necessary primitives to interact with the
Trafikinfo API. It contains a query builder that can be used to build up a
`Request` object. You can then `xml.Marshal` it and pass it on to your
favourite HTTP client to retrieve it. The API endpoint is available through
the `Endpoint` constant.

This library is under construction and currently lacks the struct types to
decode the response into.

## Usage

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	"code.dny.dev/trafikinfo"
)

func main() {
	req := trafikinfo.NewRequest().
		APIKey("YOUR_API_KEY").
		Query(
			trafikinfo.NewQuery(
				trafikinfo.WeatherStation,
				1.0,
			).Filter(
				trafikinfo.Equal("Id", "YOUR_STATION_ID"),
			),
		)

	body, err := xml.Marshal(req)
	if err != nil {
		panic(err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, trafikinfo.Endpoint, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	httpReq.Header.Set("content-type", "text/xml")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Rest of the code here, handle non-200 error responses etc.
}
```

More complete code can be found in the `examples/` directory.

Multiple queries can be passed by either passing multiple `NewQuery()` into a
single `Query()` call, or chaining `.Query()` multiple times on the result of
`NewRequest()`.

Calling `.Filter()` multiple times on a `Query` will replace the whole filter,
as a query can only have one filter block.
