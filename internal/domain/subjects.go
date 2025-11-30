package domain

import "time"

type Subjects struct{
	ID int
	NameTk string
	NameEn string
	NameRu string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
