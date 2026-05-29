package utils

import (
	"banking-system/pkg/config"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userID string, role string, cfg *config.Config) (string, error) {
	jwtKey := []byte(cfg.JWTSecret)
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// func ValidateJWT(tokenString string, cfg *config.Config) (jwt.MapClaims, error) {
// 	jwtKey := []byte(cfg.JWTSecret)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, jwt.ErrSignatureInvalid
// 		}
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims, nil
// 	}
// 	return nil, nil
// }
