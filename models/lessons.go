package models

type Lesson struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	LanguageID   int    `json:"language_id"`
	LanguageName string `json:"language_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type CreateLesson struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	LanguageID   int    `json:"language_id"`
	LanguageName string `json:"-"`
}

type UpdateLesson struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	LanguageID   int    `json:"language_id"`
	LanguageName string `json:"language_name"`
}
