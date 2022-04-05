package sim3core

import (
	"bytes"
	"html/template"
	"os"

	"github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type TileHTMLData struct {
	Type    int
	Val     int
	StrVal  string
	StrType string
}

type MapHTMLData struct {
	XArr []int
	YArr []int
	Data [][]*TileHTMLData
}

func getStrType(mt *MapTile) string {
	if mt != nil {
		if mt.Type == TILE_FOOD {
			return "food"
		}
	}

	return "tileempty"
}

func OutputMapHTML(fn string, tmplfn string, mapData *MapData, mgr *MapTileMgr) error {
	const STR = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	tpl, err := template.ParseFiles(tmplfn)
	if err != nil {
		goutils.Error("OutputMapHTML:Parse",
			zap.String("tmplfn", tmplfn),
			zap.Error(err))

		return err
	}

	params := &MapHTMLData{}

	for y := 0; y < mapData.Height; y++ {
		params.YArr = append(params.YArr, y+1)

		lstdat := []*TileHTMLData{}

		for x := 0; x < mapData.Width; x++ {
			if y == 0 {
				params.XArr = append(params.XArr, x+1)
			}

			ct := mapData.Data[y][x]
			mt := mgr.MapTile[int(ct)]

			td := &TileHTMLData{
				StrType: getStrType(mt),
			}

			if mt != nil {
				td.Type = int(mt.Type)
				td.Val = mgr.MapTileType[int(mt.Type)].Data[int(ct)]
				td.StrVal = goutils.Int2StringWithArr(td.Val, STR)
			}

			lstdat = append(lstdat, td)
		}

		params.Data = append(params.Data, lstdat)
	}

	// f, err := os.Create(fn)
	// if err != nil {
	// 	goutils.Error("OutputMapHTML:Create",
	// 		zap.String("fn", fn),
	// 		zap.Error(err))

	// 	return err
	// }
	// defer f.Close()

	// w := bufio.NewWriter(f)
	w := bytes.NewBufferString("")

	err = tpl.ExecuteTemplate(w, "map.tmpl", params)
	if err != nil {
		goutils.Error("OutputMapHTML:Execute",
			zap.Error(err))

		return err
	}

	err = os.WriteFile(fn, w.Bytes(), 0644)
	if err != nil {
		goutils.Error("OutputMapHTML:WriteFile",
			zap.String("fn", fn),
			zap.Error(err))

		return err
	}

	return nil
}
