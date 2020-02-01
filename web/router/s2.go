package router

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/cgcp/cweb"
	"github.com/suzuito/common-go/cgin"
	"github.com/suzuito/common-go/clogger"
	"github.com/suzuito/wikipedia-on-map-go/entity/model"
	"github.com/suzuito/wikipedia-on-map-go/geo"
	"github.com/suzuito/wikipedia-on-map-go/web"
	"github.com/suzuito/wikipedia-on-map-go/web/wmodel"
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

func GetCells(app web.ApplicationWeb) func(*gin.Context) {
	return func(ctx *gin.Context) {
		cweb.H(app, ctx, func(logger clogger.Logger, fcli *firestore.Client) error {
			face := cgin.DefaultQueryAsInt(ctx, "face", 1)
			level := cgin.DefaultQueryAsInt(ctx, "level", 1)
			cells := []*model.Cell{}
			cells = append(cells, *geo.GetCellsChildren(face, level)...)
			wcells := wmodel.NewCells(&cells)
			ctx.JSON(http.StatusOK, wcells)
			return nil
		}, nil)
	}
}
