package rest

import (
	"digitalbooklending/apps/domain"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vizucode/gokit/logger"
)

func (r *rest) firstBook(c *fiber.Ctx) error {

	idBook, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Error(c.UserContext(), err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, "error parse id book")
	}

	book, err := r.bookService.FirstBookById(c.UserContext(), uint(idBook))
	if err != nil {
		logger.Log.Error(c.UserContext(), err)
		return err
	}

	return r.ResponseJson(c, http.StatusOK, book, "success get book")
}

func (r *rest) listBooks(c *fiber.Ctx) error {

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		logger.Log.Error(c.UserContext(), err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, "error parse limit")
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		logger.Log.Error(c.UserContext(), err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, "error parse page")
	}

	books, err := r.bookService.GetListBooks(c.UserContext(), limit, page)
	if err != nil {
		logger.Log.Error(c.UserContext(), err)
		return err
	}

	return r.ResponseJson(c, http.StatusOK, books, "success get list books")
}

func (r *rest) updateBook(c *fiber.Ctx) error {

	ctx := c.UserContext()

	idBook, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Error(ctx, err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, "error parse id book")
	}

	var req domain.RequestBook
	if err = c.BodyParser(&req); err != nil {
		logger.Log.Error(c.UserContext(), err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, "error parse request")
	}

	_, err = r.bookService.UpdateBook(ctx, uint(idBook), req)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return r.ResponseJson(c, http.StatusOK, nil, "success update book")
}

func (r *rest) deleteBook(c *fiber.Ctx) error {

	ctx := c.UserContext()

	idBook, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Error(ctx, err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, "error parse id book")
	}

	if err = r.bookService.DeleteBook(ctx, uint(idBook)); err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return r.ResponseJson(c, http.StatusOK, nil, "success delete book")
}

func (r *rest) createBook(c *fiber.Ctx) error {

	ctx := c.UserContext()

	var req domain.RequestBook
	if err := c.BodyParser(&req); err != nil {
		logger.Log.Error(ctx, err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, "error parse request")
	}

	_, err := r.bookService.CreateBook(ctx, req)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return r.ResponseJson(c, http.StatusOK, nil, "success create book")
}
