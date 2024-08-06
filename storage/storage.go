package storage

import (
	"context"
	"learning-service/models"
)

type MainStorage interface {
	Languages()
	Lessons()
	Vocabulary()
	UserLessons()
	UserVocabulary()
}

type Languages interface {
	Create(ctx context.Context, lan *models.CreateLanguage) (*models.Language, error)
	Update(ctx context.Context, id string, lan *models.UpdateLanguage) (*models.Language, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id int) (*models.Language, error)
	List(ctx context.Context, limit, page string) ([]*models.Language, error)
}

type Lessons interface {
	Create(ctx context.Context, lesson *models.CreateLesson) (*models.Lesson, error)
	Update(ctx context.Context, id string, lesson *models.UpdateLesson) (*models.Lesson, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id int) (*models.Lesson, error)
	List(ctx context.Context, languageID int, limit, page string) ([]*models.Lesson, error)
	GetLessonBy
}
