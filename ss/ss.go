package ss

import (
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Server      string `json:"server"`
	Server_port int    `json:"server_port"`
	//Local       string `json:"local"`
	Local_port int    `json:"local_port"`
	Password   string `json:"password"`
	Method     string `json:"method"`
	Timeout    int    `json:"timeout"`
}

func NewConfig(args ...int) *Config {
	sslocal := new(Config)
	sslocal.Local_port = 1081
	if len(args) >= 1 {
		sslocal.Local_port = args[0]
	}
	sslocal.Timeout = 300
	return sslocal
}

func (this *Config) ReadConfig(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	//log.Println(bytes)
	var configs Config

	//err = json.Unmarshal(bytes, &configs)
	err = ffjson.Unmarshal(bytes, &configs)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//log.Println(configs)
	this = &configs
	return nil
}

func (this *Config) WriteConfig(filep string, filename string) error {
	bytes, err := ffjson.Marshal(*this)
	if err != nil {
		log.Fatal(err)
	}
	return ioutil.WriteFile(filepath.Join(filep, filename), bytes, os.ModePerm)
}
