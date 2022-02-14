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

var newCampaignBidModifierClientHook clientHook

// CampaignBidModifierCallOptions contains the retry settings for each method of CampaignBidModifierClient.
type CampaignBidModifierCallOptions struct {
	GetCampaignBidModifier     []gax.CallOption
	MutateCampaignBidModifiers []gax.CallOption
}

func defaultCampaignBidModifierGRPCClientOptions() []option.ClientOption {
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

func defaultCampaignBidModifierCallOptions() *CampaignBidModifierCallOptions {
	return &CampaignBidModifierCallOptions{
		GetCampaignBidModifier: []gax.CallOption{
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
		MutateCampaignBidModifiers: []gax.CallOption{
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

// internalCampaignBidModifierClient is an interface that defines the methods availaible from Google Ads API.
type internalCampaignBidModifierClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetCampaignBidModifier(context.Context, *servicespb.GetCampaignBidModifierRequest, ...gax.CallOption) (*resourcespb.CampaignBidModifier, error)
	MutateCampaignBidModifiers(context.Context, *servicespb.MutateCampaignBidModifiersRequest, ...gax.CallOption) (*servicespb.MutateCampaignBidModifiersResponse, error)
}

// CampaignBidModifierClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage campaign bid modifiers.
type CampaignBidModifierClient struct {
	// The internal transport-dependent client.
	internalClient internalCampaignBidModifierClient

	// The call options for this service.
	CallOptions *CampaignBidModifierCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CampaignBidModifierClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CampaignBidModifierClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CampaignBidModifierClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetCampaignBidModifier returns the requested campaign bid modifier in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignBidModifierClient) GetCampaignBidModifier(ctx context.Context, req *servicespb.GetCampaignBidModifierRequest, opts ...gax.CallOption) (*resourcespb.CampaignBidModifier, error) {
	return c.internalClient.GetCampaignBidModifier(ctx, req, opts...)
}

// MutateCampaignBidModifiers creates, updates, or removes campaign bid modifiers.
// Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ContextError (at )
// CriterionError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// FieldError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// MutateError (at )
// NewResourceCreationError (at )
// NotEmptyError (at )
// NullError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
func (c *CampaignBidModifierClient) MutateCampaignBidModifiers(ctx context.Context, req *servicespb.MutateCampaignBidModifiersRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignBidModifiersResponse, error) {
	return c.internalClient.MutateCampaignBidModifiers(ctx, req, opts...)
}

// campaignBidModifierGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type campaignBidModifierGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CampaignBidModifierClient
	CallOptions **CampaignBidModifierCallOptions

	// The gRPC API client.
	campaignBidModifierClient servicespb.CampaignBidModifierServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCampaignBidModifierClient creates a new campaign bid modifier service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage campaign bid modifiers.
func NewCampaignBidModifierClient(ctx context.Context, opts ...option.ClientOption) (*CampaignBidModifierClient, error) {
	clientOpts := defaultCampaignBidModifierGRPCClientOptions()
	if newCampaignBidModifierClientHook != nil {
		hookOpts, err := newCampaignBidModifierClientHook(ctx, clientHookParams{})
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
	client := CampaignBidModifierClient{CallOptions: defaultCampaignBidModifierCallOptions()}

	c := &campaignBidModifierGRPCClient{
		connPool:                  connPool,
		disableDeadlines:          disableDeadlines,
		campaignBidModifierClient: servicespb.NewCampaignBidModifierServiceClient(connPool),
		CallOptions:               &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *campaignBidModifierGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *campaignBidModifierGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *campaignBidModifierGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *campaignBidModifierGRPCClient) GetCampaignBidModifier(ctx context.Context, req *servicespb.GetCampaignBidModifierRequest, opts ...gax.CallOption) (*resourcespb.CampaignBidModifier, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCampaignBidModifier[0:len((*c.CallOptions).GetCampaignBidModifier):len((*c.CallOptions).GetCampaignBidModifier)], opts...)
	var resp *resourcespb.CampaignBidModifier
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignBidModifierClient.GetCampaignBidModifier(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *campaignBidModifierGRPCClient) MutateCampaignBidModifiers(ctx context.Context, req *servicespb.MutateCampaignBidModifiersRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignBidModifiersResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateCampaignBidModifiers[0:len((*c.CallOptions).MutateCampaignBidModifiers):len((*c.CallOptions).MutateCampaignBidModifiers)], opts...)
	var resp *servicespb.MutateCampaignBidModifiersResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignBidModifierClient.MutateCampaignBidModifiers(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
