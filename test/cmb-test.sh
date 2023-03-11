#!/usr/bin/env bash
#
# E2E test for wechat provider.

# set -x # debug
set -eo errexit

TEST_DIR=`dirname "$(realpath $0)"`
ROOT_DIR="$TEST_DIR/.."

make -f "$ROOT_DIR/Makefile" build
mkdir -p "$ROOT_DIR/test/output"

# generate wechat bills output in beancount format
"$ROOT_DIR/bin/double-entry-generator" translate \
    --provider cmb \
    --config "$ROOT_DIR/example/cmb/config.yaml" \
    --output "$ROOT_DIR/test/output/test-cmb-output.beancount" \
    "$ROOT_DIR/example/cmb/CMB.csv"

#diff -u --color \
#    "$ROOT_DIR/example/wechat/example-wechat-output.beancount" \
#    "$ROOT_DIR/test/output/test-wechat-output.beancount"

if [ $? -ne 0 ]; then
    echo "[FAIL] CMB provider output is different from expected output."
    exit 1
fi

echo "[PASS] All CMB provider tests!"
