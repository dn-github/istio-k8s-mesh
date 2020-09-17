#!/bin/sh

set -e
export GO111MODULE=on
GOOS=linux go build -o details
docker build -t devnryn/details:v1 .
docker push devnryn/details:v1
rm details
