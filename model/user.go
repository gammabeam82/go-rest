package model

import (
	"frm/request"
	"frm/security"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Password  string    `gorm:"not null;column:upassword" json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) CanChangeRole(user *User) bool {
	return u.Role == security.RoleSuperAdmin && u.ID != user.ID
}

func (u *User) CanDelete(user *User) bool {
	return u.Role == security.RoleSuperAdmin && u.ID != user.ID
}

func (u *User) CanUpdate(user *User) bool {
	return u.Role == security.RoleSuperAdmin || u.ID == user.ID
}

func (u *User) Rename(r *request.UpdateUserRequest) {
	u.Username = r.Username
}

func NewUser(c *request.CreateUserRequest) *User {
	user := &User{
		Username:  c.Username,
		Email:     c.Email,
		Password:  security.EncodePassword(c.Password),
		Role:      security.RoleUser,
		CreatedAt: time.Now(),
	}

	return user
}

type Users []User
