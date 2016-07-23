#!/bin/bash
set -e

abort()
{
    echo "An error occurred. Exiting..." >&2
    exit 1
}

# http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5
env GOOS=darwin GOARCH=amd64 go build -o devctl.mac || exit 1
env GOOS=linux GOARCH=amd64 go build -o devctl.linux || exit 1
eval "$(ssh-agent -s)" 
chmod 600 .travis/deploy_key 
ssh-add .travis/deploy_key 
git clone git@github.com:devctl/devctl.git "$HOME/deploy"
rm -rf "$HOME/deploy/deploy" "$HOME/deploy/devctl.sh"
mv devctl.* "$HOME/deploy/" 
mv install.sh "$HOME/deploy/" 
cd "$HOME/deploy" 
git config --global user.email "travis@devctl.com" 
git config --global user.name "travis" 
git add . 
git commit -m "Auto update from travis" 
git push origin master 
