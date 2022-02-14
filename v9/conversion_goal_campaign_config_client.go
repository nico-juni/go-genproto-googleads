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
	servicespb "github.com/nico-juni/go-genproto-googleads/pb/v9/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newConversionGoalCampaignConfigClientHook clientHook

// ConversionGoalCampaignConfigCallOptions contains the retry settings for each method of ConversionGoalCampaignConfigClient.
type ConversionGoalCampaignConfigCallOptions struct {
	MutateConversionGoalCampaignConfigs []gax.CallOption
}

func defaultConversionGoalCampaignConfigGRPCClientOptions() []option.ClientOption {
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

func defaultConversionGoalCampaignConfigCallOptions() *ConversionGoalCampaignConfigCallOptions {
	return &ConversionGoalCampaignConfigCallOptions{
		MutateConversionGoalCampaignConfigs: []gax.CallOption{
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

// internalConversionGoalCampaignConfigClient is an interface that defines the methods availaible from Google Ads API.
type internalConversionGoalCampaignConfigClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateConversionGoalCampaignConfigs(context.Context, *servicespb.MutateConversionGoalCampaignConfigsRequest, ...gax.CallOption) (*servicespb.MutateConversionGoalCampaignConfigsResponse, error)
}

// ConversionGoalCampaignConfigClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage conversion goal campaign config.
type ConversionGoalCampaignConfigClient struct {
	// The internal transport-dependent client.
	internalClient internalConversionGoalCampaignConfigClient

	// The call options for this service.
	CallOptions *ConversionGoalCampaignConfigCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ConversionGoalCampaignConfigClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ConversionGoalCampaignConfigClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ConversionGoalCampaignConfigClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateConversionGoalCampaignConfigs creates, updates or removes conversion goal campaign config. Operation
// statuses are returned.
func (c *ConversionGoalCampaignConfigClient) MutateConversionGoalCampaignConfigs(ctx context.Context, req *servicespb.MutateConversionGoalCampaignConfigsRequest, opts ...gax.CallOption) (*servicespb.MutateConversionGoalCampaignConfigsResponse, error) {
	return c.internalClient.MutateConversionGoalCampaignConfigs(ctx, req, opts...)
}

// conversionGoalCampaignConfigGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type conversionGoalCampaignConfigGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ConversionGoalCampaignConfigClient
	CallOptions **ConversionGoalCampaignConfigCallOptions

	// The gRPC API client.
	conversionGoalCampaignConfigClient servicespb.ConversionGoalCampaignConfigServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewConversionGoalCampaignConfigClient creates a new conversion goal campaign config service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage conversion goal campaign config.
func NewConversionGoalCampaignConfigClient(ctx context.Context, opts ...option.ClientOption) (*ConversionGoalCampaignConfigClient, error) {
	clientOpts := defaultConversionGoalCampaignConfigGRPCClientOptions()
	if newConversionGoalCampaignConfigClientHook != nil {
		hookOpts, err := newConversionGoalCampaignConfigClientHook(ctx, clientHookParams{})
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
	client := ConversionGoalCampaignConfigClient{CallOptions: defaultConversionGoalCampaignConfigCallOptions()}

	c := &conversionGoalCampaignConfigGRPCClient{
		connPool:                           connPool,
		disableDeadlines:                   disableDeadlines,
		conversionGoalCampaignConfigClient: servicespb.NewConversionGoalCampaignConfigServiceClient(connPool),
		CallOptions:                        &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *conversionGoalCampaignConfigGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *conversionGoalCampaignConfigGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *conversionGoalCampaignConfigGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *conversionGoalCampaignConfigGRPCClient) MutateConversionGoalCampaignConfigs(ctx context.Context, req *servicespb.MutateConversionGoalCampaignConfigsRequest, opts ...gax.CallOption) (*servicespb.MutateConversionGoalCampaignConfigsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateConversionGoalCampaignConfigs[0:len((*c.CallOptions).MutateConversionGoalCampaignConfigs):len((*c.CallOptions).MutateConversionGoalCampaignConfigs)], opts...)
	var resp *servicespb.MutateConversionGoalCampaignConfigsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversionGoalCampaignConfigClient.MutateConversionGoalCampaignConfigs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
