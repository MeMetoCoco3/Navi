#!/bin/bash

# SHOULD SET FAVROUTE env variable

rm main
go build main.go
# Capture the output of the Go program
main Gator
OUTPUT=$(</tmp/GatorPath)
# Change directory based on the output
echo "You are at $OUTPUT"
cd "$OUTPUT"
