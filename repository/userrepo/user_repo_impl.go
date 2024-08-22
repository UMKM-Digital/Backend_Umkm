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

//register
func (repo *AuthrepositoryImpl) RegisterRequest(user domain.Users)(domain.Users, error){
	err := repo.db.Create(&user).Error
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

//login
func (repo *AuthrepositoryImpl) FindUserByUsername(username string) (*domain.Users, error){
	user := new(domain.Users)

	if err	:= repo.db.Where("username = ?", username).Take(&user).Error; err != nil{
		return user, err
	}

	return user, nil
}

//otp
func (repo *AuthrepositoryImpl) FindUserByPhone(phone string) (*domain.Users, error) {
	user := new(domain.Users)

	if err := repo.db.Where("no_phone = ?", phone).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

//profile
func (repo *AuthrepositoryImpl) GetByID(idUser int) (domain.Users, error) {
	var user domain.Users
	if err := repo.db.Where("id = ?", idUser).Take(&user).Error; err != nil {
		return domain.Users{}, errors.New("user not found")
	}
	return user, nil
}


//send otp register
func (repo *AuthrepositoryImpl) FindUserByPhoneRegister(phone string) (*domain.Users, error) {
    user := new(domain.Users)
    result := repo.db.Where("no_phone = ?", phone).Take(user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil // Nomor telepon tidak ditemukan
        }
        return nil, result.Error // Terjadi kesalahan lain
    }
    return user, nil
}


// repository/userrepo/auth_repository.go
// func (repo *AuthrepositoryImpl) UpdateId(idUser int, user domain.Users) (domain.Users, error) {
//     if err := repo.db.Model(&domain.Users{}).Where("id = ?", idUser).Updates(user).Error; err != nil {
//         return domain.Users{}, errors.New("failed to update profile")
//     }
//     return user, nil
// }

//verivy otp

