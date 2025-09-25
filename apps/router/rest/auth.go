package rest

import (
	"digitalbooklending/apps/domain"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vizucode/gokit/logger"
)

func (r *rest) signin(c *fiber.Ctx) error {
	var req domain.SignInRequest

	if err := c.BodyParser(&req); err != nil {
		logger.Log.Error(c.UserContext(), err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, "error parse request body")
	}

	resp, err := r.authService.SignIn(c.UserContext(), domain.SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		logger.Log.Error(c.UserContext(), err)
		return err
	}

	return r.ResponseJson(c, http.StatusOK, resp, "success signin")
}

func (r *rest) signup(c *fiber.Ctx) error {
	var req domain.SignUpRequest

	if err := c.BodyParser(&req); err != nil {
		logger.Log.Error(c.UserContext(), err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, "error parse request body")
	}

	err := r.authService.SignUp(c.UserContext(), domain.SignUpRequest{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})

	if err != nil {
		logger.Log.Error(c.UserContext(), err)
		return err
	}

	return r.ResponseJson(c, http.StatusOK, nil, "success signup")
}
