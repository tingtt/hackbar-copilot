package toml

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/tingtt/options"
)

type Option struct {
	Indent *string
}

type Applier = options.Applier[Option]

type EncodeOption struct {
	Indent *string
}

func Encode(w io.Writer, i any, _options ...Applier) error {
	option := options.Create(_options...)
	e := toml.NewEncoder(w)
	if option.Indent != nil {
		e.Indent = *option.Indent
	}
	return e.Encode(i)
}

func WithIndent(indent string) Applier {
	return func(o *Option) { o.Indent = &indent }
}
