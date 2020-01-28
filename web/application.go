package web

import (
	"context"

	"github.com/suzuito/common-go/cgcp"
	"github.com/suzuito/wikipedia-on-map-go/application"
)

type ApplicationWeb interface {
	cgcp.ApplicationGCP
	application.ApplicationBase
}

type ApplicationWebImpl struct {
	*cgcp.ApplicationGCPImpl
	*application.ApplicationBaseImpl
}

func NewApplicationWebImpl(ctx context.Context) (*ApplicationWebImpl, error) {
	appGCPImpl, err := cgcp.NewApplicationGCPImpl(ctx)
	if err != nil {
		return nil, err
	}
	appBaseImpl, err := application.NewApplicationBaseImpl(ctx)
	if err != nil {
		return nil, err
	}
	return &ApplicationWebImpl{
		ApplicationGCPImpl:  appGCPImpl,
		ApplicationBaseImpl: appBaseImpl,
	}, nil
}
