#!/usr/bin/env bash

set -e

curl -L -o dob-api.new https://github.com/queeno/dob-api/releases/download/latest/dob-api &>/dev/null

DOB_API_MD5="$(md5 -q dob-api || true)"
NEW_DOB_API_MD5="$(md5 -q dob-api.new || true)"

if [ ! "$DOB_API_MD5" == "$NEW_DOB_API_MD5" ]; then
  echo "New dob-api release found! Deploying..."
  rm -f dob-api
  mv dob-api.new dob-api
else
  echo "No new release found."
  rm -f dob-api.new
fi

cd terraform

terraform init
terraform $*

# Cleanup on destroy
if [ "$1" == "destroy" ]; then
  [ -f dob-api ] && rm -f dob-api
fi
