// shuffler source code
// Author (c) 2023 Belousov Daniil

package config

import (
	"os"
	"replicator/internal/processor"
	"replicator/pkg/common"
)

type Config struct {
	Processor processor.Config
}

func Default() Config {
	return Config{
		Processor: processor.DefaultConfig(),
	}
}

func New(f string) Config {
	c := Default()
	if f != "" {
		c.MustLoadFile(f)
	}

	return c
}

func Extract(self Config, o, t string) error {
	m, ok := Marshallers()[t]
	if !ok {
		return common.NewErr(
			"unsupported configuration type to extract, available are: %s",
			ConfigTypes(),
		)
	}

	if b, err := m(self); err != nil {
		return err

	} else {
		return os.WriteFile(o, b, os.ModePerm)
	}
}
