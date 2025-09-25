package rest

import (
	"strconv"

	"digitalbooklending/apps/domain"
	"digitalbooklending/apps/middlewares"
	"digitalbooklending/apps/service"

	"github.com/gofiber/fiber/v2"
)

type rest struct {
	mw          middlewares.IMiddleware
	authService service.IAuth
}

func NewRest(
	mw middlewares.IMiddleware,
	authService service.IAuth,
) *rest {
	return &rest{
		mw:          mw,
		authService: authService,
	}
}

func (r *rest) Router(app fiber.Router) {
	app.Post("/auth/signin", r.signin)
	app.Post("/auth/signup", r.signup)
}

func (rest *rest) ResponseJson(
	ctx *fiber.Ctx,
	StatusCode int,
	data interface{},
	message string,
) error {
	return ctx.Status(StatusCode).JSON(&domain.ResponseJson{
		Status:  strconv.Itoa(StatusCode),
		Data:    data,
		Message: message,
	})
}
