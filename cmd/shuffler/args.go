// shuffler source code
// Author (c) 2023 Belousov Daniil

package main

import (
	"flag"
	"fmt"
	"os"
	"replicator/internal/config"
)

const manual = `
 -- manual -------------------------------------------------------------------
 | 'replicator' generates images by randomly selecting chunks from           |
 | the loaded 'chunks' from the source directory, which are just also images |
 | of the (important) same size, png shoud be used.                          |
 -----------------------------------------------------------------------------
`

type cliarguments struct {
	ConfigPath string
	Log        bool
	Config     config.Config
}

func (self *cliarguments) ConnectParser(p *flag.FlagSet, helpFlag *bool) {
	p.BoolVar(helpFlag, "help", false, "Print out this help message")
	p.StringVar(&self.ConfigPath, "config", "", "Flags formated in a file config")
	p.BoolVar(&self.Log, "log", self.Log, "Include logging to stdout")
	p.UintVar(&self.Config.Processor.Width, "width", self.Config.Processor.Width, "Output image width")
	p.UintVar(&self.Config.Processor.Height, "height", self.Config.Processor.Height, "Output image height")
	// p.UintVar(&self.Config.Processor.Amount, "amount", self.Config.Processor.Amount, "Amount of generated images")
	p.StringVar(&self.Config.Processor.Source, "source", self.Config.Processor.Source, "Source directory with the chunks")
	p.StringVar(&self.Config.Processor.Output, "output", self.Config.Processor.Output, "Output filename")
	// p.StringVar(&self.Config.Processor.NamingTemplate, "naming-template", self.Config.Processor.NamingTemplate, "Naming template of the generated images")
}

func FlagsParser() *flag.FlagSet {

	p := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	p.SetOutput(os.Stdout)
	p.Usage = func() {
		fmt.Println(manual)
		fmt.Fprintf(os.Stdout, "-- usage: %s [flags] -------------------------------------------------\n\n", os.Args[0])
		p.PrintDefaults()
	}

	return p
}

func readArgs() cliarguments {

	helpFlag := false
	self := cliarguments{
		Config:     config.Default(),
		ConfigPath: "",
		Log:        true,
	}

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
