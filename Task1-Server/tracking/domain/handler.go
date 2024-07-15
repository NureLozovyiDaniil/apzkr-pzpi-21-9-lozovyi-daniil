package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomiddl "github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type Handler struct {
	service Service
}

var validate = validator.New()

func NewHandler(svc Service) Handler {
	return Handler{service: svc}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()
	e.Use(
		echomiddl.AddTrailingSlash(),
		echomiddl.Logger(),
		echomiddl.Recover(),
	)

	apiv1 := e.Group("/api/v1")
	{
		apiv1.POST("/cargo/tracking-info", func(c echo.Context) error {
			var info TrackingInfo
			if err := c.Bind(&info); err != nil {
				return newErrorResponse(http.StatusBadRequest, err.Error())
			}
			if err := validate.Struct(info); err != nil {
				return newErrorResponse(http.StatusBadRequest, err.Error())
			}

			if err := h.service.AddTrackingInfo(info); err != nil {
				return newErrorResponse(http.StatusInternalServerError, err.Error())
			}

			return c.JSON(http.StatusCreated, info)
		})

		apiv1.GET("/cargo/:id/tracking-history", func(c echo.Context) error {
			id := c.Param("id")
			history, err := h.service.GetCargoTrackingHistory(id)
			if err != nil {
				return newErrorResponse(http.StatusInternalServerError, err.Error())
			}

			return c.JSON(http.StatusOK, history)
		})

		//TODO: Finish, by adding user auth token parsing, in order to get user id
		apiv1.POST("/help_requests/:id/feedback", func(c echo.Context) error {
			var feedback SubmitFeedback
			if err := c.Bind(&feedback); err != nil {
				return newErrorResponse(http.StatusBadRequest, err.Error())
			}
			if err := validate.Struct(feedback); err != nil {
				return newErrorResponse(http.StatusBadRequest, err.Error())
			}

			reqObjId, _ := primitive.ObjectIDFromHex(feedback.RequestID)

			err := h.service.SubmitFeedback(Feedback{
				RequestID: reqObjId,
				Comment:   feedback.Comment,
				Rating:    feedback.Rating,
			})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}

			return c.JSON(http.StatusCreated, feedback)
		})

		apiv1.GET("/help_requests/:id/feedback", func(c echo.Context) error {
			id := c.Param("id")
			feedback, err := h.service.GetRequestFeedback(id)
			if err != nil {
				return newErrorResponse(http.StatusInternalServerError, err.Error())
			}

			return c.JSON(http.StatusOK, feedback)
		})
	}

	return e
}
