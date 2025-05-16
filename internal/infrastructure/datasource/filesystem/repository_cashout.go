package filesystem

import (
	"fmt"
	"hackbar-copilot/internal/domain/cashout"
	"hackbar-copilot/internal/usecase/cashier"
	"io/fs"
	"iter"
	"slices"
	"strings"
	"sync"

	"github.com/tingtt/options"
)

var _ cashier.CashoutSaveLister = (*cashoutRepository)(nil)

type cashoutRepository struct {
	fs    *filesystem
	mutex *sync.RWMutex
}

const (
	cashoutFilenamePrefix = "8_cashout_"
	cashoutTomlKey        = "cashout"
)

// Latest implements cashier.CashoutSaveLister.
func (r *cashoutRepository) Latest(optionAppliers ...options.Applier[cashier.ListerOption]) iter.Seq2[cashout.Cashout, error] {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	option := options.Create(optionAppliers...)

	return func(yield func(cashout.Cashout, error) bool) {
		for filename, err := range iterateLatestCashoutTomlFiles(r.fs.read) {
			summary := cashout.Cashout{}
			err = loadFromToml(r.fs.read, filename, cashoutTomlKey, &summary)
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

func iterateLatestCashoutTomlFiles(fsR fs.FS) iter.Seq2[string /* filename */, error] {
	return func(yield func(string /* filename */, error) bool) {
		dirs, err := fs.ReadDir(fsR, ".")
		if err != nil {
			if !yield("", err) {
				return
			}
		}
		slices.Reverse(dirs)
		for _, d := range dirs {
			if d.IsDir() {
				continue
			}

			filename := d.Name()
			if strings.Contains(filename, cashoutFilenamePrefix) && strings.HasSuffix(filename, ".toml") {
				if !yield(filename, nil) {
					return
				}
			}
		}
	}
}

// Save implements cashier.CashoutSaveLister.
func (r *cashoutRepository) Save(cashout cashout.Cashout) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	filename := fmt.Sprintf(cashoutFilenamePrefix+"%d.toml", cashout.Timestamp.Unix())
	return r.fs.saveFile(filename, map[string]interface{}{cashoutTomlKey: cashout})
}
