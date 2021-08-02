package store

import (
	"frm/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(c *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(c.DatabaseUrl()))
}

func CloseConnection(conn *gorm.DB) error {
	db, err := conn.DB()

	if err != nil {
		return err
	}

	if err = db.Close(); err != nil {
		return err
	}

	return nil
}
