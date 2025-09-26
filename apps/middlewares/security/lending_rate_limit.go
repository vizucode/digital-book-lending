package security

import (
	"digitalbooklending/helpers/constants/httpstd"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (mw *security) CheckLendingRateLimit(c *fiber.Ctx) error {
	ctx := c.UserContext()

	_, userCtx, ok := ExtractUserContextFiber(c)
	if !ok {
		logger.Log.Error(ctx, errorkit.NewErrorStd(http.StatusUnauthorized, "", httpstd.Unauthorized))
		return errorkit.NewErrorStd(http.StatusUnauthorized, "", httpstd.Unauthorized)
	}

	err := mw.db.CheckLimitBorrow(ctx, uint(userCtx.Id))
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return c.Next()
}
