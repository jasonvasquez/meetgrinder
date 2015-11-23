#!/bin/bash

readonly SRC_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
readonly PACKAGES="./..."

echo -n "go test..."
output=$(go test $PACKAGES 2>&1) && echo "done." || {
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
echo "Launching app at http://$(docker-machine ip default):8081"
docker build -t meetgrinder-api .
exec docker run -p 8081:8081 --link=meetgrinder-db:db --name meetgrinder-api --rm meetgrinder-api
