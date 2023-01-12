package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	hp "github.com/hyphengolang/flyio/internal/harry-potter"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotImplemented = errors.New("not implemented")

type Character struct {
	ID      uuid.UUID
	Name    string
	Blood   string
	Species string
	Born    *time.Time
	Quote   string
	ImgURL  string
}

type Repo interface{}

type repo struct {
	conn *pgxpool.Pool
}

func (r *repo) FindMany(ctx context.Context) ([]*hp.Character, error) {
	return nil, ErrNotImplemented
}

func (r *repo) Find(ctx context.Context, key any) (*hp.Character, error) {
	return nil, ErrNotImplemented
}

func NewRepo() Repo {
	r := &repo{}
	return r
}
