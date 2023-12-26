package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"code.dny.dev/trafikinfo/internal/meta"
)

var (
	base       = "https://data.trafikverket.se"
	schemaPath = "/apb/prod/schema/%s/%s/data?annotations=true"
	trvSchemas = map[string][]string{
		// Choo choo motherfucker
		"RailCrossing":        {"1.5"},
		"ReasonCode":          {"1"},
		"TrainAnnouncement":   {"1.9"},
		"TrainMessage":        {"1.7"},
		"TrainPosition":       {"1.1"},
		"TrainStation":        {"1.4"},
		"TrainStationMessage": {"1"},
		// Road surface
		"MeasurementData100": {"1"},
		"MeasurementData20":  {"1"},
		"PavementData":       {"1"},
		"RoadData":           {"1"},
		"RoadGeometry":       {"1"},
		// Vroom vroom
		"Camera":              {"1"},
		"Icon":                {"1.1"},
		"Parking":             {"1.4"},
		"RoadCondition":       {"1.2"},
		"Situation":           {"1.5"},
		"TrafficFlow":         {"1.4"},
		"TrafficSafetyCamera": {"1"},
		"TravelTimeRoute":     {"1.5"},
		"WeatherMeasurepoint": {"2"},
		"WeatherObservation":  {"2"},
		// Toot toot
		"FerryAnnouncement": {"1.2"},
		"FerryRoute":        {"1.2"},
	}
)

func download(ctx context.Context, dir string) error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	for schema, versions := range trvSchemas {
		for _, version := range versions {
			out, err := os.Create(path.Join(dir, fmt.Sprintf("%s_%s.zip", schema, version)))
			if err != nil {
				return err
			}

			namespace := meta.FindNamespace(schema, version)
			schemaName := schema
			if namespace != "" {
				schemaName = fmt.Sprintf("%s.%s", namespace, schemaName)
			}
			url := base + fmt.Sprintf(schemaPath, schemaName, version)
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				return err
			}
			req.Header.Set("User-Agent", "trafikgen (+https://code.dny.dev/trafikinfo)")
			log.Printf("download starting: %s @ %s...\n", schema, version)
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				return fmt.Errorf("failed to fetch: %s %s, status: %d", schema, version, resp.StatusCode)
			}
			if _, err := io.Copy(out, resp.Body); err != nil {
				return err
			}
			log.Printf("download completed: %s @ %s\n", schema, version)
		}
	}

	client.CloseIdleConnections()
	return nil
}

func extract(_ context.Context, src, dst string) error {
	var files []string
	if err := filepath.WalkDir(src, func(path string, f fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			if filepath.Ext(path) == ".zip" {
				files = append(files, path)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	for _, f := range files {
		log.Printf("extracting: %s to: %s...\n", f, dst)
		archive, err := zip.OpenReader(f)
		if err != nil {
			return err
		}
		defer archive.Close()
		for _, file := range archive.Reader.File {
			reader, err := file.Open()
			if err != nil {
				return err
			}
			defer reader.Close()
			path := filepath.Join(dst, file.Name)
			_ = os.Remove(path)
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
			if file.FileInfo().IsDir() {
				continue
			}
			err = os.Remove(path)
			if err != nil {
				return err
			}
			writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer writer.Close()
			_, err = io.Copy(writer, reader)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
