#!/usr/bin/env bash

set -e
prefix=""
if [ -n "$INPUT_PREFIX" ]; then
  prefix="--prefix=$INPUT_PREFIX"
fi

strict=""
if [ "$INPUT_STRICT" == true ]; then
  strict="--strict"
fi

if [ "$INPUT_OPERATION" == "sort" ] || [ "$INPUT_OPERATION" == "highest" ]; then
  set -x
  echo "result="$( /versiontool "$INPUT_OPERATION" $INPUT_FLAGS < "$INPUT_VERSION") >> $GITHUB_OUTPUT
  exit
fi
set -x
echo "result="$(/versiontool $prefix $strict "$INPUT_OPERATION" $INPUT_FLAGS "$INPUT_VERSION") >> $GITHUB_OUTPUT
