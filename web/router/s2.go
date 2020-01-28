package router

import (
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/wikipedia-on-map-go/web"
	"github.com/suzuito/common-go/clogger"
	"github.com/suzuito/common-go/cgcp/cweb"
)

func GetCellFromFace(app web.ApplicationWeb) func(*gin.Context) {
	return func(ctx *gin.Context) {
		cweb.H(app, ctx, func(logger clogger.Logger, fcli *firestore.Client) error {
			// faceID := ctx.Param("faceID")
			// s2.CellIDFromFace(faceID)
			return nil
		}, nil)
	}
}
