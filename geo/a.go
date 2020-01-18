package geo

import "github.com/golang/geo/s2"

func GetCellChildren(lat, lng float64, lv int) *[]s2.Cell {
	ret := []s2.Cell{}
	cell := s2.CellFromLatLng(s2.LatLngFromDegrees(lat, lng))
	ret = append(ret, cell)
	currentLevel := cell.Level()
	for {
		if currentLevel == lv {
			break
		}
		currentLevel--
		parentCellID := cell.ID().Parent(currentLevel)
		cell = s2.CellFromCellID(parentCellID)
		ret = append(ret, cell)
	}
	return &ret
}

func GetCell(lat, lng float64, lv int) *s2.Cell {
	ret := GetCellChildren(lat, lng, lv)
	if len(*ret) <= 0 {
		return nil
	}
	return &(*ret)[len(*ret)-1]
}
