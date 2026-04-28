package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id int64, username, secretkey string) (string, error) {

	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(60 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	key := []byte(secretkey)
	tokenStr, err := token.SignedString(key)

	return tokenStr, err

}

func ValidateToken(tokenStr, secretkey string, withClaimsValidation bool) (int64, string, error) {

	var (
		key    = []byte(secretkey)
		claims = jwt.MapClaims{}
		token  *jwt.Token
		err    error
	)

	token, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return 0, "", err
	}

	if withClaimsValidation && !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	userID := int64(claims["id"].(float64))
	username := claims["username"].(string)
	return userID, username, nil
}
