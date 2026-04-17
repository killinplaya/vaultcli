package utils

import "vaultcli/internal/entity"

func NextID(entries []entity.Entry) int {
	if len(entries) == 0 {
		return 1
	}

	maxID := entries[0].ID

	for _, r := range entries {
		if r.ID > maxID {
			maxID = r.ID
		}
	}

	return maxID + 1
}
