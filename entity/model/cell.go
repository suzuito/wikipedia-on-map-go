package model

// Interval ...
type Interval struct {
	Lo float64
	Hi float64
}

// CellToken ...
type CellToken string

// Cell ...
type Cell struct {
	ID        CellToken
	Latitude  Interval
	Longitude Interval
	Level     int
}

// LatLng ...
type LatLng struct {
	Latitude  float64
	Longitude float64
}

// Cap ...
type Cap struct {
	Center LatLng
	Radius float64
}
