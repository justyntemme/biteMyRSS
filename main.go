package main

import (
	"github.com/justyntemme/biteMyRSS/rss"
	"github.com/justyntemme/biteMyRSS/web"
)

func main() {
	rss.LoadRSS()
	web.StartServer()

}
