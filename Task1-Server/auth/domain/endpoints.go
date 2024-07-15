package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

var validate = validator.New()

func (h *Handler) Register(c echo.Context) error {
	var user User

	if err := c.Bind(&user); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}
	if err := validate.Struct(user); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	err := h.svc.CreateUser(user)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, statusResponse{
		Status: "created",
	})
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handler) Login(c echo.Context) error {
	var login LoginRequest

	if err := c.Bind(&login); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}
	if err := validate.Struct(login); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	token, err := h.svc.GenerateToken(login.Username, login.Password)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (h *Handler) ParseToken(c echo.Context) error {
	identity, ok := c.Get("user").(*Identity)
	if !ok {
		return newErrorResponse(http.StatusUnauthorized, "cannot get user from context")
	}

	return c.JSON(http.StatusOK, identity)
}

func (h *Handler) GetUser(c echo.Context) error {
	identity, ok := c.Get("user").(*Identity)
	if !ok {
		return newErrorResponse(http.StatusUnauthorized, "cannot get user from context")
	}

	user, err := h.svc.UserById(identity.Id)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	var user UserUpdate

	identity, ok := c.Get("user").(*Identity)
	if !ok {
		return newErrorResponse(http.StatusUnauthorized, "cannot get user from context")
	}
	user.Id = identity.Id

	if err := c.Bind(&user); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}
	if err := validate.Struct(user); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	err := h.svc.UpdateUser(user)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{
		Status: "updated",
	})
}

func (h *Handler) UpdateUserByAdmin(c echo.Context) error {
	var user UserUpdate

	if err := c.Bind(&user); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}
	if err := validate.Struct(user); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}
	id := c.Param("id")
	if id == "" {
		return newErrorResponse(http.StatusBadRequest, "missing id")
	}
	user.Id = id

	err := h.svc.UpdateUser(user)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{
		Status: "updated",
	})
}

func (h *Handler) DeleteUser(c echo.Context) error {
	identity, ok := c.Get("user").(Identity)
	if !ok {
		return newErrorResponse(http.StatusUnauthorized, "cannot get user from context")
	}

	err := h.svc.DeleteUser(identity.Id)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{
		Status: "deleted",
	})
}

func (h *Handler) DeleteUserByAdmin(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return newErrorResponse(http.StatusBadRequest, "missing id")
	}

	err := h.svc.DeleteUser(id)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{
		Status: "deleted",
	})

}

func (h *Handler) BecomeVolunteer(c echo.Context) error {
	user, ok := c.Get("user").(*Identity)
	if !ok {
		return newErrorResponse(http.StatusUnauthorized, "cannot get user from context")
	}

	err := h.svc.BecomeVolunteer(user.Id)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{
		Status: "updated",
	})
}

func (h *Handler) CreateOrganization(c echo.Context) error {
	user, ok := c.Get("user").(*Identity)
	if !ok {
		return newErrorResponse(http.StatusUnauthorized, "cannot get user from context")
	}

	var org Organization

	if err := c.Bind(&org); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}
	if err := validate.Struct(org); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	err := h.svc.CreateOrganization(user.Id, org)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, statusResponse{
		Status: "created",
	})
}

type OrgMemberRequest struct {
	UserId string `json:"user_id" validate:"required"`
	OrgId  string `json:"org_id" validate:"required"`
}

func (h *Handler) JoinOrganization(c echo.Context) error {
	var req OrgMemberRequest

	if err := c.Bind(&req); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}
	if err := validate.Struct(req); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	err := h.svc.JoinOrganization(req.UserId, req.OrgId)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{
		Status: "joined",
	})
}

func (h *Handler) LeaveOrganization(c echo.Context) error {
	var req OrgMemberRequest

	if err := c.Bind(&req); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}
	if err := validate.Struct(req); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	err := h.svc.LeaveOrganization(req.UserId, req.OrgId)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, statusResponse{
		Status: "left",
	})
}

func (h *Handler) GetOrganization(c echo.Context) error {
	var id string

	if err := c.Bind(&id); err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}
	if id == "" {
		return newErrorResponse(http.StatusBadRequest, "missing id")
	}

	org, err := h.svc.GetOrganizationById(id)
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, org)
}

func (h *Handler) GetAllUsers(c echo.Context) error {
	users, err := h.svc.GetAllUsers()
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetAllOrganizations(c echo.Context) error {
	orgs, err := h.svc.GetAllOrganizations()
	if err != nil {
		return newErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, orgs)
}
