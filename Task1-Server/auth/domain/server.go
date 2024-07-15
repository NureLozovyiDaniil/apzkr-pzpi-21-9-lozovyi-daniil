package domain

import (
	"github.com/labstack/echo/v4"
	echomiddl "github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return Handler{svc: svc}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()
	e.Use(
		echomiddl.AddTrailingSlash(),
		echomiddl.Logger(),
		echomiddl.Recover(),

		echomiddl.CORSWithConfig(echomiddl.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),

		//h.Authentication,
		//commonMiddleware,
	)

	apiv1 := e.Group("/api/v1")
	{

		apiv1.POST("/register", h.Register)
		apiv1.POST("/login", h.Login)
		apiv1.GET("/identity", h.ParseToken, h.Authentication)
		apiv1.GET("/user", h.GetUser, h.Authentication)
		apiv1.PUT("/user", h.UpdateUser, h.Authentication)
		apiv1.PUT("/user/:id", h.UpdateUserByAdmin, h.Authentication, AdminAuthorization())
		apiv1.DELETE("/user", h.DeleteUser)
		apiv1.DELETE("/user/:id", h.DeleteUserByAdmin, AdminAuthorization())
		apiv1.GET("/users", h.GetAllUsers, AdminAuthorization())
		apiv1.GET("/organizations", h.GetAllOrganizations)
		apiv1.POST("/organization", h.CreateOrganization, h.Authentication)
		apiv1.POST("/organization/join", h.JoinOrganization)
		apiv1.DELETE("/organization/leave", h.LeaveOrganization)
		apiv1.GET("/organization/:id", h.GetOrganization)
		apiv1.POST("/volunteer", h.BecomeVolunteer, h.Authentication)
	}

	return e
}
