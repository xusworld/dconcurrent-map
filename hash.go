package cmap

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
)

func hash(key interface{}) uint32 {
	buff, err := toBytes(key)
	if err != nil {
		panic("toBytes() error")
		return 0
	}
	return crc32.ChecksumIEEE(buff)
}

func toBytes(key interface{}) ([]byte, error) {
	bs := make([]byte, 8)
	buff := bytes.NewBuffer(bs)
	switch v := key.(type) {
	case *string:
		return []byte(*v), nil
	case string:
		return []byte(v), nil
	case *bool:
		if *v {
			buff.WriteByte(byte(1))
		} else {
			buff.WriteByte(byte(0))
		}
		return buff.Bytes()[:1], nil
	case bool:
		if v {
			buff.WriteByte(byte(1))
		} else {
			buff.WriteByte(byte(0))
		}
		return buff.Bytes()[:1], nil
	case *int8:
		buff.WriteByte(byte(*v))
		return buff.Bytes()[:1], nil
	case int8:
		buff.WriteByte(byte(v))
		return buff.Bytes()[:1], nil
	case *uint8:
		buff.WriteByte(*v)
		return buff.Bytes()[:1], nil
	case uint8:
		buff.WriteByte(v)
		return buff.Bytes()[:1], nil
	case *int16:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:2], nil
	case int16:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:2], nil
	case *int32:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:4], nil
	case int32:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:4], nil
	case *int64:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:8], nil
	case int64:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:8], nil
	default:
		return nil, errInvalidKeyType
	}
}
