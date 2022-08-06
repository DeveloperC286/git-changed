#!/usr/bin/env sh

set -o errexit

output=$(gofmt -d -l "./src")
echo "${output}"
test -z "${output}"
