package db

import (
	"context"
	"github.com/suzuito/wikipedia-on-map-go/entity/model"
)

type Client interface {
	SetGeoLocations(
		ctx context.Context,
		level int,
		locs *[]*model.GeoLocation,
	) error
	GetGeoLocationsIncludedByCells(
		ctx context.Context,
		level int,
		cellIDs *[]string,
		values *[]*model.GeoLocation,
	) error
	Close() error
}
