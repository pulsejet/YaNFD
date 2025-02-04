// Code generated by ndn tlv codegen DO NOT EDIT.
package rdr

import (
	"encoding/binary"
	"io"
	"strings"

	enc "github.com/named-data/ndnd/std/encoding"
)

type ManifestDigestEncoder struct {
	length uint
}

type ManifestDigestParsingContext struct {
}

func (encoder *ManifestDigestEncoder) Init(value *ManifestDigest) {

	l := uint(0)
	l += 1
	l += uint(1 + enc.Nat(value.SegNo).EncodingLength())
	if value.Digest != nil {
		l += 1
		l += uint(enc.TLNum(len(value.Digest)).EncodingLength())
		l += uint(len(value.Digest))
	}
	encoder.length = l

}

func (context *ManifestDigestParsingContext) Init() {

}

func (encoder *ManifestDigestEncoder) EncodeInto(value *ManifestDigest, buf []byte) {

	pos := uint(0)

	buf[pos] = byte(204)
	pos += 1

	buf[pos] = byte(enc.Nat(value.SegNo).EncodeInto(buf[pos+1:]))
	pos += uint(1 + buf[pos])
	if value.Digest != nil {
		buf[pos] = byte(206)
		pos += 1
		pos += uint(enc.TLNum(len(value.Digest)).EncodeInto(buf[pos:]))
		copy(buf[pos:], value.Digest)
		pos += uint(len(value.Digest))
	}
}

func (encoder *ManifestDigestEncoder) Encode(value *ManifestDigest) enc.Wire {

	wire := make(enc.Wire, 1)
	wire[0] = make([]byte, encoder.length)
	buf := wire[0]
	encoder.EncodeInto(value, buf)

	return wire
}

func (context *ManifestDigestParsingContext) Parse(reader enc.ParseReader, ignoreCritical bool) (*ManifestDigest, error) {
	if reader == nil {
		return nil, enc.ErrBufferOverflow
	}

	var handled_SegNo bool = false
	var handled_Digest bool = false

	progress := -1
	_ = progress

	value := &ManifestDigest{}
	var err error
	var startPos int
	for {
		startPos = reader.Pos()
		if startPos >= reader.Length() {
			break
		}
		typ := enc.TLNum(0)
		l := enc.TLNum(0)
		typ, err = enc.ReadTLNum(reader)
		if err != nil {
			return nil, enc.ErrFailToParse{TypeNum: 0, Err: err}
		}
		l, err = enc.ReadTLNum(reader)
		if err != nil {
			return nil, enc.ErrFailToParse{TypeNum: 0, Err: err}
		}

		err = nil
		if handled := false; true {
			switch typ {
			case 204:
				if true {
					handled = true
					handled_SegNo = true
					value.SegNo = uint64(0)
					{
						for i := 0; i < int(l); i++ {
							x := byte(0)
							x, err = reader.ReadByte()
							if err != nil {
								if err == io.EOF {
									err = io.ErrUnexpectedEOF
								}
								break
							}
							value.SegNo = uint64(value.SegNo<<8) | uint64(x)
						}
					}
				}
			case 206:
				if true {
					handled = true
					handled_Digest = true
					value.Digest = make([]byte, l)
					_, err = io.ReadFull(reader, value.Digest)
				}
			default:
				if !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
					return nil, enc.ErrUnrecognizedField{TypeNum: typ}
				}
				handled = true
				err = reader.Skip(int(l))
			}
			if err == nil && !handled {
			}
			if err != nil {
				return nil, enc.ErrFailToParse{TypeNum: typ, Err: err}
			}
		}
	}

	startPos = reader.Pos()
	err = nil

	if !handled_SegNo && err == nil {
		err = enc.ErrSkipRequired{Name: "SegNo", TypeNum: 204}
	}
	if !handled_Digest && err == nil {
		value.Digest = nil
	}

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (value *ManifestDigest) Encode() enc.Wire {
	encoder := ManifestDigestEncoder{}
	encoder.Init(value)
	return encoder.Encode(value)
}

func (value *ManifestDigest) Bytes() []byte {
	return value.Encode().Join()
}

func ParseManifestDigest(reader enc.ParseReader, ignoreCritical bool) (*ManifestDigest, error) {
	context := ManifestDigestParsingContext{}
	context.Init()
	return context.Parse(reader, ignoreCritical)
}

type ManifestDataEncoder struct {
	length uint

	Entries_subencoder []struct {
		Entries_encoder ManifestDigestEncoder
	}
}

type ManifestDataParsingContext struct {
	Entries_context ManifestDigestParsingContext
}

func (encoder *ManifestDataEncoder) Init(value *ManifestData) {
	{
		Entries_l := len(value.Entries)
		encoder.Entries_subencoder = make([]struct {
			Entries_encoder ManifestDigestEncoder
		}, Entries_l)
		for i := 0; i < Entries_l; i++ {
			pseudoEncoder := &encoder.Entries_subencoder[i]
			pseudoValue := struct {
				Entries *ManifestDigest
			}{
				Entries: value.Entries[i],
			}
			{
				encoder := pseudoEncoder
				value := &pseudoValue
				if value.Entries != nil {
					encoder.Entries_encoder.Init(value.Entries)
				}
				_ = encoder
				_ = value
			}
		}
	}

	l := uint(0)
	if value.Entries != nil {
		for seq_i, seq_v := range value.Entries {
			pseudoEncoder := &encoder.Entries_subencoder[seq_i]
			pseudoValue := struct {
				Entries *ManifestDigest
			}{
				Entries: seq_v,
			}
			{
				encoder := pseudoEncoder
				value := &pseudoValue
				if value.Entries != nil {
					l += 1
					l += uint(enc.TLNum(encoder.Entries_encoder.length).EncodingLength())
					l += encoder.Entries_encoder.length
				}
				_ = encoder
				_ = value
			}
		}
	}
	encoder.length = l

}

func (context *ManifestDataParsingContext) Init() {
	context.Entries_context.Init()
}

func (encoder *ManifestDataEncoder) EncodeInto(value *ManifestData, buf []byte) {

	pos := uint(0)

	if value.Entries != nil {
		for seq_i, seq_v := range value.Entries {
			pseudoEncoder := &encoder.Entries_subencoder[seq_i]
			pseudoValue := struct {
				Entries *ManifestDigest
			}{
				Entries: seq_v,
			}
			{
				encoder := pseudoEncoder
				value := &pseudoValue
				if value.Entries != nil {
					buf[pos] = byte(202)
					pos += 1
					pos += uint(enc.TLNum(encoder.Entries_encoder.length).EncodeInto(buf[pos:]))
					if encoder.Entries_encoder.length > 0 {
						encoder.Entries_encoder.EncodeInto(value.Entries, buf[pos:])
						pos += encoder.Entries_encoder.length
					}
				}
				_ = encoder
				_ = value
			}
		}
	}
}

func (encoder *ManifestDataEncoder) Encode(value *ManifestData) enc.Wire {

	wire := make(enc.Wire, 1)
	wire[0] = make([]byte, encoder.length)
	buf := wire[0]
	encoder.EncodeInto(value, buf)

	return wire
}

func (context *ManifestDataParsingContext) Parse(reader enc.ParseReader, ignoreCritical bool) (*ManifestData, error) {
	if reader == nil {
		return nil, enc.ErrBufferOverflow
	}

	var handled_Entries bool = false

	progress := -1
	_ = progress

	value := &ManifestData{}
	var err error
	var startPos int
	for {
		startPos = reader.Pos()
		if startPos >= reader.Length() {
			break
		}
		typ := enc.TLNum(0)
		l := enc.TLNum(0)
		typ, err = enc.ReadTLNum(reader)
		if err != nil {
			return nil, enc.ErrFailToParse{TypeNum: 0, Err: err}
		}
		l, err = enc.ReadTLNum(reader)
		if err != nil {
			return nil, enc.ErrFailToParse{TypeNum: 0, Err: err}
		}

		err = nil
		if handled := false; true {
			switch typ {
			case 202:
				if true {
					handled = true
					handled_Entries = true
					if value.Entries == nil {
						value.Entries = make([]*ManifestDigest, 0)
					}
					{
						pseudoValue := struct {
							Entries *ManifestDigest
						}{}
						{
							value := &pseudoValue
							value.Entries, err = context.Entries_context.Parse(reader.Delegate(int(l)), ignoreCritical)
							_ = value
						}
						value.Entries = append(value.Entries, pseudoValue.Entries)
					}
					progress--
				}
			default:
				if !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
					return nil, enc.ErrUnrecognizedField{TypeNum: typ}
				}
				handled = true
				err = reader.Skip(int(l))
			}
			if err == nil && !handled {
			}
			if err != nil {
				return nil, enc.ErrFailToParse{TypeNum: typ, Err: err}
			}
		}
	}

	startPos = reader.Pos()
	err = nil

	if !handled_Entries && err == nil {
		// sequence - skip
	}

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (value *ManifestData) Encode() enc.Wire {
	encoder := ManifestDataEncoder{}
	encoder.Init(value)
	return encoder.Encode(value)
}

func (value *ManifestData) Bytes() []byte {
	return value.Encode().Join()
}

func ParseManifestData(reader enc.ParseReader, ignoreCritical bool) (*ManifestData, error) {
	context := ManifestDataParsingContext{}
	context.Init()
	return context.Parse(reader, ignoreCritical)
}

type MetaDataEncoder struct {
	length uint

	Name_length uint
}

type MetaDataParsingContext struct {
}

func (encoder *MetaDataEncoder) Init(value *MetaData) {
	if value.Name != nil {
		encoder.Name_length = 0
		for _, c := range value.Name {
			encoder.Name_length += uint(c.EncodingLength())
		}
	}

	l := uint(0)
	if value.Name != nil {
		l += 1
		l += uint(enc.TLNum(encoder.Name_length).EncodingLength())
		l += encoder.Name_length
	}
	if value.FinalBlockID != nil {
		l += 1
		l += uint(enc.TLNum(len(value.FinalBlockID)).EncodingLength())
		l += uint(len(value.FinalBlockID))
	}
	if value.SegmentSize != nil {
		l += 3
		l += uint(1 + enc.Nat(*value.SegmentSize).EncodingLength())
	}
	if value.Size != nil {
		l += 3
		l += uint(1 + enc.Nat(*value.Size).EncodingLength())
	}
	if value.Mode != nil {
		l += 3
		l += uint(1 + enc.Nat(*value.Mode).EncodingLength())
	}
	if value.Atime != nil {
		l += 3
		l += uint(1 + enc.Nat(*value.Atime).EncodingLength())
	}
	if value.Btime != nil {
		l += 3
		l += uint(1 + enc.Nat(*value.Btime).EncodingLength())
	}
	if value.Ctime != nil {
		l += 3
		l += uint(1 + enc.Nat(*value.Ctime).EncodingLength())
	}
	if value.Mtime != nil {
		l += 3
		l += uint(1 + enc.Nat(*value.Mtime).EncodingLength())
	}
	if value.ObjectType != nil {
		l += 3
		l += uint(enc.TLNum(len(*value.ObjectType)).EncodingLength())
		l += uint(len(*value.ObjectType))
	}
	encoder.length = l

}

func (context *MetaDataParsingContext) Init() {

}

func (encoder *MetaDataEncoder) EncodeInto(value *MetaData, buf []byte) {

	pos := uint(0)

	if value.Name != nil {
		buf[pos] = byte(7)
		pos += 1
		pos += uint(enc.TLNum(encoder.Name_length).EncodeInto(buf[pos:]))
		for _, c := range value.Name {
			pos += uint(c.EncodeInto(buf[pos:]))
		}
	}
	if value.FinalBlockID != nil {
		buf[pos] = byte(26)
		pos += 1
		pos += uint(enc.TLNum(len(value.FinalBlockID)).EncodeInto(buf[pos:]))
		copy(buf[pos:], value.FinalBlockID)
		pos += uint(len(value.FinalBlockID))
	}
	if value.SegmentSize != nil {
		buf[pos] = 253
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(62720))
		pos += 3

		buf[pos] = byte(enc.Nat(*value.SegmentSize).EncodeInto(buf[pos+1:]))
		pos += uint(1 + buf[pos])

	}
	if value.Size != nil {
		buf[pos] = 253
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(62722))
		pos += 3

		buf[pos] = byte(enc.Nat(*value.Size).EncodeInto(buf[pos+1:]))
		pos += uint(1 + buf[pos])

	}
	if value.Mode != nil {
		buf[pos] = 253
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(62724))
		pos += 3

		buf[pos] = byte(enc.Nat(*value.Mode).EncodeInto(buf[pos+1:]))
		pos += uint(1 + buf[pos])

	}
	if value.Atime != nil {
		buf[pos] = 253
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(62726))
		pos += 3

		buf[pos] = byte(enc.Nat(*value.Atime).EncodeInto(buf[pos+1:]))
		pos += uint(1 + buf[pos])

	}
	if value.Btime != nil {
		buf[pos] = 253
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(62728))
		pos += 3

		buf[pos] = byte(enc.Nat(*value.Btime).EncodeInto(buf[pos+1:]))
		pos += uint(1 + buf[pos])

	}
	if value.Ctime != nil {
		buf[pos] = 253
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(62730))
		pos += 3

		buf[pos] = byte(enc.Nat(*value.Ctime).EncodeInto(buf[pos+1:]))
		pos += uint(1 + buf[pos])

	}
	if value.Mtime != nil {
		buf[pos] = 253
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(62732))
		pos += 3

		buf[pos] = byte(enc.Nat(*value.Mtime).EncodeInto(buf[pos+1:]))
		pos += uint(1 + buf[pos])

	}
	if value.ObjectType != nil {
		buf[pos] = 253
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(62734))
		pos += 3
		pos += uint(enc.TLNum(len(*value.ObjectType)).EncodeInto(buf[pos:]))
		copy(buf[pos:], *value.ObjectType)
		pos += uint(len(*value.ObjectType))
	}
}

func (encoder *MetaDataEncoder) Encode(value *MetaData) enc.Wire {

	wire := make(enc.Wire, 1)
	wire[0] = make([]byte, encoder.length)
	buf := wire[0]
	encoder.EncodeInto(value, buf)

	return wire
}

func (context *MetaDataParsingContext) Parse(reader enc.ParseReader, ignoreCritical bool) (*MetaData, error) {
	if reader == nil {
		return nil, enc.ErrBufferOverflow
	}

	var handled_Name bool = false
	var handled_FinalBlockID bool = false
	var handled_SegmentSize bool = false
	var handled_Size bool = false
	var handled_Mode bool = false
	var handled_Atime bool = false
	var handled_Btime bool = false
	var handled_Ctime bool = false
	var handled_Mtime bool = false
	var handled_ObjectType bool = false

	progress := -1
	_ = progress

	value := &MetaData{}
	var err error
	var startPos int
	for {
		startPos = reader.Pos()
		if startPos >= reader.Length() {
			break
		}
		typ := enc.TLNum(0)
		l := enc.TLNum(0)
		typ, err = enc.ReadTLNum(reader)
		if err != nil {
			return nil, enc.ErrFailToParse{TypeNum: 0, Err: err}
		}
		l, err = enc.ReadTLNum(reader)
		if err != nil {
			return nil, enc.ErrFailToParse{TypeNum: 0, Err: err}
		}

		err = nil
		if handled := false; true {
			switch typ {
			case 7:
				if true {
					handled = true
					handled_Name = true
					value.Name, err = enc.ReadName(reader.Delegate(int(l)))
				}
			case 26:
				if true {
					handled = true
					handled_FinalBlockID = true
					value.FinalBlockID = make([]byte, l)
					_, err = io.ReadFull(reader, value.FinalBlockID)
				}
			case 62720:
				if true {
					handled = true
					handled_SegmentSize = true
					{
						tempVal := uint64(0)
						tempVal = uint64(0)
						{
							for i := 0; i < int(l); i++ {
								x := byte(0)
								x, err = reader.ReadByte()
								if err != nil {
									if err == io.EOF {
										err = io.ErrUnexpectedEOF
									}
									break
								}
								tempVal = uint64(tempVal<<8) | uint64(x)
							}
						}
						value.SegmentSize = &tempVal
					}
				}
			case 62722:
				if true {
					handled = true
					handled_Size = true
					{
						tempVal := uint64(0)
						tempVal = uint64(0)
						{
							for i := 0; i < int(l); i++ {
								x := byte(0)
								x, err = reader.ReadByte()
								if err != nil {
									if err == io.EOF {
										err = io.ErrUnexpectedEOF
									}
									break
								}
								tempVal = uint64(tempVal<<8) | uint64(x)
							}
						}
						value.Size = &tempVal
					}
				}
			case 62724:
				if true {
					handled = true
					handled_Mode = true
					{
						tempVal := uint64(0)
						tempVal = uint64(0)
						{
							for i := 0; i < int(l); i++ {
								x := byte(0)
								x, err = reader.ReadByte()
								if err != nil {
									if err == io.EOF {
										err = io.ErrUnexpectedEOF
									}
									break
								}
								tempVal = uint64(tempVal<<8) | uint64(x)
							}
						}
						value.Mode = &tempVal
					}
				}
			case 62726:
				if true {
					handled = true
					handled_Atime = true
					{
						tempVal := uint64(0)
						tempVal = uint64(0)
						{
							for i := 0; i < int(l); i++ {
								x := byte(0)
								x, err = reader.ReadByte()
								if err != nil {
									if err == io.EOF {
										err = io.ErrUnexpectedEOF
									}
									break
								}
								tempVal = uint64(tempVal<<8) | uint64(x)
							}
						}
						value.Atime = &tempVal
					}
				}
			case 62728:
				if true {
					handled = true
					handled_Btime = true
					{
						tempVal := uint64(0)
						tempVal = uint64(0)
						{
							for i := 0; i < int(l); i++ {
								x := byte(0)
								x, err = reader.ReadByte()
								if err != nil {
									if err == io.EOF {
										err = io.ErrUnexpectedEOF
									}
									break
								}
								tempVal = uint64(tempVal<<8) | uint64(x)
							}
						}
						value.Btime = &tempVal
					}
				}
			case 62730:
				if true {
					handled = true
					handled_Ctime = true
					{
						tempVal := uint64(0)
						tempVal = uint64(0)
						{
							for i := 0; i < int(l); i++ {
								x := byte(0)
								x, err = reader.ReadByte()
								if err != nil {
									if err == io.EOF {
										err = io.ErrUnexpectedEOF
									}
									break
								}
								tempVal = uint64(tempVal<<8) | uint64(x)
							}
						}
						value.Ctime = &tempVal
					}
				}
			case 62732:
				if true {
					handled = true
					handled_Mtime = true
					{
						tempVal := uint64(0)
						tempVal = uint64(0)
						{
							for i := 0; i < int(l); i++ {
								x := byte(0)
								x, err = reader.ReadByte()
								if err != nil {
									if err == io.EOF {
										err = io.ErrUnexpectedEOF
									}
									break
								}
								tempVal = uint64(tempVal<<8) | uint64(x)
							}
						}
						value.Mtime = &tempVal
					}
				}
			case 62734:
				if true {
					handled = true
					handled_ObjectType = true
					{
						var builder strings.Builder
						_, err = io.CopyN(&builder, reader, int64(l))
						if err == nil {
							tempStr := builder.String()
							value.ObjectType = &tempStr
						}
					}
				}
			default:
				if !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
					return nil, enc.ErrUnrecognizedField{TypeNum: typ}
				}
				handled = true
				err = reader.Skip(int(l))
			}
			if err == nil && !handled {
			}
			if err != nil {
				return nil, enc.ErrFailToParse{TypeNum: typ, Err: err}
			}
		}
	}

	startPos = reader.Pos()
	err = nil

	if !handled_Name && err == nil {
		value.Name = nil
	}
	if !handled_FinalBlockID && err == nil {
		value.FinalBlockID = nil
	}
	if !handled_SegmentSize && err == nil {
		value.SegmentSize = nil
	}
	if !handled_Size && err == nil {
		value.Size = nil
	}
	if !handled_Mode && err == nil {
		value.Mode = nil
	}
	if !handled_Atime && err == nil {
		value.Atime = nil
	}
	if !handled_Btime && err == nil {
		value.Btime = nil
	}
	if !handled_Ctime && err == nil {
		value.Ctime = nil
	}
	if !handled_Mtime && err == nil {
		value.Mtime = nil
	}
	if !handled_ObjectType && err == nil {
		value.ObjectType = nil
	}

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (value *MetaData) Encode() enc.Wire {
	encoder := MetaDataEncoder{}
	encoder.Init(value)
	return encoder.Encode(value)
}

func (value *MetaData) Bytes() []byte {
	return value.Encode().Join()
}

func ParseMetaData(reader enc.ParseReader, ignoreCritical bool) (*MetaData, error) {
	context := MetaDataParsingContext{}
	context.Init()
	return context.Parse(reader, ignoreCritical)
}
