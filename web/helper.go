package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/wikipedia-on-map-go/application"
	"github.com/suzuito/wikipedia-on-map-go/entity/db"
	"github.com/suzuito/wikipedia-on-map-go/slogger"
)

// HO ...
type HO struct {
	DBNotUse     bool
	LoggerNotUse bool
}

// H ...
func H(
	app application.Application,
	ctx *gin.Context,
	proc func(
		logger slogger.Logger,
		dcli db.Client,
	) error,
	opt *HO,
) {
	var logger slogger.Logger
	var dcli db.Client
	var err error
	if opt == nil || opt.LoggerNotUse == false {
		logger = app.Logger(ctx)
		defer logger.Close()
	}
	if opt == nil || opt.DBNotUse == false {
		dcli, err = app.DBClient(ctx)
		if err != nil {
			logger.Errorf("%+v", err)
			Abort(ctx, NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
			return
		}
		defer dcli.Close()
	}
	if err := proc(logger, dcli); err != nil {
		logger.Errorf("%+v", err)
	}
}
