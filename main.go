package main

import (
	"os"
	"oslinfra/cli"
)

func main() {
	os.Exit(cli.Run(os.Args))
}
