package model

import (
	"frm/request"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Role string

const (
	RoleUser       Role = "user"
	RoleAdmin      Role = "admin"
	RoleSuperAdmin Role = "super admin"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"not null" json:"username"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Password  string    `gorm:"not null;column:upassword" json:"-"`
	Role      string    `gorm:"not null" json:"role"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

type Users []User

func (u *User) Rename(r *request.UpdateUserRequest) {
	u.Username = r.Username
}

func (u *User) CanDelete(user *User) bool {
	return Role(u.Role) == RoleSuperAdmin && Role(user.Role) != RoleSuperAdmin
}

func (u *User) CanUpdate(user *User) bool {
	return Role(u.Role) == RoleSuperAdmin || u.ID == user.ID
}

func (u *User) CanChangeRole(user *User) bool {
	return Role(u.Role) == RoleSuperAdmin && Role(user.Role) != RoleSuperAdmin
}

func NewUser(c *request.CreateUserRequest) *User {
	encodedPassword, _ := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)

	user := &User{
		Username:  c.Username,
		Email:     c.Email,
		Password:  string(encodedPassword),
		Role:      string(RoleUser),
		CreatedAt: time.Now(),
	}

	return user
}
