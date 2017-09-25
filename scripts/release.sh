#!/bin/bash
set -e

export BUILD_DATE=`date +%Y-%m-%d:%H:%M:%S`
export BUILD_VERSION=`git log -1 --pretty=%B | tr " " "\n" | grep -Ei 'v[0-9]+(\.[0-9]+)*'`

export BUILDS="linux/amd64 darwin/amd64"

if [[ $BUILD_VERSION ]] 
then
  echo "Installing gox and ghr"
  go get github.com/mitchellh/gox
  go get github.com/tcnksm/ghr
  echo "Building $BUILD_VERSION"
  gox -ldflags "-X main.Version=${BUILD_VERSION} -X main.BuildDate=${BUILD_DATE}" -parallel=2 -osarch="${BUILDS}" -output "dist/devctl_{{.OS}}_{{.Arch}}_bin"

  # Pushing new version and sha to website
  pip install pyyaml --user
  git clone git@github.com:devctl/devctl.github.io.git "$HOME/devctl.github.io"
  
  git config --global user.email "caldwellbenjamin8@gmail.com" 
  git config --global user.name "benjamin-bot" 

  echo "Bundling and generating sha"
  python scripts/update_website_version.py

  echo "Pushing to Github"
  ghr -t $GITHUB_TOKEN -u benjamincaldwell -r devctl $BUILD_VERSION dist/release/


  cd "$HOME/devctl.github.io"
  git add . 
  git commit -m "Version update from circleci" 
  git diff
  git push origin master 
fi
