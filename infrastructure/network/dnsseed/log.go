// Copyright (c) 2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dnsseed

import (
	"github.com/Katkoin/katd/infrastructure/logger"
	"github.com/Katkoin/katd/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
