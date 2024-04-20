#!/bin/bash
rm -rf /tmp/katd-temp

katd --devnet --appdir=/tmp/katd-temp --profile=6061 --loglevel=debug &
katd_PID=$!
katd_KILLED=0
function killkatdIfNotKilled() {
    if [ $katd_KILLED -eq 0 ]; then
      kill $katd_PID
    fi
}
trap "killkatdIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $katd_PID

wait $katd_PID
katd_KILLED=1
katd_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "katd exit code: $katd_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $katd_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
