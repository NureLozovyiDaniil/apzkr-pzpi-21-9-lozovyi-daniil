package domain

import (
	"aid/domain/middleware"
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

		middleware.Authentication,
		//commonMiddleware,
	)

	apiv1 := e.Group("/api/v1")
	{
		apiv1.GET("/help_requests", h.GetRequests)
		apiv1.POST("/help_requests", h.CreateHelpRequest)
		apiv1.DELETE("/help_requests/:id", h.DeleteRequest, middleware.Authorization(middleware.AdminRole))
		apiv1.GET("/help_requests/:id", h.GetRequestByID)
		apiv1.PUT("/help_requests/:id", h.UpdateRequestStatus)
		apiv1.PUT("/help_requests/:id/assign", h.AssignRequestToCargo)
		apiv1.GET("/cargos/:id/help_requests", h.GetRequestsByCargo)
	}

	return e
}

func commonMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		return next(c)
	}
}
