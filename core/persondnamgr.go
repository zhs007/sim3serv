package sim3core

import (
	"github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type PersonDNAParams struct {
	WeightsGetTargetDir *goutils.MapWeights
}

type PersonDNAMgr struct {
	MapGetTargetDir map[string]FuncGetTargetDir
}

func (mgr *PersonDNAMgr) RegGetTargetDir(name string, funcGetTargetDir FuncGetTargetDir) error {
	_, isok := mgr.MapGetTargetDir[name]
	if isok {
		goutils.Error("PersonDNAMgr.RegGetTargetDir",
			zap.String("name", name),
			zap.Error(ErrInvalidFuncGetTargetDir))

		return ErrInvalidFuncGetTargetDir
	}

	mgr.MapGetTargetDir[name] = funcGetTargetDir

	return nil
}

func (mgr *PersonDNAMgr) NewPersonDNAParams() *PersonDNAParams {
	params := &PersonDNAParams{
		WeightsGetTargetDir: goutils.NewMapWeights(),
	}

	for k := range mgr.MapGetTargetDir {
		params.WeightsGetTargetDir.AddWeight(k, 1, true)
	}

	return params
}

func (mgr *PersonDNAMgr) NewPersonDNA(params *PersonDNAParams) *PersonDNA {
	dna := &PersonDNA{}

	nameGetTargetDir := params.WeightsGetTargetDir.Rand()
	dna.GetTargetDir = mgr.MapGetTargetDir[nameGetTargetDir]

	return dna
}

func newPersonDNAMgr() *PersonDNAMgr {
	mgr := &PersonDNAMgr{
		MapGetTargetDir: make(map[string]FuncGetTargetDir),
	}

	return mgr
}

var MgrPersonDNA *PersonDNAMgr

func init() {
	MgrPersonDNA = newPersonDNAMgr()

	MgrPersonDNA.RegGetTargetDir("rand", randGetTargetDir)
	MgrPersonDNA.RegGetTargetDir("left", leftGetTargetDir)
	MgrPersonDNA.RegGetTargetDir("right", rightGetTargetDir)
	MgrPersonDNA.RegGetTargetDir("up", upGetTargetDir)
	MgrPersonDNA.RegGetTargetDir("down", downGetTargetDir)
	MgrPersonDNA.RegGetTargetDir("leftdown", leftdownGetTargetDir)
	MgrPersonDNA.RegGetTargetDir("leftup", leftupGetTargetDir)
	MgrPersonDNA.RegGetTargetDir("rightdown", rightdownGetTargetDir)
	MgrPersonDNA.RegGetTargetDir("rightup", rightupGetTargetDir)
}
