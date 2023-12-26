// Package meta contains some metadata information about the schemas and
// associated helpers to extract that information.
package meta

import (
	"path"
	"strings"
)

var namespaceMapping = map[string]map[string]string{
	"FerryAnnouncement": {
		"1.2": "Ferry.TrafficInfo",
	},
	"FerryRoute": {
		"1.2": "Ferry.TrafficInfo",
	},
	"Icon": {
		"1.1": "Road.Infrastructure",
	},
	"TrafficSafetyCamera": {
		"1": "Road.Infrastructure",
	},
	"TrainPosition": {
		"1.1": "järnväg.trafikinfo",
	},
	"TrainStation": {
		"1.4": "rail.infrastructure",
	},
}

type Data struct {
	Name, Version, Namespace string
}

func (m Data) PackageVersion() string {
	s := strings.ReplaceAll(m.Version, ".", "dot")
	s = "v" + s
	return s
}

func ParseSchema(in string) Data {
	in = strings.TrimSuffix(path.Base(in), path.Ext(in))
	in = strings.TrimPrefix(in, "response_")
	elems := strings.Split(in, "_")

	name := elems[0]
	version := strings.TrimPrefix(elems[1], "v")

	return Data{
		Name:      name,
		Version:   version,
		Namespace: FindNamespace(name, version),
	}
}

func FindNamespace(name, version string) string {
	if val, ok := namespaceMapping[name]; ok {
		if nval, nok := val[version]; nok {
			return nval
		}
	}
	return ""
}
