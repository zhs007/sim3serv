package sim3core

type MapData struct {
	Width  int
	Height int
	Data   [][]MapDataType
}

func NewMapData(w int, h int) *MapData {
	mapData := &MapData{
		Width:  w,
		Height: h,
	}

	for y := 0; y < h; y++ {
		mapData.Data = append(mapData.Data, make([]MapDataType, w))
	}

	return mapData
}
