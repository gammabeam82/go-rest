package request

type CreateUserRequest struct {
	Username         string `json:"username" validate:"required,username"`
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required,password,eqfield=RepeatedPassword"`
	RepeatedPassword string `json:"repeated_password" validate:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"required,username"`
}
