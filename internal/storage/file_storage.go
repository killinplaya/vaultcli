package storage

import (
	"encoding/json"
	"os"
	"vaultcli/internal/entity"
)

type FileStorage struct {
	filePath string
}

func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{
		filePath: filePath,
	}
}

func (fs *FileStorage) Load() ([]entity.Entry, error) {
	if _, err := os.Stat(fs.filePath); err != nil {
		if os.IsNotExist(err) {
			return []entity.Entry{}, nil
		}
		return nil, err
	}

	data, err := os.ReadFile(fs.filePath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []entity.Entry{}, nil
	}

	var entries []entity.Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}

	return entries, nil
}

func (fs *FileStorage) Save(entries []entity.Entry) error {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(fs.filePath, data, 0644); err != nil {
		return err
	}

	return nil
}
