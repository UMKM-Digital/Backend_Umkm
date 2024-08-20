package entity

import "umkm/model/domain"

type UserEntity struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Picture string `json:"picture"`
	Role string `json:"role"`
}

func ToUserEntity(user domain.Users) UserEntity {
	return UserEntity{
		Id: user.IdUser,
	    Username: user.Username,
		Email:   user.Email,
		Picture: user.Picture,
		Role: user.Role,
	}
}