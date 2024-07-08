package tokens

import (
	soldiers_entity "soldiers_service/internal/entity/soldiers"
	cfg "soldiers_service/internal/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(soldiers *soldiers_entity.Soldiers, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = soldiers.ID
	claims["email"] = soldiers.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["role"] = soldiers.Role

	tokenString, err := token.SignedString([]byte(cfg.Token()))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
