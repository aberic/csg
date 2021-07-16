package utils

import (
	"encoding/binary"
	"errors"
)

func U64ToBytes(i uint64) []byte {
	var b = make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}

func BytesToU64(b []byte) (uint64, error) {
	if len(b) != 8 {
		return 0, errors.New("invalid uint64 bytes")
	}

	return binary.LittleEndian.Uint64(b), nil
}

func I64ToBytes(i int64) []byte {
	var b = make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func BytesToI64(b []byte) (int64, error) {
	if len(b) != 8 {
		return 0, errors.New("invalid int64 bytes")
	}

	return int64(binary.LittleEndian.Uint64(b)), nil
}

func IntToBytes(i int) []byte {
	var b = make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(i))
	return b
}

func BytesToInt(b []byte) (int, error) {
	if len(b) != 4 {
		return 0, errors.New("invalid int bytes")
	}

	return int(binary.LittleEndian.Uint32(b)), nil
}
