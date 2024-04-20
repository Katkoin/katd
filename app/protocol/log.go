package protocol

import (
	"github.com/Katkoin/katd/infrastructure/logger"
	"github.com/Katkoin/katd/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
