package trafikinfo

import (
	"bytes"
	"testing"

	wmp "github.com/phille97/trafikinfo/trv/weathermeasurepoint/v2dot1"
)

const exampleResponse = `
<RESPONSE>
  <RESULT>
    <WeatherMeasurepoint>
      <Id>222</Id>
      <Name>Stocksund</Name>
      <Geometry>
        <SWEREF99TM>POINT (672791.011 6586802.373)</SWEREF99TM>
        <WGS84>POINT (18.04221 59.38437)</WGS84>
      </Geometry>
      <Observation>
        <Sample>2026-03-09T21:45:03.001+01:00</Sample>
        <Weather>
          <Precipitation>no</Precipitation>
        </Weather>
        <Surface>
          <Temperature>
            <Origin>measured</Origin>
            <SensorNames>DTS12G_1</SensorNames>
            <Value>2.1</Value>
          </Temperature>
        </Surface>
        <Air>
          <Temperature>
            <Origin>measured</Origin>
            <SensorNames>HMP155_1</SensorNames>
            <Value>3</Value>
          </Temperature>
          <Dewpoint>
            <SensorNames>HMP155_1</SensorNames>
            <Value>-0.4</Value>
          </Dewpoint>
          <RelativeHumidity>
            <Origin>measured</Origin>
            <SensorNames>HMP155_1</SensorNames>
            <Value>75.9</Value>
          </RelativeHumidity>
          <VisibleDistance>
            <Origin>measured</Origin>
            <SensorNames>PWD22_1</SensorNames>
            <Value>20000</Value>
          </VisibleDistance>
        </Air>
        <Wind>
          <Height>6</Height>
          <Speed>
            <Origin>measured</Origin>
            <SensorNames>WMT700_1</SensorNames>
            <Value>0.4</Value>
          </Speed>
          <Direction>
            <Origin>measured</Origin>
            <SensorNames>WMT700_1</SensorNames>
            <Value>233</Value>
          </Direction>
        </Wind>
        <Aggregated5minutes>
          <Precipitation>
            <Rain>false</Rain>
            <Snow>false</Snow>
            <RainSum>
              <Origin>measured</Origin>
              <SensorNames>PWD22_1</SensorNames>
              <Value>0</Value>
            </RainSum>
            <SnowSum>
              <Solid>
                <Value>0</Value>
              </Solid>
              <WaterEquivalent>
                <Origin>measured</Origin>
                <SensorNames>PWD22_1</SensorNames>
                <Value>0</Value>
              </WaterEquivalent>
            </SnowSum>
            <TotalWaterEquivalent>
              <Value>0</Value>
            </TotalWaterEquivalent>
          </Precipitation>
        </Aggregated5minutes>
        <Aggregated10minutes>
          <Wind>
            <SpeedMax>
              <Origin>measured</Origin>
              <SensorNames>WMT700_1</SensorNames>
              <Value>1.5</Value>
            </SpeedMax>
            <SpeedAverage>
              <Origin>measured</Origin>
              <SensorNames>WMT700_1</SensorNames>
              <Value>0.4</Value>
            </SpeedAverage>
          </Wind>
          <Precipitation>
            <Rain>false</Rain>
            <Snow>false</Snow>
            <RainSum>
              <Origin>measured</Origin>
              <SensorNames>PWD22_1</SensorNames>
              <Value>0</Value>
            </RainSum>
            <SnowSum>
              <Solid>
                <Value>0</Value>
              </Solid>
              <WaterEquivalent>
                <Origin>measured</Origin>
                <SensorNames>PWD22_1</SensorNames>
                <Value>0</Value>
              </WaterEquivalent>
            </SnowSum>
            <TotalWaterEquivalent>
              <Value>0</Value>
            </TotalWaterEquivalent>
          </Precipitation>
        </Aggregated10minutes>
        <Aggregated30minutes>
          <Wind>
            <SpeedMax>
              <Origin>measured</Origin>
              <SensorNames>WMT700_1</SensorNames>
              <Value>1.5</Value>
            </SpeedMax>
            <SpeedAverage>
              <Origin>measured</Origin>
              <SensorNames>WMT700_1</SensorNames>
              <Value>0.5</Value>
            </SpeedAverage>
          </Wind>
          <Precipitation>
            <Rain>false</Rain>
            <Snow>false</Snow>
            <RainSum>
              <Origin>measured</Origin>
              <SensorNames>PWD22_1</SensorNames>
              <Value>0</Value>
            </RainSum>
            <SnowSum>
              <Solid>
                <Value>0</Value>
              </Solid>
              <WaterEquivalent>
                <Origin>measured</Origin>
                <SensorNames>PWD22_1</SensorNames>
                <Value>0</Value>
              </WaterEquivalent>
            </SnowSum>
            <TotalWaterEquivalent>
              <Value>0</Value>
            </TotalWaterEquivalent>
          </Precipitation>
        </Aggregated30minutes>
        <Id>381759691</Id>
      </Observation>
      <Deleted>false</Deleted>
      <ModifiedTime>2026-03-09T20:50:08.242Z</ModifiedTime>
    </WeatherMeasurepoint>
    <WeatherMeasurepoint>
      <Id>223</Id>
      <Name>Ullnasjön</Name>
      <Geometry>
        <SWEREF99TM>POINT (679184.894 6598792.808)</SWEREF99TM>
        <WGS84>POINT (18.16462 59.48921)</WGS84>
      </Geometry>
      <Observation>
        <Sample>2026-03-09T21:45:03.001+01:00</Sample>
        <Weather>
          <Precipitation>no</Precipitation>
        </Weather>
        <Surface>
          <Temperature>
            <Origin>measured</Origin>
            <SensorNames>DTS12G_1</SensorNames>
            <Value>1.7</Value>
          </Temperature>
        </Surface>
        <Air>
          <Temperature>
            <Origin>measured</Origin>
            <SensorNames>HMP155_1</SensorNames>
            <Value>2.4</Value>
          </Temperature>
          <Dewpoint>
            <SensorNames>HMP155_1</SensorNames>
            <Value>-0.1</Value>
          </Dewpoint>
          <RelativeHumidity>
            <Origin>measured</Origin>
            <SensorNames>HMP155_1</SensorNames>
            <Value>81.4</Value>
          </RelativeHumidity>
          <VisibleDistance>
            <Origin>measured</Origin>
            <SensorNames>PWD22_1</SensorNames>
            <Value>20000</Value>
          </VisibleDistance>
        </Air>
        <Wind>
          <Height>6</Height>
          <Speed>
            <Origin>measured</Origin>
            <SensorNames>WMT700_1</SensorNames>
            <Value>0.9</Value>
          </Speed>
          <Direction>
            <Origin>measured</Origin>
            <SensorNames>WMT700_1</SensorNames>
            <Value>337</Value>
          </Direction>
        </Wind>
        <Aggregated5minutes>
          <Precipitation>
            <Rain>false</Rain>
            <Snow>false</Snow>
            <RainSum>
              <Origin>measured</Origin>
              <SensorNames>PWD22_1</SensorNames>
              <Value>0</Value>
            </RainSum>
            <SnowSum>
              <Solid>
                <Value>0</Value>
              </Solid>
              <WaterEquivalent>
                <Origin>measured</Origin>
                <SensorNames>PWD22_1</SensorNames>
                <Value>0</Value>
              </WaterEquivalent>
            </SnowSum>
            <TotalWaterEquivalent>
              <Value>0</Value>
            </TotalWaterEquivalent>
          </Precipitation>
        </Aggregated5minutes>
        <Aggregated10minutes>
          <Wind>
            <SpeedMax>
              <Origin>measured</Origin>
              <SensorNames>WMT700_1</SensorNames>
              <Value>1.8</Value>
            </SpeedMax>
            <SpeedAverage>
              <Origin>measured</Origin>
              <SensorNames>WMT700_1</SensorNames>
              <Value>0.9</Value>
            </SpeedAverage>
          </Wind>
          <Precipitation>
            <Rain>false</Rain>
            <Snow>false</Snow>
            <RainSum>
              <Origin>measured</Origin>
              <SensorNames>PWD22_1</SensorNames>
              <Value>0</Value>
            </RainSum>
            <SnowSum>
              <Solid>
                <Value>0</Value>
              </Solid>
              <WaterEquivalent>
                <Origin>measured</Origin>
                <SensorNames>PWD22_1</SensorNames>
                <Value>0</Value>
              </WaterEquivalent>
            </SnowSum>
            <TotalWaterEquivalent>
              <Value>0</Value>
            </TotalWaterEquivalent>
          </Precipitation>
        </Aggregated10minutes>
        <Aggregated30minutes>
          <Wind>
            <SpeedMax>
              <Origin>measured</Origin>
              <SensorNames>WMT700_1</SensorNames>
              <Value>2.3</Value>
            </SpeedMax>
            <SpeedAverage>
              <Origin>measured</Origin>
              <SensorNames>WMT700_1</SensorNames>
              <Value>1</Value>
            </SpeedAverage>
          </Wind>
          <Precipitation>
            <Rain>false</Rain>
            <Snow>false</Snow>
            <RainSum>
              <Origin>measured</Origin>
              <SensorNames>PWD22_1</SensorNames>
              <Value>0</Value>
            </RainSum>
            <SnowSum>
              <Solid>
                <Value>0</Value>
              </Solid>
              <WaterEquivalent>
                <Origin>measured</Origin>
                <SensorNames>PWD22_1</SensorNames>
                <Value>0</Value>
              </WaterEquivalent>
            </SnowSum>
            <TotalWaterEquivalent>
              <Value>0</Value>
            </TotalWaterEquivalent>
          </Precipitation>
        </Aggregated30minutes>
        <Id>381759981</Id>
      </Observation>
      <Deleted>false</Deleted>
      <ModifiedTime>2026-03-09T20:50:09.622Z</ModifiedTime>
    </WeatherMeasurepoint>
  </RESULT>
</RESPONSE>
`

const exampleResponseWithAPIError = `
<RESPONSE>
  <RESULT>
    <ERROR>
      <SOURCE>Security</SOURCE>
      <MESSAGE>Invalid authentication</MESSAGE>
    </ERROR>
  </RESULT>
</RESPONSE>
`

func TestStreamResult(t *testing.T) {
	results := StreamResult[wmp.WeatherMeasurepoint](bytes.NewBufferString(exampleResponse))
	var data []wmp.WeatherMeasurepoint
	for result, err := range results {
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if result.Error != nil {
			t.Fatalf("unexpected API error: %s", result.Error)
		}
		data = append(data, result.Data)
	}

	if len(data) != 2 {
		t.Fatalf("expected 2 results, got %d", len(data))
	}

	if *data[0].ID() != "222" {
		t.Errorf("expected first result to have ID 222, got %s", *data[0].ID())
	}

	if *data[1].ID() != "223" {
		t.Errorf("expected second result to have ID 223, got %s", *data[1].ID())
	}
}

func TestStreamResultAPIError(t *testing.T) {
	results := StreamResult[wmp.WeatherMeasurepoint](bytes.NewBufferString(exampleResponseWithAPIError))
	count := 0
	for result, err := range results {
		count++
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if result.Error == nil {
			t.Fatalf("expected API error, got nil")
		}
		if result.Error.Source != "Security" {
			t.Errorf("expected error source to be Security, got %s", result.Error.Source)
		}
		if result.Error.Message != "Invalid authentication" {
			t.Errorf("expected error message to be Invalid authentication, got %s", result.Error.Message)
		}
	}

	if count != 1 {
		t.Fatalf("expected 1 result, got %d", count)
	}
}
