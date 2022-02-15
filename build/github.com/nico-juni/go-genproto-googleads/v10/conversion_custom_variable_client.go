// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package googleads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicespb "github.com/nico-juni/go-genproto-googleads/pb/v10/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newConversionCustomVariableClientHook clientHook

// ConversionCustomVariableCallOptions contains the retry settings for each method of ConversionCustomVariableClient.
type ConversionCustomVariableCallOptions struct {
	MutateConversionCustomVariables []gax.CallOption
}

func defaultConversionCustomVariableGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultConversionCustomVariableCallOptions() *ConversionCustomVariableCallOptions {
	return &ConversionCustomVariableCallOptions{
		MutateConversionCustomVariables: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalConversionCustomVariableClient is an interface that defines the methods availaible from Google Ads API.
type internalConversionCustomVariableClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateConversionCustomVariables(context.Context, *servicespb.MutateConversionCustomVariablesRequest, ...gax.CallOption) (*servicespb.MutateConversionCustomVariablesResponse, error)
}

// ConversionCustomVariableClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage conversion custom variables.
type ConversionCustomVariableClient struct {
	// The internal transport-dependent client.
	internalClient internalConversionCustomVariableClient

	// The call options for this service.
	CallOptions *ConversionCustomVariableCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ConversionCustomVariableClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ConversionCustomVariableClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ConversionCustomVariableClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateConversionCustomVariables creates or updates conversion custom variables. Operation statuses are
// returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ConversionCustomVariableError (at )
// DatabaseError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *ConversionCustomVariableClient) MutateConversionCustomVariables(ctx context.Context, req *servicespb.MutateConversionCustomVariablesRequest, opts ...gax.CallOption) (*servicespb.MutateConversionCustomVariablesResponse, error) {
	return c.internalClient.MutateConversionCustomVariables(ctx, req, opts...)
}

// conversionCustomVariableGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type conversionCustomVariableGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ConversionCustomVariableClient
	CallOptions **ConversionCustomVariableCallOptions

	// The gRPC API client.
	conversionCustomVariableClient servicespb.ConversionCustomVariableServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewConversionCustomVariableClient creates a new conversion custom variable service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage conversion custom variables.
func NewConversionCustomVariableClient(ctx context.Context, opts ...option.ClientOption) (*ConversionCustomVariableClient, error) {
	clientOpts := defaultConversionCustomVariableGRPCClientOptions()
	if newConversionCustomVariableClientHook != nil {
		hookOpts, err := newConversionCustomVariableClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ConversionCustomVariableClient{CallOptions: defaultConversionCustomVariableCallOptions()}

	c := &conversionCustomVariableGRPCClient{
		connPool:                       connPool,
		disableDeadlines:               disableDeadlines,
		conversionCustomVariableClient: servicespb.NewConversionCustomVariableServiceClient(connPool),
		CallOptions:                    &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *conversionCustomVariableGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *conversionCustomVariableGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *conversionCustomVariableGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *conversionCustomVariableGRPCClient) MutateConversionCustomVariables(ctx context.Context, req *servicespb.MutateConversionCustomVariablesRequest, opts ...gax.CallOption) (*servicespb.MutateConversionCustomVariablesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateConversionCustomVariables[0:len((*c.CallOptions).MutateConversionCustomVariables):len((*c.CallOptions).MutateConversionCustomVariables)], opts...)
	var resp *servicespb.MutateConversionCustomVariablesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversionCustomVariableClient.MutateConversionCustomVariables(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
