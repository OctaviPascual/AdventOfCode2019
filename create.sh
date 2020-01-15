#!/usr/bin/env bash

set -euxo pipefail

function _help() {
    echo "\
start: tool to create new day
Usage: ./create.sh [DAY]
"
}

DAY="${1}"
PUZZLE_URL="https://adventofcode.com/2019/day/${DAY}/input"

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
