#!/bin/bash
rm -rf /tmp/katd-temp

katd --simnet --appdir=/tmp/katd-temp --profile=6061 &
katd_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $katd_PID

wait $katd_PID
katd_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "katd exit code: $katd_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $katd_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
