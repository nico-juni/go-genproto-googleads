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

var newCampaignCriterionClientHook clientHook

// CampaignCriterionCallOptions contains the retry settings for each method of CampaignCriterionClient.
type CampaignCriterionCallOptions struct {
	MutateCampaignCriteria []gax.CallOption
}

func defaultCampaignCriterionGRPCClientOptions() []option.ClientOption {
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

func defaultCampaignCriterionCallOptions() *CampaignCriterionCallOptions {
	return &CampaignCriterionCallOptions{
		MutateCampaignCriteria: []gax.CallOption{
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

// internalCampaignCriterionClient is an interface that defines the methods availaible from Google Ads API.
type internalCampaignCriterionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCampaignCriteria(context.Context, *servicespb.MutateCampaignCriteriaRequest, ...gax.CallOption) (*servicespb.MutateCampaignCriteriaResponse, error)
}

// CampaignCriterionClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage campaign criteria.
type CampaignCriterionClient struct {
	// The internal transport-dependent client.
	internalClient internalCampaignCriterionClient

	// The call options for this service.
	CallOptions *CampaignCriterionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CampaignCriterionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CampaignCriterionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CampaignCriterionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCampaignCriteria creates, updates, or removes criteria. Operation statuses are returned.
//
// List of thrown errors:
// AdxError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// CampaignCriterionError (at )
// CollectionSizeError (at )
// ContextError (at )
// CriterionError (at )
// DatabaseError (at )
// DistinctError (at )
// FieldError (at )
// FieldMaskError (at )
// FunctionError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// MutateError (at )
// NewResourceCreationError (at )
// NotEmptyError (at )
// NullError (at )
// OperationAccessDeniedError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RegionCodeError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
func (c *CampaignCriterionClient) MutateCampaignCriteria(ctx context.Context, req *servicespb.MutateCampaignCriteriaRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignCriteriaResponse, error) {
	return c.internalClient.MutateCampaignCriteria(ctx, req, opts...)
}

// campaignCriterionGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type campaignCriterionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CampaignCriterionClient
	CallOptions **CampaignCriterionCallOptions

	// The gRPC API client.
	campaignCriterionClient servicespb.CampaignCriterionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCampaignCriterionClient creates a new campaign criterion service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage campaign criteria.
func NewCampaignCriterionClient(ctx context.Context, opts ...option.ClientOption) (*CampaignCriterionClient, error) {
	clientOpts := defaultCampaignCriterionGRPCClientOptions()
	if newCampaignCriterionClientHook != nil {
		hookOpts, err := newCampaignCriterionClientHook(ctx, clientHookParams{})
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
	client := CampaignCriterionClient{CallOptions: defaultCampaignCriterionCallOptions()}

	c := &campaignCriterionGRPCClient{
		connPool:                connPool,
		disableDeadlines:        disableDeadlines,
		campaignCriterionClient: servicespb.NewCampaignCriterionServiceClient(connPool),
		CallOptions:             &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *campaignCriterionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *campaignCriterionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *campaignCriterionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *campaignCriterionGRPCClient) MutateCampaignCriteria(ctx context.Context, req *servicespb.MutateCampaignCriteriaRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignCriteriaResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateCampaignCriteria[0:len((*c.CallOptions).MutateCampaignCriteria):len((*c.CallOptions).MutateCampaignCriteria)], opts...)
	var resp *servicespb.MutateCampaignCriteriaResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignCriterionClient.MutateCampaignCriteria(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
