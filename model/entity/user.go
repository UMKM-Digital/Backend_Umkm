package entity

import "umkm/model/domain"

type UserEntity struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func ToUserEntity(user domain.Users) UserEntity {
	return UserEntity{
		Id: user.IdUser,
	    Username: user.Username,
		Email:   user.Email,
	}
}

func ToUserListEntity(buyers []domain.Users) []UserEntity {
	var result []UserEntity
	for _, buyer := range buyers {
		result = append(result, ToUserEntity(buyer))
	}
	return result
}