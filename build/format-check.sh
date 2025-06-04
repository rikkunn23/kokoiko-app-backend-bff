#!/bin/sh

unformatted=$(gofmt -l .)
if [ -n "$unformatted" ]; then
  echo "Unformatted files detected:"
  echo "$unformatted"
  exit 1
fi
