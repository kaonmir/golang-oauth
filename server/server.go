package server

import (
	"log"

	"github.com/kaonmir/OAuth/config"
)

func Init() {
	r := NewRouter()
	port := config.Env().Port
	log.Print(port)
	r.Run(port)
}
