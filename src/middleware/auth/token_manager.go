package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ivanauliaa/go-appoinment/src/utils"
	"github.com/labstack/echo/v4"
)

type Claims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

func GenerateAccessToken(id uint) string {
	accessTokenAge, _ := time.ParseDuration(os.Getenv(("ACCESS_TOKEN_AGE")))
	accessTokenSecretKey := os.Getenv("ACCESS_TOKEN_KEY")

	return generateToken(id, accessTokenAge, accessTokenSecretKey)
}

func GenerateRefreshToken(id uint) string {
	refreshTokenAge, _ := time.ParseDuration(os.Getenv(("REFRESH_TOKEN_AGE")))
	refreshTokenSecretKey := os.Getenv("REFRESH_TOKEN_KEY")

	return generateToken(id, refreshTokenAge, refreshTokenSecretKey)
}

func VerifyRefreshToken(refreshToken string) (uint, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_KEY")), nil
	})
	if err != nil {
		return 0, fmt.Errorf("invalid refresh token")
	}

	return uint(claims["id"].(float64)), nil
}

func generateToken(id uint, expirationTime time.Duration, secretKey string) string {
	claims := &Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}

func GetAuthCredential(accessToken string) (uint, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
	})
	if err != nil {
		return 0, fmt.Errorf("invalid access token")
	}

	return uint(claims["id"].(float64)), nil
}

func JWTErrorChecker(err error, c echo.Context) error {
	return c.JSON(utils.ClientErrorResponse(http.StatusUnauthorized, "unauthorized"))
}
