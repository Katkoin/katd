package handshake

import (
	"github.com/Katkoin/katd/infrastructure/logger"
	"github.com/katkoin/katdd/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
