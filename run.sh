#!/bin/bash

readonly SRC_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
readonly PACKAGES=". ./types"

echo "Building packages: $PACKAGES"

echo -n "go get......."
output=$(go get $PACKAGES 2>&1) && echo "done." || {
  echo "error."
  echo "$output"
  exit 1
}

echo -n "go install..."
output=$(go install $PACKAGES 2>&1) && echo "done." || {
  echo "error."
  echo "$output"
  exit 1
}

echo "Running app..."
exec $GOPATH/bin/meetgrinder-api
