package blockrelay

import (
	"github.com/Katkoin/katd/app/appmessage"
	peerpkg "github.com/katkoin/katdd/app/protocol/peer"
	"github.com/katkoin/katdd/domain"
	"github.com/katkoin/katdd/infrastructure/config"
	"github.com/katkoin/katdd/infrastructure/network/netadapter/router"
)

// SendVirtualSelectedParentInvContext is the interface for the context needed for the SendVirtualSelectedParentInv flow.
type SendVirtualSelectedParentInvContext interface {
	Domain() domain.Domain
	Config() *config.Config
}

// SendVirtualSelectedParentInv sends a peer the selected parent hash of the virtual
func SendVirtualSelectedParentInv(context SendVirtualSelectedParentInvContext,
	outgoingRoute *router.Route, peer *peerpkg.Peer) error {

	virtualSelectedParent, err := context.Domain().Consensus().GetVirtualSelectedParent()
	if err != nil {
		return err
	}

	if virtualSelectedParent.Equal(context.Config().NetParams().GenesisHash) {
		log.Debugf("Skipping sending the virtual selected parent hash to peer %s because it's the genesis", peer)
		return nil
	}

	log.Debugf("Sending virtual selected parent hash %s to peer %s", virtualSelectedParent, peer)

	virtualSelectedParentInv := appmessage.NewMsgInvBlock(virtualSelectedParent)
	return outgoingRoute.Enqueue(virtualSelectedParentInv)
}