package filesystem

import (
	"fmt"
	"hackbar-copilot/internal/domain/cashout"
	"io/fs"
	"iter"
	"slices"
	"strings"

	"github.com/tingtt/options"
)

var _ cashout.Repository = (*cashoutRepository)(nil)

type cashoutRepository struct {
	fs *filesystem
}

// Latest implements ordersummary.Repository.
func (o *cashoutRepository) Latest(optionAppliers ...options.Applier[cashout.ListerOption]) iter.Seq2[cashout.Cashout, error] {
	option := options.Create(optionAppliers...)

	return func(yield func(cashout.Cashout, error) bool) {
		dirs, err := fs.ReadDir(o.fs.read, ".")
		if err != nil {
			if !yield(cashout.Cashout{}, err) {
				return
			}
		}
		slices.Reverse(dirs)
		for _, d := range dirs {
			if d.IsDir() {
				continue
			}

			filename := d.Name()
			if strings.Contains(filename, "6_orders_summarized_") && strings.HasSuffix(filename, ".toml") {
				file, err := o.fs.read.Open(filename)
				if err != nil {
					if !yield(cashout.Cashout{}, err) {
						return
					}
				}
				defer file.Close()

				summary := cashout.Cashout{}
				err = loadFromToml(o.fs.read, filename, "ordersummary", &summary)
				if err != nil {
					if !yield(cashout.Cashout{}, err) {
						return
					}
				}

				if option.Since != nil && summary.Timestamp.Before(*option.Since) {
					break
				}
				if option.Until != nil && summary.Timestamp.After(*option.Until) {
					continue
				}
				if !yield(summary, nil) {
					return
				}
			}
		}
	}
}

// Save implements ordersummary.Repository.
func (o *cashoutRepository) Save(s cashout.Cashout) error {
	filename := fmt.Sprintf("6_orders_summarized_%d.toml", s.Timestamp.Unix())
	return o.fs.saveFile(filename, map[string]interface{}{"ordersummary": s})
}
