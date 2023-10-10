package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"code.dny.dev/trafikinfo"
)

func main() {
	req, err := trafikinfo.NewRequest().
		APIKey("YOUR_API_KEY").
		Query(
			trafikinfo.NewQuery(
				trafikinfo.WeatherMeasurepoint,
				2.0,
			).Filter(
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

	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Println("Invalied credentials")
		os.Exit(1)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		var e trafikinfo.APIError
		err := json.Unmarshal(data, &e)
		if err != nil {
			fmt.Println("Could not decode API error response")
			os.Exit(1)
		}
		fmt.Println(e.Response.Result[0].Error.Message)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("got status code: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
		io.Copy(io.Discard, resp.Body)
		os.Exit(1)
	}

	type respMsg struct {
		Resonpse struct {
			Result []struct {
				MeasurePoint []trafikinfo.WeatherMeasurepoint2Dot0 `json:"WeatherMeasurepoint"`
				Info         trafikinfo.Info                       `json:"INFO"`
			}
		} `json:"RESPONSE"`
	}

	fmt.Println(string(data))

	var c respMsg
	err = json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", c)
}
