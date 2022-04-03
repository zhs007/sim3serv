package sim3core

// 地图数据
type MapDataType int16

// tile类型
type TileType int

const (
	TILE_FOOD TileType = 1
)

func String2TileType(str string) (TileType, error) {
	switch str {
	case "TILE_FOOD":
		return TILE_FOOD, nil
	}

	return 0, ErrInvalidTileType
}
