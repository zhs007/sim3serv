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

type MapTileMgr struct {
	MapTile map[int]*MapTile
}

func LoadMapTileMgr(fn string) (*MapTileMgr, error) {
	mgr := &MapTileMgr{
		MapTile: make(map[int]*MapTile),
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
						zap.Error(err))

					return err
				}
				mt.GrowingPeriod = int(i64)
			}
		}

		mgr.MapTile[mt.ID] = mt

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
