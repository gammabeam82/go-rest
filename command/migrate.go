package main

import (
	"frm/di"
	"frm/model"
	"frm/store"
	"gorm.io/gorm"
	"log"
)

func main() {
	container, err := di.BuildContainer()

	if err != nil {
		log.Fatal(err)
	}

	err = container.Provide(func(conn *gorm.DB) {
		defer func() {
			if err = store.CloseConnection(conn); err != nil {
				log.Fatal(err)
			}
		}()

		if err = conn.AutoMigrate(&model.User{}); err != nil {
			log.Fatal(err)
		}

	})

	if err != nil {
		log.Fatal(err)
	}
}
