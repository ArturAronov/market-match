package utils

import "encoding/base64"

func Base64ToByte(input string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(input)
}
