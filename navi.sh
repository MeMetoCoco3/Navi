#!/bin/bash

export FAVROUTE="$NAVIPATH/favorites/favs.json"
if [ $# -eq 0 ]; then
  naviGO
elif [ "$1" == "gator" ] || [ "$1" == "fv" ]; then
  naviGO "$1" "$2"
  OUTPUT=$(cat /tmp/GatorPath)
  cd "$OUTPUT"
  echo "You are at $OUTPUT"
else
  naviGO "$1" "$2"
fi
