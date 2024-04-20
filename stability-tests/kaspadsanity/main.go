package main

import (
	"fmt"
	"os"

	"github.com/Katkoin/katd/stability-tests/common"
	"github.com/katkoin/katdd/util/profiling"

	"github.com/katkoin/katdd/util/panics"
	"github.com/pkg/errors"
)

func main() {
	defer panics.HandlePanic(log, "katdsanity-main", nil)
	err := parseConfig()
	if err != nil {
		panic(errors.Wrap(err, "error in parseConfig"))
	}
	defer backendLog.Close()
	common.UseLogger(backendLog, log.Level())

	cfg := activeConfig()
	if cfg.Profile != "" {
		profiling.Start(cfg.Profile, log)
	}

	argsChan := readArgs()
	failures, err := commandLoop(argsChan)
	if err != nil {
		panic(errors.Wrap(err, "error in commandLoop"))
	}

	if len(failures) > 0 {
		fmt.Fprintf(os.Stderr, "FAILED:\n")
		for _, failure := range failures {
			fmt.Fprintln(os.Stderr, failure)
		}
		backendLog.Close()
		os.Exit(1)
	}

	log.Infof("All tests have passed")
}
