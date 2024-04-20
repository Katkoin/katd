#!/bin/bash

APPDIR=/tmp/katd-temp
katd_RPC_PORT=29587

rm -rf "${APPDIR}"

katd --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${katd_RPC_PORT}" --profile=6061 &
katd_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${katd_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $katd_PID

wait $katd_PID
katd_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "katd exit code: $katd_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $katd_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
