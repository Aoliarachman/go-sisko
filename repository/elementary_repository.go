package repository

import (
	"context"
	"database/sql"
	"go-sisko/model/domain"
)

type ElementaryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, elementary domain.Elementary) domain.Elementary
	Update(ctx context.Context, tx *sql.Tx, elementary domain.Elementary) domain.Elementary
	Delete(ctx context.Context, tx *sql.Tx, elementary domain.Elementary)
	FindById(ctx context.Context, tx *sql.Tx, elementaryId int) (domain.Elementary, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Elementary
}
