package Model

import (
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"

	// "github.com/go-playground/validator/v10"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	translator *ut.Translator
	validate   *validator.Validate
	validat    *validator.Validate
)

func init() {
	translator, validat = initializeTranslator()
	validate = initializeValidator(validat)
}

func initializeTranslator() (*ut.Translator, *validator.Validate) {
	translator := en.New()
	validate := validator.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatalf("failed to register translations: %v", err)
	}
	return &trans, validate
}

func initializeValidator(v *validator.Validate) *validator.Validate {
	// v := validator.New()

	_ = v.RegisterTranslation("required", *translator, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field()) //Field()
		return t
	})

	// Register other custom validations here

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return v
}

// Validate performs validation using the initialized validator.
func Validate(v interface{}) error {
	if err := validate.Struct(v); err != nil {
		return err
	}
	return nil
}
