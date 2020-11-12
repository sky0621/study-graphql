#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
echo "${SCRIPT_DIR}"
cd "${SCRIPT_DIR}" && cd ../db

# https://github.com/rubenv/sql-migrate
#go get -v github.com/rubenv/sql-migrate/...
sql-migrate new -env="local" $@
