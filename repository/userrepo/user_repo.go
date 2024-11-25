package userrepo

import (
	"umkm/model/domain"
)

type AuthUserRepo interface {
	RegisterRequest(user domain.Users)(domain.Users, error)
	FindUserByUsername(username string) (*domain.Users, error)
	FindUserByPhone(phone string) (*domain.Users, error)
	GetByID(idUser int) (domain.Users, error)
	UpdateId(idUser int, user domain.Users) (domain.Users, error)
	FindUserByPhoneRegister(phone string)(*domain.Users, error)
	CekInPassword(userId int) (*domain.Users, error)
	UpdatePassword(userId int, newPassword string) error
	FindOrCreateUserByGoogleID(googleID string, email string, nama string, picture string) (*domain.Users, error)
	ChangePassword(email string) (*domain.Users, error)
	ListUser()([]domain.Users, error)
	CountUserByGenderWithPercentage()  (map[string]map[string]float64, error)
	CountUserByStudy() (map[string]int64, error)
	CountUserByAge() (map[string]int64, error)
	DeleteUser(iduser int) error
}
	