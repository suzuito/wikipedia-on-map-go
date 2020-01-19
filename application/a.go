package application

import (
	"context"

	"github.com/suzuito/wikipedia-on-map-go/entity/db"
	"github.com/suzuito/wikipedia-on-map-go/env"
	"github.com/suzuito/wikipedia-on-map-go/slogger"
)

type Application interface {
	DBClient(ctx context.Context) (db.Client, error)
	Logger(ctx context.Context) slogger.Logger
	IndexLevel() int
}

type ApplicationBase struct {
	indexLevel int
}

func NewApplicationBase() (*ApplicationBase, error) {
	indexLevel := env.GetenvAsInt("INDEX_LEVEL", 15)
	return &ApplicationBase{
		indexLevel: indexLevel,
	}, nil
}

func (a *ApplicationBase) IndexLevel() int {
	return a.indexLevel
}
