// shuffler source code
// Author (c) 2023 Belousov Daniil

package config

import (
	"fmt"
	"reflect"

	"github.com/goccy/go-json"
	"github.com/goccy/go-yaml"
)

type configMarshaller func(v interface{}) ([]byte, error)

func Marshallers() map[string]configMarshaller {
	return map[string]configMarshaller{
		"json": json.Marshal,
		"yaml": yaml.Marshal,
		"yml":  yaml.Marshal,
	}
}

func ConfigTypes() string {
	return fmt.Sprintf(
		"%v", reflect.ValueOf(Marshallers()).MapKeys(),
	)
}
