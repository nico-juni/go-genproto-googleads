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
	resourcespb "github.com/nico-juni/go-genproto-googleads/pb/v9/resources"
	servicespb "github.com/nico-juni/go-genproto-googleads/pb/v9/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newCustomerLabelClientHook clientHook

// CustomerLabelCallOptions contains the retry settings for each method of CustomerLabelClient.
type CustomerLabelCallOptions struct {
	GetCustomerLabel     []gax.CallOption
	MutateCustomerLabels []gax.CallOption
}

func defaultCustomerLabelGRPCClientOptions() []option.ClientOption {
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

func defaultCustomerLabelCallOptions() *CustomerLabelCallOptions {
	return &CustomerLabelCallOptions{
		GetCustomerLabel: []gax.CallOption{
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
		MutateCustomerLabels: []gax.CallOption{
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

// internalCustomerLabelClient is an interface that defines the methods availaible from Google Ads API.
type internalCustomerLabelClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetCustomerLabel(context.Context, *servicespb.GetCustomerLabelRequest, ...gax.CallOption) (*resourcespb.CustomerLabel, error)
	MutateCustomerLabels(context.Context, *servicespb.MutateCustomerLabelsRequest, ...gax.CallOption) (*servicespb.MutateCustomerLabelsResponse, error)
}

// CustomerLabelClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage labels on customers.
type CustomerLabelClient struct {
	// The internal transport-dependent client.
	internalClient internalCustomerLabelClient

	// The call options for this service.
	CallOptions *CustomerLabelCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CustomerLabelClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CustomerLabelClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CustomerLabelClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetCustomerLabel returns the requested customer-label relationship in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CustomerLabelClient) GetCustomerLabel(ctx context.Context, req *servicespb.GetCustomerLabelRequest, opts ...gax.CallOption) (*resourcespb.CustomerLabel, error) {
	return c.internalClient.GetCustomerLabel(ctx, req, opts...)
}

// MutateCustomerLabels creates and removes customer-label relationships.
// Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// HeaderError (at )
// InternalError (at )
// LabelError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *CustomerLabelClient) MutateCustomerLabels(ctx context.Context, req *servicespb.MutateCustomerLabelsRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerLabelsResponse, error) {
	return c.internalClient.MutateCustomerLabels(ctx, req, opts...)
}

// customerLabelGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type customerLabelGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CustomerLabelClient
	CallOptions **CustomerLabelCallOptions

	// The gRPC API client.
	customerLabelClient servicespb.CustomerLabelServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCustomerLabelClient creates a new customer label service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage labels on customers.
func NewCustomerLabelClient(ctx context.Context, opts ...option.ClientOption) (*CustomerLabelClient, error) {
	clientOpts := defaultCustomerLabelGRPCClientOptions()
	if newCustomerLabelClientHook != nil {
		hookOpts, err := newCustomerLabelClientHook(ctx, clientHookParams{})
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
	client := CustomerLabelClient{CallOptions: defaultCustomerLabelCallOptions()}

	c := &customerLabelGRPCClient{
		connPool:            connPool,
		disableDeadlines:    disableDeadlines,
		customerLabelClient: servicespb.NewCustomerLabelServiceClient(connPool),
		CallOptions:         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *customerLabelGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *customerLabelGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *customerLabelGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *customerLabelGRPCClient) GetCustomerLabel(ctx context.Context, req *servicespb.GetCustomerLabelRequest, opts ...gax.CallOption) (*resourcespb.CustomerLabel, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCustomerLabel[0:len((*c.CallOptions).GetCustomerLabel):len((*c.CallOptions).GetCustomerLabel)], opts...)
	var resp *resourcespb.CustomerLabel
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customerLabelClient.GetCustomerLabel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *customerLabelGRPCClient) MutateCustomerLabels(ctx context.Context, req *servicespb.MutateCustomerLabelsRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerLabelsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateCustomerLabels[0:len((*c.CallOptions).MutateCustomerLabels):len((*c.CallOptions).MutateCustomerLabels)], opts...)
	var resp *servicespb.MutateCustomerLabelsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customerLabelClient.MutateCustomerLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
