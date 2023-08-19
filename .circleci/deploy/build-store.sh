#!/bin/bash

set -e

mkdir -p /tmp/workspace/artifacts

CGO_ENABLED=0 GOOS=linux go build -a -o statistico-odds-checker ./lambda/main.go

zip /tmp/workspace/artifacts/statistico-odds-checker.zip statistico-odds-checker
