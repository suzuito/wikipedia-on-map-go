package router

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/cgcp/cweb"
	"github.com/suzuito/common-go/cgin"
	"github.com/suzuito/common-go/clogger"
	"github.com/suzuito/wikipedia-on-map-go/entity/model"
	"github.com/suzuito/wikipedia-on-map-go/gcp/gdb"
	"github.com/suzuito/wikipedia-on-map-go/geo"
	"github.com/suzuito/wikipedia-on-map-go/web"
	"github.com/suzuito/wikipedia-on-map-go/web/wmodel"
)

func GetGeoLocations(app web.ApplicationWeb) func(*gin.Context) {
	return func(ctx *gin.Context) {
		cweb.H(app, ctx, func(logger clogger.Logger, fcli *firestore.Client) error {
			lat := cgin.DefaultQueryAsFloat64(ctx, "lat", 0)
			lng := cgin.DefaultQueryAsFloat64(ctx, "lng", 0)
			radius := cgin.DefaultQueryAsFloat64(ctx, "radius", 10)
			_, cells := geo.GetConvexCellsByLatLng2(
				lat, lng,
				app.IndexLevel(),
				radius,
			)
			cellTokenIDs := []string{}
			for _, cell := range *cells {
				cellTokenIDs = append(cellTokenIDs, cell.ID().ToToken())
			}
			dcli := gdb.NewClientFirestore(fcli)
			locs := []*model.GeoLocation{}
			if err := dcli.GetGeoLocationsIncludedByCells(
				ctx,
				app.IndexLevel(),
				&cellTokenIDs,
				&locs,
			); err != nil {
				cweb.Abort(ctx, cweb.NewHTTPError(500, err.Error(), err))
				return err
			}
			ctx.JSON(http.StatusOK, wmodel.NewGeoLocations(&locs))
			return nil
		}, nil)
	}
}

func GetGeoCap(app web.ApplicationWeb) func(*gin.Context) {
	return func(ctx *gin.Context) {
		cweb.H(app, ctx, func(logger clogger.Logger, fcli *firestore.Client) error {
			lat := cgin.DefaultQueryAsFloat64(ctx, "lat", 0)
			lng := cgin.DefaultQueryAsFloat64(ctx, "lng", 0)
			radius := cgin.DefaultQueryAsFloat64(ctx, "radius", 10)
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

func GetGeoCellsConvex(app web.ApplicationWeb) func(*gin.Context) {
	return func(ctx *gin.Context) {
		cweb.H(app, ctx, func(logger clogger.Logger, fcli *firestore.Client) error {
			lat := cgin.DefaultQueryAsFloat64(ctx, "lat", 0)
			lng := cgin.DefaultQueryAsFloat64(ctx, "lng", 0)
			radius := cgin.DefaultQueryAsFloat64(ctx, "radius", 10)
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

func GetGeoCellsChildren(app web.ApplicationWeb) func(*gin.Context) {
	return func(ctx *gin.Context) {
		cweb.H(app, ctx, func(logger clogger.Logger, fcli *firestore.Client) error {
			lat := cgin.DefaultQueryAsFloat64(ctx, "lat", 0)
			lng := cgin.DefaultQueryAsFloat64(ctx, "lng", 0)
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
