package mysql

import (
	"context"
	"digitalbooklending/apps/models"
	"digitalbooklending/helpers/constants/httpstd"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *mysql) CreateLendingRecord(ctx context.Context, lendingRecord models.LendingRecord) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		var book models.Book
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", lendingRecord.BookId).
			First(&book).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errorkit.NewErrorStd(http.StatusNotFound, "", "book not found")
			}
			logger.Log.Error(ctx, err)
			return errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
		}

		if book.Quantity <= 0 {
			return errorkit.NewErrorStd(http.StatusBadRequest, "", "book is out of stock")
		}

		var currentLanding models.LendingRecord
		_ = tx.Model(&models.LendingRecord{}).
			Where("book_id = ?", lendingRecord.BookId).
			Where("user_id = ?", lendingRecord.UserId).
			Where("return_date IS NULL").
			First(&currentLanding)

		if currentLanding.Id > 0 {
			return errorkit.NewErrorStd(http.StatusBadRequest, "", "book is already borrowed")
		}

		lendingRecord.BorrowDate = time.Now()
		lendingRecord.ReturnDate = nil
		if err := tx.Create(&lendingRecord).Error; err != nil {
			logger.Log.Error(ctx, err)
			return errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
		}

		if err := tx.Model(&models.Book{}).
			Where("id = ?", lendingRecord.BookId).
			Update("quantity", gorm.Expr("quantity - ?", 1)).Error; err != nil {
			logger.Log.Error(ctx, err)
			return errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
		}

		// 5. Buat audit log
		auditPayload := models.AuditLog{
			Action:   "borrow_book",
			Entity:   book.Title,
			UserId:   lendingRecord.UserId,
			Details:  fmt.Sprintf("peminjaman buku %s, stok sisa %d", book.Title, book.Quantity-1),
			EntityId: book.Id,
		}
		if err := tx.Create(&auditPayload).Error; err != nil {
			logger.Log.Error(ctx, err)
			return errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
		}

		return nil
	})
}

func (r *mysql) FirstLendingRecordByUserId(ctx context.Context, userId uint) (models.LendingRecord, error) {
	var lendingRecord models.LendingRecord
	result := r.db.WithContext(ctx).Where("user_id = ?", userId).First(&lendingRecord)
	if result.Error != nil {
		return lendingRecord, result.Error
	}
	return lendingRecord, nil
}

func (r *mysql) CountUserLendingRecords(ctx context.Context, userId uint) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&models.LendingRecord{}).Where("user_id = ?", userId).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (r *mysql) UpdateLendingRecord(ctx context.Context, lendingRecord models.LendingRecord) error {
	now := time.Now()

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&models.LendingRecord{}).
			Where("book_id = ?", lendingRecord.BookId).
			Where("user_id = ?", lendingRecord.UserId).
			Where("return_date IS NULL"). // biar nggak return dua kali
			Update("return_date", now)

		if result.Error != nil {
			logger.Log.Error(ctx, result.Error)
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return errorkit.NewErrorStd(http.StatusNotFound, "", httpstd.NotFound)
			}
			return errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
		}
		if result.RowsAffected == 0 {
			return errorkit.NewErrorStd(http.StatusBadRequest, "", "no active lending record found")
		}

		result = tx.Model(&models.Book{}).
			Where("id = ?", lendingRecord.BookId).
			Update("quantity", gorm.Expr("quantity + ?", 1))

		if result.Error != nil {
			logger.Log.Error(ctx, result.Error)
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return errorkit.NewErrorStd(http.StatusNotFound, "", httpstd.NotFound)
			}
			return errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
		}
		if result.RowsAffected == 0 {
			return errorkit.NewErrorStd(http.StatusBadRequest, "", "book not found")
		}

		// first book
		var book models.Book
		tx.Model(&models.Book{}).
			Where("id = ?", lendingRecord.BookId).
			First(&book)

		auditPayload := models.AuditLog{
			Action:   "return_book",
			Entity:   book.Title,
			UserId:   lendingRecord.UserId,
			Details:  fmt.Sprintf("pengembalian buku %s sebanyak %d", book.Title, book.Quantity),
			EntityId: book.Id,
		}

		// create audit log
		result = tx.Create(&auditPayload)
		if result.Error != nil {
			logger.Log.Error(ctx, result.Error)
			return errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
		}

		return nil // commit
	})
}
