package postgres

import (
	"context"

	"github.com/2004942/library/internal/domain"
	"github.com/2004942/library/pkg/connection"
)

var _ SubjectRepository = (*subjectRepository)(nil)

type SubjectRepository interface {
	Create(ctx context.Context, subject domain.Subjects) (int, error)
	Update(ctx context.Context, subject domain.Subjects) error
	
}

type subjectRepository struct {
	psqlDB connection.DB
}

func NewSubjectRepository(psqlDB connection.DB) *subjectRepository {
	return &subjectRepository{psqlDB: psqlDB}
}

func (r *subjectRepository) Create(ctx context.Context, subject domain.Subjects) (int, error) {
	var id int

	query := `INSERT INTO subject (name_tk, name_en, name_ru) VALUES ($1, $2, $3) RETURNING id`

	err := r.psqlDB.QueryRow(ctx, query, subject.NameTk, subject.NameEn, subject.NameEn).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *subjectRepository) Update(ctx context.Context, subject domain.Subjects) error{
	query := `UPDATE subject SET name_tk = $1, name_en = $2, name_ru =$3 WHERE id=$4`
	
	_, err := r.psqlDB.Exec(ctx, query, subject.NameTk, subject.NameEn, subject.NameEn, subject.ID)
	if err != nil{
		return err
	}
	return nil
}