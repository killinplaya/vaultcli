package entity

import (
	"strings"
	"time"
)

type Entry struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Login       string   `json:"login"`
	Password    string   `json:"password"`
	URL         string   `json:"url"`
	Notes       string   `json:"notes"`
	Tags        []string `json:"tags"`
	TimeCreated string   `json:"time_created"`
	TimeUpdated string   `json:"time_updated"`
}

func NewEntry(name, login, password, url, notes string, tags []string) *Entry {
	// 1. Очистка строк от лишних пробелов
	name = strings.TrimSpace(name)
	login = strings.TrimSpace(login)
	password = strings.TrimSpace(password)
	url = strings.TrimSpace(url)
	notes = strings.TrimSpace(notes)

	// 2. Нормализация тегов
	cleanTags := make([]string, 0, len(tags))
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		cleanTags = append(cleanTags, tag)
	}

	// 3. Установка времени при пустых значениях
	now := time.Now().Format(time.RFC3339)

	// 4. Сборка и возвращение
	return &Entry{
		Name:        name,
		Login:       login,
		Password:    password,
		URL:         url,
		Notes:       notes,
		Tags:        cleanTags,
		TimeCreated: now,
		TimeUpdated: now,
	}
}
