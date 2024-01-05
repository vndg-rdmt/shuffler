// shuffler source code
// Author (c) 2023 Belousov Daniil

package main

import (
	"fmt"
	"os"
	"replicator/internal/config"
)

func main() {
	a := readArgs()
	if err := config.Extract(config.Default(), a.Output, a.ConfigType); err != nil {
		fmt.Printf("Error: %s\nUse %s flag to show help", err.Error(), "--help")
		os.Exit(2)
	}
}
