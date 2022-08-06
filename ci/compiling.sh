#!/usr/bin/env sh

set -o errexit

go build -o git-changed "./src/"
