package model

import (
	"encoding/base64"
	"net/url"
)

type GeoLocation struct {
	ID          string
	Latitude    float64
	Longitude   float64
	Name        string
	CellIDToken string
}

func NewGeoLocation(name string, lat, lng float64, cellIDToken string) *GeoLocation {
	return &GeoLocation{
		ID: base64.StdEncoding.EncodeToString([]byte(
			url.QueryEscape(name),
		)),
		Latitude:    lat,
		Longitude:   lng,
		Name:        name,
		CellIDToken: cellIDToken,
	}
}
