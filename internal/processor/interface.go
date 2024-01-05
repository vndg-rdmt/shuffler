// shuffler source code
// Author (c) 2023 Belousov Daniil

package processor

import (
	"image"

	"go.uber.org/zap"
)

type Proccesor interface {
	Generate(ck []Chunk) error
	LoadSources() ([]Chunk, error)
}

type Chunk struct {
	Name  string
	Image image.Image
}

type ChunkSize struct {
	W int
	H int
}

type Config struct {
	Source string `json:"source" yaml:"source"`
	Width  uint   `json:"width" yaml:"width"`
	Height uint   `json:"height" yaml:"height"`
	Output string `json:"output" yaml:"output"`
	// Amount uint   `json:"amount" yaml:"amount"`
	// NamingTemplate string `json:"naming_template" yaml:"naming_template"`
}

func DefaultConfig() Config {
	return Config{
		Width:  2,
		Height: 2,
		Output: "replicator.output",
	}
}

func New(conf Config, log *zap.Logger) Proccesor {
	return &implementation{
		config: conf,
		logger: log,
	}
}
