package consensus

import (
	"github.com/Katkoin/katd/infrastructure/logger"
	"github.com/katkoin/katdd/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
