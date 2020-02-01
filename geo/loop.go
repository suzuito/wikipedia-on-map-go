package geo

import (
	"github.com/golang/geo/s2"
	"github.com/suzuito/wikipedia-on-map-go/entity/model"
)

func NewLoopFromS2Loop(loop *s2.Loop) *model.Loop {
	ret := model.Loop{}
	for _, v := range loop.Vertices() {
		ll := s2.LatLngFromPoint(v)
		ret.Points = append(ret.Points, *NewLatLngFromS2LatLng(ll))
	}
	return &ret
}

func NewLoopFromS2Cell(cell *s2.Cell) *model.Loop {
	return NewLoopFromS2Loop(s2.LoopFromCell(*cell))
}
