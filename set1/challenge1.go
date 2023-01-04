package set1

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var b64Table = []rune{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
	'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
	'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
	'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/',
}

var mask byte = 0b111111

func HexToBytes(s string) ([]byte, error) {
	if len(s)%2 != 0 {
		return nil, errors.New("Invalid hex string, odd bytes")
	}

	var b []byte

	for i := 0; i < len(s); {
		n, e := strconv.ParseInt(s[i:i+2], 16, 8)
		if e != nil {
			return nil, errors.Wrap(e, "Invalid byte")
		}

		b = append(b, byte(n&0xff))
		i += 2
	}

	return b, nil
}

func ToBase64(h string) (string, error) {
	var b []byte
	var err error

	if b, err = HexToBytes(h); err != nil {
		return "", err
	}

	var res strings.Builder
	var i = 0

	for i < len(b) {
		// first 6 bits from b[i]
		firstByte := byte(b[i]) >> 0x2

		// last 2 bits from b[i] left || first 4 bits from b[i + 1]
		secondByte := ((byte(b[i]) & 0x3) << 0x4) | (byte(b[i+1]) >> 0x4)

		// last 4 bits from b[i + 1] || first 2 bits from b[i + 2]
		thirdByte := ((byte(b[i+1]) & 0xf) << 0x4) | (byte(b[i+2]) >> 0x6)

		// last 6 bits from b[i + 2]
		fouthByte := byte(b[i+2]) & mask

		res.WriteRune(b64Table[firstByte])
		res.WriteRune(b64Table[secondByte])
		res.WriteRune(b64Table[thirdByte])
		res.WriteRune(b64Table[fouthByte])

		i += 3
	}

	if len(b)%3 == 1 {
		// first 6 bits from b[i]
		firstByte := byte(b[i]) >> 0x2

		// last 2 bits from b[i] left shifted by 4 for total of 6 bits
		secondByte := (byte(b[i]) & 0x3) << 0x4

		res.WriteRune(b64Table[firstByte])
		res.WriteRune(b64Table[secondByte])
		res.WriteRune('=')
		res.WriteRune('=')
	} else if len(b)%3 == 2 {
		// first 6 bits from b[i]
		firstByte := byte(b[i]) >> 0x2

		// last 2 bits from b[i] || first 4 bits from b[i + 1]
		secondByte := ((byte(b[i]) & 0x3) << 0x4) | (byte(b[i+1]) >> 0x4)

		// last 4 bits from b[i + 1] left shifted by 2 for total of 6 bits
		thirdByte := (byte(b[i+1]) & 0xf) << 0x2

		res.WriteRune(b64Table[firstByte])
		res.WriteRune(b64Table[secondByte])
		res.WriteRune(b64Table[thirdByte])
		res.WriteRune('=')
	}

	return res.String(), nil
}
