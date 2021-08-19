package di

import (
	"frm/bus"
	"frm/config"
	"frm/controller"
	"frm/handler"
	"frm/repository"
	"frm/router"
	"frm/security"
	"frm/store"
	"frm/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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

	_ = container.Provide(func() *validator.Validate {
		return validation.NewValidator()
	})

	_ = container.Provide(func(db *gorm.DB) *repository.UserRepository {
		return repository.NewUserRepository(db)
	})

	_ = container.Provide(func(c *config.Config, r *repository.UserRepository) *security.Authenticator {
		return security.NewAuthenticator(c, r)
	})

	_ = container.Provide(func() *bus.EventBus {
		b := bus.NewEventBus()

		b.Register(bus.UserSubscriber{})

		return b
	})

	_ = container.Provide(func(r *repository.UserRepository, v *validator.Validate, b *bus.EventBus) *handler.UserHandler {
		return handler.NewUserHandler(r, v, b)
	})

	_ = container.Provide(func(c *config.Config, a *security.Authenticator, h *handler.UserHandler) *mux.Router {
		return router.NewRouter(
			c,
			controller.NewIndexController(),
			controller.NewSecurityController(a),
			controller.NewUserController(h),
		)
	})

	return container, nil
}
