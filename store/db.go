package store

import (
	"frm/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(c *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(c.DatabaseUrl()))
}
