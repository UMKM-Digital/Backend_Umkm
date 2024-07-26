package helper

import (
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	
)

func GetAuthId(c echo.Context) (int, error) {
	claims := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims)
	authID, err := strconv.Atoi(claims.ID)

	if err != nil {
		return -1, err
	}

	return authID, nil
}