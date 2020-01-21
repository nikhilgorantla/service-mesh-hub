// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/v1/rbac.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for RbacMode
func (this *RbacMode) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RbacMode
func (this *RbacMode) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RbacMode_Disable
func (this *RbacMode_Disable) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RbacMode_Disable
func (this *RbacMode_Disable) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RbacMode_Enable
func (this *RbacMode_Enable) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RbacMode_Enable
func (this *RbacMode_Enable) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RbacStatus
func (this *RbacStatus) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RbacStatus
func (this *RbacStatus) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RbacMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	RbacUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
