package main

import (
	"flag"
	"fmt"
	"frm/di"
	"frm/model"
	"frm/repository"
	"frm/request"
	"frm/store"
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

	req := &request.CreateUserRequest{
		Username:         username,
		Email:            email,
		Password:         password,
		RepeatedPassword: password,
	}

	if ok, err := req.Validate(); !ok {
		log.Fatal(err)
	}

	user := model.NewUser(req)

	err = container.Invoke(func(repo *repository.UserRepository) {
		defer func() {
			if err = store.CloseConnection(repo.GetConnection()); err != nil {
				log.Fatal(err)
			}
		}()

		if err = repo.Create(user); err != nil {
			log.Fatal(err)
		}
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("User %s was successfully created", req.Username))
}
