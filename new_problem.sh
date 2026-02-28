#!/usr/bin/env bash

set -e

if [ -z "$1" ]; then
  echo "Usage: ./new_problem.sh <problem_name>"
  echo "Example: ./new_problem.sh remove_element"
  exit 1
fi

RAW_NAME="$1"
BASE_DIR="Problems"
PROBLEM_DIR="$BASE_DIR/$RAW_NAME"

if [ -d "$PROBLEM_DIR" ]; then
  echo "❌ Problem '$RAW_NAME' already exists"
  exit 1
fi

# Convert snake_case → lowerCamelCase
toLowerCamel() {
  local input="$1"
  local output=""
  local capitalize=false

  for (( i=0; i<${#input}; i++ )); do
    char="${input:$i:1}"
    if [ "$char" = "_" ]; then
      capitalize=true
    else
      if $capitalize; then
        output+="${char^^}"
        capitalize=false
      else
        output+="$char"
      fi
    fi
  done

  echo "$output"
}

# Convert snake_case → UpperCamelCase
toUpperCamel() {
  local lc
  lc=$(toLowerCamel "$1")
  echo "${lc^}"
}

FUNC_NAME=$(toLowerCamel "$RAW_NAME")
TEST_NAME="Test$(toUpperCamel "$RAW_NAME")"

echo "Creating problem: $RAW_NAME"
echo "Function name: $FUNC_NAME"
echo "Test name: $TEST_NAME"

mkdir -p "$PROBLEM_DIR"

# Go source
cat <<EOF > "$PROBLEM_DIR/$RAW_NAME.go"
package $RAW_NAME

func $FUNC_NAME() {
	// TODO: implement
}
EOF

# Go test
cat <<EOF > "$PROBLEM_DIR/${RAW_NAME}_test.go"
package $RAW_NAME

import "testing"

func $TEST_NAME(t *testing.T) {
	t.Fatal("not implemented")
}
EOF

# Markdown
cat <<EOF > "$PROBLEM_DIR/question.md"
# $RAW_NAME

## Problem
<!-- Paste problem description here -->

## Notes
<!-- Edge cases, thoughts -->
EOF

echo "✅ Problem '$RAW_NAME' created"