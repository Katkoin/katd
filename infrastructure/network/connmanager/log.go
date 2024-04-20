package connmanager

import (
	"github.com/Katkoin/katd/infrastructure/logger"
	"github.com/Katkoin/katd/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
