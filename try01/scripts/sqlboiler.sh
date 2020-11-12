#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
echo "${SCRIPT_DIR}"
cd "${SCRIPT_DIR}" && cd ../src

rm -rf ./backend/boiled/*

sqlboiler --wipe psql
