package routes

type User struct {
	User_name     string `json:"user_name" validate:"required"`
	User_email    string `json:"user_email" validate:"required"`
	User_password string `json:"user_password" validate:"required"`
}
