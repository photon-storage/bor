package main

import (
	"os"

	"github.com/photon-storage/bor/internal/cli"
	"github.com/photon-storage/bor/params"
)

func main() {
	params.UpdateBorInfo()
	os.Exit(cli.Run(os.Args[1:]))
}
