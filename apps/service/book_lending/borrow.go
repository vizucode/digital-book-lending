package booklending

import (
	"context"
	"digitalbooklending/apps/domain"
	"digitalbooklending/apps/middlewares/security"
	"digitalbooklending/apps/models"
	"net/http"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (s *bookLending) BorrowBooks(ctx context.Context, req domain.RequestLending) (err error) {

	userCtx, ok := security.ExtractUserContext(ctx)
	if !ok {
		return errorkit.NewErrorStd(http.StatusUnauthorized, "", "user not authenticated")
	}

	if err = s.validator.VarCtx(ctx, req.BookId, "required"); err != nil {
		return err
	}

	err = s.db.CreateLendingRecord(ctx, models.LendingRecord{
		UserId: uint(userCtx.Id),
		BookId: req.BookId,
	})

	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return nil
}
