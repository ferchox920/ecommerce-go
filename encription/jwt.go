package encryption

import (
	"github.com/fercho920/ecommerce-go/models"
	"github.com/golang-jwt/jwt/v5"
)

func SignedLoginToken(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name":  u.FirstName,
	})

	jwtString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return jwtString, nil
}
