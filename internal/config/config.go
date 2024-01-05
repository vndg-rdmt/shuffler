// shuffler source code
// Author (c) 2023 Belousov Daniil

package config

import (
	"os"
	"replicator/pkg/common"

	"github.com/goccy/go-json"
)

func (self *Config) MustLoadFile(f string) {
	b := common.Must(os.ReadFile(f))
	common.NoErr(json.Unmarshal(b, self))
}
