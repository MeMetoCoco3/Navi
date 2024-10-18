#!/bin/bash

# Capture the output of the Go program
go run viws.go >cd

# Pass the captured output to the 'main goto' command
DIR=$(main goto "$OUTPUT")

# Print and use the output
echo -e "\n$DIR\n"

# Change directory based on the output
cd "$DIR"
