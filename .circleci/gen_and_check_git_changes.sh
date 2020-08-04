#!/bin/bash

set -e

sudo apt-get install ragel

make gen

if [[ `git status --porcelain` ]]; then
  echo "Uncommited changes"
  exit 1
fi
