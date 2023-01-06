package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/prateeknischal/cryptopals/set1"
)

func main() {
	var set, challenge int
	flag.IntVar(&set, "set", 0, "Set number to execute")
	flag.IntVar(&challenge, "challenge", 0, "Challenge number to execute")

	flag.Parse()

	if set == 0 || challenge == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if set == 1 {
		if challenge == 1 {
			var s string
			fmt.Scanf("%s", &s)
			v, _ := set1.ToBase64(s)
			fmt.Println(v)
			return
		}

		if challenge == 2 {
			var a, b string
			fmt.Scanf("%s %s", &a, &b)
			v, _ := set1.Xor(a, b)
			fmt.Println(v)
		}

		if challenge == 3 {
			var a string
			fmt.Scanf("%s", &a)

			res, _ := set1.SingleByteXor(a)
			fmt.Println(res)
		}

		if challenge == 4 {
			var f string
			var r io.Reader
			var err error

			fmt.Scanf("%s", &f)

			if r, err = os.Open(f); err != nil {
				os.Exit(1)
			}

			fmt.Println(set1.FindXoredString(r))
		}

		if challenge == 5 {
			p := "Burning 'em, if you ain't quick and nimble" + "\n" + "I go crazy when I hear a cymbal"
			k := "ICE"

			res, err := set1.RepeatingKeyXor(p, k)
			if err != nil {
				panic(err)
			}

			fmt.Println(res)
		}
	}
}
