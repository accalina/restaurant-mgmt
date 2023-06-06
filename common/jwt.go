package common

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/accalina/restaurant-mgmt/exception"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, roles string) string {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)

	claims := jwt.MapClaims{
		"username": username,
		"role":     roles,
		"exp":      time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)
	return tokenSigned
}

func removeBearerFromHeader(authorization string) string {
	const bearerPrefix = "Bearer "
	if len(authorization) > len(bearerPrefix) && authorization[:len(bearerPrefix)] == bearerPrefix {
		return authorization[len(bearerPrefix):]
	}
	return ""
}

func ParseJwt(authorization string) (string, string, error) {
	if authorization == "" {
		return "", "", errors.New("authentication needed")
	}
	tokenString := removeBearerFromHeader(authorization)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return "", "", err
	}
	if !token.Valid {
		return "", "", errors.New("invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	expirationTime := int64(claims["exp"].(float64))
	currentTime := time.Now().Unix()
	if currentTime > expirationTime {
		return "", "", errors.New("token has expired")
	}
	username := claims["username"].(string)
	role := claims["role"].(string)
	return username, role, nil
}
