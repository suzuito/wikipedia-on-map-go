package wmodel

import "github.com/suzuito/wikipedia-on-map-go/entity/model"

type Loop struct {
	Points []LatLng `json:"points"`
}

func NewLoop(l *model.Loop) *Loop {
	ret := Loop{}
	for _, v := range l.Points {
		ret.Points = append(ret.Points, *NewLatLng(&v))
	}
	return &ret
}
