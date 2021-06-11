package response

import (
	"gin-DevOps/model"
)

type RegisterUserResponse struct {
	User model.User `json:"user"`
}

type UserListResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
    Email string `json:"email"`
	Phone string `json:"phone"`
	Group string `json:"group"`
}

