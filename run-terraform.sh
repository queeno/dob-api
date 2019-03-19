#!/usr/bin/env bash

set -e

function is_binary_old_or_absent(){
  [ ! -f dob-api ] && return 0
  RELEASE_MD5_URL="$(curl -s https://api.github.com/repos/queeno/dob-api/releases/latest | jq -r '.assets[] | select(.name == "dob-api.md5") | .browser_download_url')"
  RELEASE_MD5="$(curl -L -s $RELEASE_MD5_URL 2>/dev/null)"
  CURRENT_MD5="$(openssl md5 -hex dob-api  | awk '{ print $2 }')"
  echo "Remote release: $RELEASE_MD5"
  echo "Local release: $CURRENT_MD5"
  [ "$RELEASE_MD5" == "$CURRENT_MD5" ] && return 1
  return 0
}

if is_binary_old_or_absent; then
  echo "New dob-api release found! Deploying..."
  RELEASE_URL="$(curl -s https://api.github.com/repos/queeno/dob-api/releases/latest | jq -r '.assets[] | select(.name == "dob-api") | .browser_download_url')"
  rm -f dob-api
  curl -Lso dob-api $RELEASE_URL
else
  echo "No new release found."
fi

chmod +x dob-api

cd terraform

terraform init
terraform $*

# Cleanup on destroy
if [ "$1" == "destroy" ]; then
  [ -f dob-api ] && rm -f dob-api
fi
