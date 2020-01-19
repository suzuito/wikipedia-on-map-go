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
		cellTokenIDs []string,
		locs *[]*model.GeoLocation,
	) error
	Close() error
}
