package ss

import (
	"path"
	"testing"
)

func Test_FindFreeNode(t *testing.T) {
	configs := make([]Config, 3)
	ParseFreeNode(&configs, 1080, "http://www.ishadowsocks.com")
	for _, n := range configs {
		t.Log(n)
	}
	WriteFiles(&configs, path.Base("./"))
}
