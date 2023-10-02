package brpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (client *Client) connect() (*grpc.ClientConn, error) {
	creds := insecure.NewCredentials()
	if client.creds != nil {
		creds = client.creds
	}

	options := client.opts
	options = append(options, grpc.WithTransportCredentials(creds))

	// set no proxy
	if client.noProxy {
		options = append(options, grpc.WithNoProxy())
	}

	ctx, cancel := context.WithTimeout(context.Background(), client.timeout)
	defer cancel()

	connection, err := grpc.DialContext(ctx, client.host, options...)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
