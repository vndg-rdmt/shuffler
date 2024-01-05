// shuffler source code
// Author (c) 2023 Belousov Daniil

package main

import (
	"fmt"
	"os"
	"replicator/internal/processor"
)

func main() {
	a := readArgs()

	if a.ConfigPath != "" {
		a.Config.MustLoadFile(a.ConfigPath)
	}

	p := processor.New(
		a.Config.Processor,
		logger(),
	)

	ck, err := p.LoadSources()
	if err != nil {
		fmt.Printf(fmt.Sprintf("errors loading sources: %s", err.Error()))
		os.Exit(2)
	}

	if err := p.Generate(ck); err != nil {
		fmt.Printf(fmt.Sprintf("errors generatin image: %s", err.Error()))
		os.Exit(2)
	}

	fmt.Println("image generated")
}
