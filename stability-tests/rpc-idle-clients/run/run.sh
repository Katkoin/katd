#!/bin/bash
rm -rf /tmp/katd-temp

NUM_CLIENTS=128
katd --devnet --appdir=/tmp/katd-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
katd_PID=$!
katd_KILLED=0
function killkatdIfNotKilled() {
  if [ $katd_KILLED -eq 0 ]; then
    kill $katd_PID
  fi
}
trap "killkatdIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $katd_PID

wait $katd_PID
katd_EXIT_CODE=$?
katd_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "katd exit code: $katd_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $katd_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
