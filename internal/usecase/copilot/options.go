package copilot

import "github.com/tingtt/options"

type queryOption struct {
	filterByName *string
}

type QueryOptionApplier = options.Applier[queryOption]

func WithFilterByName(name string) QueryOptionApplier {
	return func(o *queryOption) {
		o.filterByName = &name
	}
}
