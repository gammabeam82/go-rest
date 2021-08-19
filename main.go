package main

import (
	"frm/config"
	"frm/di"
	"frm/server"
	"frm/store"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
)

func main() {
	container, err := di.BuildContainer()

	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(func(r *mux.Router, c *config.Config, db *gorm.DB) {
		defer func(conn *gorm.DB) {
			_ = store.CloseConnection(conn)
		}(db)

		server.Run(r, c.HttpPort())
	})

	if err != nil {
		log.Fatal(err)
	}
}
