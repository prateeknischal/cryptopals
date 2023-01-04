package main

import (
	"flag"
	"fmt"
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
			v, err := set1.ToBase64(s)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(v)
			return
		}
	}
}
