// Copyright 2021 Google LLC
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
// source: google/ads/googleads/v9/resources/ad_parameter.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
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

// An ad parameter that is used to update numeric values (such as prices or
// inventory levels) in any text line of an ad (including URLs). There can
// be a maximum of two AdParameters per ad group criterion. (One with
// parameter_index = 1 and one with parameter_index = 2.)
// In the ad the parameters are referenced by a placeholder of the form
// "{param#:value}". E.g. "{param1:$17}"
type AdParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the ad parameter.
	// Ad parameter resource names have the form:
	//
	// `customers/{customer_id}/adParameters/{ad_group_id}~{criterion_id}~{parameter_index}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The ad group criterion that this ad parameter belongs to.
	AdGroupCriterion *string `protobuf:"bytes,5,opt,name=ad_group_criterion,json=adGroupCriterion,proto3,oneof" json:"ad_group_criterion,omitempty"`
	// Immutable. The unique index of this ad parameter. Must be either 1 or 2.
	ParameterIndex *int64 `protobuf:"varint,6,opt,name=parameter_index,json=parameterIndex,proto3,oneof" json:"parameter_index,omitempty"`
	// Numeric value to insert into the ad text. The following restrictions
	//  apply:
	//  - Can use comma or period as a separator, with an optional period or
	//    comma (respectively) for fractional values. For example, 1,000,000.00
	//    and 2.000.000,10 are valid.
	//  - Can be prepended or appended with a currency symbol. For example,
	//    $99.99 is valid.
	//  - Can be prepended or appended with a currency code. For example, 99.99USD
	//    and EUR200 are valid.
	//  - Can use '%'. For example, 1.0% and 1,0% are valid.
	//  - Can use plus or minus. For example, -10.99 and 25+ are valid.
	//  - Can use '/' between two numbers. For example 4/1 and 0.95/0.45 are
	//    valid.
	InsertionText *string `protobuf:"bytes,7,opt,name=insertion_text,json=insertionText,proto3,oneof" json:"insertion_text,omitempty"`
}

func (x *AdParameter) Reset() {
	*x = AdParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_resources_ad_parameter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdParameter) ProtoMessage() {}

func (x *AdParameter) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_resources_ad_parameter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdParameter.ProtoReflect.Descriptor instead.
func (*AdParameter) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDescGZIP(), []int{0}
}

func (x *AdParameter) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AdParameter) GetAdGroupCriterion() string {
	if x != nil && x.AdGroupCriterion != nil {
		return *x.AdGroupCriterion
	}
	return ""
}

func (x *AdParameter) GetParameterIndex() int64 {
	if x != nil && x.ParameterIndex != nil {
		return *x.ParameterIndex
	}
	return 0
}

func (x *AdParameter) GetInsertionText() string {
	if x != nil && x.InsertionText != nil {
		return *x.InsertionText
	}
	return ""
}

var File_google_ads_googleads_v9_resources_ad_parameter_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x03, 0x0a, 0x0b, 0x41, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x05, 0xfa,
	0x41, 0x26, 0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x64, 0x0a, 0x12, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x31, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2b, 0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x0f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x01, 0x52, 0x0e, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12,
	0x2a, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x3a, 0x7e, 0xea, 0x41, 0x7b,
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x53, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x61, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x61, 0x64,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x63, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x7d, 0x42, 0x15, 0x0a, 0x13, 0x5f,
	0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x6f, 0x6e, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x69, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x42, 0xfd, 0x01, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x42, 0x10, 0x41, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x39, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x39, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDescData = file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDesc
)

func file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDescData)
	})
	return file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDescData
}

var file_google_ads_googleads_v9_resources_ad_parameter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v9_resources_ad_parameter_proto_goTypes = []interface{}{
	(*AdParameter)(nil), // 0: google.ads.googleads.v9.resources.AdParameter
}
var file_google_ads_googleads_v9_resources_ad_parameter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v9_resources_ad_parameter_proto_init() }
func file_google_ads_googleads_v9_resources_ad_parameter_proto_init() {
	if File_google_ads_googleads_v9_resources_ad_parameter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v9_resources_ad_parameter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdParameter); i {
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
	file_google_ads_googleads_v9_resources_ad_parameter_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v9_resources_ad_parameter_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v9_resources_ad_parameter_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v9_resources_ad_parameter_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v9_resources_ad_parameter_proto = out.File
	file_google_ads_googleads_v9_resources_ad_parameter_proto_rawDesc = nil
	file_google_ads_googleads_v9_resources_ad_parameter_proto_goTypes = nil
	file_google_ads_googleads_v9_resources_ad_parameter_proto_depIdxs = nil
}
