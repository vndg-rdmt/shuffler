// shuffler source code
// Author (c) 2023 Belousov Daniil

package common

import (
	"errors"
	"fmt"
)

func MustExecute[T any](f func() (T, error)) T {
	if res, err := f(); err != nil {
		panic(err)
	} else {
		return res
	}
}

func Must[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}

	return res
}

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func NewErr(format string, a ...any) error {
	return errors.New(
		fmt.Sprintf(format, a...),
	)
}
