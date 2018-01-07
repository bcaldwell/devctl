#!/bin/bash

_devctl_run_command=`type -a devctl | awk 'NF{ print $NF }' | tail -n 1`

if [ -n "$DEVCTL_RUN_DEV" ]; then
  local _devctl_path
  _devctl_path=`devctl --dryrun cd github.com/bcaldwell/devctl | awk 'NF{ print $NF }'`
  _devctl_run_command="go run ${_devctl_path}/main.go"
fi

devctl() {
  case "$1" in
    load-dev)
      export DEVCTL_RUN_DEV=true
      local devctl_path
      devctl_path=`devctl --dryrun cd github.com/bcaldwell/devctl | awk 'NF{ print $NF }'`
      # shellcheck disable=SC1090
      source "${devctl_path}/devctl.sh"
      _devctl_echo_info "Loaded dev devctl"
      return
      ;;
    load-system)
      unset DEVCTL_RUN_DEV
      # shellcheck disable=SC1091
      source "${HOME}/.devctl/devctl.sh"
      _devctl_echo_info "Loaded system devctl"
      return
      ;;
  esac

  fd="$(mktemp /tmp/devctl-fd-XXXXX)"
  rm -f "${fd}"

  eval "${_devctl_run_command} $*" 8>"${fd}"

  while builtin read -r line; do
    eval "${line}"
  done < "${fd}"

  rm -f "${fd}"
}

_devctl_check_error(){
  if [[ $1 -ne 0 ]]; then
    _devctl_echo_fail "$2 failed"
  else
    _devctl_echo_success "$2 successful"
  fi
}

NC='\x1b[0m'
GREEN='\x1b[32m'
RED='\x1b[31m'
BLUE='\x1b[94m'
YELLOW='\x1b[33m'

_devctl_echo_success() {
  echo -e "${GREEN}✔ ${NC} $1"
}

_devctl_echo_fail() {
    echo -e "${RED}✗ ${NC} $1"
}

_devctl_echo_info() {
    echo -e "${BLUE}🐧 ${NC} $1"
}

_devctl_echo_warning() {
    echo -e "${YELLOW}⚠ ${NC} $1"
}
