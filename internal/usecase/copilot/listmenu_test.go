package copilot

import (
	"hackbar-copilot/internal/domain/menu/menutest"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_copilot_ListMenu(t *testing.T) {
	t.Parallel()

	t.Run("will return sorted menu groups", func(t *testing.T) {
		t.Parallel()

		menuGroups := menutest.ExampleGroupsIter
		want := menutest.ExampleGroups
		sort.Slice(want, func(i, j int) bool {
			return want[i].Name < want[j].Name
		})

		menuSaveLister := new(MockMenu)
		menuSaveLister.On("All").Return(menuGroups, nil)

		c := &copilot{menu: menuSaveLister}

		got, err := c.ListMenu(SortMenuGroupByName())
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}
