#!/bin/bash

# SHOULD SET FAVROUTE env variable
export FAVROUTE=$(pwd)/favorites/favs.json

if [ $# -eq 0 ]; then
  navy
  exit 0
fi

navy $1

if [ "$1" = "gator" ]; then
  OUTPUT=$(</tmp/GatorPath)
  echo "You are at $OUTPUT"
  cd "$OUTPUT"
  exit 1
fi
