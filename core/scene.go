package sim3core

type Scene struct {
	MapData *MapData

	MgrPersonDNA *PersonDNAMgr
}

func NewScene(mapData *MapData, mgrPersonDNA *PersonDNAMgr) *Scene {
	return &Scene{
		MapData:      mapData,
		MgrPersonDNA: mgrPersonDNA,
	}
}
