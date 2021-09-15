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
  echo "::set-output name=result::"$( /versiontool "$INPUT_OPERATION" $INPUT_FLAGS < "$INPUT_VERSION")
fi
set -x
echo "::set-output name=result::"$(/versiontool $prefix $strict "$INPUT_OPERATION" $INPUT_FLAGS "$INPUT_VERSION")
