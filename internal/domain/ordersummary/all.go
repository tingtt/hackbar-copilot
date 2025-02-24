package ordersummary

import (
	"iter"
)

// All implements SummarizeLister.
// Subtle: this method shadows the method (Repository).All of summarizeLister.Repository.
func (s *summarizeLister) All() iter.Seq2[Summary, error] {
	return func(yield func(Summary, error) bool) {
		panic("unimplemented")
	}
}
