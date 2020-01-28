package application

import (
	"context"
	"os"
	"strings"

	env "github.com/suzuito/common-env"
)

type ApplicationBase interface {
	IndexLevel() int
	AllowedOrigins() []string
}

type ApplicationBaseImpl struct {
	indexLevel int
}

func NewApplicationBaseImpl(ctx context.Context) (*ApplicationBaseImpl, error) {
	indexLevel := env.GetenvAsInt("INDEX_LEVEL", 15)
	return &ApplicationBaseImpl{
		indexLevel: indexLevel,
	}, nil
}

func (a *ApplicationBaseImpl) IndexLevel() int {
	return a.indexLevel
}

func (a *ApplicationBaseImpl) AllowedOrigins() []string {
	return strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
}
