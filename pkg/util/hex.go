package util

import (
	"encoding/hex"
	"strings"
)

// delete the 0x from the front
func RemoveZeroHex(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return Hex2Bytes(s)
}

func Hex2Bytes(str string) []byte {
	h, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return h
}

func EncodeHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func EncodeHexWith0x(bytes []byte) string {
	return "0x" + EncodeHex(bytes)
}

func DecodeHexString(hexString string) ([]byte, error) {
	if strings.HasPrefix(hexString, "0x") || strings.HasPrefix(hexString, "0X") {
		hexString = hexString[2:]
	}
	if len(hexString)%2 != 0 {
		hexString = "0" + hexString
	}
	return hex.DecodeString(hexString)
}

func HasHexPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func IsHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}
