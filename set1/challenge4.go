package set1

import (
	"fmt"
	"io"
)

func FindXoredString(r io.Reader) string {
	var line string
	var maxScore = 0.0
	var res string
	for {
		n, err := fmt.Fscanf(r, "%s", &line)
		if err != nil || n == 0 {
			break
		}
		s, score := SingleByteXor(line)
		if score > maxScore {
			maxScore = score
			res = s
		}
	}

	return res
}
