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
	Level     int      `json:"level"`
	Center    LatLng   `json:"center"`
	BoundLoop *Loop    `json:"boundLoop"`
	Face      int      `json:"face"`
}

func NewCell(
	c *model.Cell,
) *Cell {
	return &Cell{
		ID:        string(c.ID),
		Latitude:  *NewInterval(&c.Latitude),
		Longitude: *NewInterval(&c.Longitude),
		Level:     c.Level,
		Center:    *NewLatLng(&c.Center),
		BoundLoop: NewLoop(c.BoundLoop),
		Face:      c.Face,
	}
}

func NewCells(
	cs *[]*model.Cell,
) *[]*Cell {
	ret := []*Cell{}
	for _, c := range *cs {
		ret = append(ret, NewCell(c))
	}
	return &ret
}

// LatLng ...
type LatLng struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func NewLatLng(ll *model.LatLng) *LatLng {
	return &LatLng{
		Latitude:  ll.Latitude,
		Longitude: ll.Longitude,
	}
}

// Cap ...
type Cap struct {
	Center LatLng  `json:"center"`
	Radius float64 `json:"radius"`
}

func NewCap(c *model.Cap) *Cap {
	return &Cap{
		Center: *NewLatLng(&c.Center),
		Radius: c.Radius,
	}
}

// GeoLocation ...
type GeoLocation struct {
	ID          string  `json:"id"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Name        string  `json:"name"`
	CellIDToken string  `json:"cellIdToken"`
}

func NewGeoLocation(g *model.GeoLocation) *GeoLocation {
	return &GeoLocation{
		ID:          g.ID,
		Latitude:    g.Latitude,
		Longitude:   g.Longitude,
		Name:        g.Name,
		CellIDToken: g.CellIDToken,
	}
}

func NewGeoLocations(gs *[]*model.GeoLocation) *[]*GeoLocation {
	ret := []*GeoLocation{}
	for _, g := range *gs {
		ret = append(ret, NewGeoLocation(g))
	}
	return &ret
}
