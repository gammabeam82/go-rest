package di

import (
	"frm/config"
	"frm/repository"
	"frm/store"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func BuildContainer() (*dig.Container, error) {
	var err error

	container := dig.New()

	err = container.Provide(func() (*config.Config, error) {
		return config.NewConfig()
	})

	if err != nil {
		return nil, err
	}

	err = container.Provide(func(c *config.Config) (*gorm.DB, error) {
		return store.NewConnection(c)
	})

	if err != nil {
		return nil, err
	}

	err = container.Provide(func(db *gorm.DB) *repository.UserRepository {
		return repository.NewUserRepository(db)
	})

	if err != nil {
		return nil, err
	}

	return container, nil
}
