#!/bin/bash

readonly SRC_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo -n "go get......."
go get ./... > /dev/null 2>&1 && echo "done." || {
  echo "error."
  exit 1
}

echo -n "go install..."
go install ./... > /dev/null 2>&1 && echo "done." || {
  echo "error."
  exit 1
}

echo "Running app..."
$GOPATH/bin/meetgrinder
