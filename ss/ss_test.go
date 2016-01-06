package ss

import (
	"testing"
)

func TestConfig_WriteConfig(t *testing.T) {
	sslocal := NewConfig(1080)
	sslocal.Server = "127.0.0.1"
	sslocal.Server_port = 443
	sslocal.Timeout = 300
	sslocal.Method = "aes-256-cfb"
	sslocal.Password = "as23dffx"
	sslocal.WriteConfig("config1.json")
}

func TestConfig_ReadConfig(t *testing.T) {
	sslocal := NewConfig()
	sslocal.ReadConfig("config1.json")
	if sslocal.Method != "aes-256-cfb" {
		t.Error(sslocal)
	}
}
