package validation

import (
	"encoding/json"
	"errors"

	resterr "github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptBr_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ptBr := pt_BR.New()
		unt := ut.New(ptBr, ptBr)
		transl, _ = unt.GetTranslator("pt_BR")
		ptBr_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateError(validation_err error) *resterr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return resterr.NewBadRequestError(jsonErr.Error())
	} else if errors.As(validation_err, &jsonValidationError) {
		var causes []resterr.Causes
		for _, err := range jsonValidationError {
			cause := resterr.Causes{
				Field:   err.Field(),
				Message: err.Translate(transl),
			}
			causes = append(causes, cause)
		}
		return resterr.NewBadRequestValidationError("Alguns campos são inválidos", causes)
	} else {
		return resterr.NewBadRequestError("Erro ao tentar converter campos")
	}
}
