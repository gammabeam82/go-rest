package main

import (
	"flag"
	"frm/di"
	"frm/model"
	"frm/repository"
	"frm/request"
	"frm/store"
	"github.com/go-playground/validator/v10"
	"log"
)

func main() {
	var username, email, password string

	container, err := di.BuildContainer()

	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&username, "username", "", "")
	flag.StringVar(&email, "email", "", "")
	flag.StringVar(&password, "password", "", "")

	flag.Parse()

	err = container.Invoke(func(repo *repository.UserRepository, v *validator.Validate) {
		req := &request.CreateUserRequest{
			Username:         username,
			Email:            email,
			Password:         password,
			RepeatedPassword: password,
		}

		err = v.Struct(req)

		if err != nil {
			log.Fatal(err)
		}

		user := model.NewUser(req)

		defer func() {
			if err = store.CloseConnection(repo.GetConnection()); err != nil {
				log.Fatal(err)
			}
		}()

		if err = repo.Create(user); err != nil {
			log.Fatal(err)
		}

		log.Printf("User %s was successfully created\n", req.Username)
	})

	if err != nil {
		log.Fatal(err)
	}
}
