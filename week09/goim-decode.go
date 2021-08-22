package week09

import (
	"encoding/binary"
	"errors"
)

func Decode(buf []byte) (IGoimDecoder, error) {
	decoder := &Decoder{}

	decoder.packetLength = binary.BigEndian.Uint32(buf[_packOffset : _packOffset+_packSize])
	decoder.headerLength = binary.BigEndian.Uint16(buf[_headerOffset : _headerOffset+_headerSize])
	decoder.version = binary.BigEndian.Uint16(buf[_verOffset : _verOffset+_verSize])
	decoder.operation = binary.BigEndian.Uint32(buf[_opOffset : _opOffset+_opSize])
	decoder.sequence = binary.BigEndian.Uint32(buf[_seqOffset : _seqOffset+_seqSize])

	if decoder.packetLength > _maxPackSize {
		return nil, errors.New("error package length")
	}

	if _bodyLength := int(decoder.packetLength - uint32(decoder.headerLength)); _bodyLength > 0 {
		decoder.body = buf[_bodyOffset : _bodyOffset+_bodyLength]
	}

	return decoder, nil
}

type IGoimDecoder interface {
	PacketLength() uint32
	HeaderLength() uint16
	Version() uint16
	Operation() uint32
	Sequence() uint32
	Body() []byte
}

func (d *Decoder) PacketLength() uint32 {
	return d.packetLength
}

func (d *Decoder) HeaderLength() uint16 {
	return d.headerLength
}

func (d *Decoder) Version() uint16 {
	return d.version
}

func (d *Decoder) Operation() uint32 {
	return d.operation
}

func (d *Decoder) Sequence() uint32 {
	return d.sequence
}

func (d *Decoder) Body() []byte {
	if d.body == nil {
		return []byte{}
	} else {
		return d.body
	}
}

// https://github.com/Terry-Mao/goim/blob/e742c99ad76e626d5f6df8b33bc47ca005501980/api/protocol/protocol.go#L25
const (
	// MaxBodySize max body size
	MaxBodySize = uint32(1 << 12)
)
const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_maxPackSize   = MaxBodySize + uint32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_bodyOffset   = _seqOffset + _seqSize
)

type Decoder struct {
	packetLength uint32
	headerLength uint16
	version      uint16
	operation    uint32
	sequence     uint32
	body         []byte
}
