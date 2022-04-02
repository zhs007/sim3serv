package sim3core

type BasicTile struct {
	Type TileType
}

// GetType
func (tile *BasicTile) GetType() TileType {
	return tile.Type
}
