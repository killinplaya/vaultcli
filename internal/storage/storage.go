package storage

import "vaultcli/internal/entity"

type Storage interface {
	Load() ([]entity.Entry, error)
	Save(entries []entity.Entry) error
}
