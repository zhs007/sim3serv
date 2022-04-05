package sim3core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhs007/goutils"
)

func Test_OutputMapHTML(t *testing.T) {
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

	err = OutputMapHTML("../tmpl/map001.html", "../tmpl/map.tmpl", md, mgr)
	assert.NoError(t, err)

	t.Logf("Test_OutputMapHTML OK")
}
