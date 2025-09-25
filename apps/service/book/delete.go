package book

import (
	"context"

	"github.com/vizucode/gokit/logger"
)

func (uc *book) DeleteBook(ctx context.Context, id uint) (err error) {

	if err = uc.db.DeleteBook(ctx, id); err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return nil
}
