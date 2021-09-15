#!/usr/bin/env bash

prefix=""
if [ -n "$INPUT_prefix" ]; then
  prefix="--prefix=$INPUT_prefix"
fi

strict=""
if [ "$INPUT_strict" == true ]; then
  strict="--strict"
fi

if [ "$INPUT_operation" == "sort" ] || [ "$INPUT_operation" == "highest" ]; then
  set -x
  echo "::set-output name=result::"$( echo "$INPUT_version" | ./versiontool "$INPUT_operation" $INPUT_flags)
  exit $?
fi
set -x
echo "::set-output name=result::"$(./versiontool $prefix $strict "$INPUT_operation" $INPUT_flags "$INPUT_version")
exit $?
