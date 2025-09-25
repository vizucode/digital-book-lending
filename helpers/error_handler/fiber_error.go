package errorhandler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"digitalbooklending/helpers/constants/rpcstd"
	"digitalbooklending/helpers/translator"

	"github.com/vizucode/gokit/utils/errorkit"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type stdRespError struct {
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
}

// ErrorHandler handles errors and returns an appropriate response.
func FiberErrHandler(ctx *fiber.Ctx, err error) error {

	// Check if the error is a validation error
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, val := range validationErr {
			switch strings.ToLower(val.Tag()) {
			case "required":
				// translate
				targetLang := ctx.Locals("lang").(string)
				sourceMessage := strings.ToLower(val.Field()) + " is required"
				messsage, _ := translator.Translate(sourceMessage, "en", targetLang)

				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    messsage,
				})
			case "validate_user_contact":
				// translate
				targetLang := ctx.Locals("lang").(string)
				sourceMessage := strings.ToLower(val.Field()) + " must be a valid email address or phone number"
				messsage, _ := translator.Translate(sourceMessage, "en", targetLang)

				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    messsage,
				})
			case "password_regex_validator":
				// translate
				targetLang := ctx.Locals("lang").(string)
				sourceMessage := strings.ToLower(val.Field()) + " character must at least 1 uppercase letter, 1 lowercase letter, have a number and 1 special character"
				messsage, _ := translator.Translate(sourceMessage, "en", targetLang)

				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    messsage,
				})
			case "min":
				// translate
				targetLang := ctx.Locals("lang").(string)
				sourceMessage := strings.ToLower(val.Field()) + " must be at least " + val.Param() + " characters long"
				messsage, _ := translator.Translate(sourceMessage, "en", targetLang)

				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    messsage,
				})
			case "indonesia_phone":
				// translate
				targetLang := ctx.Locals("lang").(string)
				sourceMessage := "is not a valid Indonesian phone number format, please enter a valid Indonesian phone number"
				messsage, _ := translator.Translate(sourceMessage, "en", targetLang)

				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    messsage,
				})
			case "email":
				// translate
				targetLang := ctx.Locals("lang").(string)
				sourceMessage := strings.ToLower("email address is not valid, please enter a valid email address")
				messsage, _ := translator.Translate(sourceMessage, "en", targetLang)

				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    messsage,
				})
			case "max":
				// translate
				targetLang := ctx.Locals("lang").(string)
				sourceMessage := fmt.Sprintf("the length of %s must be at most %s characters", strings.ToLower(val.Field()), val.Param())
				messsage, _ := translator.Translate(sourceMessage, "en", targetLang)

				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    messsage,
				})
			}
		}
	}

	// check std erro
	var stdError *errorkit.ErrorStd
	ok := errors.As(err, &stdError)
	if ok {
		return ctx.Status(stdError.HttpStatusCode).JSON(stdRespError{
			StatusCode: stdError.ErrorCode(),
			Message:    strings.ToLower(stdError.Error()),
		})
	}

	if err != nil {
		// translate
		return ctx.Status(http.StatusNotFound).JSON(stdRespError{
			StatusCode: "404" + rpcstd.NOT_FOUND,
			Message:    "data was not found",
		})
	}

	return nil
}
