#!/bin/bash
set -e

curlsource() {
    f=$(mktemp -t curlsource-XXXX)
    curl -o "$f" -s -L "$1"
    source "$f"
    rm -f "$f"
}

# curlsource to use functions from devctl.sh
curlsource https://raw.githubusercontent.com/devctl/devctl/master/devctl.sh

logged_in_user="$(whoami)"
shell=$SHELL

_devctl_install() {
  if [ -f "/opt/devctl/devctl" ] && [ -f "/opt/devctl/devctl.sh" ]; then
    _devctl_echo_success "devctl already installed. Use devctl update to install latest version"
  else
    sudo mkdir -p "${install_location}/devctl"
    sudo chown "${logged_in_user}" "${install_location}/devctl"

    local latest_version=$(wget -qO- "${website}/dl/latest")

    _devctl_echo_info "Installing devctl ${latest_version}"

    _devctl_install_version "${latest_version}"

    _devctl_echo_success "Successfully installed devctl"
  fi

  case "${shell}" in
    */zsh)
      _devctl_setup_profile "$HOME/.zshrc"
      ;;
    *)
      _devctl_echo_fail "No :P"
      ;;
  esac
}

_devctl_setup_profile() {
  local rcfile
  rcfile=$1
  touch "${rcfile}"
  if grep -q /opt/devctl/devctl.sh "${rcfile}"; then
    _devctl_echo_success "shell already set up for dev"
    return
  fi

  echo -e "\n# added by devctl command\n[ -f /opt/devctl/devctl.sh ] && source /opt/devctl/devctl.sh" >> "${rcfile}"
  _devctl_echo_success "shell set up for devctl"
  _devctl_echo_info "added a line to the end of ${rcfile}"
}

_devctl_source() {
  _devctl_echo_info "Sourcing devctl"
  source /opt/devctl/devctl.sh
}

_devctl_install_dependencies() {
  _devctl_echo_info "Install dependencies"
  devctl setup
}

_devctl_install
_devctl_source
_devctl_install_dependencies
