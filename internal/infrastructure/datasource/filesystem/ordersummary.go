package filesystem

import (
	"fmt"
	"hackbar-copilot/internal/domain/ordersummary"
	"io/fs"
	"iter"
	"slices"
	"strings"

	"github.com/tingtt/options"
)

var _ ordersummary.Repository = (*orderSummaryRepository)(nil)

type orderSummaryRepository struct {
	fs *filesystem
}

// Latest implements ordersummary.Repository.
func (o *orderSummaryRepository) Latest(optionAppliers ...options.Applier[ordersummary.ListerOption]) iter.Seq2[ordersummary.Summary, error] {
	option := options.Create(optionAppliers...)

	return func(yield func(ordersummary.Summary, error) bool) {
		dirs, err := fs.ReadDir(o.fs.read, ".")
		if err != nil {
			if !yield(ordersummary.Summary{}, err) {
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
					if !yield(ordersummary.Summary{}, err) {
						return
					}
				}
				defer file.Close()

				summary := ordersummary.Summary{}
				err = loadFromToml(o.fs.read, filename, "ordersummary", &summary)
				if err != nil {
					if !yield(ordersummary.Summary{}, err) {
						return
					}
				}

				if option.Since != nil && summary.Timestamp.Before(*option.Since) {
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
func (o *orderSummaryRepository) Save(s ordersummary.Summary) error {
	filename := fmt.Sprintf("6_orders_summarized_%d.toml", s.Timestamp.Unix())
	return o.fs.saveFile(filename, map[string]interface{}{"ordersummary": s})
}
