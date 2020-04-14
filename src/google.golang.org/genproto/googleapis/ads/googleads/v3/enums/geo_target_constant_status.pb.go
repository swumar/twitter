// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/geo_target_constant_status.proto

package enums

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

// The possible statuses of a geo target constant.
type GeoTargetConstantStatusEnum_GeoTargetConstantStatus int32

const (
	// No value has been specified.
	GeoTargetConstantStatusEnum_UNSPECIFIED GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	GeoTargetConstantStatusEnum_UNKNOWN GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 1
	// The geo target constant is valid.
	GeoTargetConstantStatusEnum_ENABLED GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 2
	// The geo target constant is obsolete and will be removed.
	GeoTargetConstantStatusEnum_REMOVAL_PLANNED GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 3
)

var GeoTargetConstantStatusEnum_GeoTargetConstantStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVAL_PLANNED",
}

var GeoTargetConstantStatusEnum_GeoTargetConstantStatus_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"ENABLED":         2,
	"REMOVAL_PLANNED": 3,
}

func (x GeoTargetConstantStatusEnum_GeoTargetConstantStatus) String() string {
	return proto.EnumName(GeoTargetConstantStatusEnum_GeoTargetConstantStatus_name, int32(x))
}

func (GeoTargetConstantStatusEnum_GeoTargetConstantStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c37ed64745a799b7, []int{0, 0}
}

// Container for describing the status of a geo target constant.
type GeoTargetConstantStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoTargetConstantStatusEnum) Reset()         { *m = GeoTargetConstantStatusEnum{} }
func (m *GeoTargetConstantStatusEnum) String() string { return proto.CompactTextString(m) }
func (*GeoTargetConstantStatusEnum) ProtoMessage()    {}
func (*GeoTargetConstantStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c37ed64745a799b7, []int{0}
}

func (m *GeoTargetConstantStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoTargetConstantStatusEnum.Unmarshal(m, b)
}
func (m *GeoTargetConstantStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoTargetConstantStatusEnum.Marshal(b, m, deterministic)
}
func (m *GeoTargetConstantStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoTargetConstantStatusEnum.Merge(m, src)
}
func (m *GeoTargetConstantStatusEnum) XXX_Size() int {
	return xxx_messageInfo_GeoTargetConstantStatusEnum.Size(m)
}
func (m *GeoTargetConstantStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoTargetConstantStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GeoTargetConstantStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.GeoTargetConstantStatusEnum_GeoTargetConstantStatus", GeoTargetConstantStatusEnum_GeoTargetConstantStatus_name, GeoTargetConstantStatusEnum_GeoTargetConstantStatus_value)
	proto.RegisterType((*GeoTargetConstantStatusEnum)(nil), "google.ads.googleads.v3.enums.GeoTargetConstantStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/geo_target_constant_status.proto", fileDescriptor_c37ed64745a799b7)
}

var fileDescriptor_c37ed64745a799b7 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x6b, 0xc2, 0x30,
	0x18, 0x9d, 0x15, 0x36, 0x88, 0x07, 0x4b, 0x77, 0x18, 0x6c, 0x7a, 0xd0, 0x1f, 0x90, 0x1e, 0x7a,
	0xcb, 0x60, 0x90, 0x6a, 0x26, 0x32, 0x17, 0xcb, 0x9c, 0x1d, 0x1b, 0x85, 0x92, 0xd9, 0x10, 0x04,
	0xcd, 0x27, 0x26, 0xca, 0x7e, 0xcf, 0x8e, 0xfb, 0x29, 0xfb, 0x29, 0xbb, 0xed, 0x1f, 0x8c, 0xa6,
	0xea, 0xad, 0xbb, 0x84, 0x97, 0xbc, 0xef, 0xbd, 0xbc, 0xef, 0xa1, 0x3b, 0x05, 0xa0, 0x56, 0x32,
	0x14, 0x85, 0x09, 0x2b, 0x58, 0xa2, 0x7d, 0x14, 0x4a, 0xbd, 0x5b, 0x9b, 0x50, 0x49, 0xc8, 0xad,
	0xd8, 0x2a, 0x69, 0xf3, 0x05, 0x68, 0x63, 0x85, 0xb6, 0xb9, 0xb1, 0xc2, 0xee, 0x0c, 0xde, 0x6c,
	0xc1, 0x42, 0xd0, 0xad, 0x44, 0x58, 0x14, 0x06, 0x9f, 0xf4, 0x78, 0x1f, 0x61, 0xa7, 0xbf, 0xee,
	0x1c, 0xed, 0x37, 0xcb, 0x50, 0x68, 0x0d, 0x56, 0xd8, 0x25, 0xe8, 0x83, 0xb8, 0xff, 0x81, 0x6e,
	0x46, 0x12, 0x9e, 0x9d, 0xff, 0xe0, 0x60, 0x3f, 0x73, 0xee, 0x4c, 0xef, 0xd6, 0xfd, 0x57, 0x74,
	0x55, 0x43, 0x07, 0x6d, 0xd4, 0x9a, 0xf3, 0x59, 0xc2, 0x06, 0xe3, 0xfb, 0x31, 0x1b, 0xfa, 0x67,
	0x41, 0x0b, 0x5d, 0xcc, 0xf9, 0x03, 0x9f, 0xbe, 0x70, 0xbf, 0x51, 0x5e, 0x18, 0xa7, 0xf1, 0x84,
	0x0d, 0x7d, 0x2f, 0xb8, 0x44, 0xed, 0x27, 0xf6, 0x38, 0x4d, 0xe9, 0x24, 0x4f, 0x26, 0x94, 0x73,
	0x36, 0xf4, 0x9b, 0xf1, 0x6f, 0x03, 0xf5, 0x16, 0xb0, 0xc6, 0xff, 0xa6, 0x8f, 0x3b, 0x35, 0xdf,
	0x27, 0x65, 0xfa, 0xa4, 0xf1, 0x16, 0x1f, 0xe4, 0x0a, 0x56, 0x42, 0x2b, 0x0c, 0x5b, 0x15, 0x2a,
	0xa9, 0xdd, 0x6e, 0xc7, 0x32, 0x37, 0x4b, 0x53, 0xd3, 0xed, 0xad, 0x3b, 0x3f, 0xbd, 0xe6, 0x88,
	0xd2, 0x2f, 0xaf, 0x3b, 0xaa, 0xac, 0x68, 0x61, 0x70, 0x05, 0x4b, 0x94, 0x46, 0xb8, 0x6c, 0xc2,
	0x7c, 0x1f, 0xf9, 0x8c, 0x16, 0x26, 0x3b, 0xf1, 0x59, 0x1a, 0x65, 0x8e, 0xff, 0xf1, 0x7a, 0xd5,
	0x23, 0x21, 0xb4, 0x30, 0x84, 0x9c, 0x26, 0x08, 0x49, 0x23, 0x42, 0xdc, 0xcc, 0xfb, 0xb9, 0x0b,
	0x16, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x6d, 0x83, 0x46, 0xf3, 0x01, 0x00, 0x00,
}