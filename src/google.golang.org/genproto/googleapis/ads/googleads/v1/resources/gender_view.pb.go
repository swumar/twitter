// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/gender_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A gender view.
type GenderView struct {
	// Output only. The resource name of the gender view.
	// Gender view resource names have the form:
	//
	// `customers/{customer_id}/genderViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenderView) Reset()         { *m = GenderView{} }
func (m *GenderView) String() string { return proto.CompactTextString(m) }
func (*GenderView) ProtoMessage()    {}
func (*GenderView) Descriptor() ([]byte, []int) {
	return fileDescriptor_1caf9d1d919debe3, []int{0}
}

func (m *GenderView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenderView.Unmarshal(m, b)
}
func (m *GenderView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenderView.Marshal(b, m, deterministic)
}
func (m *GenderView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenderView.Merge(m, src)
}
func (m *GenderView) XXX_Size() int {
	return xxx_messageInfo_GenderView.Size(m)
}
func (m *GenderView) XXX_DiscardUnknown() {
	xxx_messageInfo_GenderView.DiscardUnknown(m)
}

var xxx_messageInfo_GenderView proto.InternalMessageInfo

func (m *GenderView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GenderView)(nil), "google.ads.googleads.v1.resources.GenderView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/gender_view.proto", fileDescriptor_1caf9d1d919debe3)
}

var fileDescriptor_1caf9d1d919debe3 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xe9, 0x06, 0x2f, 0xbc, 0x45, 0x11, 0x76, 0xd2, 0x21, 0xe8, 0x94, 0x81, 0x20, 0x24,
	0x94, 0xdd, 0xe2, 0x29, 0xbb, 0x0c, 0x3c, 0xc8, 0x18, 0x58, 0x44, 0x0a, 0x23, 0x6b, 0x62, 0x0d,
	0xac, 0xf9, 0x8f, 0xa4, 0xeb, 0x0e, 0x63, 0x5f, 0xc6, 0xa3, 0x9f, 0xc0, 0xcf, 0xe0, 0xa7, 0xf0,
	0xbc, 0x8f, 0x20, 0x1e, 0xa4, 0x4b, 0x93, 0xee, 0xa4, 0xde, 0x1e, 0xf8, 0xff, 0x9e, 0x27, 0x4f,
	0x9f, 0x86, 0x83, 0x0c, 0x20, 0x9b, 0x0b, 0xcc, 0xb8, 0xc1, 0x56, 0x56, 0xaa, 0x8c, 0xb0, 0x16,
	0x06, 0x96, 0x3a, 0x15, 0x06, 0x67, 0x42, 0x71, 0xa1, 0xa7, 0xa5, 0x14, 0x2b, 0xb4, 0xd0, 0x50,
	0x40, 0xa7, 0x67, 0x49, 0xc4, 0xb8, 0x41, 0xde, 0x84, 0xca, 0x08, 0x79, 0x53, 0xf7, 0xcc, 0xe5,
	0x2e, 0x24, 0x7e, 0x92, 0x62, 0xce, 0xa7, 0x33, 0xf1, 0xcc, 0x4a, 0x09, 0xda, 0x66, 0x74, 0x4f,
	0xf6, 0x00, 0x67, 0xab, 0x4f, 0xa7, 0x7b, 0x27, 0xa6, 0x14, 0x14, 0xac, 0x90, 0xa0, 0x8c, 0xbd,
	0x5e, 0xbc, 0x05, 0x61, 0x38, 0xda, 0x55, 0x8a, 0xa5, 0x58, 0x75, 0xc6, 0xe1, 0xa1, 0xb3, 0x4f,
	0x15, 0xcb, 0xc5, 0x71, 0x70, 0x1e, 0x5c, 0xfd, 0x1f, 0x5e, 0x7f, 0xd0, 0xf6, 0x27, 0xed, 0x87,
	0x97, 0x4d, 0xbf, 0x5a, 0x2d, 0xa4, 0x41, 0x29, 0xe4, 0xb8, 0xc9, 0x98, 0x1c, 0xb8, 0x84, 0x3b,
	0x96, 0x0b, 0xf2, 0xb0, 0xa5, 0xf7, 0x7f, 0xf2, 0x75, 0x50, 0xba, 0x34, 0x05, 0xe4, 0x42, 0x1b,
	0xbc, 0x76, 0x72, 0x53, 0xef, 0x55, 0x01, 0x06, 0xaf, 0xf7, 0xc6, 0xdb, 0x0c, 0xbf, 0x82, 0xb0,
	0x9f, 0x42, 0x8e, 0x7e, 0x9d, 0x6f, 0x78, 0xd4, 0xbc, 0x32, 0xae, 0xbe, 0x7a, 0x1c, 0x3c, 0xde,
	0xd6, 0xae, 0x0c, 0xe6, 0x4c, 0x65, 0x08, 0x74, 0x56, 0xbd, 0xb4, 0xdb, 0x04, 0x37, 0x15, 0x7f,
	0xf8, 0x91, 0x37, 0x5e, 0xbd, 0xb4, 0xda, 0x23, 0x4a, 0x5f, 0x5b, 0xbd, 0x91, 0x8d, 0xa4, 0xdc,
	0x20, 0x2b, 0x2b, 0x15, 0x47, 0x68, 0xe2, 0xc8, 0x77, 0xc7, 0x24, 0x94, 0x9b, 0xc4, 0x33, 0x49,
	0x1c, 0x25, 0x9e, 0xd9, 0xb6, 0xfa, 0xf6, 0x40, 0x08, 0xe5, 0x86, 0x10, 0x4f, 0x11, 0x12, 0x47,
	0x84, 0x78, 0x6e, 0xf6, 0x6f, 0x57, 0x76, 0xf0, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x9c, 0x9c,
	0x7c, 0x74, 0x02, 0x00, 0x00,
}