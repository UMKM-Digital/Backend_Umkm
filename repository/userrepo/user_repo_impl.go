package userrepo

import (
	"errors"
	"umkm/model/domain"

	"gorm.io/gorm"
)

type AuthrepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) *AuthrepositoryImpl{
	return &AuthrepositoryImpl{db:db}
}

//regoster
func (repo *AuthrepositoryImpl) RegisterRequest(user domain.Users)(domain.Users, error){
	err := repo.db.Create(&user).Error
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

//login
func (repo *AuthrepositoryImpl) FindUserByEmail(email string) (*domain.Users, error){
	user := new(domain.Users)

	if err	:= repo.db.Where("email = ?", email).Take(&user).Error; err != nil{
		return user, err
	}

	return user, nil
}

//logout
func (repo *AuthrepositoryImpl) GetSeller(Id int) (domain.Users, error){
	var userData domain.Users

	err := repo.db.First(&userData, "id = ?", Id).Error

	if err != nil {
		return domain.Users{}, errors.New("user tidak ditemukan")
	}

	return userData, nil
}