package main

import (
	"frm/config"
	"frm/di"
	"frm/server"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	container, err := di.BuildContainer()

	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(func(r *mux.Router, c *config.Config) {
		server.Run(r, c.HttpPort())
	})

	if err != nil {
		log.Fatal(err)
	}
}
