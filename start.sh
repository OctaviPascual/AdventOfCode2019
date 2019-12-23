#! /usr/bin/env bash

DAY="${1}"
PUZZLE_URL="https://adventofcode.com/2019/day/${DAY}/input"

curl "${PUZZLE_URL}" -H "cookie: session=${AOC_SESSION_COOKIE}" -o day.txt 2>/dev/null
