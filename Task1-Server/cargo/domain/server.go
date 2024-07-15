package domain

import (
	"github.com/labstack/echo/v4"
	echomiddl "github.com/labstack/echo/v4/middleware"
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

		//middleware.Authentication,
		commonMiddleware,
	)

	apiv1 := e.Group("/api/v1")
	{
		apiv1.POST("/cargo", h.CreateCargo)
		apiv1.GET("/cargo/:id", h.GetCargoByID)
		apiv1.PUT("/cargo/:id/status", h.UpdateCargoStatus)
		apiv1.POST("/cargo/:id/assign-volunteer", h.AssignVolunteerToCargo)
		apiv1.POST("/cargo/:id/steps", h.AddDeliverySteps)
		apiv1.PUT("/cargo/:id/step", h.CompleteDeliveryStep)
		apiv1.GET("/cargo/rfid/:rfid", h.GetCargoByRFID)
	}

	return e
}

func commonMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		return next(c)
	}
}
