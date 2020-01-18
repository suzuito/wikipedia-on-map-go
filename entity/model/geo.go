package model

type GeoLocation struct {
	ID          string
	Latitude    float64
	Longitude   float64
	Name        string
	CellIDToken string
}

func NewGeoLocation(name string, lat, lng float64, cellIDToken string) *GeoLocation {
	return &GeoLocation{
		Latitude:    lat,
		Longitude:   lng,
		Name:        name,
		CellIDToken: cellIDToken,
	}
}
