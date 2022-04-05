package sim3core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhs007/goutils"
)

func Test_LoadGenMapParams(t *testing.T) {
	goutils.InitLogger("test", "", "debug", true, ".")

	params, err := LoadGenMapParams("../gamedata/genmap.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, params)

	t.Logf("Test_LoadGenMapParams OK")
}

func Test_genMap(t *testing.T) {
	goutils.InitLogger("test", "", "debug", true, ".")

	mgr, err := LoadMapTileMgr("../gamedata/maptile.csv")
	assert.NoError(t, err)
	assert.NotNil(t, mgr)

	params, err := LoadGenMapParams("../gamedata/genmap.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, params)

	md, err := genMap(mgr, 100, 100, params)
	assert.NoError(t, err)
	assert.NotNil(t, md)

	t.Logf("Test_genMap OK")
}
