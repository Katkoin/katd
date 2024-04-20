package main

import (
	"github.com/Katkoin/katd/infrastructure/logger"
	"github.com/katkoin/katdd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
