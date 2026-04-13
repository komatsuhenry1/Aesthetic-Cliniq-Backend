package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string, userRole string, clinicId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":       userId,
		"role":      userRole,
		"clinic_id": clinicId,
		"exp":       time.Now().Add(time.Hour * 8).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func GetUserIdAndRole(c *gin.Context) (string, string) {
	claims, exists := c.Get("claims")
	if !exists {
		return "", ""
	}
	userId, ok := claims.(jwt.MapClaims)["sub"].(string)
	if !ok {
		return "", ""
	}
	userRole, ok := claims.(jwt.MapClaims)["role"].(string)
	if !ok {
		return "", ""
	}
	return userId, userRole
}

