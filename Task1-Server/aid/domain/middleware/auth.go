package middleware

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	UserRole = iota + 1
	VolunteerRole
	ManagerRole
	AdminRole
)

type Identity struct {
	ID   string `json:"user_id"`
	Role int    `json:"role"`
}

func Authenticate(req *http.Request) (*Identity, error) {
	var u Identity

	client := &http.Client{}

	sendReq, err := http.NewRequest("GET", "http://localhost:5000/api/v1/identity", nil)
	if err != nil {
		return nil, err
	}

	sendReq.Header.Set("Authorization", req.Header.Get("Authorization"))

	resp, err := client.Do(sendReq)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&u)

	return &u, nil
}

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := Authenticate(c.Request())
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized"+err.Error())
		}

		c.Set("user", user)
		return next(c)
	}
}

func Authorization(requiredRoles ...int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*Identity)

			for _, role := range requiredRoles {
				if user.Role >= role {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "Forbidden, role not allowed")
		}
	}
}
