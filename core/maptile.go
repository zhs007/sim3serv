package sim3core

import (
	"github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type MapTile struct {
	ID            int
	Name          string
	Type          TileType
	MaxVal        int
	Growth        int
	ChillDown     int
	GrowingPeriod int
}

type MapTileTypeData struct {
	Data map[int]int
}

func newMapTileTypeData() *MapTileTypeData {
	return &MapTileTypeData{
		Data: make(map[int]int),
	}
}

func (mttd *MapTileTypeData) ins(id int) {
	_, isok := mttd.Data[id]
	if !isok {
		mttd.Data[id] = len(mttd.Data)
	}
}

type MapTileMgr struct {
	MapTile     map[int]*MapTile
	MapTileType map[int]*MapTileTypeData
}

func LoadMapTileMgr(fn string) (*MapTileMgr, error) {
	mgr := &MapTileMgr{
		MapTile:     make(map[int]*MapTile),
		MapTileType: make(map[int]*MapTileTypeData),
	}

	err := goutils.LoadCSVFile(fn, func(i int, row []string) bool {
		return i == 0
	}, func(i int, row []string, mapHeader map[int]string) error {
		mt := &MapTile{}

		for j, str := range row {
			switch mapHeader[j] {
			case "id":
				i64, err := goutils.String2Int64(str)
				if err != nil {
					goutils.Error("LoadMapTileMgr:String2Int64",
						zap.Int("row", i),
						zap.Int("col", j),
						zap.String("headname", mapHeader[j]),
						zap.String("v", str),
						zap.String("fn", fn),
						zap.Error(err))

					return err
				}
				mt.ID = int(i64)
			case "name":
				mt.Name = str
			case "type":
				t, err := String2TileType(str)
				if err != nil {
					goutils.Error("LoadMapTileMgr:String2Int64",
						zap.Int("row", i),
						zap.Int("col", j),
						zap.String("headname", mapHeader[j]),
						zap.String("v", str),
						zap.String("fn", fn),
						zap.Error(err))

					return err
				}
				mt.Type = t
			case "maxval":
				i64, err := goutils.String2Int64(str)
				if err != nil {
					goutils.Error("LoadMapTileMgr:String2Int64",
						zap.Int("row", i),
						zap.Int("col", j),
						zap.String("headname", mapHeader[j]),
						zap.String("v", str),
						zap.String("fn", fn),
						zap.Error(err))

					return err
				}
				mt.MaxVal = int(i64)
			case "growth":
				i64, err := goutils.String2Int64(str)
				if err != nil {
					goutils.Error("LoadMapTileMgr:String2Int64",
						zap.Int("row", i),
						zap.Int("col", j),
						zap.String("headname", mapHeader[j]),
						zap.String("v", str),
						zap.String("fn", fn),
						zap.Error(err))

					return err
				}
				mt.Growth = int(i64)
			case "chilldown":
				i64, err := goutils.String2Int64(str)
				if err != nil {
					goutils.Error("LoadMapTileMgr:String2Int64",
						zap.Int("row", i),
						zap.Int("col", j),
						zap.String("headname", mapHeader[j]),
						zap.String("v", str),
						zap.String("fn", fn),
						zap.Error(err))

					return err
				}
				mt.ChillDown = int(i64)
			case "growingperiod":
				i64, err := goutils.String2Int64(str)
				if err != nil {
					goutils.Error("LoadMapTileMgr:String2Int64",
						zap.Int("row", i),
						zap.Int("col", j),
						zap.String("headname", mapHeader[j]),
						zap.String("v", str),
						zap.String("fn", fn),
						zap.Error(err))

					return err
				}
				mt.GrowingPeriod = int(i64)
			}
		}

		err := mgr.insMapTileType(mt)
		if err != nil {
			goutils.Error("LoadMapTileMgr:LoadCSVFile:load",
				zap.String("fn", fn),
				zap.Int("row", i),
				zap.Error(err))

			return err
		}

		return nil
	})
	if err != nil {
		goutils.Error("LoadMapTileMgr:LoadCSVFile",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	return mgr, nil
}

func (mgr *MapTileMgr) GenMap(w, h int, params *GenMapParams) (*MapData, error) {
	return genMap(mgr, w, h, params)
}

func (mgr *MapTileMgr) insMapTileType(mt *MapTile) error {
	_, isok := mgr.MapTile[mt.ID]
	if isok {
		goutils.Error("MapTileMgr:insMapTileType",
			zap.Int("id", mt.ID),
			zap.Error(ErrInvalidMapTileID))

		return ErrInvalidMapTileID
	}

	mgr.MapTile[mt.ID] = mt

	mttd, isok := mgr.MapTileType[int(mt.Type)]
	if isok {
		mttd.ins(mt.ID)
	} else {
		mttd := newMapTileTypeData()

		mttd.ins(mt.ID)

		mgr.MapTileType[int(mt.Type)] = mttd
	}

	return nil
}
