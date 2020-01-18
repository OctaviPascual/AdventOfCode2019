#!/usr/bin/env bash

set -euo pipefail

function _help() {
    echo "\
Tool to bootstrap a new day
Usage: ./bootstrap.sh DAY
Note that AOC_SESSION_COOKIE environment variable must be set to download the input
"
}

DAY="${1:-}"

if ! [[ $DAY =~ ^[1-9][0-9]*$ ]] ; then
    _help
    echo "error: DAY must be a number"
    exit 1
fi

if ! ((${DAY} >= 1 && ${DAY} <= 25)); then
    _help
    echo "error: DAY must be between 1 and 25"
    exit 1
fi

set -x

PUZZLE_URL="https://adventofcode.com/2019/day/${DAY}/input"

# Append a 0 at the beginning of the day if it's less than 10
if (( ${DAY} < 10 )); then
    DAY=0${DAY}
fi

DIRNAME="day${DAY}"
PUZZLE_FILE="${DIRNAME}/day${DAY}.txt"
GO_FILE="${DIRNAME}/day${DAY}.go"
GO_TEST_FILE="${DIRNAME}/day${DAY}_test.go"

mkdir "${DIRNAME}"
curl "${PUZZLE_URL}" -H "cookie: session=${AOC_SESSION_COOKIE}" -o "${PUZZLE_FILE}" 2>/dev/null
chmod 0444 "${PUZZLE_FILE}"

cp "bootstrap/dayXX.template" "${GO_FILE}"
chmod 0644 "${GO_FILE}"
sed -i "" "s/XX/${DAY}/g" "${GO_FILE}"

cp "bootstrap/dayXX_test.template" "${GO_TEST_FILE}"
chmod 0644 "${GO_TEST_FILE}"
sed -i "" "s/XX/${DAY}/g" "${GO_TEST_FILE}"
