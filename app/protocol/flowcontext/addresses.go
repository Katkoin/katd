package flowcontext

import (
	"github.com/katkoin/katd/infrastructure/network/addressmanager"
)

// AddressManager returns the address manager associated to the flow context.
func (f *FlowContext) AddressManager() *addressmanager.AddressManager {
	return f.addressManager
}