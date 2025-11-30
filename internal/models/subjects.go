package models

import (
	"time"

	"github.com/2004942/library/internal/domain"
)
type Subjects struct{
	ID int `db:"id"`
	NameTk string `db:"name_tk"`
	NameEn string `db:"nmae_en"`
	NamaeRu string `db:"name_ru"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

func (s *Subjects) ToDomain() domain.Subjects{
	return domain.Subjects{
		ID: s.ID,
		NameTk: s.NameTk,
		NameEn: s.NameEn,
		NameRu: s.NamaeRu,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		DeletedAt: s.DeletedAt,
	}
}