package rpchandlers

import (
	"github.com/Katkoin/katd/app/appmessage"
	"github.com/katkoin/katdd/app/rpc/rpccontext"
	"github.com/katkoin/katdd/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
