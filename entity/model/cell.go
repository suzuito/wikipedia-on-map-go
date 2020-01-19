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
}
