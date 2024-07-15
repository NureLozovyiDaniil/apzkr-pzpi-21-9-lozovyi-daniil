package domain

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func Authenticate(h *Handler, c echo.Context) (*Identity, error) {
	header := c.Request().Header.Get("Authorization")
	if header == "" {
		return nil, newErrorResponse(http.StatusUnauthorized, "empty header")
	}
	// Перевіряємо чи токен відповідає формату
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return nil, newErrorResponse(http.StatusUnauthorized, "wrong header type")
	}
	// Спроба отримати користувача з токену
	user, err := h.svc.ParseToken(headerParts[1])
	if err != nil {
		return nil, newErrorResponse(http.StatusUnauthorized, err.Error())
	}

	return &user, nil
}

func (h *Handler) Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := Authenticate(h, c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized"+err.Error())
		}

		c.Set("user", user)
		return next(c)
	}
}

func OwnershipAuthorization() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*Identity)
			if !ok {
				return newErrorResponse(http.StatusUnauthorized, "cannot get user from context")
			}

			id := c.Param("id")
			if user.Id != id {
				return newErrorResponse(http.StatusForbidden, "you do not have permission to access this resource")
			}

			return next(c)
		}
	}
}

func AdminAuthorization() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*Identity)
			if !ok {
				return newErrorResponse(http.StatusUnauthorized, "cannot get user from context")
			}

			if user.Role == RoleAdmin {
				return newErrorResponse(http.StatusForbidden, "you do not have permission to access this resource")
			}

			return next(c)
		}
	}
}
