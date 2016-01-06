package ss

import (
	"testing"
)

func Test_FindFreeNode(t *testing.T) {
	configs := make([]Config, 3)
	ParseFreeNode(&configs, "http://www.ishadowsocks.com")
	for _, n := range configs {
		t.Log(n)
	}
	WriteFiles(&configs)
}
