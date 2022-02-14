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

var newAdGroupAdLabelClientHook clientHook

// AdGroupAdLabelCallOptions contains the retry settings for each method of AdGroupAdLabelClient.
type AdGroupAdLabelCallOptions struct {
	GetAdGroupAdLabel     []gax.CallOption
	MutateAdGroupAdLabels []gax.CallOption
}

func defaultAdGroupAdLabelGRPCClientOptions() []option.ClientOption {
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

func defaultAdGroupAdLabelCallOptions() *AdGroupAdLabelCallOptions {
	return &AdGroupAdLabelCallOptions{
		GetAdGroupAdLabel: []gax.CallOption{
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
		MutateAdGroupAdLabels: []gax.CallOption{
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

// internalAdGroupAdLabelClient is an interface that defines the methods availaible from Google Ads API.
type internalAdGroupAdLabelClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetAdGroupAdLabel(context.Context, *servicespb.GetAdGroupAdLabelRequest, ...gax.CallOption) (*resourcespb.AdGroupAdLabel, error)
	MutateAdGroupAdLabels(context.Context, *servicespb.MutateAdGroupAdLabelsRequest, ...gax.CallOption) (*servicespb.MutateAdGroupAdLabelsResponse, error)
}

// AdGroupAdLabelClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage labels on ad group ads.
type AdGroupAdLabelClient struct {
	// The internal transport-dependent client.
	internalClient internalAdGroupAdLabelClient

	// The call options for this service.
	CallOptions *AdGroupAdLabelCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupAdLabelClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupAdLabelClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AdGroupAdLabelClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetAdGroupAdLabel returns the requested ad group ad label in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *AdGroupAdLabelClient) GetAdGroupAdLabel(ctx context.Context, req *servicespb.GetAdGroupAdLabelRequest, opts ...gax.CallOption) (*resourcespb.AdGroupAdLabel, error) {
	return c.internalClient.GetAdGroupAdLabel(ctx, req, opts...)
}

// MutateAdGroupAdLabels creates and removes ad group ad labels.
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
// NewResourceCreationError (at )
// QuotaError (at )
// RequestError (at )
func (c *AdGroupAdLabelClient) MutateAdGroupAdLabels(ctx context.Context, req *servicespb.MutateAdGroupAdLabelsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupAdLabelsResponse, error) {
	return c.internalClient.MutateAdGroupAdLabels(ctx, req, opts...)
}

// adGroupAdLabelGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGroupAdLabelGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AdGroupAdLabelClient
	CallOptions **AdGroupAdLabelCallOptions

	// The gRPC API client.
	adGroupAdLabelClient servicespb.AdGroupAdLabelServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAdGroupAdLabelClient creates a new ad group ad label service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage labels on ad group ads.
func NewAdGroupAdLabelClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupAdLabelClient, error) {
	clientOpts := defaultAdGroupAdLabelGRPCClientOptions()
	if newAdGroupAdLabelClientHook != nil {
		hookOpts, err := newAdGroupAdLabelClientHook(ctx, clientHookParams{})
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
	client := AdGroupAdLabelClient{CallOptions: defaultAdGroupAdLabelCallOptions()}

	c := &adGroupAdLabelGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		adGroupAdLabelClient: servicespb.NewAdGroupAdLabelServiceClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *adGroupAdLabelGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGroupAdLabelGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGroupAdLabelGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGroupAdLabelGRPCClient) GetAdGroupAdLabel(ctx context.Context, req *servicespb.GetAdGroupAdLabelRequest, opts ...gax.CallOption) (*resourcespb.AdGroupAdLabel, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAdGroupAdLabel[0:len((*c.CallOptions).GetAdGroupAdLabel):len((*c.CallOptions).GetAdGroupAdLabel)], opts...)
	var resp *resourcespb.AdGroupAdLabel
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupAdLabelClient.GetAdGroupAdLabel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *adGroupAdLabelGRPCClient) MutateAdGroupAdLabels(ctx context.Context, req *servicespb.MutateAdGroupAdLabelsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupAdLabelsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateAdGroupAdLabels[0:len((*c.CallOptions).MutateAdGroupAdLabels):len((*c.CallOptions).MutateAdGroupAdLabels)], opts...)
	var resp *servicespb.MutateAdGroupAdLabelsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupAdLabelClient.MutateAdGroupAdLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
