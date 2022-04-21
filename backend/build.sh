#!/bin/sh

# shellcheck disable=SC2112
function main() {
  docker build -t learning:1.0 -f build/dockerfile/learning.dockerfile .
  docker save learning:1.0 -o build/images/learning.tar
  scp build/images/learning.tar guest@121.199.167.227:/home/guest/learning/images
}

main
