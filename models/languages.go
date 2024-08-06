package models

import "time"

type Language struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LanguageCode string    `json:"language_code"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateLanguage struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	LanguageCode string    `json:"language_code"`
	CreatedAt    time.Time `json:"-"`
}

type UpdateLanguage struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	LanguageCode string `json:"language_code"`
}
