// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/resources/customer_client.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
	enums "github.com/nico-juni/go-genproto-googleads/pb/v10/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A link between the given customer and a client customer. CustomerClients only
// exist for manager customers. All direct and indirect client customers are
// included, as well as the manager itself.
type CustomerClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the customer client.
	// CustomerClient resource names have the form:
	// `customers/{customer_id}/customerClients/{client_customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The resource name of the client-customer which is linked to
	// the given customer. Read only.
	ClientCustomer *string `protobuf:"bytes,12,opt,name=client_customer,json=clientCustomer,proto3,oneof" json:"client_customer,omitempty"`
	// Output only. Specifies whether this is a
	// [hidden account](https://support.google.com/google-ads/answer/7519830).
	// Read only.
	Hidden *bool `protobuf:"varint,13,opt,name=hidden,proto3,oneof" json:"hidden,omitempty"`
	// Output only. Distance between given customer and client. For self link, the level value
	// will be 0. Read only.
	Level *int64 `protobuf:"varint,14,opt,name=level,proto3,oneof" json:"level,omitempty"`
	// Output only. Common Locale Data Repository (CLDR) string representation of the
	// time zone of the client, e.g. America/Los_Angeles. Read only.
	TimeZone *string `protobuf:"bytes,15,opt,name=time_zone,json=timeZone,proto3,oneof" json:"time_zone,omitempty"`
	// Output only. Identifies if the client is a test account. Read only.
	TestAccount *bool `protobuf:"varint,16,opt,name=test_account,json=testAccount,proto3,oneof" json:"test_account,omitempty"`
	// Output only. Identifies if the client is a manager. Read only.
	Manager *bool `protobuf:"varint,17,opt,name=manager,proto3,oneof" json:"manager,omitempty"`
	// Output only. Descriptive name for the client. Read only.
	DescriptiveName *string `protobuf:"bytes,18,opt,name=descriptive_name,json=descriptiveName,proto3,oneof" json:"descriptive_name,omitempty"`
	// Output only. Currency code (e.g. 'USD', 'EUR') for the client. Read only.
	CurrencyCode *string `protobuf:"bytes,19,opt,name=currency_code,json=currencyCode,proto3,oneof" json:"currency_code,omitempty"`
	// Output only. The ID of the client customer. Read only.
	Id *int64 `protobuf:"varint,20,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Output only. The resource names of the labels owned by the requesting customer that are
	// applied to the client customer.
	// Label resource names have the form:
	//
	// `customers/{customer_id}/labels/{label_id}`
	AppliedLabels []string `protobuf:"bytes,21,rep,name=applied_labels,json=appliedLabels,proto3" json:"applied_labels,omitempty"`
	// Output only. The status of the client customer. Read only.
	Status enums.CustomerStatusEnum_CustomerStatus `protobuf:"varint,22,opt,name=status,proto3,enum=google.ads.googleads.v10.enums.CustomerStatusEnum_CustomerStatus" json:"status,omitempty"`
}

func (x *CustomerClient) Reset() {
	*x = CustomerClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_customer_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerClient) ProtoMessage() {}

func (x *CustomerClient) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_customer_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerClient.ProtoReflect.Descriptor instead.
func (*CustomerClient) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_customer_client_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerClient) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CustomerClient) GetClientCustomer() string {
	if x != nil && x.ClientCustomer != nil {
		return *x.ClientCustomer
	}
	return ""
}

func (x *CustomerClient) GetHidden() bool {
	if x != nil && x.Hidden != nil {
		return *x.Hidden
	}
	return false
}

func (x *CustomerClient) GetLevel() int64 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *CustomerClient) GetTimeZone() string {
	if x != nil && x.TimeZone != nil {
		return *x.TimeZone
	}
	return ""
}

func (x *CustomerClient) GetTestAccount() bool {
	if x != nil && x.TestAccount != nil {
		return *x.TestAccount
	}
	return false
}

func (x *CustomerClient) GetManager() bool {
	if x != nil && x.Manager != nil {
		return *x.Manager
	}
	return false
}

func (x *CustomerClient) GetDescriptiveName() string {
	if x != nil && x.DescriptiveName != nil {
		return *x.DescriptiveName
	}
	return ""
}

func (x *CustomerClient) GetCurrencyCode() string {
	if x != nil && x.CurrencyCode != nil {
		return *x.CurrencyCode
	}
	return ""
}

func (x *CustomerClient) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *CustomerClient) GetAppliedLabels() []string {
	if x != nil {
		return x.AppliedLabels
	}
	return nil
}

func (x *CustomerClient) GetStatus() enums.CustomerStatusEnum_CustomerStatus {
	if x != nil {
		return x.Status
	}
	return enums.CustomerStatusEnum_UNSPECIFIED
}

var File_google_ads_googleads_v10_resources_customer_client_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_resources_customer_client_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x34,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x07, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x54, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x29,
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x57, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x20, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x1e, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x25, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0c, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x05, 0x52, 0x07,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x10, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x06, 0x52, 0x0f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2d, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x07, 0x52, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x18,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48,
	0x08, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x4d, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x26, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x20, 0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x5e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x6a, 0xea, 0x41, 0x67, 0x0a, 0x27, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x7d, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x68, 0x69, 0x64, 0x64,
	0x65, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x85, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x42, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca,
	0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x30, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_resources_customer_client_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_resources_customer_client_proto_rawDescData = file_google_ads_googleads_v10_resources_customer_client_proto_rawDesc
)

func file_google_ads_googleads_v10_resources_customer_client_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_resources_customer_client_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_resources_customer_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_resources_customer_client_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_resources_customer_client_proto_rawDescData
}

var file_google_ads_googleads_v10_resources_customer_client_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_resources_customer_client_proto_goTypes = []interface{}{
	(*CustomerClient)(nil),                       // 0: google.ads.googleads.v10.resources.CustomerClient
	(enums.CustomerStatusEnum_CustomerStatus)(0), // 1: google.ads.googleads.v10.enums.CustomerStatusEnum.CustomerStatus
}
var file_google_ads_googleads_v10_resources_customer_client_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.resources.CustomerClient.status:type_name -> google.ads.googleads.v10.enums.CustomerStatusEnum.CustomerStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_resources_customer_client_proto_init() }
func file_google_ads_googleads_v10_resources_customer_client_proto_init() {
	if File_google_ads_googleads_v10_resources_customer_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_resources_customer_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerClient); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_ads_googleads_v10_resources_customer_client_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_resources_customer_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_resources_customer_client_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_resources_customer_client_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_resources_customer_client_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_resources_customer_client_proto = out.File
	file_google_ads_googleads_v10_resources_customer_client_proto_rawDesc = nil
	file_google_ads_googleads_v10_resources_customer_client_proto_goTypes = nil
	file_google_ads_googleads_v10_resources_customer_client_proto_depIdxs = nil
}
