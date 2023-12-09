#!/bin/bash

set -e

mkdir "day$1"
cp "day1/Makefile" "day$1/."
echo -e "package main"> "day$1/main.go"
touch "day$1/README.md"