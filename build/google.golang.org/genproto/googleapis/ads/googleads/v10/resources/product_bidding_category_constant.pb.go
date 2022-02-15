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
// source: google/ads/googleads/v10/resources/product_bidding_category_constant.proto

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

// A Product Bidding Category.
type ProductBiddingCategoryConstant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the product bidding category.
	// Product bidding category resource names have the form:
	//
	// `productBiddingCategoryConstants/{country_code}~{level}~{id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. ID of the product bidding category.
	//
	// This ID is equivalent to the google_product_category ID as described in
	// this article: https://support.google.com/merchants/answer/6324436.
	Id *int64 `protobuf:"varint,10,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Output only. Two-letter upper-case country code of the product bidding category.
	CountryCode *string `protobuf:"bytes,11,opt,name=country_code,json=countryCode,proto3,oneof" json:"country_code,omitempty"`
	// Output only. Resource name of the parent product bidding category.
	ProductBiddingCategoryConstantParent *string `protobuf:"bytes,12,opt,name=product_bidding_category_constant_parent,json=productBiddingCategoryConstantParent,proto3,oneof" json:"product_bidding_category_constant_parent,omitempty"`
	// Output only. Level of the product bidding category.
	Level enums.ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel `protobuf:"varint,5,opt,name=level,proto3,enum=google.ads.googleads.v10.enums.ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel" json:"level,omitempty"`
	// Output only. Status of the product bidding category.
	Status enums.ProductBiddingCategoryStatusEnum_ProductBiddingCategoryStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v10.enums.ProductBiddingCategoryStatusEnum_ProductBiddingCategoryStatus" json:"status,omitempty"`
	// Output only. Language code of the product bidding category.
	LanguageCode *string `protobuf:"bytes,13,opt,name=language_code,json=languageCode,proto3,oneof" json:"language_code,omitempty"`
	// Output only. Display value of the product bidding category localized according to
	// language_code.
	LocalizedName *string `protobuf:"bytes,14,opt,name=localized_name,json=localizedName,proto3,oneof" json:"localized_name,omitempty"`
}

func (x *ProductBiddingCategoryConstant) Reset() {
	*x = ProductBiddingCategoryConstant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductBiddingCategoryConstant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductBiddingCategoryConstant) ProtoMessage() {}

func (x *ProductBiddingCategoryConstant) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductBiddingCategoryConstant.ProtoReflect.Descriptor instead.
func (*ProductBiddingCategoryConstant) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDescGZIP(), []int{0}
}

func (x *ProductBiddingCategoryConstant) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ProductBiddingCategoryConstant) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ProductBiddingCategoryConstant) GetCountryCode() string {
	if x != nil && x.CountryCode != nil {
		return *x.CountryCode
	}
	return ""
}

func (x *ProductBiddingCategoryConstant) GetProductBiddingCategoryConstantParent() string {
	if x != nil && x.ProductBiddingCategoryConstantParent != nil {
		return *x.ProductBiddingCategoryConstantParent
	}
	return ""
}

func (x *ProductBiddingCategoryConstant) GetLevel() enums.ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel {
	if x != nil {
		return x.Level
	}
	return enums.ProductBiddingCategoryLevelEnum_UNSPECIFIED
}

func (x *ProductBiddingCategoryConstant) GetStatus() enums.ProductBiddingCategoryStatusEnum_ProductBiddingCategoryStatus {
	if x != nil {
		return x.Status
	}
	return enums.ProductBiddingCategoryStatusEnum_UNSPECIFIED
}

func (x *ProductBiddingCategoryConstant) GetLanguageCode() string {
	if x != nil && x.LanguageCode != nil {
		return *x.LanguageCode
	}
	return ""
}

func (x *ProductBiddingCategoryConstant) GetLocalizedName() string {
	if x != nil && x.LocalizedName != nil {
		return *x.LocalizedName
	}
	return ""
}

var File_google_ads_googleads_v10_resources_product_bidding_category_constant_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x62, 0x69, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x1a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x62, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x07, 0x0a, 0x1e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x64, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x3f, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x39, 0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x9c, 0x01, 0x0a, 0x28, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3f, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x39,
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x48, 0x02, 0x52, 0x24, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x76, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x5b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x69, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x69, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x7a, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x79, 0xea, 0x41, 0x76, 0x0a, 0x37, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x7d, 0x7e, 0x7b, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x7d, 0x7e, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2b, 0x0a, 0x29, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x95, 0x02,
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x23, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x31, 0x30, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDescData = file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDesc
)

func file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDescData
}

var file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_goTypes = []interface{}{
	(*ProductBiddingCategoryConstant)(nil),                                   // 0: google.ads.googleads.v10.resources.ProductBiddingCategoryConstant
	(enums.ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel)(0),   // 1: google.ads.googleads.v10.enums.ProductBiddingCategoryLevelEnum.ProductBiddingCategoryLevel
	(enums.ProductBiddingCategoryStatusEnum_ProductBiddingCategoryStatus)(0), // 2: google.ads.googleads.v10.enums.ProductBiddingCategoryStatusEnum.ProductBiddingCategoryStatus
}
var file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.resources.ProductBiddingCategoryConstant.level:type_name -> google.ads.googleads.v10.enums.ProductBiddingCategoryLevelEnum.ProductBiddingCategoryLevel
	2, // 1: google.ads.googleads.v10.resources.ProductBiddingCategoryConstant.status:type_name -> google.ads.googleads.v10.enums.ProductBiddingCategoryStatusEnum.ProductBiddingCategoryStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_init() }
func file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_init() {
	if File_google_ads_googleads_v10_resources_product_bidding_category_constant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductBiddingCategoryConstant); i {
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
	file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_resources_product_bidding_category_constant_proto = out.File
	file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_rawDesc = nil
	file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_goTypes = nil
	file_google_ads_googleads_v10_resources_product_bidding_category_constant_proto_depIdxs = nil
}
