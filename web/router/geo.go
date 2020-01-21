package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/wikipedia-on-map-go/application"
	"github.com/suzuito/wikipedia-on-map-go/entity/db"
	"github.com/suzuito/wikipedia-on-map-go/entity/model"
	"github.com/suzuito/wikipedia-on-map-go/geo"
	"github.com/suzuito/wikipedia-on-map-go/slogger"
	"github.com/suzuito/wikipedia-on-map-go/web"
	"github.com/suzuito/wikipedia-on-map-go/web/wmodel"
)

func GetGeoLocations(app application.Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		web.H(app, ctx, func(logger slogger.Logger, dcli db.Client) error {
			lat := QueryFloat64(ctx, "lat", 0)
			lng := QueryFloat64(ctx, "lng", 0)
			radius := QueryFloat64(ctx, "radius", 10)
			_, cells := geo.GetConvexCellsByLatLng2(
				lat, lng,
				app.IndexLevel(),
				radius,
			)
			cellTokenIDs := []string{}
			for _, cell := range *cells {
				cellTokenIDs = append(cellTokenIDs, cell.ID().ToToken())
			}
			locs := []*model.GeoLocation{}
			if err := dcli.GetGeoLocationsIncludedByCells(
				ctx,
				app.IndexLevel(),
				&cellTokenIDs,
				&locs,
			); err != nil {
				web.Abort(ctx, web.NewHTTPError(500, err.Error(), err))
				return err
			}
			ctx.JSON(http.StatusOK, wmodel.NewGeoLocations(&locs))
			return nil
		}, nil)
	}
}

func GetGeoCap(app application.Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		web.H(app, ctx, func(logger slogger.Logger, dcli db.Client) error {
			lat := QueryFloat64(ctx, "lat", 0)
			lng := QueryFloat64(ctx, "lng", 0)
			radius := QueryFloat64(ctx, "radius", 10)
			cap := wmodel.NewCap(
				geo.NewCapFromS2Cap(
					geo.GetCap(lat, lng, radius),
				),
			)
			ctx.JSON(http.StatusOK, cap)
			return nil
		}, nil)
	}
}

func GetGeoCellsConvex(app application.Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		web.H(app, ctx, func(logger slogger.Logger, dcli db.Client) error {
			lat := QueryFloat64(ctx, "lat", 0)
			lng := QueryFloat64(ctx, "lng", 0)
			radius := QueryFloat64(ctx, "radius", 10)
			_, cellsResp := geo.GetConvexCellsByLatLng2(
				lat, lng,
				app.IndexLevel(), radius,
			)
			mcells := geo.NewCellsFromS2Cells(
				cellsResp,
			)
			ctx.JSON(
				http.StatusOK,
				wmodel.NewCells(mcells),
			)
			return nil
		}, nil)
	}
}

func GetGeoCellsChildren(app application.Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		web.H(app, ctx, func(logger slogger.Logger, dcli db.Client) error {
			lat := QueryFloat64(ctx, "lat", 0)
			lng := QueryFloat64(ctx, "lng", 0)
			cells := geo.NewCellsFromS2Cells(
				geo.GetCellChildren(lat, lng, app.IndexLevel()),
			)
			respCells := []*wmodel.Cell{}
			for _, cell := range *cells {
				respCells = append(
					respCells,
					wmodel.NewCell(
						cell,
					),
				)
			}
			ctx.JSON(http.StatusOK, respCells)
			return nil
		}, nil)
	}
}
