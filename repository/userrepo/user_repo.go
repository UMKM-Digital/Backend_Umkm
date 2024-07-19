package userrepo

import (
	
	"umkm/model/domain"
)

type AuthUserRepo interface {
	RegisterRequest(user domain.Users)(domain.Users, error)
	FindUserByEmail(email string) (*domain.Users, error)
	
}