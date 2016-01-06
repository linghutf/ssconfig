package ss

import (
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"log"
	"os"
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

func NewConfig(port ...int) *Config {
	sslocal := new(Config)
	if len(port) >= 1 {
		sslocal.Local_port = port[0]
	}
	sslocal.Local_port = 1081
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

func (this *Config) WriteConfig(filename string) error {
	bytes, err := ffjson.Marshal(*this)
	if err != nil {
		log.Fatal(err)
	}
	return ioutil.WriteFile(filename, bytes, os.ModePerm)
}
