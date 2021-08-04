package model

import (
	"frm/request"
	"frm/service"
	"time"
)

const (
	RoleUser       = "user"
	RoleAdmin      = "admin"
	RoleSuperAdmin = "super admin"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Password  string    `gorm:"not null;column:upassword" json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) Rename(r *request.UpdateUserRequest) {
	u.Username = r.Username
}

func NewUser(c *request.CreateUserRequest) *User {
	user := &User{
		Username:  c.Username,
		Email:     c.Email,
		Password:  service.EncodePassword(c.Password),
		Role:      RoleUser,
		CreatedAt: time.Now(),
	}

	return user
}

type Users []User
