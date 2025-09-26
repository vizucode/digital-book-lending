package rest

import (
	"strconv"

	"digitalbooklending/apps/domain"
	"digitalbooklending/apps/middlewares"
	"digitalbooklending/apps/service"

	"github.com/gofiber/fiber/v2"
)

type rest struct {
	mw             middlewares.IMiddleware
	authService    service.IAuth
	bookService    service.IBook
	lendingService service.IBookLending
}

func NewRest(
	mw middlewares.IMiddleware,
	authService service.IAuth,
	bookService service.IBook,
	lendingService service.IBookLending,
) *rest {
	return &rest{
		mw:             mw,
		authService:    authService,
		bookService:    bookService,
		lendingService: lendingService,
	}
}

func (r *rest) Router(app fiber.Router) {
	app.Post("/auth/signin", r.signin)
	app.Post("/auth/signup", r.signup)

	app.Post("/books", r.mw.AuthMiddleware, r.createBook)
	app.Get("/books", r.listBooks)
	app.Get("/books/:id", r.mw.AuthMiddleware, r.firstBook)
	app.Delete("/books/:id", r.mw.AuthMiddleware, r.deleteBook)
	app.Put("/books/:id", r.mw.AuthMiddleware, r.updateBook)

	app.Post("/borrow", r.mw.AuthMiddleware, r.borrowBook)
	app.Post("/return", r.mw.AuthMiddleware, r.returnBooks)
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
