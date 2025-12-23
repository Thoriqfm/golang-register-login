package jwt

import (
	"errors"
	"golang-register-login/entity"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Interface interface {
	CreateToken(userID uuid.UUID, isAdmin bool) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, error)
	GetLoginUser(c *gin.Context) (*entity.User, error)
}

type JsonWebToken struct {
	SecretKey   string
	ExpiredTime time.Duration
}

type Claims struct {
	UserID  uuid.UUID
	isAdmin bool
	jwt.RegisteredClaims
}

func Init() Interface {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	expiredTime, err := strconv.Atoi(os.Getenv("JWT_EXP_TIME"))

	if err != nil {
		log.Fatalf("Failed to load JWT expired time: %v", err)
	}

	return &JsonWebToken{
		SecretKey:   secretKey,
		ExpiredTime: time.Duration(expiredTime) * time.Hour,
	}
}

func (j *JsonWebToken) CreateToken(userID uuid.UUID, isAdmin bool) (string, error) {
	claims := &Claims{
		UserID:  userID,
		isAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiredTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JsonWebToken) ValidateToken(tokenString string) (uuid.UUID, error) {
	var (
		claim  Claims
		userID uuid.UUID
	)

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return userID, err
	}

	if !token.Valid {
		return userID, errors.New("invalid token")
	}

	userID = claim.UserID

	return userID, nil
}

func (j *JsonWebToken) GetLoginUser(c *gin.Context) (*entity.User, error) {
	user, ok := c.Get("user")
	if !ok {
		return &entity.User{}, errors.New("user not found")
	}

	return user.(*entity.User), nil
}
