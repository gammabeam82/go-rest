package main

import (
	"frm/di"
	"gorm.io/gorm"
	"log"
)

func main() {
	c, err := di.BuildContainer()

	if err != nil {
		log.Fatal(err)
	}

	err = c.Invoke(func(db *gorm.DB) {
		conn, _ := db.DB()

		err = conn.Ping()
	})

	if err != nil {
		log.Fatal(err)
	}
}
