#!/bin/sh

if [ "$ENVIRONMENT" = 'development' ]; then
    echo "Starting in dev env..."
    exec go run status-api.go
else
    echo "Building and starting..."
    go build -o statusApi .
    exec ./statusApi
fi