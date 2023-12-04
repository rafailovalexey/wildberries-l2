#!/bin/bash

if [ "$#" -lt 1 ]; then
  echo "Usage: $0 <MOCKS_OUTPUT_DIRECTORY> <MOCKS_FILES>"
  exit 1
fi

MOCKS_OUTPUT_DIRECTORY="$1"
MOCKS_FILES="${*:2}"

for mock_file in $MOCKS_FILES; do

  DIRECTORY=$(dirname "$mock_file")
  FILENAME=$(basename "$mock_file")

  EXTENSION="${FILENAME##*.}"
  FILENAME_WITHOUT_EXTENSIONS="${FILENAME%.*}"

  OUTPUT_PATH="$DIRECTORY/$MOCKS_OUTPUT_DIRECTORY/${FILENAME_WITHOUT_EXTENSIONS}_mock.$EXTENSION"

  mkdir -p "$DIRECTORY/$MOCKS_OUTPUT_DIRECTORY"

  echo "Generating mock file for $mock_file"

  mockgen -source="$mock_file" -destination="$OUTPUT_PATH"

done
