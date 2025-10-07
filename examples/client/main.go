// An example on how to use the trafikinfo library to request weather data.
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/phille97/trafikinfo"
	wmp "github.com/phille97/trafikinfo/trv/weathermeasurepoint/v2dot1"
)

func main() {
	req, err := trafikinfo.NewRequest().
		APIKey("YOUR_API_KEY").
		Query(
			trafikinfo.NewQuery(wmp.T()).Filter(
				trafikinfo.Equal("Name", "YOUR_STATION_NAME"),
			),
		).Build()
	if err != nil {
		panic(err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, trafikinfo.Endpoint, bytes.NewBuffer(req))
	if err != nil {
		panic(err)
	}

	httpReq.Header.Set("content-type", "text/xml")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read response data")
		os.Exit(1)
	}

	var res wmp.Response
	if err := xml.Unmarshal(data, &res); err != nil {
		fmt.Println("Could not decode API response")
		fmt.Println(err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK || res.HasErrors() {
		fmt.Println("Status code", resp.StatusCode, "with errors", res.ErrorMsg())
		os.Exit(1)
	}

	fmt.Println(string(data))
	fmt.Printf("%+v\n", res)
}
