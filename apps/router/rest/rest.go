package rest

import (
	"strconv"

	"digitalbooklending/apps/domain"
	"digitalbooklending/apps/middlewares"

	"github.com/gofiber/fiber/v2"
)

type rest struct {
	mw middlewares.IMiddleware
}

func NewRest(
	mw middlewares.IMiddleware,
) *rest {
	return &rest{
		mw: mw,
	}
}

func (r *rest) Router(app fiber.Router) {

}

func (rest *rest) ResponseJson(
	ctx *fiber.Ctx,
	StatusCode int,
	data interface{},
	message string,
) error {
	return ctx.Status(StatusCode).JSON(&domain.ResponseJson{
		StatusCode: strconv.Itoa(StatusCode),
		Data:       data,
		Message:    message,
	})
}
