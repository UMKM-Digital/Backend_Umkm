package userservice

import (
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/web"
	"umkm/repository/userrepo"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	authrepository userrepo.AuthUserRepo
}

func Newauthservice(authrepository userrepo.AuthUserRepo) *AuthServiceImpl{
	return &AuthServiceImpl{
		authrepository: authrepository,
	}
}

func (service *AuthServiceImpl) RegisterRequest(user web.RegisterRequest)(map[string]interface{}, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if errHash != nil {
		return nil, errHash
	}

	user.Password = string(passHash)
	newUser := domain.Users{
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Role: user.Role,
	}

	saveUser, errSaveUser := service.authrepository.RegisterRequest(newUser)
	if errSaveUser != nil {
		return nil, errSaveUser
	}

	return helper.ResponseToJson{"username": saveUser.Username, "email": saveUser.Email}, nil
}
