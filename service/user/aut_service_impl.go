package userservice

import (
	"time"
	"strconv"
	"errors"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/web"
	"umkm/repository/userrepo"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	authrepository userrepo.AuthUserRepo
	tokenUseCase   helper.TokenUseCase
}

func Newauthservice(authrepository userrepo.AuthUserRepo, token helper.TokenUseCase) *AuthServiceImpl{
	return &AuthServiceImpl{
		authrepository: authrepository,
		tokenUseCase: token,
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


func (service *AuthServiceImpl) LoginRequest(email string, password string) (map[string]interface{}, error) {
	user, getUserErr := service.authrepository.FindUserByEmail(email)
	if getUserErr != nil {
		return nil, errors.New("email not found")
	}

	if checkPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); checkPassword != nil {
		return nil, errors.New("incorrect password")
	}

	claims := helper.JwtCustomClaims{
		ID:    strconv.Itoa(user.IdUser),
		Name:  user.Username,
		Email: user.Email,
	}

	token, tokenErr := service.tokenUseCase.GenerateAccessToken(claims)
	if tokenErr != nil {
		return nil, tokenErr
	}

	// Hitung waktu kedaluwarsa token
	expirationTime := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	return map[string]interface{}{
		"token":      token,
		"expires_at": expirationTime, // Sertakan waktu kedaluwarsa yang sebenarnya
	}, nil
}

