package bus

import (
	"frm/model"
	"log"
)

const (
	userCreated = "user.created"
	userUpdated = "user.updated"
)

type UserCreated struct {
	data *model.User
}

func NewUserCreated(u *model.User) *UserCreated {
	return &UserCreated{
		data: u,
	}
}

func (u UserCreated) Name() string {
	return userCreated
}

func (u UserCreated) Data() interface{} {
	return u.data
}

type UserSubscriber struct {
}

func (u UserSubscriber) Supports() Event {
	return UserCreated{}
}

func (u UserSubscriber) Handle(e Event) {
	user, _ := e.Data().(*model.User)

	log.Printf("New user: %s", user.Email)

	//do something important here :)
}
