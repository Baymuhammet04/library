package models

import (
	"time"

	"github.com/2004942/library/internal/domain"
)

type Classes struct{
	ID int `db:"id"`
	NameTk string `db:"name_tk"`
	NameEn string `db:"nmae_en"`
	NmaeRu string `db:"name_ru"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

func(c *Classes) ToDomain() domain.Classes{
	return domain.Classes{
		ID: c.ID,
		NameTk: c.NameTk,
		NameEn: c.NameEn,
		NmaeRu: c.NmaeRu,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt,
	}
}