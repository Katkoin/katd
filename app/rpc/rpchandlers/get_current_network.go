package rpchandlers

import (
	"github.com/Katkoin/katd/app/appmessage"
	"github.com/Katkoin/katd/app/rpc/rpccontext"
	"github.com/Katkoin/katd/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
