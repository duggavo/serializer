// Package serializer implements a simple and useful tool
// for serialization/deserialization of binary data
package serializer

import (
	"encoding/binary"
	"math/big"
	"runtime"
	"strconv"
	"strings"
)

type Serializer struct {
	Data   []byte
	Endian binary.AppendByteOrder
}

func (s *Serializer) AddUint8(n uint8) {
	s.Data = append(s.Data, n)
}
func (s *Serializer) AddUint16(n uint16) {
	s.Data = s.Endian.AppendUint16(s.Data, n)
}
func (s *Serializer) AddUint32(n uint32) {
	s.Data = s.Endian.AppendUint32(s.Data, n)
}
func (s *Serializer) AddUint64(n uint64) {
	s.Data = s.Endian.AppendUint64(s.Data, n)
}

func (s *Serializer) AddUvarint(n uint64) {
	s.Data = binary.AppendUvarint(s.Data, n)
}

func (s *Serializer) AddFixedByteArray(a []byte, length int) {
	if len(a) != length {
		panic("invalid length")
	}
	s.Data = append(s.Data, a...)
}
func (s *Serializer) AddByteSlice(a []byte) {
	s.Data = append(binary.AppendUvarint(s.Data, uint64(len(a))), a...)
}
func (s *Serializer) AddString(a string) {
	s.AddByteSlice([]byte(a))
}
func (s *Serializer) AddBigInt(n *big.Int) {
	bin := n.Bytes()

	s.Data = append(binary.AppendUvarint(s.Data, uint64(len(bin))), bin...)
}
func (s *Serializer) AddBool(b bool) {
	if b {
		s.Data = append(s.Data, 0)
	} else {
		s.Data = append(s.Data, 1)
	}
}

func getCaller() string {
	_, file, line, _ := runtime.Caller(2)
	fileSpl := strings.Split(file, "/")
	return fileSpl[len(fileSpl)-1] + ":" + strconv.FormatInt(int64(line), 10)
}
