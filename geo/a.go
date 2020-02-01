package geo

import (
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
	"github.com/suzuito/wikipedia-on-map-go/entity/model"
)

const kEarthRadiusKm = 6371.01

// https://github.com/google/s2geometry/blob/344e8603bb9b53854fb5dcb37f98ca60b0053631/src/s2/s2testing.cc#L162-L164
func kmToAngle(km float64) s1.Angle {
	return s1.Angle(km / kEarthRadiusKm)
}

func angleToKm(a s1.Angle) float64 {
	return kEarthRadiusKm * float64(a)
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

func GetCap(lat, lng float64, radius float64) *s2.Cap {
	cap := s2.CapFromCenterAngle(
		s2.PointFromLatLng(
			s2.LatLngFromDegrees(
				lat,
				lng,
			),
		),
		kmToAngle(radius),
	)
	return &cap
}

func GetConvexCellsByLatLng(lat, lng float64, lv int, radius float64) (*s2.Cap, *[]*s2.Cell) {
	cap := GetCap(lat, lng, radius)
	cellIDs := cap.CellUnionBound()
	cells := []*s2.Cell{}
	for _, cellID := range cellIDs {
		cell := s2.CellFromCellID(cellID)
		cells = append(cells, &cell)
	}
	return cap, &cells
}

func GetConvexCellsByLatLng2(
	lat, lng float64,
	lv int, radius float64,
) (*s2.Cap, *[]*s2.Cell) {
	cellsRespMap := make(map[string]*s2.Cell)
	cellsResp := []*s2.Cell{}
	cap, cells := GetConvexCellsByLatLng(
		lat,
		lng,
		lv,
		radius,
	)
	for _, cell := range *cells {
		cellsTemp := GetCellsSpecifiedLevel(
			cell,
			lv,
		)
		for _, cellTemp := range *cellsTemp {
			if !cap.ContainsCell(*cellTemp) && !cap.IntersectsCell(*cellTemp) {
				continue
			}
			cellsRespMap[cellTemp.ID().ToToken()] = cellTemp
		}
	}
	for _, cellTemp := range cellsRespMap {
		cellsResp = append(cellsResp, cellTemp)
	}
	return cap, &cellsResp
}

func GetCellsSpecifiedLevel(c *s2.Cell, lv int) *[]*s2.Cell {
	ret := []*s2.Cell{}
	if lv == c.Level() {
		ret = append(ret, c)
	} else if lv < c.Level() {
		cc := s2.CellFromCellID(c.ID().Parent(lv))
		ret = append(
			ret,
			&cc,
		)
	} else if lv > c.Level() {
		cellIDBegin := c.ID().ChildBeginAtLevel(lv)
		cellIDEnd := c.ID().ChildEndAtLevel(lv)
		for cellID := cellIDBegin; cellID != cellIDEnd; cellID = cellID.Next() {
			cc := s2.CellFromCellID(cellID)
			ret = append(
				ret,
				&cc,
			)
		}
	}
	return &ret
}

func getCellParent(c *s2.Cell, lv int) *[]*s2.Cell {
	for {
		c.ID().Parent(lv)
	}
}

func NewCellFromS2Cell(c *s2.Cell) *model.Cell {
	ret := model.Cell{}
	ret.Center = *NewLatLngFromS2LatLng(s2.LatLngFromPoint(
		c.Center(),
	))
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
	ret.BoundLoop = NewLoopFromS2Cell(c)
	ret.Level = c.Level()
	return &ret
}

func NewCellsFromS2Cells(cs *[]*s2.Cell) *[]*model.Cell {
	ret := []*model.Cell{}
	for _, c := range *cs {
		ret = append(ret, NewCellFromS2Cell(c))
	}
	return &ret
}

func NewLatLngFromS2LatLng(ll s2.LatLng) *model.LatLng {
	return &model.LatLng{
		Latitude:  ll.Lat.Degrees(),
		Longitude: ll.Lng.Degrees(),
	}
}

func NewCapFromS2Cap(c *s2.Cap) *model.Cap {
	return &model.Cap{
		Center: *NewLatLngFromS2LatLng(
			s2.LatLngFromPoint(
				c.Center(),
			),
		),
		Radius: angleToKm(c.Radius()),
	}
}

func getCellsChildren(face, lv int) *[]*s2.Cell {
	cellIDFace := s2.CellIDFromFace(face)
	cellIDBegin := cellIDFace.ChildBeginAtLevel(lv)
	cellIDEnd := cellIDFace.ChildEndAtLevel(lv)
	returned := []*s2.Cell{}
	for cellID := cellIDBegin; cellID != cellIDEnd; cellID = cellID.Next() {
		cell := s2.CellFromCellID(cellID)
		returned = append(returned, &cell)
	}
	return &returned
}

func GetCellsChildren(face, lv int) *[]*model.Cell {
	cells := getCellsChildren(face, lv)
	return NewCellsFromS2Cells(cells)
}
