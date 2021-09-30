package middleware

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJwt struct {
	SecretJwT string
	Expired   int64
}

func (jwtConf *ConfigJwt) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJwT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return c.JSON(http.StatusForbidden, e)
		}),
	}
}

func (jwtConf *ConfigJwt) GenerateToken(userID int) string {
	claims := JWTCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(jwtConf.Expired)).Unix(),
		},
	}
	initToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := initToken.SignedString([]byte(jwtConf.SecretJwT))
	return token
}

func GetUser(c echo.Context) *JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)
	return claims
}
