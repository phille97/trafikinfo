package trafikinfo

import (
	"fmt"
	"time"
)

type Geometry struct {
	SWEREF99TM *string `json:"SWEREF99TM,omitempty"`
	WGS84      *string `json:"WGS84,omitempty"`
}

type Road struct {
	Temperature       *float64 `json:"Temp,omitempty"`
	TemperatureIconID *string  `json:"TempIconId,omitempty"`
}

type Info struct {
	LastModified *LastModified `json:"LASTMODIFIED,omitempty"`
	LastChangeID *string       `json:"LASTCHANGEID,omitempty"`
	EvalResult   []any         `json:"EVALRESULT,omitempty"`
	SSEURL       *string       `json:"SSEURL,omitempty"`
}

type LastModified struct {
	DateTime *time.Time `json:"_attr_datetime,omitempty"`
}

// CountyNumber is a numerical ID assigned to a county in Sweden
type CountyNumber uint

const (
	CountyStockholm CountyNumber = iota + 1
	CountyStockholmDeprecated
	CountyUppsala
	CountySodermanland
	CountyOstergotland
	CountyJonkoping
	CountyKronobergs
	CountyKalmar
	CountyGotland
	CountyBlekinge
	_
	CountySkane
	CountyHallands
	CountyVastraGotaland
	_
	_
	CountyVarmland
	CountyOrebro
	CountyVastmanland
	CountyDalarna
	CountyGavleborg
	CountyVasternorrland
	CountyJamtland
	CountyVasterbotten
	CountyNorrbotten
)

func (c CountyNumber) String() string {
	switch c {
	case CountyStockholm, CountyStockholmDeprecated:
		return "Stockholms län"
	case CountyUppsala:
		return "Uppsala län"
	case CountySodermanland:
		return "Södermanlands län"
	case CountyOstergotland:
		return "Östergötlands län"
	case CountyJonkoping:
		return "Jönköpings län"
	case CountyKronobergs:
		return "Kronobergs län"
	case CountyKalmar:
		return "Kalmar län"
	case CountyGotland:
		return "Gotlands län"
	case CountyBlekinge:
		return "Blekinge län"
	case CountySkane:
		return "Skåne län"
	case CountyHallands:
		return "Hallands län"
	case CountyVastraGotaland:
		return "Västra Götalands län"
	case CountyVarmland:
		return "Värmlands län"
	case CountyOrebro:
		return "Örebro län"
	case CountyVastmanland:
		return "Västmanlands län"
	case CountyDalarna:
		return "Dalarnas län"
	case CountyGavleborg:
		return "Gävleborgs län"
	case CountyVasternorrland:
		return "Västernorrlands län"
	case CountyJamtland:
		return "Jämtlands län"
	case CountyVasterbotten:
		return "Västerbottens län"
	case CountyNorrbotten:
		return "Norrbottens län"
	default:
		return fmt.Sprintf("Okänt län (%d)", c)
	}
}
