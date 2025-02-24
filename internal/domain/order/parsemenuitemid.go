package order

import (
	"fmt"
)

func ParseMenuItemID(s string) (MenuItemID, error) {
	var m MenuItemID
	_, err := fmt.Sscanf(s, "%s-%s", &m.GroupName, &m.ItemName)
	if err != nil {
		return MenuItemID{}, fmt.Errorf("failed to parse MenuItemID: %w", err)
	}
	return m, nil
}
