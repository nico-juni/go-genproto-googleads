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

var newBiddingSeasonalityAdjustmentClientHook clientHook

// BiddingSeasonalityAdjustmentCallOptions contains the retry settings for each method of BiddingSeasonalityAdjustmentClient.
type BiddingSeasonalityAdjustmentCallOptions struct {
	GetBiddingSeasonalityAdjustment     []gax.CallOption
	MutateBiddingSeasonalityAdjustments []gax.CallOption
}

func defaultBiddingSeasonalityAdjustmentGRPCClientOptions() []option.ClientOption {
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

func defaultBiddingSeasonalityAdjustmentCallOptions() *BiddingSeasonalityAdjustmentCallOptions {
	return &BiddingSeasonalityAdjustmentCallOptions{
		GetBiddingSeasonalityAdjustment: []gax.CallOption{
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
		MutateBiddingSeasonalityAdjustments: []gax.CallOption{
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

// internalBiddingSeasonalityAdjustmentClient is an interface that defines the methods availaible from Google Ads API.
type internalBiddingSeasonalityAdjustmentClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetBiddingSeasonalityAdjustment(context.Context, *servicespb.GetBiddingSeasonalityAdjustmentRequest, ...gax.CallOption) (*resourcespb.BiddingSeasonalityAdjustment, error)
	MutateBiddingSeasonalityAdjustments(context.Context, *servicespb.MutateBiddingSeasonalityAdjustmentsRequest, ...gax.CallOption) (*servicespb.MutateBiddingSeasonalityAdjustmentsResponse, error)
}

// BiddingSeasonalityAdjustmentClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage bidding seasonality adjustments.
type BiddingSeasonalityAdjustmentClient struct {
	// The internal transport-dependent client.
	internalClient internalBiddingSeasonalityAdjustmentClient

	// The call options for this service.
	CallOptions *BiddingSeasonalityAdjustmentCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BiddingSeasonalityAdjustmentClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BiddingSeasonalityAdjustmentClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *BiddingSeasonalityAdjustmentClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetBiddingSeasonalityAdjustment returns the requested seasonality adjustment in full detail.
func (c *BiddingSeasonalityAdjustmentClient) GetBiddingSeasonalityAdjustment(ctx context.Context, req *servicespb.GetBiddingSeasonalityAdjustmentRequest, opts ...gax.CallOption) (*resourcespb.BiddingSeasonalityAdjustment, error) {
	return c.internalClient.GetBiddingSeasonalityAdjustment(ctx, req, opts...)
}

// MutateBiddingSeasonalityAdjustments creates, updates, or removes seasonality adjustments.
// Operation statuses are returned.
func (c *BiddingSeasonalityAdjustmentClient) MutateBiddingSeasonalityAdjustments(ctx context.Context, req *servicespb.MutateBiddingSeasonalityAdjustmentsRequest, opts ...gax.CallOption) (*servicespb.MutateBiddingSeasonalityAdjustmentsResponse, error) {
	return c.internalClient.MutateBiddingSeasonalityAdjustments(ctx, req, opts...)
}

// biddingSeasonalityAdjustmentGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type biddingSeasonalityAdjustmentGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing BiddingSeasonalityAdjustmentClient
	CallOptions **BiddingSeasonalityAdjustmentCallOptions

	// The gRPC API client.
	biddingSeasonalityAdjustmentClient servicespb.BiddingSeasonalityAdjustmentServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBiddingSeasonalityAdjustmentClient creates a new bidding seasonality adjustment service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage bidding seasonality adjustments.
func NewBiddingSeasonalityAdjustmentClient(ctx context.Context, opts ...option.ClientOption) (*BiddingSeasonalityAdjustmentClient, error) {
	clientOpts := defaultBiddingSeasonalityAdjustmentGRPCClientOptions()
	if newBiddingSeasonalityAdjustmentClientHook != nil {
		hookOpts, err := newBiddingSeasonalityAdjustmentClientHook(ctx, clientHookParams{})
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
	client := BiddingSeasonalityAdjustmentClient{CallOptions: defaultBiddingSeasonalityAdjustmentCallOptions()}

	c := &biddingSeasonalityAdjustmentGRPCClient{
		connPool:                           connPool,
		disableDeadlines:                   disableDeadlines,
		biddingSeasonalityAdjustmentClient: servicespb.NewBiddingSeasonalityAdjustmentServiceClient(connPool),
		CallOptions:                        &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *biddingSeasonalityAdjustmentGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *biddingSeasonalityAdjustmentGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *biddingSeasonalityAdjustmentGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *biddingSeasonalityAdjustmentGRPCClient) GetBiddingSeasonalityAdjustment(ctx context.Context, req *servicespb.GetBiddingSeasonalityAdjustmentRequest, opts ...gax.CallOption) (*resourcespb.BiddingSeasonalityAdjustment, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetBiddingSeasonalityAdjustment[0:len((*c.CallOptions).GetBiddingSeasonalityAdjustment):len((*c.CallOptions).GetBiddingSeasonalityAdjustment)], opts...)
	var resp *resourcespb.BiddingSeasonalityAdjustment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.biddingSeasonalityAdjustmentClient.GetBiddingSeasonalityAdjustment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *biddingSeasonalityAdjustmentGRPCClient) MutateBiddingSeasonalityAdjustments(ctx context.Context, req *servicespb.MutateBiddingSeasonalityAdjustmentsRequest, opts ...gax.CallOption) (*servicespb.MutateBiddingSeasonalityAdjustmentsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateBiddingSeasonalityAdjustments[0:len((*c.CallOptions).MutateBiddingSeasonalityAdjustments):len((*c.CallOptions).MutateBiddingSeasonalityAdjustments)], opts...)
	var resp *servicespb.MutateBiddingSeasonalityAdjustmentsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.biddingSeasonalityAdjustmentClient.MutateBiddingSeasonalityAdjustments(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
