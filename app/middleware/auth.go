package middleware

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
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
		"admin",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.Expired))).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJwT))

	return token
}

func GetUser(c echo.Context) *JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)
	return claims
}
