package set1

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

var letterDistribution = map[rune]float64{
	'E': 11.1607,
	'M': 3.0129,
	'A': 8.4966,
	'H': 3.0034,
	'R': 7.5809,
	'G': 2.4705,
	'I': 7.5448,
	'B': 2.0720,
	'O': 7.1635,
	'F': 1.8121,
	'T': 6.9509,
	'Y': 1.7779,
	'N': 6.6544,
	'W': 1.2899,
	'S': 5.7351,
	'K': 1.1016,
	'L': 5.4893,
	'V': 1.0074,
	'C': 4.5388,
	'X': 0.2902,
	'U': 3.6308,
	'Z': 0.2722,
	'D': 3.3844,
	'J': 0.1965,
	'P': 3.1671,
	'Q': 0.1962,
	'e': 11.1607,
	'm': 3.0129,
	'a': 8.4966,
	'h': 3.0034,
	'r': 7.5809,
	'g': 2.4705,
	'i': 7.5448,
	'b': 2.0720,
	'o': 7.1635,
	'f': 1.8121,
	't': 6.9509,
	'y': 1.7779,
	'n': 6.6544,
	'w': 1.2899,
	's': 5.7351,
	'k': 1.1016,
	'l': 5.4893,
	'v': 1.0074,
	'c': 4.5388,
	'x': 0.2902,
	'u': 3.6308,
	'z': 0.2722,
	'd': 3.3844,
	'j': 0.1965,
	'p': 3.1671,
	'q': 0.1962,
}

func distance(b []byte) float64 {
	var f = make(map[rune]float64)
	for _, x := range b {
		f[rune(x)] += 1.0 / float64(len(b))
	}

	dist := 0.0

	for k, v := range letterDistribution {
		d := f[k]

		dist += math.Abs(v - d)
	}

	return dist
}

func hexToString(b []byte) string {
	var s strings.Builder
	for i := 0; i < len(b); i++ {
		s.WriteRune(rune(b[i]))
	}

	return s.String()
}

func SingleByteXor(s string) string {
	var m = make(map[string]float64)

	for i := 1; i <= 255; i++ {
		cipherText, _ := HexToBytes(s)
		plainText := make([]byte, len(cipherText))

		for j := 0; j < len(cipherText); j++ {
			plainText[j] = cipherText[j] ^ byte(i)
		}

		hasNonPrintable := false
		for _, x := range plainText {
			if !unicode.IsGraphic(rune(x)) {
				hasNonPrintable = true
				break
			}
		}

		if hasNonPrintable {
			continue
		}

		m[hexToString(plainText)] = distance(plainText)
	}

	var res string
	var mx = math.Inf(1)
	for k, v := range m {
		if v < mx {
			res = k
			mx = v
		}
	}

	fmt.Println(mx)
	return res
}
