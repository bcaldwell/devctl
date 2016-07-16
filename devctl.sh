#!/bin/bash
devctl_dir="$(dirname "$0:A")"
devctl() {
  fd="$(mktemp /tmp/devctl-fd-XXXXX)"

  rm -f "${fd}"

  go run "${devctl_dir}"/main.go "$@" 8>"${fd}"

  cat "${fd}"

  while read -r line
  do
    cd "${line}"
  done < "${fd}"

  rm -f "${fd}"
}
