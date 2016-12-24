// Code generated by protoc-gen-go.
// source: mqttaudio.proto
// DO NOT EDIT!

/*
Package icd is a generated protocol buffer package.

It is generated from these files:
	mqttaudio.proto

It has these top-level messages:
	AudioData
	Error
*/
package icd

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type STATUS int32

const (
	STATUS_FAILURE             STATUS = 1
	STATUS_LIMITED_OPERATIONAL STATUS = 2
	STATUS_OPERATIONAL         STATUS = 3
)

var STATUS_name = map[int32]string{
	1: "FAILURE",
	2: "LIMITED_OPERATIONAL",
	3: "OPERATIONAL",
}
var STATUS_value = map[string]int32{
	"FAILURE":             1,
	"LIMITED_OPERATIONAL": 2,
	"OPERATIONAL":         3,
}

func (x STATUS) Enum() *STATUS {
	p := new(STATUS)
	*p = x
	return p
}
func (x STATUS) String() string {
	return proto.EnumName(STATUS_name, int32(x))
}
func (x *STATUS) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(STATUS_value, data, "STATUS")
	if err != nil {
		return err
	}
	*x = STATUS(value)
	return nil
}
func (STATUS) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ERROR_LEVEL int32

const (
	ERROR_LEVEL_TRACE   ERROR_LEVEL = 1
	ERROR_LEVEL_DEBUG   ERROR_LEVEL = 2
	ERROR_LEVEL_INFO    ERROR_LEVEL = 3
	ERROR_LEVEL_WARNING ERROR_LEVEL = 4
	ERROR_LEVEL_ERROR   ERROR_LEVEL = 5
)

var ERROR_LEVEL_name = map[int32]string{
	1: "TRACE",
	2: "DEBUG",
	3: "INFO",
	4: "WARNING",
	5: "ERROR",
}
var ERROR_LEVEL_value = map[string]int32{
	"TRACE":   1,
	"DEBUG":   2,
	"INFO":    3,
	"WARNING": 4,
	"ERROR":   5,
}

func (x ERROR_LEVEL) Enum() *ERROR_LEVEL {
	p := new(ERROR_LEVEL)
	*p = x
	return p
}
func (x ERROR_LEVEL) String() string {
	return proto.EnumName(ERROR_LEVEL_name, int32(x))
}
func (x *ERROR_LEVEL) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ERROR_LEVEL_value, data, "ERROR_LEVEL")
	if err != nil {
		return err
	}
	*x = ERROR_LEVEL(value)
	return nil
}
func (ERROR_LEVEL) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AudioData struct {
	Channels         *int32  `protobuf:"varint,1,opt,name=channels" json:"channels,omitempty"`
	FrameLength      *int32  `protobuf:"varint,2,opt,name=frame_length" json:"frame_length,omitempty"`
	SamplingRate     *int32  `protobuf:"varint,3,opt,name=sampling_rate" json:"sampling_rate,omitempty"`
	Bitrate          *int32  `protobuf:"varint,4,opt,name=bitrate" json:"bitrate,omitempty"`
	Audio            []int32 `protobuf:"zigzag32,6,rep,packed,name=audio" json:"audio,omitempty"`
	UserId           *string `protobuf:"bytes,7,opt,name=user_id" json:"user_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AudioData) Reset()                    { *m = AudioData{} }
func (m *AudioData) String() string            { return proto.CompactTextString(m) }
func (*AudioData) ProtoMessage()               {}
func (*AudioData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AudioData) GetChannels() int32 {
	if m != nil && m.Channels != nil {
		return *m.Channels
	}
	return 0
}

func (m *AudioData) GetFrameLength() int32 {
	if m != nil && m.FrameLength != nil {
		return *m.FrameLength
	}
	return 0
}

func (m *AudioData) GetSamplingRate() int32 {
	if m != nil && m.SamplingRate != nil {
		return *m.SamplingRate
	}
	return 0
}

func (m *AudioData) GetBitrate() int32 {
	if m != nil && m.Bitrate != nil {
		return *m.Bitrate
	}
	return 0
}

func (m *AudioData) GetAudio() []int32 {
	if m != nil {
		return m.Audio
	}
	return nil
}

func (m *AudioData) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

type Error struct {
	Msg              *string      `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Level            *ERROR_LEVEL `protobuf:"varint,2,opt,name=level,enum=icd.ERROR_LEVEL" json:"level,omitempty"`
	Description      *string      `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Error) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *Error) GetLevel() ERROR_LEVEL {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return ERROR_LEVEL_TRACE
}

func (m *Error) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*AudioData)(nil), "icd.AudioData")
	proto.RegisterType((*Error)(nil), "icd.Error")
	proto.RegisterEnum("icd.STATUS", STATUS_name, STATUS_value)
	proto.RegisterEnum("icd.ERROR_LEVEL", ERROR_LEVEL_name, ERROR_LEVEL_value)
}

var fileDescriptor0 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8e, 0x3d, 0x6f, 0xf2, 0x30,
	0x14, 0x85, 0x05, 0x21, 0x40, 0x6e, 0xde, 0xb7, 0x98, 0xd0, 0xaa, 0xd9, 0x8a, 0x98, 0x10, 0x03,
	0x43, 0xff, 0x40, 0x65, 0x8a, 0x41, 0xae, 0xd2, 0xa4, 0x32, 0xa1, 0x1d, 0x23, 0x97, 0xb8, 0x60,
	0x29, 0x1f, 0xd4, 0x31, 0x5d, 0xfb, 0xd7, 0xeb, 0x78, 0x62, 0xbb, 0x7a, 0x74, 0xee, 0x79, 0x0e,
	0x8c, 0xca, 0x6f, 0xad, 0xf9, 0x25, 0x97, 0xf5, 0xf2, 0xac, 0x6a, 0x5d, 0x07, 0x8e, 0x3c, 0xe4,
	0xb3, 0x5f, 0xf0, 0x70, 0xcb, 0xd6, 0x5c, 0xf3, 0x00, 0xc1, 0xf0, 0x70, 0xe2, 0x55, 0x25, 0x8a,
	0x26, 0xec, 0x4c, 0x3b, 0x73, 0x37, 0xb8, 0x85, 0x7f, 0x5f, 0x8a, 0x97, 0x22, 0x2b, 0x44, 0x75,
	0xd4, 0xa7, 0xb0, 0x6b, 0xe9, 0x1d, 0xfc, 0x6f, 0x78, 0x79, 0x2e, 0x64, 0x75, 0xcc, 0x14, 0xd7,
	0x22, 0x74, 0x2c, 0x1e, 0xc1, 0xe0, 0x53, 0x6a, 0x0b, 0x7a, 0x16, 0x8c, 0xc1, 0xb5, 0xc2, 0xb0,
	0x3f, 0x75, 0xe6, 0xe3, 0x55, 0x17, 0x75, 0xda, 0xcc, 0xa5, 0x11, 0x2a, 0x93, 0x79, 0x38, 0x30,
	0x19, 0x6f, 0xf6, 0x02, 0x2e, 0x51, 0xaa, 0x56, 0x81, 0x0f, 0x4e, 0xd9, 0x1c, 0xad, 0xd7, 0x0b,
	0x1e, 0xc0, 0x2d, 0xc4, 0x8f, 0x28, 0xac, 0xf0, 0xe6, 0x11, 0x2d, 0xcd, 0xd6, 0x25, 0x61, 0x2c,
	0x61, 0x59, 0x44, 0xde, 0x49, 0x14, 0x4c, 0xc0, 0xcf, 0x45, 0x73, 0x50, 0xf2, 0xac, 0x65, 0x5d,
	0xd9, 0x01, 0xde, 0xe2, 0x09, 0xfa, 0xbb, 0x14, 0xa7, 0xfb, 0x9d, 0x29, 0x1b, 0x6c, 0x30, 0x8d,
	0xf6, 0x8c, 0x18, 0xe7, 0x3d, 0x4c, 0x22, 0xfa, 0x4a, 0x53, 0xb2, 0xce, 0x92, 0x37, 0xc2, 0x70,
	0x4a, 0x93, 0x18, 0x47, 0xa8, 0x6b, 0xc6, 0xf8, 0xd7, 0xc0, 0x59, 0x10, 0xf0, 0xaf, 0x25, 0x1e,
	0xb8, 0x29, 0xc3, 0xcf, 0x6d, 0x87, 0x39, 0xd7, 0x64, 0xb5, 0xdf, 0x9a, 0xaf, 0x21, 0xf4, 0x68,
	0xbc, 0x49, 0x90, 0xd3, 0x5a, 0x3e, 0x30, 0x8b, 0x69, 0xbc, 0x45, 0xbd, 0x36, 0x61, 0x7f, 0x91,
	0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x19, 0x47, 0x12, 0x38, 0x6b, 0x01, 0x00, 0x00,
}
