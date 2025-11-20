package domain

import "time"

type Subjects struct{
	ID int
	NameTk string
	NameEn string
	NmaeRu string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
