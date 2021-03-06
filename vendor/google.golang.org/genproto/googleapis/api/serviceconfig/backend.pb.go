// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/backend.proto
// DO NOT EDIT!

package google_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `Backend` defines the backend configuration for a service.
type Backend struct {
	// A list of API backend rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*BackendRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *Backend) Reset()                    { *m = Backend{} }
func (m *Backend) String() string            { return proto.CompactTextString(m) }
func (*Backend) ProtoMessage()               {}
func (*Backend) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Backend) GetRules() []*BackendRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A backend rule provides configuration for an individual API element.
type BackendRule struct {
	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// The address of the API backend.
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	// The number of seconds to wait for a response from a request.  The
	// default depends on the deployment context.
	Deadline float64 `protobuf:"fixed64,3,opt,name=deadline" json:"deadline,omitempty"`
}

func (m *BackendRule) Reset()                    { *m = BackendRule{} }
func (m *BackendRule) String() string            { return proto.CompactTextString(m) }
func (*BackendRule) ProtoMessage()               {}
func (*BackendRule) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterType((*Backend)(nil), "google.api.Backend")
	proto.RegisterType((*BackendRule)(nil), "google.api.BackendRule")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/backend.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0xcf, 0x4e, 0x03, 0x21,
	0x10, 0xc6, 0x43, 0xab, 0x56, 0xa7, 0xc6, 0x03, 0x17, 0x89, 0x27, 0xd3, 0x8b, 0xbd, 0x08, 0x89,
	0x5e, 0xbc, 0xba, 0x89, 0x31, 0xde, 0x36, 0xbc, 0x80, 0xa1, 0x30, 0x12, 0x22, 0x32, 0x0d, 0x54,
	0x1f, 0xc8, 0x27, 0x95, 0xfd, 0xe3, 0x76, 0x2f, 0x24, 0x1f, 0xbf, 0x1f, 0xc3, 0x7c, 0xf0, 0xe2,
	0x89, 0x7c, 0x44, 0xe9, 0x29, 0x9a, 0xe4, 0x25, 0x65, 0xaf, 0x3c, 0xa6, 0x7d, 0xa6, 0x03, 0xa9,
	0x01, 0x99, 0x7d, 0x28, 0xaa, 0x1e, 0xaa, 0x60, 0xfe, 0x09, 0x16, 0x2d, 0xa5, 0x8f, 0xe0, 0xd5,
	0xce, 0xd8, 0x4f, 0x4c, 0x4e, 0xf6, 0x2a, 0x87, 0x71, 0x4c, 0xf5, 0x36, 0x4f, 0xb0, 0x6a, 0x06,
	0xc8, 0xef, 0xe1, 0x34, 0x7f, 0x47, 0x2c, 0x82, 0xdd, 0x2e, 0xb7, 0xeb, 0x87, 0x6b, 0x79, 0xd4,
	0xe4, 0xe8, 0xe8, 0xca, 0xf5, 0x60, 0x6d, 0xde, 0x61, 0x3d, 0xbb, 0xe5, 0x37, 0x70, 0x5e, 0x30,
	0xa2, 0x3d, 0x50, 0xae, 0x03, 0xd8, 0xf6, 0x42, 0x4f, 0x99, 0x0b, 0x58, 0x19, 0xe7, 0x32, 0x96,
	0x22, 0x16, 0x3d, 0xfa, 0x8f, 0xdd, 0x2b, 0x87, 0xc6, 0xc5, 0x90, 0x50, 0x2c, 0x2b, 0x62, 0x7a,
	0xca, 0xcd, 0x1d, 0x5c, 0x59, 0xfa, 0x9a, 0x6d, 0xd1, 0x5c, 0x8e, 0x1f, 0xb6, 0x5d, 0x8d, 0x96,
	0xfd, 0x2e, 0x4e, 0x5e, 0x9f, 0xdb, 0xb7, 0xdd, 0x59, 0x5f, 0xeb, 0xf1, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x1b, 0xf2, 0x31, 0x3a, 0x1f, 0x01, 0x00, 0x00,
}
