package brpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

func (client *Client) Credentials(creds credentials.TransportCredentials) *Client {
	client.creds = creds
	return client
}

func (client *Client) Options(dialOptions ...grpc.DialOption) *Client {
	client.opts = append(client.opts, dialOptions...)
	return client
}

func (client *Client) Timeout(duration time.Duration) *Client {
	client.timeout = duration
	return client
}

func (client *Client) NoProxy() *Client {
	client.noProxy = true
	return client
}

func (client *Client) SslTrust() *Client {
	client.sslCheck = true
	return client
}

func (client *Client) Connection() (*grpc.ClientConn, error) {
	connection, err := client.connect()
	if err != nil {
		return nil, err
	}

	return connection, nil
}
