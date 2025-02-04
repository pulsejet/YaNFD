package spec_2022

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"time"

	enc "github.com/named-data/ndnd/std/encoding"
	"github.com/named-data/ndnd/std/ndn"
	"github.com/named-data/ndnd/std/utils"
)

const TimeFmt = "20060102T150405" // ISO 8601 time format

func _() {
	// Trait for Signature of Data
	var _ ndn.Signature = &Data{}
	// Trait for Signature of Interest
	var _ ndn.Signature = &Interest{}
	// Trait for Data of Data
	var _ ndn.Data = &Data{}
	// Trait for Interest of Interest
	var _ ndn.Interest = &Interest{}
}

type Spec struct{}

func (d *Data) SigType() ndn.SigType {
	if d.SignatureInfo == nil {
		return ndn.SignatureNone
	} else {
		return ndn.SigType(d.SignatureInfo.SignatureType)
	}
}

func (d *Data) KeyName() enc.Name {
	if d.SignatureInfo == nil || d.SignatureInfo.KeyLocator == nil {
		return nil
	} else {
		return d.SignatureInfo.KeyLocator.Name
	}
}

func (d *Data) SigNonce() []byte {
	return nil
}

func (d *Data) SigTime() *time.Time {
	return nil
}

func (d *Data) SetSigTime(t *time.Time) error {
	if d.SignatureInfo == nil {
		d.SignatureInfo = &SignatureInfo{}
	}
	if t == nil {
		d.SignatureInfo.SignatureTime = nil
	} else {
		d.SignatureInfo.SignatureTime = utils.IdPtr(time.Duration(t.UnixMilli()) * time.Millisecond)
	}
	return nil
}

func (d *Data) SigSeqNum() *uint64 {
	return nil
}

func (d *Data) Validity() (notBefore, notAfter *time.Time) {
	if d.SignatureInfo != nil && d.SignatureInfo.ValidityPeriod != nil {
		notBefore, err := time.Parse(TimeFmt, d.SignatureInfo.ValidityPeriod.NotBefore)
		if err != nil {
			return nil, nil
		}
		notAfter, err := time.Parse(TimeFmt, d.SignatureInfo.ValidityPeriod.NotAfter)
		if err != nil {
			return nil, nil
		}
		return &notBefore, &notAfter
	} else {
		return nil, nil
	}
}

func (d *Data) SigValue() []byte {
	if d.SignatureValue == nil {
		return nil
	} else {
		return d.SignatureValue.Join()
	}
}

func (d *Data) Signature() ndn.Signature {
	return d
}

func (d *Data) Name() enc.Name {
	return d.NameV
}

func (d *Data) ContentType() *ndn.ContentType {
	if d.MetaInfo != nil && d.MetaInfo.ContentType != nil {
		ret := ndn.ContentType(*d.MetaInfo.ContentType)
		return &ret
	} else {
		return nil
	}
}

func (d *Data) Freshness() *time.Duration {
	if d.MetaInfo != nil {
		return d.MetaInfo.FreshnessPeriod
	} else {
		return nil
	}
}

func (d *Data) FinalBlockID() *enc.Component {
	if d.MetaInfo != nil && d.MetaInfo.FinalBlockID != nil {
		ret, err := enc.ReadComponent(enc.NewBufferReader(d.MetaInfo.FinalBlockID))
		if err == nil {
			return &ret
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (d *Data) Content() enc.Wire {
	return d.ContentV
}

func (t *Interest) SigType() ndn.SigType {
	if t.SignatureInfo == nil {
		return ndn.SignatureNone
	} else {
		return ndn.SigType(t.SignatureInfo.SignatureType)
	}
}

func (t *Interest) KeyName() enc.Name {
	if t.SignatureInfo == nil || t.SignatureInfo.KeyLocator == nil {
		return nil
	} else {
		return t.SignatureInfo.KeyLocator.Name
	}
}

func (t *Interest) SigNonce() []byte {
	if t.SignatureInfo != nil {
		return t.SignatureInfo.SignatureNonce
	} else {
		return nil
	}
}

func (t *Interest) SigTime() *time.Time {
	if t.SignatureInfo != nil && t.SignatureInfo.SignatureTime != nil {
		return utils.IdPtr(time.UnixMilli(t.SignatureInfo.SignatureTime.Milliseconds()))
	} else {
		return nil
	}
}

func (t *Interest) SigSeqNum() *uint64 {
	if t.SignatureInfo != nil {
		return t.SignatureInfo.SignatureSeqNum
	} else {
		return nil
	}
}

func (t *Interest) Validity() (notBefore, notAfter *time.Time) {
	return nil, nil
}

func (t *Interest) SigValue() []byte {
	return t.SignatureValue.Join()
}

func (t *Interest) Signature() ndn.Signature {
	return t
}

func (t *Interest) Name() enc.Name {
	return t.NameV
}

func (t *Interest) CanBePrefix() bool {
	return t.CanBePrefixV
}

func (t *Interest) MustBeFresh() bool {
	return t.MustBeFreshV
}

func (t *Interest) ForwardingHint() []enc.Name {
	if t.ForwardingHintV == nil {
		return nil
	}
	return t.ForwardingHintV.Names
}

func (t *Interest) Nonce() *uint64 {
	if t.NonceV == nil {
		return nil
	} else {
		return utils.IdPtr(uint64(*t.NonceV))
	}
}

func (t *Interest) Lifetime() *time.Duration {
	return t.InterestLifetimeV
}

func (t *Interest) HopLimit() *uint {
	if t.HopLimitV == nil {
		return nil
	} else {
		return utils.IdPtr(uint(*t.HopLimitV))
	}
}

func (t *Interest) AppParam() enc.Wire {
	return t.ApplicationParameters
}

func (Spec) MakeData(name enc.Name, config *ndn.DataConfig, content enc.Wire, signer ndn.Signer) (*ndn.EncodedData, error) {
	// Create Data packet.
	if name == nil {
		return nil, ndn.ErrInvalidValue{Item: "Data.Name", Value: nil}
	}
	if config == nil {
		return nil, ndn.ErrInvalidValue{Item: "Data.DataConfig", Value: nil}
	}
	finalBlock := []byte(nil)
	if config.FinalBlockID != nil {
		finalBlock = config.FinalBlockID.Bytes()
	}
	data := &Data{
		NameV: name,
		MetaInfo: &MetaInfo{
			ContentType:     utils.ConvIntPtr[ndn.ContentType, uint64](config.ContentType),
			FreshnessPeriod: config.Freshness,
			FinalBlockID:    finalBlock,
		},
		ContentV:       content,
		SignatureInfo:  nil,
		SignatureValue: nil,
	}
	packet := &Packet{
		Data: data,
	}

	// Fill-in SignatureInfo.
	estSigLen := 0
	if signer != nil && signer.Type() != ndn.SignatureNone {
		estSigLen = int(signer.EstimateSize())

		data.SignatureInfo = &SignatureInfo{
			SignatureType: uint64(signer.Type()),
		}

		if key := signer.KeyLocator(); key != nil {
			data.SignatureInfo.KeyLocator = &KeyLocator{Name: key}
		}

		if config.SigNotBefore != nil && config.SigNotAfter != nil {
			data.SignatureInfo.ValidityPeriod = &ValidityPeriod{
				NotBefore: config.SigNotBefore.UTC().Format(TimeFmt),
				NotAfter:  config.SigNotAfter.UTC().Format(TimeFmt),
			}
		}
	}

	// Encode packet.
	encoder := PacketEncoder{
		Data_encoder: DataEncoder{
			SignatureValue_estLen: uint(estSigLen),
		},
	}

	encoder.Init(packet)
	wire := encoder.Encode(packet)
	if wire == nil {
		return nil, ndn.ErrFailedToEncode
	}
	sigCovered := enc.Wire(nil)
	if estSigLen > 0 {
		// Compute signature
		sigCovered = encoder.Data_encoder.sigCovered

		// Since PacketEncoder only adds a TL, Data_encoder.SignatureValue_wireIdx is still valid
		sigVal, err := signer.Sign(sigCovered)
		if err != nil {
			return nil, err
		}

		if len(sigVal) > estSigLen {
			return nil, ndn.ErrNotSupported{Item: "Signature value cannot be longer than estimated length"}
		}
		wire[encoder.Data_encoder.SignatureValue_wireIdx] = sigVal

		// Fix SignatureValue length
		// This does not handle the case where the signature value is so much shorter than
		// the estimated length that the length field needs to be shrunk.
		// The signer needs to provide a reasonable estimate, hopefully exact.
		buf := wire[encoder.Data_encoder.SignatureValue_wireIdx-1]
		buf[len(buf)-1] = byte(len(sigVal))
		shrink := estSigLen - len(sigVal)
		wire[0] = enc.ShrinkLength(wire[0], shrink)
	}
	return &ndn.EncodedData{
		Wire:       wire,
		SigCovered: sigCovered,
		Config:     config,
	}, nil
}

func (Spec) ReadData(reader enc.ParseReader) (ndn.Data, enc.Wire, error) {
	context := PacketParsingContext{}
	context.Init()
	ret, err := context.Parse(reader, false)
	if err != nil {
		return nil, nil, err
	}
	if ret.Data == nil {
		return nil, nil, ndn.ErrWrongType
	}
	if ret.Data.NameV == nil {
		return nil, nil, ndn.ErrInvalidValue{Item: "Data.Name", Value: nil}
	}
	return ret.Data, context.Data_context.sigCovered, nil
}

func (Spec) MakeInterest(name enc.Name, config *ndn.InterestConfig, appParam enc.Wire, signer ndn.Signer) (*ndn.EncodedInterest, error) {
	// Create Interest packet.
	if name == nil {
		return nil, ndn.ErrInvalidValue{Item: "Interest.Name", Value: nil}
	}
	if config == nil {
		return nil, ndn.ErrInvalidValue{Item: "Interest.DataConfig", Value: nil}
	}
	forwardingHint := (*Links)(nil)
	if config.ForwardingHint != nil {
		forwardingHint = &Links{
			Names: config.ForwardingHint,
		}
	}
	interest := &Interest{
		NameV:                 name,
		CanBePrefixV:          config.CanBePrefix,
		MustBeFreshV:          config.MustBeFresh,
		ForwardingHintV:       forwardingHint,
		NonceV:                utils.ConvIntPtr[uint64, uint32](config.Nonce),
		InterestLifetimeV:     config.Lifetime,
		HopLimitV:             utils.ConvIntPtr[uint, byte](config.HopLimit),
		ApplicationParameters: appParam,
		SignatureInfo:         nil,
		SignatureValue:        nil,
	}
	packet := &Packet{
		Interest: interest,
	}

	needDigest := appParam != nil
	estSigLen := 0

	// Fill-in SignatureInfo.
	if signer != nil && signer.Type() != ndn.SignatureNone {
		if !needDigest {
			return nil, ndn.ErrInvalidValue{Item: "Interest.ApplicationParameters", Value: nil}
		}

		estSigLen = int(signer.EstimateSize())

		interest.SignatureInfo = &SignatureInfo{
			SignatureType:   uint64(signer.Type()),
			SignatureNonce:  config.SigNonce,
			SignatureTime:   config.SigTime,
			SignatureSeqNum: config.SigSeqNo,
		}

		if key := signer.KeyLocator(); key != nil {
			interest.SignatureInfo.KeyLocator = &KeyLocator{Name: key}
		}
	}

	// Encode packet.
	encoder := PacketEncoder{
		Interest_encoder: InterestEncoder{
			SignatureValue_estLen: uint(estSigLen),
			NameV_needDigest:      needDigest,
		},
	}
	ecdr := &encoder.Interest_encoder
	encoder.Init(packet)
	wire := encoder.Encode(packet)
	if wire == nil {
		return nil, ndn.ErrFailedToEncode
	}
	sigVal := []byte(nil)
	err := error(nil)
	sigCovered := enc.Wire(nil)
	if estSigLen > 0 {
		// Compute signature
		sigCovered = ecdr.sigCovered
		if ecdr.SignatureValue_wireIdx < 0 {
			return nil, enc.ErrUnexpected{Err: errors.New("SignatureValue is not correctly set")}
		}

		// Since PacketEncoder only adds a TL, Interest_encoder.SignatureValue_wireIdx is still valid
		sigVal, err = signer.Sign(sigCovered)
		if err != nil {
			return nil, err
		}

		if uint(len(sigVal)) > ecdr.SignatureValue_estLen {
			return nil, ndn.ErrNotSupported{Item: "Signature value cannot be longer than estimated length"}
		}

		// Fix SignatureValue length
		wire[ecdr.SignatureValue_wireIdx] = sigVal
		buf := wire[ecdr.SignatureValue_wireIdx-1]
		buf[len(buf)-1] = byte(len(sigVal))

		// Don't fix packet length for now, as it may cause trouble
	}
	finalName := packet.Interest.NameV
	if needDigest {
		// Compute digest
		// assert ecdr.NameV_wireIdx == 0
		buf := wire[0]
		_, s1 := enc.ParseTLNum(buf)
		_, s2 := enc.ParseTLNum(buf[s1:])
		// Add the offset by Interest TL
		digestPos := ecdr.NameV_pos + uint(s1+s2)
		digestBuf := buf[digestPos : digestPos+32]
		// Set the digest of final name
		finalName[len(finalName)-1].Val = digestBuf
		// Due to no copy, digest coveres AppParam type(1B) + len + wire[1:]
		appParamLen := enc.TLNum(appParam.Length()).EncodingLength()
		digestCovered := wire[1:]
		// Compute sha256 hash
		h := sha256.New()
		h.Write(wire[0][len(wire[0])-appParamLen-1:])
		for _, buf := range digestCovered {
			_, err = h.Write(buf)
			if err != nil {
				return nil, enc.ErrUnexpected{Err: err}
			}
		}
		copy(digestBuf, h.Sum(nil))
	}

	// Fix packet length
	shrink := estSigLen - len(sigVal)
	if shrink > 0 {
		wire[0] = enc.ShrinkLength(wire[0], shrink)
	} else if shrink < 0 {
		return nil, ndn.ErrNotSupported{Item: "Too long signature value is not supported"}
	}

	return &ndn.EncodedInterest{
		Wire:       wire,
		SigCovered: sigCovered,
		FinalName:  finalName,
		Config:     config,
	}, nil
}

func checkInterest(val *Interest, context *InterestParsingContext) error {
	if val.NameV == nil {
		return ndn.ErrInvalidValue{Item: "Interest.Name", Value: nil}
	}
	if val.SignatureValue != nil && val.ApplicationParameters == nil {
		return enc.ErrIncorrectDigest
	}
	if val.ApplicationParameters != nil {
		// Check digest
		name := val.NameV
		if len(name) == 0 || name.At(-1).Typ != enc.TypeParametersSha256DigestComponent {
			return enc.ErrIncorrectDigest
		}
		digestCovered := context.digestCovered
		h := sha256.New()
		for _, buf := range digestCovered {
			_, err := h.Write(buf)
			if err != nil {
				return enc.ErrUnexpected{Err: err}
			}
		}
		digestBuf := h.Sum(nil)
		if !bytes.Equal(name.At(-1).Val, digestBuf) {
			return enc.ErrIncorrectDigest
		}
	}
	return nil
}

func (Spec) ReadInterest(reader enc.ParseReader) (ndn.Interest, enc.Wire, error) {
	context := PacketParsingContext{}
	context.Init()
	pkt, err := context.Parse(reader, false)
	if err != nil {
		return nil, nil, err
	}
	if pkt.Interest == nil {
		return nil, nil, ndn.ErrWrongType
	}

	err = checkInterest(pkt.Interest, &context.Interest_context)
	if err != nil {
		return nil, nil, err
	}

	return pkt.Interest, context.Interest_context.sigCovered, nil
}

// ReadPacket parses a packet from the reader.
//
//	Precondition: reader contains only one TLV.
//	Postcondition: exactly one of Interest, Data, or LpPacket is returned.
//
// If precondition is not met, then postcondition is not required to hold. But the call won't crash.
func ReadPacket(reader enc.ParseReader) (*Packet, *PacketParsingContext, error) {
	context := &PacketParsingContext{}
	context.Init()
	ret, err := context.Parse(reader, false)
	if err != nil {
		return nil, nil, err
	}
	if ret.Data != nil {
		if ret.Data.NameV == nil {
			return nil, nil, ndn.ErrInvalidValue{Item: "Data.Name", Value: nil}
		}
	} else if ret.Interest != nil {
		err = checkInterest(ret.Interest, &context.Interest_context)
		if err != nil {
			return nil, nil, err
		}
	} else if ret.LpPacket != nil {
		// As a client we shouldn't receive IDLE packets
		if ret.LpPacket.Fragment == nil {
			return nil, nil, ndn.ErrInvalidValue{Item: "LpPacket.Fragment", Value: nil}
		}
	} else {
		return nil, nil, ndn.ErrWrongType
	}
	return ret, context, nil
}

func (c InterestParsingContext) SigCovered() enc.Wire {
	return c.sigCovered
}

func (c DataParsingContext) SigCovered() enc.Wire {
	return c.sigCovered
}
