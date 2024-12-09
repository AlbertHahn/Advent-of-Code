#!/bin/bash

COOKIE=$1
YEAR=$2
DAY=$3

curl -b session=${COOKIE} https://adventofcode.com/${YEAR}/day/${DAY}/input > ./${YEAR}/${DAY}/input.txt