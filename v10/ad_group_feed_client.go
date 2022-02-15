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

var newAdGroupFeedClientHook clientHook

// AdGroupFeedCallOptions contains the retry settings for each method of AdGroupFeedClient.
type AdGroupFeedCallOptions struct {
	MutateAdGroupFeeds []gax.CallOption
}

func defaultAdGroupFeedGRPCClientOptions() []option.ClientOption {
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

func defaultAdGroupFeedCallOptions() *AdGroupFeedCallOptions {
	return &AdGroupFeedCallOptions{
		MutateAdGroupFeeds: []gax.CallOption{
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

// internalAdGroupFeedClient is an interface that defines the methods availaible from Google Ads API.
type internalAdGroupFeedClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateAdGroupFeeds(context.Context, *servicespb.MutateAdGroupFeedsRequest, ...gax.CallOption) (*servicespb.MutateAdGroupFeedsResponse, error)
}

// AdGroupFeedClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage ad group feeds.
type AdGroupFeedClient struct {
	// The internal transport-dependent client.
	internalClient internalAdGroupFeedClient

	// The call options for this service.
	CallOptions *AdGroupFeedCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupFeedClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupFeedClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AdGroupFeedClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateAdGroupFeeds creates, updates, or removes ad group feeds. Operation statuses are
// returned.
//
// List of thrown errors:
// AdGroupFeedError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// DatabaseError (at )
// DistinctError (at )
// FieldError (at )
// FunctionError (at )
// FunctionParsingError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// MutateError (at )
// NotEmptyError (at )
// NullError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
func (c *AdGroupFeedClient) MutateAdGroupFeeds(ctx context.Context, req *servicespb.MutateAdGroupFeedsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupFeedsResponse, error) {
	return c.internalClient.MutateAdGroupFeeds(ctx, req, opts...)
}

// adGroupFeedGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGroupFeedGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AdGroupFeedClient
	CallOptions **AdGroupFeedCallOptions

	// The gRPC API client.
	adGroupFeedClient servicespb.AdGroupFeedServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAdGroupFeedClient creates a new ad group feed service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage ad group feeds.
func NewAdGroupFeedClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupFeedClient, error) {
	clientOpts := defaultAdGroupFeedGRPCClientOptions()
	if newAdGroupFeedClientHook != nil {
		hookOpts, err := newAdGroupFeedClientHook(ctx, clientHookParams{})
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
	client := AdGroupFeedClient{CallOptions: defaultAdGroupFeedCallOptions()}

	c := &adGroupFeedGRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		adGroupFeedClient: servicespb.NewAdGroupFeedServiceClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *adGroupFeedGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGroupFeedGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGroupFeedGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGroupFeedGRPCClient) MutateAdGroupFeeds(ctx context.Context, req *servicespb.MutateAdGroupFeedsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupFeedsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateAdGroupFeeds[0:len((*c.CallOptions).MutateAdGroupFeeds):len((*c.CallOptions).MutateAdGroupFeeds)], opts...)
	var resp *servicespb.MutateAdGroupFeedsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupFeedClient.MutateAdGroupFeeds(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
