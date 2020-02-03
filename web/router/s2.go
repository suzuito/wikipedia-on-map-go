package router

import (
	"net/http"
	"strconv"
	"strings"

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
			faces := []int{}
			facesStr := ctx.DefaultQuery("faces", "")
			if facesStr == "" {
				faces = []int{0, 1, 2, 3, 4, 5}
			} else {
				facesStr2 := strings.Split(facesStr, ",")
				for _, faceStr2 := range facesStr2 {
					i, err := strconv.Atoi(faceStr2)
					if err != nil {
						continue
					}
					faces = append(faces, i)
				}
			}
			level := cgin.DefaultQueryAsInt(ctx, "level", 1)
			cells := []*model.Cell{}
			for _, face := range faces {
				cells = append(cells, *geo.GetCellsChildren(face, level)...)
			}
			wcells := wmodel.NewCells(&cells)
			ctx.JSON(http.StatusOK, wcells)
			return nil
		}, nil)
	}
}

func GetCellsAll(app web.ApplicationWeb) func(*gin.Context) {
	return func(ctx *gin.Context) {
		cweb.H(app, ctx, func(logger clogger.Logger, fcli *firestore.Client) error {
			level := cgin.DefaultQueryAsInt(ctx, "level", 1)
			cells := []*model.Cell{}
			cells = append(cells, *geo.GetCellsChildren(0, level)...)
			cells = append(cells, *geo.GetCellsChildren(1, level)...)
			cells = append(cells, *geo.GetCellsChildren(2, level)...)
			cells = append(cells, *geo.GetCellsChildren(3, level)...)
			cells = append(cells, *geo.GetCellsChildren(4, level)...)
			cells = append(cells, *geo.GetCellsChildren(5, level)...)
			wcells := wmodel.NewCells(&cells)
			ctx.JSON(http.StatusOK, wcells)
			return nil
		}, nil)
	}
}
