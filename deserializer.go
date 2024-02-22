package serializer

import (
	"encoding/binary"
	"fmt"
)

type Deserializer struct {
	Data  []byte
	Error error
}

func (s *Deserializer) ReadUint8() uint8 {
	if s.Error != nil {
		return 0
	}
	if len(s.Data) < 1 {
		s.Error = fmt.Errorf(getCaller() + " invalid length")
		return 0
	}
	b := s.Data[0]
	s.Data = s.Data[1:]
	return b
}
func (s *Deserializer) ReadUint16() uint16 {
	if s.Error != nil {
		return 0
	}
	if len(s.Data) < 2 {
		s.Error = fmt.Errorf(getCaller() + " invalid length")
		return 0
	}
	b := s.Data[:2]
	s.Data = s.Data[2:]
	return binary.LittleEndian.Uint16(b)
}
func (s *Deserializer) ReadUint32() uint32 {
	if s.Error != nil {
		return 0
	}
	if len(s.Data) < 4 {
		s.Error = fmt.Errorf(getCaller() + " invalid length")
		return 0
	}
	b := s.Data[:4]
	s.Data = s.Data[4:]
	return binary.LittleEndian.Uint32(b)
}
func (s *Deserializer) ReadUint64() uint64 {
	if s.Error != nil {
		return 0
	}
	if len(s.Data) < 8 {
		s.Error = fmt.Errorf(getCaller() + " invalid length")
		return 0
	}
	b := s.Data[:8]
	s.Data = s.Data[8:]
	return binary.LittleEndian.Uint64(b)
}
func (s *Deserializer) ReadUvarint() uint64 {
	if s.Error != nil {
		return 0
	}
	if len(s.Data) < 1 {
		s.Error = fmt.Errorf(getCaller() + " invalid length")
		return 0
	}
	d, x := binary.Uvarint(s.Data)
	if x < 0 {
		s.Error = fmt.Errorf(getCaller() + " invalid uvarint")
		return 0
	}
	s.Data = s.Data[x:]
	return d
}

func (s *Deserializer) ReadFixedByteArray(length int) []byte {
	if s.Error != nil {
		return []byte{}
	}
	if len(s.Data) < length {
		s.Error = fmt.Errorf(getCaller() + " invalid length")
		return []byte{}
	}
	b := s.Data[:length]
	s.Data = s.Data[length:]
	return b
}
func (s *Deserializer) ReadByteSlice() []byte {
	if s.Error != nil {
		return []byte{}
	}
	if len(s.Data) < 1 {
		s.Error = fmt.Errorf(getCaller() + " invalid length")
		return []byte{}
	}
	length, read := binary.Uvarint(s.Data)
	if read < 0 {
		s.Error = fmt.Errorf(getCaller() + " invalid uvarint length")
		return []byte{}
	}
	s.Data = s.Data[read:]
	if len(s.Data) < int(length) {
		s.Error = fmt.Errorf(getCaller() + " invalid binary length")
		return []byte{}
	}

	b := s.Data[:length]
	s.Data = s.Data[length:]
	return b
}
func (s *Deserializer) ReadString() string {
	return string(s.ReadByteSlice())
}

func (s *Deserializer) ReadBool() bool {
	if s.Error != nil {
		return false
	}
	if len(s.Data) < 1 {
		s.Error = fmt.Errorf(getCaller() + " invalid length")
		return false
	}
	b := s.Data[0]
	s.Data = s.Data[1:]
	if b == 1 {
		return true
	} else if b == 0 {
		return false
	} else {
		s.Error = fmt.Errorf(getCaller() + " invalid boolean value")
		return false
	}
}
