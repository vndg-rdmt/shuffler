// shuffler source code
// Author (c) 2023 Belousov Daniil

package processor

import (
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"replicator/pkg/common"

	"go.uber.org/zap"
)

type implementation struct {
	config Config
	logger *zap.Logger
}

func (self *implementation) LoadSources() ([]Chunk, error) {
	self.logger.Info(
		"loading resources",
		zap.String("from", self.config.Source),
	)

	i, err := os.Stat(self.config.Source)
	if err != nil {
		self.logger.Error(
			"cannot get fs stat of the source",
			zap.String("source", self.config.Source),
			zap.Error(err),
		)
		return nil, err
	}

	if !i.IsDir() {
		self.logger.Error(
			"source is not a directory",
			zap.String("source", self.config.Source),
		)
		return nil, common.NewErr("sources '%s' directory not found", i.Name())
	}

	self.logger.Info(
		"reading sources directory",
		zap.String("source", self.config.Source),
	)
	de, err := os.ReadDir(self.config.Source)
	if err != nil {
		self.logger.Error(
			"cannot read sources directory",
			zap.String("source", self.config.Source),
			zap.Error(err),
		)
		return nil, err
	}

	self.logger.Info(
		"loading chunks",
		zap.Int("amount", len(de)),
	)
	var chunksBuffer = make([]Chunk, 0, len(de))
	for index, v := range de {

		self.logger.Info(
			"loading source chunk",
			zap.Int("index", index),
		)
		i, err := v.Info()
		if err != nil {
			self.logger.Error(
				"cannot get fs info",
				zap.String("name", i.Name()),
				zap.Error(err),
			)
			return nil, err
		}

		if i.IsDir() {
			self.logger.Info(
				"skipping directory",
				zap.String("name", i.Name()),
			)
			continue
		}

		fl, err := os.Open(filepath.Join(self.config.Source, v.Name()))
		if err != nil {
			self.logger.Error(
				"cannot open file",
				zap.String("name", v.Name()),
				zap.Error(err),
			)
			return nil, err
		}
		defer fl.Close()

		img, _, err := image.Decode(fl)
		if err != nil {
			self.logger.Error(
				"cannot decode image, skipping",
				zap.String("name", v.Name()),
				zap.Error(err),
			)
			continue
		}

		chunksBuffer = append(
			chunksBuffer,
			Chunk{
				Name:  fl.Name(),
				Image: img,
			},
		)
		self.logger.Info(
			"chunk loaded",
			zap.Int("index", index),
			zap.String("name", i.Name()),
		)
	}

	self.logger.Info(
		"chunks loaded",
		zap.Int("amount", len(chunksBuffer)),
	)
	return chunksBuffer, nil
}

func (self *implementation) Generate(ck []Chunk) error {
	self.logger.Info(
		"image generation init",
	)

	sizes, err := self.NewChunkSize(ck)
	if err != nil {
		self.logger.Error(
			"cannot initiate chunk size",
			zap.Error(err),
		)
		return err
	}
	self.logger.Info(
		"chunk size initiated",
		zap.Int("width", sizes.W),
		zap.Int("height", sizes.H),
	)

	canvas := self.newCanvas(sizes)
	self.logger.Info(
		"canvas created",
		zap.Int("width", canvas.Rect.Dx()),
		zap.Int("height", canvas.Rect.Dy()),
	)

	self.drawImage(ck, sizes, canvas)
	self.logger.Info(
		"image drawn",
	)

	var imgname string = self.config.Output
	outputFile, err := os.Create(imgname)
	if err != nil {
		self.logger.Error(
			"cannot create output image",
			zap.String("name", imgname),
			zap.Error(err),
		)
		return err
	}
	defer outputFile.Close()

	if err := png.Encode(outputFile, canvas); err != nil {
		self.logger.Error(
			"cannot encode image to output file",
			zap.Error(err),
		)
		return err
	}

	self.logger.Info(
		"image created",
		zap.String("saved to", imgname),
	)
	return nil
}

func (self *implementation) newCanvas(s ChunkSize) *image.RGBA {
	return image.NewRGBA(
		image.Rect(
			0, 0, int(self.config.Width)*s.W, int(self.config.Height)*s.H,
		),
	)
}

func (self *implementation) drawImage(ck []Chunk, sizes ChunkSize, canvas *image.RGBA) {
	randomChunk := self.randomSelector(ck)

	for y := 0; y < int(self.config.Height); y++ {
		for x := 0; x < int(self.config.Width); x++ {
			c := randomChunk()
			draw.Draw(
				canvas,
				image.Rect(
					x*sizes.W,
					y*sizes.H,
					x*sizes.W+sizes.W,
					y*sizes.H+sizes.H,
				),
				c.Image,
				image.Point{0, 0},
				draw.Src,
			)
			self.logger.Info(
				"added chunk",
				zap.String("name", c.Name),
			)
		}
	}
}
