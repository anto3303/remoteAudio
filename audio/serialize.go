package audio

import (
	"fmt"

	sbAudio "github.com/dh1tw/remoteAudio/sb_audio"
	"github.com/gogo/protobuf/proto"
	"github.com/hraban/opus"
)

// struct will all repetitive variables for serialization of
// audio packets
type serializer struct {
	*AudioDevice
	opusEncoder        *opus.Encoder
	opusBuffer         []byte
	wireSamplingrate   float64
	wireOutputChannels int
	framesPerBufferI   int32 // framesPerBuffer
	samplingRateI      int32 // samplingRate
	channelsI          int32 // output channels
	bitrateI           int32 // bitrate
	userID             string
}

func (s *serializer) SerializeOpusAudioMsg(in []float32) ([]byte, error) {

	fmt.Println("in bytes to encode:", len(in))
	len, err := s.opusEncoder.EncodeFloat32(in, s.opusBuffer)
	if err != nil {
		return nil, err
	}

	msg := sbAudioDataPool.Get().(*sbAudio.AudioData)
	defer sbAudioDataPool.Put(msg)

	msg.Channels = s.channelsI
	msg.FrameLength = s.framesPerBufferI
	msg.SamplingRate = s.samplingRateI
	msg.Bitrate = s.bitrateI
	msg.AudioRaw = s.opusBuffer[:len]
	msg.Codec = sbAudio.Codec_OPUS

	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// SerializeAudioMsg serializes audio frames in a protocol buffers with the
// corresponding meta data. The amount of audio channels and sampingrate can
// be specified.
func (s *serializer) SerializeAudioMsg(in []float32) ([]byte, error) {

	var resampledAudio []float32
	var audioToWire []int32

	// if necessary resample the audio and / or adjust the channels
	if (s.wireSamplingrate != s.Samplingrate) || (s.wireOutputChannels != s.Channels) {
		ratio := s.wireSamplingrate / s.Samplingrate // output samplerate / input samplerate
		var err error
		// cases: device MONO & output MONO  and device STEREO & output STEREO
		resampledAudio, err = s.Converter.Process(in, ratio, false)
		if err != nil {
			return nil, err
		}

		// audio device is STEREO but over the wire we want MONO
		if s.channelsI == MONO && s.Channels == STEREO {
			reduced := make([]float32, 0, len(resampledAudio)/2)
			// chop of the right channel
			for i := 0; i < len(resampledAudio); i += 2 {
				reduced = append(reduced, resampledAudio[i])
			}
			resampledAudio = reduced
		} else if s.channelsI == STEREO && s.Channels == MONO {
			// audio device is MONO but over the wire we want STEREO
			// doesn't make much sense
			expanded := make([]float32, 0, len(resampledAudio)*2)
			// left channel = right channel
			for _, sample := range resampledAudio {
				expanded = append(expanded, sample)
				expanded = append(expanded, sample)
			}
			resampledAudio = expanded
		}
	}

	// convert the data to int32
	if len(resampledAudio) > 0 { // in case we had to resample
		audioToWire = make([]int32, 0, len(resampledAudio))
		for _, sample := range resampledAudio {
			audioToWire = append(audioToWire, int32(sample*bitMapToInt32[s.bitrateI]))
		}
	} else { // otherwise just take the data from the sound card buffer
		audioToWire = make([]int32, 0, len(in))
		for _, sample := range in {
			audioToWire = append(audioToWire, int32(sample*bitMapToInt32[s.bitrateI]))
		}
	}

	msg := sbAudioDataPool.Get().(*sbAudio.AudioData)
	defer sbAudioDataPool.Put(msg)

	msg.Channels = s.channelsI
	msg.FrameLength = s.framesPerBufferI
	msg.SamplingRate = s.samplingRateI
	msg.Bitrate = s.bitrateI
	msg.AudioPacked = audioToWire

	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	// fmt.Println(len(data))

	return data, nil
}
