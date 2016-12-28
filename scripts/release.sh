#!/bin/bash
set -e

BUILD_DATE=`date '+%A %W %Y %X'`
BUILD_VERSION=`git log -1 --pretty=%B | tr " " "\n" | grep -Ei 'v[0-9]+(\.[0-9]+)*'`

if [[ $BUILD_VERSION ]] 
then
  echo "Installing gox and ghr"
  go get github.com/mitchellh/gox
  go get github.com/tcnksm/ghr
  echo "Building $BUILD_VERSION"
  gox -parallel=2 -osarch="linux/amd64 darwin/amd64" -output "dist/ncd_{{.OS}}_{{.Arch}}"
  echo "Pushing to Github"
  ghr -t $GITHUB_TOKEN -u $CIRCLE_PROJECT_USERNAME -r $CIRCLE_PROJECT_REPONAME $BUILD_VERSION dist/
fi