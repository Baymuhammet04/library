package usecase
import (
	"context"

	"github.com/2004942/library/internal/domain"
	"github.com/2004942/library/internal/repository/postgres"
)

var _ SubjectUC = (*subjectUC)(nil)

type SubjectUC interface {
	Create(ctx context.Context, subject domain.Subjects) (int, error)
	Update(ctx context.Context, subject domain.Subjects) error
}

type subjectUC struct {
	repo postgres.SubjectRepository
}

func NewSubjectUC(repo postgres.SubjectRepository) *subjectUC {
	return &subjectUC{repo: repo}
}

func (u *subjectUC) Create(ctx context.Context, subject domain.Subjects) (int, error) {
	id, err := u.repo.Create(ctx, subject)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func(u *subjectUC) Update(ctx context.Context, subject domain.Subjects) error{
	err := u.repo.Update(ctx, subject)
	if err!= nil{
		return err
	}
	
	return nil
}