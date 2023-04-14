package middleware

import (
	"context"
	"errors"
	"net/http"
	"practice/config"
	"practice/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type JWTAuth struct {
	jwtKey []byte
	JWTTTL int
}

func NewJWTAuth(cfg *config.Config) *JWTAuth {
	return &JWTAuth{
		jwtKey: []byte(cfg.JWTKey),
		JWTTTL: cfg.JWTTTL,
	}
}

func (m *JWTAuth) GenerateJWT(username, userRole string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(m.JWTTTL) * time.Second)
	claims := models.JWTClaim{
		Username: username,
		UserRole: userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(m.jwtKey)
}

func (m *JWTAuth) GenerateJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			return err
		}

		userData := c.Get(string(models.ContextKey)).(*models.ContextUserData)
		token, err := m.GenerateJWT(userData.Usename, userData.UserRole)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, token)
	}
}

func (m *JWTAuth) ValidateToken(signedToken string) (*models.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return m.jwtKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.JWTClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func (m *JWTAuth) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := extractToken(c.Request())

		claims, err := m.ValidateToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}

		ctx := context.WithValue(c.Request().Context(), models.ContextKey, &models.ContextUserData{
			Usename:  claims.Username,
			UserRole: claims.UserRole,
		})

		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}
