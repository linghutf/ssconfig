package main

import (
	"./ss"
	"flag"
)

var port *int = flag.Int("p", 1080, "Use: -p <local port>.")

func main() {
	flag.Parse()
	ss.ScrapUsefulNode(*port)
}
