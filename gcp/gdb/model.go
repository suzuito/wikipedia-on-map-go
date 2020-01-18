package gdb

import "github.com/suzuito/wikipedia-on-map-go/entity/model"

type GeoLocation struct {
	ID          string  `firestore:"id"`
	Latitude    float64 `firestore:"latitude"`
	Longitude   float64 `firestore:"longitude"`
	Name        string  `firestore:"name"`
	CellIDToken string  `firestore:"cellIDToken"`
}

func NewGeoLocationFromGeoLocation(geo *model.GeoLocation) *GeoLocation {
	return &GeoLocation{
		ID:          geo.ID,
		Name:        geo.Name,
		Latitude:    geo.Latitude,
		Longitude:   geo.Longitude,
		CellIDToken: geo.CellIDToken,
	}
}
