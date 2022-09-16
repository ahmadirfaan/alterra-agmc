package utils

import (
	"alterra-agmc-day3/app"
	"alterra-agmc-day3/models/database"
	"errors"
	"github.com/golang-jwt/jwt"
	"strconv"
	"strings"
	"time"
)

// This function to generates Access Token
func GenerateToken(user database.User) (*string, string, error) {
	application := app.Init()
	jwtSecretKey := application.Config.JWTSecret

	// Generates Access Token Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["userId"] = user.Id
	claims["exp"] = time.Now().UTC().Add(time.Hour * 24 * 7).Unix()
	format := time.Unix(time.Now().UTC().Add(time.Hour*24*7).Unix(), 0).Format(time.RFC3339)
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return nil, format, err
	}
	return &t, format, err
}

func ExtractToken(header string) (string, error) {
	application := app.Init()
	tokenString := header
	tokenString = strings.ReplaceAll(tokenString, " ", "")
	tokenString = strings.ReplaceAll(tokenString, "Bearer", "")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("500|unexpected signing method")
		}
		return []byte(application.Config.JWTSecret), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)
	return strconv.Itoa(int(userId)), err
}
