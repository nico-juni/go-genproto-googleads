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

var newFeedItemSetClientHook clientHook

// FeedItemSetCallOptions contains the retry settings for each method of FeedItemSetClient.
type FeedItemSetCallOptions struct {
	GetFeedItemSet     []gax.CallOption
	MutateFeedItemSets []gax.CallOption
}

func defaultFeedItemSetGRPCClientOptions() []option.ClientOption {
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

func defaultFeedItemSetCallOptions() *FeedItemSetCallOptions {
	return &FeedItemSetCallOptions{
		GetFeedItemSet: []gax.CallOption{
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
		MutateFeedItemSets: []gax.CallOption{
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

// internalFeedItemSetClient is an interface that defines the methods availaible from Google Ads API.
type internalFeedItemSetClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetFeedItemSet(context.Context, *servicespb.GetFeedItemSetRequest, ...gax.CallOption) (*resourcespb.FeedItemSet, error)
	MutateFeedItemSets(context.Context, *servicespb.MutateFeedItemSetsRequest, ...gax.CallOption) (*servicespb.MutateFeedItemSetsResponse, error)
}

// FeedItemSetClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage feed Item Set
type FeedItemSetClient struct {
	// The internal transport-dependent client.
	internalClient internalFeedItemSetClient

	// The call options for this service.
	CallOptions *FeedItemSetCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FeedItemSetClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FeedItemSetClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *FeedItemSetClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetFeedItemSet returns the requested feed item set in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *FeedItemSetClient) GetFeedItemSet(ctx context.Context, req *servicespb.GetFeedItemSetRequest, opts ...gax.CallOption) (*resourcespb.FeedItemSet, error) {
	return c.internalClient.GetFeedItemSet(ctx, req, opts...)
}

// MutateFeedItemSets creates, updates or removes feed item sets. Operation statuses are
// returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *FeedItemSetClient) MutateFeedItemSets(ctx context.Context, req *servicespb.MutateFeedItemSetsRequest, opts ...gax.CallOption) (*servicespb.MutateFeedItemSetsResponse, error) {
	return c.internalClient.MutateFeedItemSets(ctx, req, opts...)
}

// feedItemSetGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type feedItemSetGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing FeedItemSetClient
	CallOptions **FeedItemSetCallOptions

	// The gRPC API client.
	feedItemSetClient servicespb.FeedItemSetServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewFeedItemSetClient creates a new feed item set service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage feed Item Set
func NewFeedItemSetClient(ctx context.Context, opts ...option.ClientOption) (*FeedItemSetClient, error) {
	clientOpts := defaultFeedItemSetGRPCClientOptions()
	if newFeedItemSetClientHook != nil {
		hookOpts, err := newFeedItemSetClientHook(ctx, clientHookParams{})
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
	client := FeedItemSetClient{CallOptions: defaultFeedItemSetCallOptions()}

	c := &feedItemSetGRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		feedItemSetClient: servicespb.NewFeedItemSetServiceClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *feedItemSetGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *feedItemSetGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *feedItemSetGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *feedItemSetGRPCClient) GetFeedItemSet(ctx context.Context, req *servicespb.GetFeedItemSetRequest, opts ...gax.CallOption) (*resourcespb.FeedItemSet, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetFeedItemSet[0:len((*c.CallOptions).GetFeedItemSet):len((*c.CallOptions).GetFeedItemSet)], opts...)
	var resp *resourcespb.FeedItemSet
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.feedItemSetClient.GetFeedItemSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *feedItemSetGRPCClient) MutateFeedItemSets(ctx context.Context, req *servicespb.MutateFeedItemSetsRequest, opts ...gax.CallOption) (*servicespb.MutateFeedItemSetsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateFeedItemSets[0:len((*c.CallOptions).MutateFeedItemSets):len((*c.CallOptions).MutateFeedItemSets)], opts...)
	var resp *servicespb.MutateFeedItemSetsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.feedItemSetClient.MutateFeedItemSets(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
