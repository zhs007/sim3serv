package sim3core

import (
	"math/rand"
	"os"

	"github.com/zhs007/goutils"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type GenMapParams struct {
	MapTileWeights map[int]float32 `yaml:"tileWeights"`
}

func (params *GenMapParams) RandAMapTile() int {
	cr := rand.Float32()
	for mt, v := range params.MapTileWeights {
		if cr < v {
			return mt
		}

		cr -= v
	}

	return 0
}

func (params *GenMapParams) check() error {
	var cr float32
	for _, v := range params.MapTileWeights {
		cr += v
	}

	if cr > 1 {
		goutils.Error("GenMapParams:check",
			zap.Float32("totalval", cr),
			zap.Error(ErrInvalidGenMapParams))

		return ErrInvalidGenMapParams
	}

	return nil
}

func LoadGenMapParams(fn string) (*GenMapParams, error) {
	dat, err := os.ReadFile(fn)
	if err != nil {
		goutils.Error("LoadGenMapParams:ReadFile",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	params := &GenMapParams{}

	err = yaml.Unmarshal(dat, &params)
	if err != nil {
		goutils.Error("LoadGenMapParams:Unmarshal",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	err = params.check()
	if err != nil {
		goutils.Error("LoadGenMapParams:check",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	return params, nil
}

func genMap(mgr *MapTileMgr, w, h int, params *GenMapParams) (*MapData, error) {
	md := NewMapData(w, h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ct := params.RandAMapTile()
			// if ct == -1 {
			// 	goutils.Error("genMap:RandAMapTile",
			// 		zap.Int("x", x),
			// 		zap.Int("y", y),
			// 		zap.Error(ErrInvalidRandResult))

			// 	return nil, ErrInvalidRandResult
			// }

			md.Data[y][x] = MapDataType(ct)
		}
	}

	return md, nil
}
