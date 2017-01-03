#!/bin/bash

website="https://devctl.github.io"
install_location="/opt"

devctl_dir="$(dirname "$0:A")"
binary_file="devctl"

if [ -n "$BASH_SOURCE" ]; then
  devctl_dir="$(dirname "$BASH_SOURCE")"
fi

#if [[ "$OSTYPE" == "linux-gnu" ]]; then
#  binary_file="devctl"
#fi

run_command=""

if [ "${devctl_dir}" != "/opt/devctl" ]; then
    run_command="go run"
    binary_file="main.go"
fi

devctl() {
  # local _devctl_verbose=1
  # while getopts 'abf:v' flag; do
  #   case "${flag}" in
  #     v) _devctl_verbose=0 ;;
  #   esac
  # done

  case "$1" in
    load-dev)
      local devctl_path
      devctl_path="$(devctl cd github.com/benjamincaldwell/devctl && pwd)"
      # shellcheck disable=SC1090
      source "${devctl_path}/devctl.sh"
      _devctl_echo_info "Loaded dev devctl"
      return
      ;;
    load-system)
      # shellcheck disable=SC1091
      source "/opt/devctl/devctl.sh"
      _devctl_echo_info "Loaded system devctl"
      return
      ;;
    update)
      if [ "${devctl_dir}" != "/opt/devctl" ]; then
        _devctl_echo_fail "Refuse to update dev version. Run devctl load-system first"
      else
        _devctl_check_update
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

_devctl_check_update(){
  _devctl_echo_info "checking for updates"
  local latest_version=$(wget -qO- "${website}/dl/latest")
  local installed=$(devctl version | cut -c 2-)
  if _devctl_check_version "${installed}" "${latest_version}"
  then
    _devctl_echo_success "Already up to date"
  else
    _devctl_echo_info "Downloading update"
    {
      _devctl_install_version "${latest_version}"
    } >/dev/null 2>&1
    _devctl_check_error $? "Update"
    # shellcheck disable=SC1091
    source "/opt/devctl/devctl.sh"
  fi
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


_devctl_install_version() {
  {
    local system_name=$(_devctl_system_detector)
    local tar_file_name="/tmp/devctl.tar.gz"
    wget "https://github.com/devctl/devctl/releases/download/v${1}/devctl_${system_name}.tar.gz" -O "${tar_file_name}"
    
    local remote_hash=$(wget -qO- "${website}/dl/sha/${system_name}" | grep "${1}" | perl -e 'if (<> =~ /$1:([a-fA-F\d]{32})/g) {print "$1"} else {print <>}')
    if _devctl_verify_hash "${tar_file_name}" "${remote_hash}"
    then
      if [[ -d "${install_location}/devctl" ]] && [ "$(ls -A ${install_location}/devctl)" ]; then
        /bin/rm -r "${install_location}"/devctl/*
      fi
      tar -zxvf "${tar_file_name}" -C "${install_location}" --keep-newer-files
      chmod +x "${install_location}/devctl/devctl"
      /bin/rm -r "${tar_file_name}"
    else
      return
    fi
  }
  return $?
}

_devctl_verify_hash() {
  #0 is true and 1 is false
  local hash=$(openssl dgst -sha256 "${1}" | cut -d ' ' -f 2)
  if [ "${hash}" == "${2}" ]
  then
    return 0
  fi
  return 1
}

_devctl_system_detector() {
  local os=$(uname -s | tr '[:upper:]' '[:lower:]')
  local arch
  case $(uname -m) in
    x86_64)
      arch="amd64"
      ;;
    *"arm"*)
      arch="arm"
      ;;
    *)
      arch="unsupported"
      ;;
  esac
  echo "${os}_${arch}"
}

_devctl_check_version() {
    local installed=${1} latest=${2}
    local winner=$(echo -e "${installed}\n${latest}" | sed '/^$/d' | sort -nr | head -1)
    [[ "$winner" = "$installed" ]] && return 0
    return 1
}
