package ss

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"path"
	"strconv"
	"strings"
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ParseFreeNode(configs *[]Config, port int, url string) {
	doc, err := goquery.NewDocument(url)
	CheckErr(err)

	freeNodeCnts := 3
	infoCnts := 4

	//find 3 free nodes
	cont_sel := doc.Find("#free").Find(".col-lg-4.text-center").Nodes[:freeNodeCnts]
	for i := 0; i < freeNodeCnts; i++ {
		config := make([]string, infoCnts)
		infoNodes := goquery.NewDocumentFromNode(cont_sel[i]).Find("h4").Nodes[:infoCnts]
		//log.Println(infoNodes)
		for j := 0; j < infoCnts; j++ {
			s := goquery.NewDocumentFromNode(infoNodes[j]).Text()
			config[j] = s[strings.Index(s, ":")+1:]
		}
		//log.Println(config)
		var perr error
		ss := NewConfig(port)
		ss.Server = config[0]
		ss.Server_port, perr = strconv.Atoi(config[1])
		CheckErr(perr)
		ss.Password = config[2]
		ss.Method = config[3]

		(*configs)[i] = *ss
	}
}

func WriteFiles(configs *[]Config, filepath string) bool {
	for i, n := range *configs {
		filename := "config" + strconv.Itoa(i) + ".json"
		if n.WriteConfig(filepath, filename) != nil {
			return false
		}
	}
	return true
}

func ScrapUsefulNode(port int) {
	configs := make([]Config, 3)
	ParseFreeNode(&configs, port, "http://www.ishadowsocks.com")
	if WriteFiles(&configs, path.Base("./")) {
		log.Println("[successfully updated].")
	} else {
		log.Println("error occured.")
	}
}
