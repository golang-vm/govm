
package encoding

import (
	"io"
)

const (
	NEG1 = ^(uint64)(0)
)

func Uint24ToInt24(v uint32)(int32){
	if v & (1 << 23) != 0 {
		v |= 0xff << 24
	}
	return (int32)(v)
}

func Uint8ToBytes(v uint8, buf []byte)([]byte){
	if buf == nil {
		buf = make([]byte, 1)
	}
	buf[0] = (byte)(v)
	return buf
}

func Uint16ToBytes(v uint16, buf []byte)([]byte){
	if buf == nil {
		buf = make([]byte, 2)
	}
	buf[0] = (byte)(v >> 8)
	buf[1] = (byte)(v)
	return buf
}

func Uint24ToBytes(v uint32, buf []byte)([]byte){
	if buf == nil {
		buf = make([]byte, 3)
	}
	buf[0] = (byte)(v >> 16)
	buf[1] = (byte)(v >> 8)
	buf[2] = (byte)(v)
	return buf
}

func Uint32ToBytes(v uint32, buf []byte)([]byte){
	if buf == nil {
		buf = make([]byte, 4)
	}
	buf[0] = (byte)(v >> 24)
	buf[1] = (byte)(v >> 16)
	buf[2] = (byte)(v >> 8)
	buf[3] = (byte)(v)
	return buf
}

func Uint64ToBytes(v uint64, buf []byte)([]byte){
	if buf == nil {
		buf = make([]byte, 8)
	}
	buf[0] = (byte)(v >> 56)
	buf[1] = (byte)(v >> 48)
	buf[2] = (byte)(v >> 40)
	buf[3] = (byte)(v >> 32)
	buf[4] = (byte)(v >> 24)
	buf[5] = (byte)(v >> 16)
	buf[6] = (byte)(v >> 8)
	buf[7] = (byte)(v)
	return buf
}

func BytesToUint8(buf []byte)(uint8){
	_ = buf[0]
	return (uint8)(buf[0])
}

func BytesToUint16(buf []byte)(uint16){
	_ = buf[1]
	return ((uint16)(buf[0]) << 8) | (uint16)(buf[1])
}

func BytesToUint24(buf []byte)(uint32){
	_ = buf[2]
	return ((uint32)(buf[0]) << 16) | ((uint32)(buf[1]) << 8) | (uint32)(buf[2])
}

func BytesToUint32(buf []byte)(uint32){
	_ = buf[3]
	return ((uint32)(buf[0]) << 24) | ((uint32)(buf[1]) << 16) | ((uint32)(buf[2]) << 8) | (uint32)(buf[3])
}

func BytesToUint64(buf []byte)(uint64){
	_ = buf[7]
	return ((uint64)(buf[0]) << 56) | ((uint64)(buf[1]) << 48) | ((uint64)(buf[2]) << 40) | ((uint64)(buf[3]) << 32) |
					((uint64)(buf[4]) << 24) | ((uint64)(buf[5]) << 16) | ((uint64)(buf[6]) << 8) | (uint64)(buf[7])
}

func ReadUint8(r io.Reader)(v uint8, err error){
	var buf [1]byte
	if _, err = io.ReadFull(r, buf[:]); err != nil {
		return
	}
	v = buf[0]
	return
}

func ReadUint16(r io.Reader)(v uint16, err error){
	var buf [2]byte
	if _, err = io.ReadFull(r, buf[:]); err != nil {
		return
	}
	v = BytesToUint16(buf[:])
	return
}

func ReadUint24(r io.Reader)(v uint32, err error){
	var buf [3]byte
	if _, err = io.ReadFull(r, buf[:]); err != nil {
		return
	}
	v = BytesToUint24(buf[:])
	return
}

func ReadUint32(r io.Reader)(v uint32, err error){
	var buf [4]byte
	if _, err = io.ReadFull(r, buf[:]); err != nil {
		return
	}
	v = BytesToUint32(buf[:])
	return
}

func ReadUint64(r io.Reader)(v uint64, err error){
	var buf [8]byte
	if _, err = io.ReadFull(r, buf[:]); err != nil {
		return
	}
	v = BytesToUint64(buf[:])
	return
}

func ReadByte2(r io.Reader)(a, b byte, err error){
	var buf [2]byte
	if _, err = io.ReadFull(r, buf[:]); err != nil {
		return
	}
	a, b = buf[0], buf[1]
	return
}

func ReadByte3(r io.Reader)(a, b, c byte, err error){
	var buf [3]byte
	if _, err = io.ReadFull(r, buf[:]); err != nil {
		return
	}
	a, b, c = buf[0], buf[1], buf[2]
	return
}

func ReadByte4(r io.Reader)(a, b, c, d byte, err error){
	var buf [4]byte
	if _, err = io.ReadFull(r, buf[:]); err != nil {
		return
	}
	a, b, c, d = buf[0], buf[1], buf[2], buf[3]
	return
}

func ReadByteAndUint16(r io.Reader)(a byte, v uint16, err error){
	var v0 uint32
	if v0, err = ReadUint24(r); err != nil {
		return
	}
	a, v = (byte)(v0 >> 16), (uint16)(v0)
	return
}

func MustReadUint8(r io.Reader)(v uint8){
	var err error
	v, err = ReadUint8(r)
	if err != nil { panic(err) }
	return
}

func MustReadUint16(r io.Reader)(v uint16){
	var err error
	v, err = ReadUint16(r)
	if err != nil { panic(err) }
	return
}

func MustReadUint24(r io.Reader)(v uint32){
	var err error
	v, err = ReadUint24(r)
	if err != nil { panic(err) }
	return
}

func MustReadUint32(r io.Reader)(v uint32){
	var err error
	v, err = ReadUint32(r)
	if err != nil { panic(err) }
	return
}

func MustReadUint64(r io.Reader)(v uint64){
	var err error
	v, err = ReadUint64(r)
	if err != nil { panic(err) }
	return
}
