package client

import (
	"context"
	"time"

	"github.com/Katkoin/katd/cmd/katwallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/Katkoin/katd/cmd/katwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the katwalletd server, and returns the client instance
func Connect(address string) (pb.KatwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("katwallet daemon is not running, start it with `katwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewKatwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
