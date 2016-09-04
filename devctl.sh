#!/bin/bash
devctl_dir="$(dirname "$0:A")"
binary_file="devctl.mac"

if [ -n "$BASH_SOURCE" ]; then
  devctl_dir="$(dirname "$BASH_SOURCE")"
fi

if [[ "$OSTYPE" == "linux-gnu" ]]; then
  binary_file="devctl.linux"
fi

run_command=""

if [ "${devctl_dir}" != "/opt/devctl" ]; then
    run_command="go run"
    binary_file="main.go"
fi

devctl() {
   case "$1" in
    load-dev)
      local devctl_path
      devctl_path="$(devctl cd devctl && pwd)"
      # shellcheck disable=SC1090
      source "${devctl_path}/devctl.sh"
      devctl_echo_info "Loaded dev devctl"
      return
      ;;
    load-system)
      # shellcheck disable=SC1091
      source "/opt/devctl/devctl.sh"
      devctl_echo_info "Loaded system devctl"
      return
      ;;
    update)
      if [ "${devctl_dir}" != "/opt/devctl" ]; then
        devctl_echo_fail "Refuse to update dev version. Run devctl load-system first"
      elif ! (git -C "${devctl_dir}" config remote.origin.url | grep -q github.com/devctl/devctl); then
        devctl_echo_fail "Version is not a github.com clone"
      else
        devctl_check_update
      fi
      return
      ;;
  esac

  fd="$(mktemp /tmp/devctl-fd-XXXXX)"

  rm -f "${fd}"

  eval "${run_command} ${devctl_dir}/${binary_file} $*" 8>"${fd}"

  while builtin read -r line; do
    eval "${line}"
  done < "${fd}"

  rm -f "${fd}"
}

devctl_check_update(){
  devctl_echo_info "checking for updates"
  {
    git -C /opt/devctl fetch
    git -C /opt/devctl reset --hard origin/master
  } >/dev/null 2>&1
  devctl_check_error $? "Update"
  # shellcheck disable=SC1091
  source "/opt/devctl/devctl.sh"
}

devctl_check_error(){
  if [[ $1 -ne 0 ]]; then
    devctl_echo_fail "$2 failed"
  else
    devctl_echo_success "$2 successful"
  fi
}

NC='\x1b[0m'
GREEN='\x1b[32m'
RED='\x1b[31m'
BLUE='\x1b[94m'
YELLOW='\x1b[33m'

devctl_echo_success() {
  echo -e "${GREEN}✔ ${NC} $1"
}

devctl_echo_fail() {
    echo -e "${RED}✗ ${NC} $1"
}

devctl_echo_info() {
    echo -e "${BLUE}🐧 ${NC} $1"
}

devctl_echo_warning() {
    echo -e "${YELLOW}⚠ ${NC} $1"
}
