package domain

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) CreateCargo(c echo.Context) error {
	var cargo Cargo

	if err := c.Bind(&cargo); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	err := h.svc.CreateCargo(cargo)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, statusResponse{
		Status: "created",
	})
}

func (h *Handler) GetCargoByID(c echo.Context) error {
	id := c.Param("id")

	cargo, err := h.svc.GetCargoByID(id)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cargo)
}

func (h *Handler) UpdateCargoStatus(c echo.Context) error {
	id := c.Param("id")
	status := c.QueryParam("status")

	err := h.svc.UpdateCargoStatus(id, status)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{
		Status: "updated",
	})
}

func (h *Handler) AssignVolunteerToCargo(c echo.Context) error {
	cargoId := c.Param("id")
	volId := c.QueryParam("volunteer_id")

	err := h.svc.AssignVolunteerToCargo(cargoId, volId)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{
		Status: "updated",
	})
}

func (h *Handler) AddDeliverySteps(c echo.Context) error {
	cargoId := c.Param("id")
	var steps []DeliveryStep

	if err := c.Bind(&steps); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	err := h.svc.AddDeliverySteps(cargoId, steps)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, statusResponse{
		Status: "created",
	})
}

func (h *Handler) CompleteDeliveryStep(c echo.Context) error {
	cargoId := c.Param("id")
	stepId := c.QueryParam("step_id")
	photo := c.QueryParam("photo")
	rfidScanned := c.QueryParam("rfid_scanned")
	lossDescription := c.QueryParam("loss_description")

	err := h.svc.CompleteDeliveryStep(cargoId, stepId, photo, rfidScanned == "true", lossDescription)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{
		Status: "updated",
	})
}

func (h *Handler) GetCargoByRFID(c echo.Context) error {
	rfid := c.Param("rfid")

	cargo, err := h.svc.GetCargoByRFID(rfid)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cargo)
}
