package main

import (
	"github.com/kaonmir/OAuth/config"
	"github.com/kaonmir/OAuth/db"
	"github.com/kaonmir/OAuth/server"
)

func main() {
	config.Init()
	db.Init()
	server.Init()
}
