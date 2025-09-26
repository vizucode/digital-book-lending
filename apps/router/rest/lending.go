package rest

import (
	"digitalbooklending/apps/domain"
	"digitalbooklending/apps/middlewares/security"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vizucode/gokit/logger"
)

func (r *rest) borrowBook(c *fiber.Ctx) error {

	ctx, _, ok := security.ExtractUserContextFiber(c)
	if !ok {
		logger.Log.Error(ctx, "unauthorized")
		return r.ResponseJson(c, http.StatusUnauthorized, nil, "user unauthorized")
	}

	var req domain.RequestLending
	if err := c.BodyParser(&req); err != nil {
		logger.Log.Error(ctx, err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, "failed to parse request body")
	}

	if err := r.lendingService.BorrowBooks(ctx, req); err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return r.ResponseJson(c, http.StatusOK, nil, "book borrowed successfully")
}

func (r *rest) returnBooks(c *fiber.Ctx) error {

	ctx, _, ok := security.ExtractUserContextFiber(c)
	if !ok {
		logger.Log.Error(ctx, "unauthorized")
		return r.ResponseJson(c, http.StatusUnauthorized, nil, "user unauthorized")
	}

	var req domain.RequestReturnBook
	if err := c.BodyParser(&req); err != nil {
		logger.Log.Error(ctx, err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, "failed to parse request body")
	}

	if err := r.lendingService.ReturnBooks(ctx, req); err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return r.ResponseJson(c, http.StatusOK, nil, "book returned successfully")
}
