package sim3core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhs007/goutils"
)

func Test_LoadMapTileMgr(t *testing.T) {
	goutils.InitLogger("test", "", "debug", true, ".")

	mgr, err := LoadMapTileMgr("../gamedata/maptile.csv")
	assert.NoError(t, err)
	assert.NotNil(t, mgr)

	t.Logf("Test_LoadMapTileMgr OK")
}
