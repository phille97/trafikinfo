package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"

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

	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Println("Invalied credentials")
		os.Exit(1)
	}

	if resp.StatusCode == http.StatusBadRequest {
		type errmsg struct {
			Response struct {
				Result []struct {
					Error struct {
						Message string `json:"MESSAGE"`
					} `json:"ERROR"`
				} `json:"RESULT"`
			} `json:"RESPONSE"`
		}
		var e errmsg
		d := json.NewDecoder(resp.Body)
		err := d.Decode(&e)
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

	c, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(c))
}
