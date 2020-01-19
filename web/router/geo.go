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
			lat := QueryFloat64(ctx, "lat", -1)
			lng := QueryFloat64(ctx, "lng", -1)
			radius := QueryFloat64(ctx, "radius", 10)
			if lat < 0 {
				web.Abort(ctx, web.NewHTTPError(http.StatusBadRequest, "Set positive float value as lat", nil))
				return nil
			}
			if lng < 0 {
				web.Abort(ctx, web.NewHTTPError(http.StatusBadRequest, "Set positive float value as lng", nil))
				return nil
			}
			cells := geo.GetConvexCellsByLatLng(
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
				cellTokenIDs,
				&locs,
			); err != nil {
				web.Abort(ctx, web.NewHTTPError(500, err.Error(), err))
				return err
			}
			ctx.JSON(http.StatusOK, locs)
			return nil
		}, nil)
	}
}

func GetGeoCap(app application.Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		web.H(app, ctx, func(logger slogger.Logger, dcli db.Client) error {
			lat := QueryFloat64(ctx, "lat", -1)
			lng := QueryFloat64(ctx, "lng", -1)
			radius := QueryFloat64(ctx, "radius", 10)
			if lat < 0 {
				web.Abort(ctx, web.NewHTTPError(http.StatusBadRequest, "Set positive float value as lat", nil))
				return nil
			}
			if lng < 0 {
				web.Abort(ctx, web.NewHTTPError(http.StatusBadRequest, "Set positive float value as lng", nil))
				return nil
			}
			cells := geo.NewCellsFromS2Cells(
				geo.GetConvexCellsByLatLng(
					lat,
					lng,
					app.IndexLevel(),
					radius,
				),
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

func GetGeoCellsConvex(app application.Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		web.H(app, ctx, func(logger slogger.Logger, dcli db.Client) error {
			lat := QueryFloat64(ctx, "lat", -1)
			lng := QueryFloat64(ctx, "lng", -1)
			radius := QueryFloat64(ctx, "radius", 10)
			if lat < 0 {
				web.Abort(ctx, web.NewHTTPError(http.StatusBadRequest, "Set positive float value as lat", nil))
				return nil
			}
			if lng < 0 {
				web.Abort(ctx, web.NewHTTPError(http.StatusBadRequest, "Set positive float value as lng", nil))
				return nil
			}
			cells := geo.NewCellsFromS2Cells(
				geo.GetConvexCellsByLatLng(
					lat,
					lng,
					app.IndexLevel(),
					radius,
				),
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

func GetGeoCellsChildren(app application.Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		web.H(app, ctx, func(logger slogger.Logger, dcli db.Client) error {
			lat := QueryFloat64(ctx, "lat", -1)
			lng := QueryFloat64(ctx, "lng", -1)
			if lat < 0 {
				web.Abort(ctx, web.NewHTTPError(http.StatusBadRequest, "Set positive float value as lat", nil))
				return nil
			}
			if lng < 0 {
				web.Abort(ctx, web.NewHTTPError(http.StatusBadRequest, "Set positive float value as lng", nil))
				return nil
			}
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
