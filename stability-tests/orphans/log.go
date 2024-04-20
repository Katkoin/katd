package main

import (
	"github.com/Katkoin/katd/infrastructure/logger"
	"github.com/Katkoin/katd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("ORPH")
	spawn      = panics.GoroutineWrapperFunc(log)
)
