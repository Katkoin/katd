package prefixmanager

import (
	"github.com/Katkoin/katd/infrastructure/logger"
	"github.com/Katkoin/katd/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
