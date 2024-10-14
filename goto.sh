#!/bin/bash

DIR=$(main goto "$1")

# -e to allow \n
echo -e "\n$DIR\n"
cd "$DIR"
