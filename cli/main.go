package cli

import (
	"flag"
	"fmt"
)

func Run(args []string) int {
	flag.Usage = func() {
		fmt.Printf("Usage of %s\n", args[0])
		// TODO: add usage
		flag.PrintDefaults()
	}
	flag.Parse()
	return 0
}
