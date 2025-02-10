#!/bin/bash

set -e

mkdir -p /tmp/workspace/docker-cache

docker save -o /tmp/workspace/docker-cache/statisticooddschecker_console.tar statistico-odds-checker-console:latest
