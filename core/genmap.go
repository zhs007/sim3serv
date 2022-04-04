package sim3core

import (
	"math/rand"

	"github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type GenMapParams struct {
	MapTileWeights map[int]int
	TotalWeight    int
}

func (params *GenMapParams) Rebuild() {
	params.TotalWeight = 0

	for _, v := range params.MapTileWeights {
		params.TotalWeight += v
	}
}

func (params *GenMapParams) RandAMapTile() int {
	cr := rand.Int() % params.TotalWeight
	for mt, v := range params.MapTileWeights {
		if cr < v {
			return mt
		}
	}

	return -1
}

func genMap(mgr *MapTileMgr, w, h int, params *GenMapParams) (*MapData, error) {
	md := NewMapData(w, h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ct := params.RandAMapTile()
			if ct == -1 {
				goutils.Error("genMap:RandAMapTile",
					zap.Int("x", x),
					zap.Int("y", y),
					zap.Error(ErrInvalidRandResult))

				return nil, ErrInvalidRandResult
			}

			md.Data[y][x] = MapDataType(ct)
		}
	}

	return md, nil
}
