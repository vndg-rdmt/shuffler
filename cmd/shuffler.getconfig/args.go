// shuffler source code
// Author (c) 2023 Belousov Daniil

package main

import (
	"flag"
	"fmt"
	"os"
	"replicator/internal/config"
)

type cliarguments struct {
	Output     string
	ConfigType string
}

func defaults() cliarguments {
	return cliarguments{
		Output:     "replicator.config",
		ConfigType: "json",
	}
}

func (self *cliarguments) ConnectParser(p *flag.FlagSet, helpFlag *bool) {
	p.BoolVar(helpFlag, "help", false, "Print out this help message")
	p.StringVar(&self.Output, "output", self.Output, "Output name of the file")
	p.StringVar(&self.ConfigType, "type", self.ConfigType,
		fmt.Sprintf("type of the output configuration file from: %s", config.ConfigTypes()),
	)
}

func FlagsParser() *flag.FlagSet {
	p := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	p.SetOutput(os.Stdout)
	p.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
		p.PrintDefaults()
	}

	return p
}

func readArgs() cliarguments {

	helpFlag := false
	self := defaults()

	p := FlagsParser()
	self.ConnectParser(p, &helpFlag)

	if err := p.Parse(os.Args[1:]); err != nil {
		p.Usage()
		os.Exit(1)
	}

	if helpFlag {
		p.Usage()
		os.Exit(0)
	}

	return self
}
