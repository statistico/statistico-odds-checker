#!/bin/bash

set -e

mkdir -p /tmp/workspace/artifacts

CGO_ENABLED=0 GOOS=linux go build -a -o /tmp/workspace/artifacts/statistico-odds-checker ./lambda/main.go

cd /tmp/workspace/artifacts

zip statistico-odds-checker.zip statistico-odds-checker
