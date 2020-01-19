package geo

import (
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
	"github.com/suzuito/wikipedia-on-map-go/entity/model"
)

// https://github.com/google/s2geometry/blob/344e8603bb9b53854fb5dcb37f98ca60b0053631/src/s2/s2testing.cc#L162-L164
func kmToAngle(km float64) s1.Angle {
	kEarthRadiusKm := 6371.01
	return s1.Angle(km / kEarthRadiusKm)
}

func GetCellChildren(lat, lng float64, lv int) *[]*s2.Cell {
	ret := []*s2.Cell{}
	cell := s2.CellFromLatLng(s2.LatLngFromDegrees(lat, lng))
	ret = append(ret, &cell)
	currentLevel := cell.Level()
	for {
		if currentLevel == lv {
			break
		}
		currentLevel--
		parentCellID := cell.ID().Parent(currentLevel)
		cellr := s2.CellFromCellID(parentCellID)
		ret = append(ret, &cellr)
	}
	return &ret
}

func GetCell(lat, lng float64, lv int) *s2.Cell {
	ret := GetCellChildren(lat, lng, lv)
	if len(*ret) <= 0 {
		return nil
	}
	return (*ret)[len(*ret)-1]
}

func GetConvexCellsByLatLng(lat, lng float64, lv int, radius float64) *[]*s2.Cell {
	cap := s2.CapFromCenterAngle(
		s2.PointFromLatLng(
			s2.LatLngFromDegrees(
				lat,
				lng,
			),
		),
		kmToAngle(radius),
	)
	cellIDs := cap.CellUnionBound()
	cells := []*s2.Cell{}
	for _, cellID := range cellIDs {
		cell := s2.CellFromCellID(cellID)
		cells = append(cells, &cell)
	}
	return &cells
}

func NewCellFromS2Cell(c *s2.Cell) *model.Cell {
	ret := model.Cell{}
	cellID := c.ID()
	rect := c.RectBound()
	lo := rect.Lo()
	hi := rect.Hi()
	ret.ID = model.CellToken(cellID.ToToken())
	ret.Latitude = model.Interval{
		Lo: lo.Lat.Degrees(),
		Hi: hi.Lat.Degrees(),
	}
	ret.Longitude = model.Interval{
		Lo: lo.Lng.Degrees(),
		Hi: hi.Lng.Degrees(),
	}
	return &ret
}

func NewCellsFromS2Cells(cs *[]*s2.Cell) *[]*model.Cell {
	ret := []*model.Cell{}
	for _, c := range *cs {
		ret = append(ret, NewCellFromS2Cell(c))
	}
	return &ret
}
