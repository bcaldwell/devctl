#!/bin/bash
set -e

export BUILD_DATE=`date +%Y-%m-%d:%H:%M:%S`
export BUILD_VERSION=`git log -1 --pretty=%B | tr " " "\n" | grep -Ei 'v[0-9]+(\.[0-9]+)*'`

if [[ $BUILD_VERSION ]] 
then
  echo "Installing gox and ghr"
  go get github.com/mitchellh/gox
  go get github.com/tcnksm/ghr
  echo "Building $BUILD_VERSION"
  gox -ldflags "-X main.Version=${BUILD_VERSION} -X main.BuildDate=${BUILD_DATE}" -parallel=2 -osarch="linux/amd64 darwin/amd64" -output "dist/devctl_{{.OS}}_{{.Arch}}_bin"
  
  echo "Creating release tarball"
  cd dist
  mkdir release
  mkdir devctl
  mv devctl_darwin_amd64_bin devctl/devctl
  cp ../devctl.sh devctl/
  tar -czvf release/devctl_darwin_amd64.tar.gz devctl

  mkdir devctl_linux_amd64
  mv devctl_linux_amd64_bin devctl/devctl
  cp ../devctl.sh devctl/
  tar -czvf release/devctl_linux_amd64.tar.gz devctl
  cd ..

  echo "Pushing to Github"
  ghr -t $GITHUB_TOKEN -u devctl -r devctl $BUILD_VERSION dist/release/

  # Generating md5 hash
  export DARWIN_AMD64_MD5=$(md5sum dist/release/devctl_darwin_amd64.tar.gz | cut -d ' ' -f 1)
  export LINUX_AMD64_MD5=$(md5sum dist/release/devctl_linux_amd64.tar.gz | cut -d ' ' -f 1)

  # Pushing new version and sha to website
  pip install pyyaml
  git clone git@github.com:devctl/devctl.github.io.git "$HOME/devctl.github.io"
  
  git config --global user.email "caldwellbenjamin8@gmail.com" 
  git config --global user.name "benjamin-bot" 

  python scripts/update_website_version.py

  cd "$HOME/devctl.github.io"
  git add . 
  git commit -m "Version update from circleci" 
  git diff
  git push origin master 
fi
