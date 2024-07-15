package domain

import (
	"aid/domain/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var validate = validator.New()

func (h *Handler) CreateHelpRequest(c echo.Context) error {
	user, ok := c.Get("user").(*middleware.Identity)
	if !ok {
		return newErrorResponse(http.StatusUnauthorized, "cannot get user from context")
	}
	userObjID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}
	var request HelpRequest
	if err := c.Bind(&request); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}
	request.UserID = userObjID
	if err := validate.Struct(request); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	if err := h.svc.CreateHelpRequest(request); err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, request)
}

func (h *Handler) DeleteRequest(c echo.Context) error {
	id := c.Param("id")
	if err := h.svc.DeleteRequest(id); err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{Status: "deleted"})
}

func (h *Handler) GetRequests(c echo.Context) error {
	requests, err := h.svc.GetRequests()
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, requests)
}

func (h *Handler) GetRequestByID(c echo.Context) error {
	id := c.Param("id")
	request, err := h.svc.GetRequestByID(id)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, request)
}

func (h *Handler) UpdateRequestStatus(c echo.Context) error {
	id := c.Param("id")
	status := c.QueryParam("status")
	if status == "" {
		return newErrorResponse(http.StatusBadRequest, "status is required")
	}

	if err := h.svc.UpdateRequestStatus(id, status); err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{Status: "updated"})
}

func (h *Handler) AssignRequestToCargo(c echo.Context) error {
	requestID := c.Param("id")
	cargoID := c.QueryParam("cargo_id")
	if cargoID == "" {
		return newErrorResponse(http.StatusBadRequest, "cargo_id is required")
	}

	if err := h.svc.AssignRequestToCargo(requestID, cargoID); err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{Status: "assigned"})
}

func (h *Handler) GetRequestsByCargo(c echo.Context) error {
	cargoID := c.Param("id")
	requests, err := h.svc.GetRequestsByCargo(cargoID)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, requests)
}
