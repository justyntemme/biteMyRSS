package main

import (
	"github.com/justyntemme/biteMyRSS/config"
	"github.com/justyntemme/biteMyRSS/rss"
	"github.com/justyntemme/biteMyRSS/web"
)

func main() {
	Configuration := new(config.Configuration)
	config.ParseConfig("config.toml", Configuration)
	rss.LoadRSS()
	web.StartServer(":8080")
}
