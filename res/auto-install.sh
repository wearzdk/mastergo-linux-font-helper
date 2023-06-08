#!/usr/bin/env bash

echo -e "\n\n"

arch=$(uname -m);
if [ "$arch" == "x86_64" ]; then
  arch="amd64";
elif [ "$arch" == "i386" ]; then
  arch="386";
elif [ "$arch" == "aarch64" ]; then
  arch="arm64";
else
  echo "不受支持的CPU架构: $arch";
  exit 1;
fi

get_latest_release() {
  curl -Ls --silent "https://github.com/zskzskabcd/mastergo-linux-font-helper/releases/latest" | perl -ne 'print "$1\n" if /v([0-9]{1,3}\.[0-9]{1,3})/' | head -1;
}

get_latest_release_link_download() {
  local latest=$(get_latest_release);
  echo "https://github.com/zskzskabcd/mastergo-linux-font-helper/releases/download/v${latest}/mastergo-font-linux-${arch}.tar.gz";
}

download() {
  local link=$(get_latest_release_link_download);
  mkdir /tmp/mastergo-font-linux-amd64-install;
  cd /tmp/mastergo-font-linux-amd64-install;
  rm -rf ./mastergo-font-linux*;
  wget "$link";
}

install() {
  local file="mastergo-font-linux-${arch}.tar.gz";
  cd /tmp/mastergo-font-linux-amd64-install;
  tar -xzvf "$file";
  bash install.sh;
}

download;
install;
