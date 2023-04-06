#!/usr/bin/env bash

set -eu

count=$(echo $(( $(grep -E "\[Exercise\]|</summary>" -c README.md ))))

echo "There are $count questions and exercises"

sed -i "s/currently \*\*[0-9]*\*\*/currently \*\*$count\\**/" README.md
