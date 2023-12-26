package main

import (
	"bufio"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"code.dny.dev/trafikinfo/internal/meta"
	"code.dny.dev/trafikinfo/internal/tree"
	"code.dny.dev/trafikinfo/internal/xsd"
)

func generate(_ context.Context, dir string, output string) error {
	schemas := []string{}

	if err := filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			schemas = append(schemas, path)
		}
		return nil
	}); err != nil {
		return err
	}

	for _, schema := range schemas {
		mdata := meta.ParseSchema(schema)

		r, err := os.Open(schema)
		if err != nil {
			return err
		}
		data, err := parse(r, mdata)
		if err != nil {
			return err
		}

		fmt.Println("walking XSD for ", mdata.Name, mdata.Version)
		render := tree.WalkXSD(data)

		intPath := filepath.Join(output, "internal", "trv", strings.ToLower(mdata.Name), mdata.PackageVersion())
		if err := os.MkdirAll(intPath, os.ModePerm); err != nil {
			return err
		}
		intPkg, err := os.Create(filepath.Join(intPath, "zz_gen.go"))
		if err != nil {
			return err
		}
		intBuf := bufio.NewWriter(intPkg)

		// Write the internal packages
		fmt.Println("rendering internal for", mdata.Name, mdata.Version)
		render.InternalTemplate(intBuf)

		if err := intBuf.Flush(); err != nil {
			return err
		}
		if err := intPkg.Sync(); err != nil {
			return err
		}
		if err := intPkg.Close(); err != nil {
			return err
		}

		pubPath := filepath.Join(output, "trv", strings.ToLower(mdata.Name), mdata.PackageVersion())
		if err := os.MkdirAll(pubPath, os.ModePerm); err != nil {
			return err
		}
		pubPkg, err := os.Create(filepath.Join(pubPath, "zz_gen.go"))
		if err != nil {
			return err
		}
		pubBuf := bufio.NewWriter(pubPkg)

		// Write the public packages
		fmt.Println("rendering public for", mdata.Name, mdata.Version)
		render.PublicTemplate(pubBuf)

		if err := pubBuf.Flush(); err != nil {
			return err
		}
		if err := pubPkg.Sync(); err != nil {
			return err
		}
		if err := pubPkg.Close(); err != nil {
			return err
		}
	}

	return nil
}

func parse(r io.Reader, meta meta.Data) (*xsd.Schema, error) {
	schema := xsd.Schema{
		Meta: meta,
	}
	dec := xml.NewDecoder(r)

	if err := dec.Decode(&schema); err != nil {
		return nil, err
	}
	return &schema, nil
}
