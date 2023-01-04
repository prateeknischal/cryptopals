package set1

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func BytesToHex(b []byte) string {
	var s strings.Builder

	for i := 0; i < len(b); i++ {
		s.WriteString(fmt.Sprintf("%x", b[i]))
	}

	return s.String()
}

func Xor(a, b string) (string, error) {
	var aBytes, bBytes []byte
	var err error

	if len(a) != len(b) {
		return "", errors.New("hex strings are not of same length")
	}

	if aBytes, err = HexToBytes(a); err != nil {
		return "", err
	}

	if bBytes, err = HexToBytes(b); err != nil {
		return "", err
	}

	xorBytes := make([]byte, len(aBytes))

	for i := 0; i < len(aBytes); i++ {
		xorBytes[i] = aBytes[i] ^ bBytes[i]
	}

	return BytesToHex(xorBytes), nil
}
