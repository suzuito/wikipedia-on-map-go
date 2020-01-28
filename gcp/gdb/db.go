package gdb

import (
	"context"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/common-go/werror"
	"github.com/suzuito/wikipedia-on-map-go/entity/model"
	"google.golang.org/api/iterator"
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
			NewGeoLocationFromGeoLocation(value),
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

func (c *ClientFirestore) GetGeoLocationsIncludedByCells(
	ctx context.Context,
	level int,
	cellIDs *[]string,
	values *[]*model.GeoLocation,
) error {
	crefCell := c.cli.
		Collection(GeoLocationLevelCol).
		Doc(strconv.Itoa(level)).
		Collection(GeoLocationCellCol)
	for _, cellID := range *cellIDs {
		drefCell := crefCell.Doc(cellID)
		crefLoc := drefCell.Collection(GeoLocationCol)
		iter := crefLoc.Documents(ctx)
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return werror.New(err)
			}
			loc := model.GeoLocation{}
			if err := doc.DataTo(&loc); err != nil {
				return werror.New(err)
			}
			*values = append(*values, &loc)
		}
	}
	return nil
}

func (c *ClientFirestore) Close() error {
	return c.cli.Close()
}
