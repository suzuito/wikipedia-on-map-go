package gdb

import (
	"context"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/wikipedia-on-map-go/entity/model"
	"github.com/suzuito/wikipedia-on-map-go/gcp/db"
	"github.com/suzuito/wikipedia-on-map-go/werror"
)

var (
	GeoLocationLevelCol = "geolocation_levels"
	GeoLocationCellCol  = "geolocation_cells"
	GeoLocationCol      = "geolocations"
)

// ClientFirestore ...
type ClientFirestore struct {
	cli *firestore.Client
}

// NewClientFirestore ...
func NewClientFirestore(cli *firestore.Client) *ClientFirestore {
	return &ClientFirestore{
		cli: cli,
	}
}

func (c *ClientFirestore) SetGeoLocations(
	ctx context.Context,
	level int,
	values *[]*model.GeoLocation,
) error {
	max := 450
	i := 0
	crefCell := c.cli.
		Collection(GeoLocationLevelCol).
		Doc(strconv.Itoa(level)).
		Collection(GeoLocationCellCol)
	batch := c.cli.Batch()
	for _, value := range *values {
		drefCell := crefCell.
			Doc(value.CellIDToken).
			Collection(GeoLocationCol)
		drefGeo := drefCell.Doc(value.ID)
		batch.Set(
			drefGeo,
			db.NewGeoLocationFromGeoLocation(value),
		)
		i++
		if i >= max {
			if _, err := batch.Commit(ctx); err != nil {
				return werror.New(err)
			}
			i = 0
			batch = c.cli.Batch()
		}
	}
	if _, err := batch.Commit(ctx); err != nil {
		return werror.New(err)
	}
	return nil
}
