#!/bin/bash

name="blog"
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags=prod -tags=msgaudit -ldflags "-s -w" -o $name main.go


