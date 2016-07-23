#!/bin/bash
set -e

logged_in_user="$(whoami)"
shell=$SHELL


clone_devctl() {
  if [[ -d "/opt/devctl/.git" ]]; then
    echo_success "already have devctl"
  else
    local git_url
    git_url="https://github.com/devctl/devctl.git"

    # Very intentionally do the `git clone` as the logged in user so ssh keys aren't an issue.
    sudo mkdir -p /opt/devctl
    sudo chown "${logged_in_user}" /opt/devctl
    echo_info "Cloning devctl/devctl into /opt/devctl"
    git -C /opt/devctl clone "${git_url}" .
    if [[ $? -ne 0 ]]; then
      echo_fail "Git clone failed"
      exit 1
    fi

    echo_success "cloned devctl/devctl"
  fi

  case "${shell}" in
    */zsh)
      # Pretty much every zsh user just uses ~/.zshrc so we won't worry about
      # all that file detection stuff we do with bash.
      install_shell_shim "$HOME/.zshrc"
      ;;
    *)
      echo_fail "No :P"
      ;;
  esac
}

install_shell_shim() {
  local rcfile
  rcfile=$1
  touch "${rcfile}"
  if grep -q /opt/devctl/devctl.sh "${rcfile}"; then
    echo_success "shell already set up for dev"
    return
  fi

  echo -e "\n# added by devctl command\n[ -f /opt/devctl/devctl.sh ] && source /opt/devctl/devctl.sh" >> "${rcfile}"
  echo_success "shell set up for devctl"
  echo_info "added a line to the end of ${rcfile}"
}

NC='\x1b[0m'
GREEN='\x1b[32m'
RED='\x1b[31m'
BLUE='\x1b[94m'
YELLOW='\x1b[33m'

echo_success() {
  echo -e "${GREEN}✔ ${NC} $1"
}

echo_fail() {
    echo -e "${RED}✗ ${NC} $1"
}

echo_info() {
    echo -e "${BLUE}🐧 ${NC} $1"
}

echo_warning() {
    echo -e "${YELLOW}⚠ ${NC} $1"
}

clone_devctl