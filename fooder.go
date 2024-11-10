package main

import (
	"os"

	"github.com/happymanju/fooder/cli"
)

func main() {
	os.Exit(cli.CLI(os.Args[1:]))
}
