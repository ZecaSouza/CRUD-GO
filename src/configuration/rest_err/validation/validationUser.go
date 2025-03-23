package validation

import (
	"encoding/json"
	"errors"

	"github.com/ZecaSouza/CRUD-GO/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enLocale := en.New()
		unt := ut.New(enLocale, enLocale)
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonError *json.UnmarshalTypeError
	var validationErrors validator.ValidationErrors

	if errors.As(validationErr, &jsonError) {
		return rest_err.NewBadRequestError("invalid field type")
	} else if errors.As(validationErr, &validationErrors) {
		errorsCauses := []rest_err.Causes{}

		for _, err := range validationErrors {
			cause := rest_err.Causes{
				Message: err.Translate(transl),
				Field:   err.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestError("invalid input")
	}

	return rest_err.NewBadRequestError("error while validating input")
}
