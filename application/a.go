package application

import (
	"context"

	"github.com/suzuito/wikipedia-on-map-go/entity/db"
)

type Application interface {
	DBClient(ctx context.Context) (db.Client, error)
}

type ApplicationBase struct{}

func NewApplicationBase() (*ApplicationBase, error) {
	return &ApplicationBase{}, nil
}
