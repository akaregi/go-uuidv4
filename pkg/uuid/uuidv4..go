package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

type UUID [16]byte

func NewUUID() (UUID, error) {
	var uuid UUID

	_, err := io.ReadFull(rand.Reader, uuid[:])

	if err != nil {
		return uuid, err
	}

	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return uuid, nil
}

func Stringify(src UUID) string {
	// Expected [36]byte
	// [len([16]byte <- UUID]) * 2]byte + [4]byte
	buffer := make([]byte, hex.EncodedLen(len(src))+4)

	// Insert the hyphen
	// RRRRRRRR |  -  |   RRRR   |  -   |   4RRR    |  -   |   rRRR    |  -   | RRRRRRRRRRRR
	// [0]~~[7] | [8] | [9]~[12] | [13] | [14]~[17] | [18] | [19]~[22] | [23] | [24]~~~~[35]
	buffer[8] = '-'
	buffer[13] = '-'
	buffer[18] = '-'
	buffer[23] = '-'

	// Insert encoded bytes
	hex.Encode(buffer[:8], src[:4])
	hex.Encode(buffer[9:13], src[4:6])
	hex.Encode(buffer[14:18], src[6:8])
	hex.Encode(buffer[19:23], src[8:10])
	hex.Encode(buffer[24:], src[10:])

	return string(buffer)
}
