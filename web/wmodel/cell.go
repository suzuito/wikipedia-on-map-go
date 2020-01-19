package wmodel

import "github.com/suzuito/wikipedia-on-map-go/entity/model"

// Interval ...
type Interval struct {
	Lo float64 `json:"lo"`
	Hi float64 `json:"hi"`
}

func NewInterval(
	i *model.Interval,
) *Interval {
	return &Interval{
		Lo: i.Lo,
		Hi: i.Hi,
	}
}

// Cell ...
type Cell struct {
	ID        string   `json:"id"`
	Latitude  Interval `json:"latitude"`
	Longitude Interval `json:"longitude"`
}

func NewCell(
	c *model.Cell,
) *Cell {
	return &Cell{
		ID:        string(c.ID),
		Latitude:  *NewInterval(&c.Latitude),
		Longitude: *NewInterval(&c.Longitude),
	}
}
