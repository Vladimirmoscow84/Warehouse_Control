package auth

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

// UserClaims хранит роль и индентификатор пользователя
type UserClaims struct {
	UserID int
	Role   string
}

// authentificator - интерфейс авторизации
type Authentificator interface {
	GenerateToken(userID int, role string, ttl time.Duration) (string, error)
	CheckToken(tokenStr string) (*UserClaims, error)
}

// JWTAuth реализует интерфейс с JWT
type JWTAuth struct {
	secret []byte
}

// New - конструктор нового JWTAuth
func New(secret string) *JWTAuth {
	return &JWTAuth{secret: []byte(secret)}
}

// GenereateToken генерирует токен
func (j *JWTAuth) GenerateToken(userID int, role string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"exp":    time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.secret)
}

// CheckToken проверят токен и возвращает UserClaims
func (j *JWTAuth) CheckToken(tokenStr string) (*UserClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil || !token.Valid {
		log.Println("[auth] invalid token")
		return nil, errors.New("[auth] invalid token")
	}

	rawID, ok := claims["user_id"]
	if !ok {
		log.Println("[auth] user_id missing in token")
		return nil, errors.New("[auth] user_id missing in token")
	}

	var userID int
	switch v := rawID.(type) {
	case float64:
		userID = int(v)
	case int:
		userID = v
	default:
		log.Println("[auth] wrong user_id format")
		return nil, errors.New("[auth] wrong user_id format")
	}

	role, ok := claims["role"].(string)
	if !ok {
		log.Println("[auth] role not found in token")
		return nil, errors.New("[auth]role not found in token")
	}
	return &UserClaims{
		UserID: int(userID),
		Role:   role,
	}, nil
}
