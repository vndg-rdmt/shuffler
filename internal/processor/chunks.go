// shuffler source code
// Author (c) 2023 Belousov Daniil

package processor

import (
	"math/rand"
	"replicator/pkg/common"
	"time"
)

func (self *implementation) NewChunkSize(ck []Chunk) (ChunkSize, error) {
	s, err := self.loadedChunkSize(ck)
	if err != nil {
		return ChunkSize{}, err
	}

	return s, self.verifyChunkSizes(ck, s)
}

// Returns chunk size of a sources
func (self *implementation) loadedChunkSize(ck []Chunk) (ChunkSize, error) {
	if len(ck) < 1 {
		return ChunkSize{}, common.NewErr("not resoure chunks are loaded")
	}

	return ChunkSize{
		W: ck[0].Image.Bounds().Dx(),
		H: ck[0].Image.Bounds().Dy(),
	}, nil
}

func (self *implementation) verifyChunkSizes(ck []Chunk, s ChunkSize) error {
	for _, v := range ck {
		if v.Image.Bounds().Dx() != s.W || v.Image.Bounds().Dy() != s.H {
			return common.NewErr(
				"source image chunks shoud have the same sizes, image '%s' width and height are %d %d, but expected %d and %d",
				v.Name,
				v.Image.Bounds().Dx(),
				v.Image.Bounds().Dy(),
				s.W,
				s.H,
			)
		}
	}

	return nil
}

func (sefl *implementation) randomSelector(ck []Chunk) func() Chunk {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func() Chunk {
		return ck[r.Intn(len(ck))]
	}
}
