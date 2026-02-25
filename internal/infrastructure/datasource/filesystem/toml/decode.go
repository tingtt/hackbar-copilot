package toml

import (
	"io"

	"github.com/BurntSushi/toml"
)

func Decode(r io.Reader, i any) error {
	_, err := toml.NewDecoder(r).Decode(i)
	return err
}
