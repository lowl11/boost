package brpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	host     string
	creds    credentials.TransportCredentials
	opts     []grpc.DialOption
	timeout  time.Duration
	noProxy  bool
	sslCheck bool
}

func New(host string) *Client {
	return &Client{
		host:    host,
		opts:    make([]grpc.DialOption, 0),
		timeout: time.Second * 30,
	}
}

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
