package security

import (
	"context"

	"digitalbooklending/apps/domain"
	contextkeys "digitalbooklending/helpers/constants/context_keys"

	"github.com/gofiber/fiber/v2"
)

func ExtractUserContextFiber(c *fiber.Ctx) (responseCtx context.Context, resp domain.UserContext, ok bool) {
	resultUserContex, ok := c.Locals(contextkeys.UserContext).(domain.UserContext)

	// set context with user info
	ctx := context.WithValue(c.UserContext(), contextkeys.UserContext, resultUserContex)
	return ctx, resultUserContex, ok
}

func ExtractUserContext(ctx context.Context) (resp domain.UserContext, ok bool) {
	resp, ok = ctx.Value(contextkeys.UserContext).(domain.UserContext)
	return resp, ok
}
