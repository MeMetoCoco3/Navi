#!/bin/bash

OUTPUT=$(cat /tmp/GatorPath)
echo "You are at $OUTPUT" # Display the current location
cd "$OUTPUT"
