package userrepo

import (
	"umkm/model/domain"
)

type AuthUserRepo interface {
	RegisterRequest(user domain.Users)(domain.Users, error)
	FindUserByUsername(username string) (*domain.Users, error)
	FindUserByPhone(phone string) (*domain.Users, error)
	GetByID(idUser int) (domain.Users, error)
	// UpdateId(idUser int, user domain.Users) (domain.Users, error)
	FindUserByPhoneRegister(phone string)(*domain.Users, error)
}
	