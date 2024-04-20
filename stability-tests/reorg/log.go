package main

import (
	"github.com/Katkoin/katd/infrastructure/logger"
	"github.com/katkoin/katdd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RORG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
