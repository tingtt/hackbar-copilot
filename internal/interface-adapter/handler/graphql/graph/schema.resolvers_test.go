package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolver_Mutation(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()
		assert.NotNil(t, (&Resolver{}).Mutation())
	})
}

func TestResolver_Query(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()
		assert.NotNil(t, (&Resolver{}).Query())
	})
}
