package gcp

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/suzuito/wikipedia-on-map-go/application"
	"github.com/suzuito/wikipedia-on-map-go/entity/db"
	"github.com/suzuito/wikipedia-on-map-go/gcp/gdb"
	"github.com/suzuito/wikipedia-on-map-go/slogger"
)

type ApplicationGCP struct {
	*application.ApplicationBase
	appFirebase *firebase.App
}

func NewApplicationGCP(ctx context.Context) (*ApplicationGCP, error) {
	app := ApplicationGCP{}
	appBase, err := application.NewApplicationBase()
	if err != nil {
		return nil, err
	}
	app.ApplicationBase = appBase
	appFirebase, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	app.appFirebase = appFirebase
	return &app, nil
}

func (a *ApplicationGCP) DBClient(ctx context.Context) (db.Client, error) {
	cli, err := a.appFirebase.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return gdb.NewClientFirestore(cli), nil
}
func (a *ApplicationGCP) Logger(ctx context.Context) slogger.Logger {
	return &slogger.LoggerPrint{}
}
