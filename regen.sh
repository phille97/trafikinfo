#!/bin/bash

rm -rf regen.schemas || true
rm -rf regen.output || true

go run ./cmd/trafikgen download -download-dir=regen.schemas

go run ./cmd/trafikgen generate -schema-dir=regen.schemas -output-dir=regen.output
rm -rf regen.schemas

go fmt ./regen.output/...

rm -rf trv/*/
mv regen.output/trv/* trv/

rm -rf internal/trv/*/
mv regen.output/internal/trv/* internal/trv/

rm -rf regen.output