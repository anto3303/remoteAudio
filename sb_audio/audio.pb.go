// Code generated by protoc-gen-gogo.
// source: audio.proto
// DO NOT EDIT!

/*
	Package shackbus_audio is a generated protocol buffer package.

	It is generated from these files:
		audio.proto

	It has these top-level messages:
		AudioData
		ClientRequest
		ServerResponse
*/
package shackbus_audio

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Codec int32

const (
	Codec_NONE Codec = 0
	Codec_OPUS Codec = 1
	Codec_PCM  Codec = 2
)

var Codec_name = map[int32]string{
	0: "NONE",
	1: "OPUS",
	2: "PCM",
}
var Codec_value = map[string]int32{
	"NONE": 0,
	"OPUS": 1,
	"PCM":  2,
}

func (x Codec) Enum() *Codec {
	p := new(Codec)
	*p = x
	return p
}
func (x Codec) String() string {
	return proto.EnumName(Codec_name, int32(x))
}
func (x *Codec) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Codec_value, data, "Codec")
	if err != nil {
		return err
	}
	*x = Codec(value)
	return nil
}
func (Codec) EnumDescriptor() ([]byte, []int) { return fileDescriptorAudio, []int{0} }

type AudioData struct {
	Codec            *Codec  `protobuf:"varint,1,opt,name=codec,enum=shackbus.audio.Codec" json:"codec,omitempty"`
	Channels         *int32  `protobuf:"varint,2,opt,name=channels" json:"channels,omitempty"`
	FrameLength      *int32  `protobuf:"varint,3,opt,name=frame_length,json=frameLength" json:"frame_length,omitempty"`
	SamplingRate     *int32  `protobuf:"varint,4,opt,name=sampling_rate,json=samplingRate" json:"sampling_rate,omitempty"`
	BitDepth         *int32  `protobuf:"varint,5,opt,name=bit_depth,json=bitDepth" json:"bit_depth,omitempty"`
	AudioRaw         []byte  `protobuf:"bytes,6,opt,name=audio_raw,json=audioRaw" json:"audio_raw,omitempty"`
	AudioPacked      []int32 `protobuf:"zigzag32,7,rep,packed,name=audio_packed,json=audioPacked" json:"audio_packed,omitempty"`
	UserId           *string `protobuf:"bytes,8,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AudioData) Reset()                    { *m = AudioData{} }
func (m *AudioData) String() string            { return proto.CompactTextString(m) }
func (*AudioData) ProtoMessage()               {}
func (*AudioData) Descriptor() ([]byte, []int) { return fileDescriptorAudio, []int{0} }

func (m *AudioData) GetCodec() Codec {
	if m != nil && m.Codec != nil {
		return *m.Codec
	}
	return Codec_NONE
}

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

func (m *AudioData) GetBitDepth() int32 {
	if m != nil && m.BitDepth != nil {
		return *m.BitDepth
	}
	return 0
}

func (m *AudioData) GetAudioRaw() []byte {
	if m != nil {
		return m.AudioRaw
	}
	return nil
}

func (m *AudioData) GetAudioPacked() []int32 {
	if m != nil {
		return m.AudioPacked
	}
	return nil
}

func (m *AudioData) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

type ClientRequest struct {
	RxAudioOn        *bool  `protobuf:"varint,1,opt,name=rx_audio_on,json=rxAudioOn" json:"rx_audio_on,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ClientRequest) Reset()                    { *m = ClientRequest{} }
func (m *ClientRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientRequest) ProtoMessage()               {}
func (*ClientRequest) Descriptor() ([]byte, []int) { return fileDescriptorAudio, []int{1} }

func (m *ClientRequest) GetRxAudioOn() bool {
	if m != nil && m.RxAudioOn != nil {
		return *m.RxAudioOn
	}
	return false
}

type ServerResponse struct {
	Online           *bool   `protobuf:"varint,1,opt,name=online" json:"online,omitempty"`
	LastSeen         *int64  `protobuf:"varint,2,opt,name=last_seen,json=lastSeen" json:"last_seen,omitempty"`
	RxAudioOn        *bool   `protobuf:"varint,3,opt,name=rx_audio_on,json=rxAudioOn" json:"rx_audio_on,omitempty"`
	TxUser           *string `protobuf:"bytes,4,opt,name=tx_user,json=txUser" json:"tx_user,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ServerResponse) Reset()                    { *m = ServerResponse{} }
func (m *ServerResponse) String() string            { return proto.CompactTextString(m) }
func (*ServerResponse) ProtoMessage()               {}
func (*ServerResponse) Descriptor() ([]byte, []int) { return fileDescriptorAudio, []int{2} }

func (m *ServerResponse) GetOnline() bool {
	if m != nil && m.Online != nil {
		return *m.Online
	}
	return false
}

func (m *ServerResponse) GetLastSeen() int64 {
	if m != nil && m.LastSeen != nil {
		return *m.LastSeen
	}
	return 0
}

func (m *ServerResponse) GetRxAudioOn() bool {
	if m != nil && m.RxAudioOn != nil {
		return *m.RxAudioOn
	}
	return false
}

func (m *ServerResponse) GetTxUser() string {
	if m != nil && m.TxUser != nil {
		return *m.TxUser
	}
	return ""
}

func init() {
	proto.RegisterType((*AudioData)(nil), "shackbus.audio.AudioData")
	proto.RegisterType((*ClientRequest)(nil), "shackbus.audio.ClientRequest")
	proto.RegisterType((*ServerResponse)(nil), "shackbus.audio.ServerResponse")
	proto.RegisterEnum("shackbus.audio.Codec", Codec_name, Codec_value)
}
func (m *AudioData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AudioData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Codec != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAudio(dAtA, i, uint64(*m.Codec))
	}
	if m.Channels != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAudio(dAtA, i, uint64(*m.Channels))
	}
	if m.FrameLength != nil {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAudio(dAtA, i, uint64(*m.FrameLength))
	}
	if m.SamplingRate != nil {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAudio(dAtA, i, uint64(*m.SamplingRate))
	}
	if m.BitDepth != nil {
		dAtA[i] = 0x28
		i++
		i = encodeVarintAudio(dAtA, i, uint64(*m.BitDepth))
	}
	if m.AudioRaw != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintAudio(dAtA, i, uint64(len(m.AudioRaw)))
		i += copy(dAtA[i:], m.AudioRaw)
	}
	if len(m.AudioPacked) > 0 {
		dAtA1 := make([]byte, len(m.AudioPacked)*5)
		var j2 int
		for _, num := range m.AudioPacked {
			x3 := (uint32(num) << 1) ^ uint32((num >> 31))
			for x3 >= 1<<7 {
				dAtA1[j2] = uint8(uint64(x3)&0x7f | 0x80)
				j2++
				x3 >>= 7
			}
			dAtA1[j2] = uint8(x3)
			j2++
		}
		dAtA[i] = 0x3a
		i++
		i = encodeVarintAudio(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA1[:j2])
	}
	if m.UserId != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintAudio(dAtA, i, uint64(len(*m.UserId)))
		i += copy(dAtA[i:], *m.UserId)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ClientRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RxAudioOn != nil {
		dAtA[i] = 0x8
		i++
		if *m.RxAudioOn {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ServerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServerResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Online != nil {
		dAtA[i] = 0x8
		i++
		if *m.Online {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.LastSeen != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAudio(dAtA, i, uint64(*m.LastSeen))
	}
	if m.RxAudioOn != nil {
		dAtA[i] = 0x18
		i++
		if *m.RxAudioOn {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.TxUser != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAudio(dAtA, i, uint64(len(*m.TxUser)))
		i += copy(dAtA[i:], *m.TxUser)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Audio(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Audio(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAudio(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AudioData) Size() (n int) {
	var l int
	_ = l
	if m.Codec != nil {
		n += 1 + sovAudio(uint64(*m.Codec))
	}
	if m.Channels != nil {
		n += 1 + sovAudio(uint64(*m.Channels))
	}
	if m.FrameLength != nil {
		n += 1 + sovAudio(uint64(*m.FrameLength))
	}
	if m.SamplingRate != nil {
		n += 1 + sovAudio(uint64(*m.SamplingRate))
	}
	if m.BitDepth != nil {
		n += 1 + sovAudio(uint64(*m.BitDepth))
	}
	if m.AudioRaw != nil {
		l = len(m.AudioRaw)
		n += 1 + l + sovAudio(uint64(l))
	}
	if len(m.AudioPacked) > 0 {
		l = 0
		for _, e := range m.AudioPacked {
			l += sozAudio(uint64(e))
		}
		n += 1 + sovAudio(uint64(l)) + l
	}
	if m.UserId != nil {
		l = len(*m.UserId)
		n += 1 + l + sovAudio(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ClientRequest) Size() (n int) {
	var l int
	_ = l
	if m.RxAudioOn != nil {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ServerResponse) Size() (n int) {
	var l int
	_ = l
	if m.Online != nil {
		n += 2
	}
	if m.LastSeen != nil {
		n += 1 + sovAudio(uint64(*m.LastSeen))
	}
	if m.RxAudioOn != nil {
		n += 2
	}
	if m.TxUser != nil {
		l = len(*m.TxUser)
		n += 1 + l + sovAudio(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAudio(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAudio(x uint64) (n int) {
	return sovAudio(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AudioData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAudio
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AudioData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AudioData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codec", wireType)
			}
			var v Codec
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (Codec(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Codec = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channels", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Channels = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrameLength", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FrameLength = &v
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SamplingRate", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SamplingRate = &v
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitDepth", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BitDepth = &v
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AudioRaw", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAudio
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AudioRaw = append(m.AudioRaw[:0], dAtA[iNdEx:postIndex]...)
			if m.AudioRaw == nil {
				m.AudioRaw = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAudio
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthAudio
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAudio
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
					m.AudioPacked = append(m.AudioPacked, v)
				}
			} else if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAudio
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
				m.AudioPacked = append(m.AudioPacked, v)
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field AudioPacked", wireType)
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAudio
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.UserId = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAudio(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAudio
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClientRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAudio
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RxAudioOn", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.RxAudioOn = &b
		default:
			iNdEx = preIndex
			skippy, err := skipAudio(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAudio
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAudio
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Online", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Online = &b
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSeen", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LastSeen = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RxAudioOn", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.RxAudioOn = &b
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxUser", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAudio
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.TxUser = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAudio(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAudio
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAudio(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAudio
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAudio
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAudio
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAudio
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAudio(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAudio = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAudio   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("audio.proto", fileDescriptorAudio) }

var fileDescriptorAudio = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x91, 0xd1, 0x8a, 0xd3, 0x40,
	0x18, 0x85, 0x77, 0x9a, 0x6d, 0x9b, 0xfc, 0xed, 0x96, 0x38, 0xa0, 0x0e, 0x0a, 0x21, 0x56, 0x85,
	0xa0, 0x50, 0xc1, 0x37, 0x70, 0xbb, 0x5e, 0x08, 0xba, 0x2d, 0x53, 0xf6, 0x7a, 0x98, 0x26, 0xbf,
	0x9b, 0xb0, 0xd9, 0x49, 0x9c, 0x99, 0xb8, 0xb9, 0xf2, 0x31, 0xc4, 0x47, 0xf2, 0xd2, 0x47, 0x90,
	0xfa, 0x22, 0x32, 0x93, 0x55, 0xd8, 0xbd, 0x9b, 0xf3, 0x9d, 0x43, 0xfe, 0xc3, 0x09, 0xcc, 0x64,
	0x57, 0x54, 0xcd, 0xaa, 0xd5, 0x8d, 0x6d, 0xe8, 0xc2, 0x94, 0x32, 0xbf, 0xda, 0x77, 0x66, 0xe5,
	0xe9, 0xf2, 0xfb, 0x08, 0xa2, 0x77, 0xee, 0x75, 0x26, 0xad, 0xa4, 0xaf, 0x61, 0x9c, 0x37, 0x05,
	0xe6, 0x8c, 0xa4, 0x24, 0x5b, 0xbc, 0x7d, 0xb8, 0xba, 0x9b, 0x5e, 0xad, 0x9d, 0xc9, 0x87, 0x0c,
	0x7d, 0x02, 0x61, 0x5e, 0x4a, 0xa5, 0xb0, 0x36, 0x6c, 0x94, 0x92, 0x6c, 0xcc, 0xff, 0x6b, 0xfa,
	0x0c, 0xe6, 0x9f, 0xb5, 0xbc, 0x46, 0x51, 0xa3, 0xba, 0xb4, 0x25, 0x0b, 0xbc, 0x3f, 0xf3, 0xec,
	0xa3, 0x47, 0xf4, 0x39, 0x9c, 0x18, 0x79, 0xdd, 0xd6, 0x95, 0xba, 0x14, 0x5a, 0x5a, 0x64, 0xc7,
	0x3e, 0x33, 0xff, 0x07, 0xb9, 0xb4, 0x48, 0x9f, 0x42, 0xb4, 0xaf, 0xac, 0x28, 0xb0, 0xb5, 0x25,
	0x1b, 0x0f, 0x47, 0xf6, 0x95, 0x3d, 0x73, 0xda, 0x99, 0xbe, 0x96, 0xd0, 0xf2, 0x86, 0x4d, 0x52,
	0x92, 0xcd, 0x79, 0xe8, 0x01, 0x97, 0x37, 0xf4, 0x25, 0xcc, 0x07, 0xb3, 0x95, 0xf9, 0x15, 0x16,
	0x6c, 0x9a, 0x06, 0xd9, 0x83, 0xd3, 0x51, 0x4c, 0xf8, 0xb0, 0xc7, 0xd6, 0x63, 0xfa, 0x18, 0xa6,
	0x9d, 0x41, 0x2d, 0xaa, 0x82, 0x85, 0x29, 0xc9, 0x22, 0x3e, 0x71, 0xf2, 0x43, 0xb1, 0x7c, 0x03,
	0x27, 0xeb, 0xba, 0x42, 0x65, 0x39, 0x7e, 0xe9, 0xd0, 0x58, 0x9a, 0xc0, 0x4c, 0xf7, 0x62, 0xf8,
	0x66, 0xa3, 0xfc, 0x42, 0x21, 0x8f, 0x74, 0xef, 0xd7, 0xdb, 0xa8, 0xe5, 0x37, 0x58, 0xec, 0x50,
	0x7f, 0x45, 0xcd, 0xd1, 0xb4, 0x8d, 0x32, 0x48, 0x1f, 0xc1, 0xa4, 0x51, 0x75, 0xa5, 0xf0, 0x36,
	0x7c, 0xab, 0x5c, 0xef, 0x5a, 0x1a, 0x2b, 0x0c, 0xa2, 0xf2, 0xcb, 0x05, 0x3c, 0x74, 0x60, 0x87,
	0xa8, 0xee, 0x9f, 0x09, 0xee, 0x9d, 0x71, 0x85, 0x6d, 0x2f, 0x5c, 0x49, 0x3f, 0x58, 0xc4, 0x27,
	0xb6, 0xbf, 0x30, 0xa8, 0x5f, 0xbd, 0x80, 0xb1, 0xff, 0x3d, 0x34, 0x84, 0xe3, 0xf3, 0xcd, 0xf9,
	0xfb, 0xf8, 0xc8, 0xbd, 0x36, 0xdb, 0x8b, 0x5d, 0x4c, 0xe8, 0x14, 0x82, 0xed, 0xfa, 0x53, 0x3c,
	0x3a, 0x8d, 0x7f, 0x1e, 0x12, 0xf2, 0xeb, 0x90, 0x90, 0xdf, 0x87, 0x84, 0xfc, 0xf8, 0x93, 0x1c,
	0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xbe, 0x7a, 0x1a, 0x1f, 0x02, 0x00, 0x00,
}
