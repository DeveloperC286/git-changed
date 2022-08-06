#!/usr/bin/env sh

set -o errexit

gofmt -w -s "./src"
