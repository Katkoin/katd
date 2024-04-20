#!/bin/bash
rm -rf /tmp/katd-temp

katd --devnet --appdir=/tmp/katd-temp --profile=6061 --loglevel=debug &
katd_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $katd_PID

wait $katd_PID
katd_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "katd exit code: $katd_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $katd_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
