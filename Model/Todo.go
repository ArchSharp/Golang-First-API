package Model

import (
	// uuid "github.com/jackc/pgtype/ext/gofrs-uuid"

	// "log"
	"log"
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"

	// "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"

	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	"gorm.io/gorm"
)

//	type Todo struct {
//		ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
//		Item      string  `json:"item"`
//		Owner     *string `json:"owner,omitempty"`
//		Completed bool    `json:"completed"`
//	}
type Todo struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Item      string    `json:"item" validate:"required,item"` // custom validation
	Owner     *string   `json:"owner,omitempty" validate:"required,min=2,max=100"`
	Completed bool      `json:"completed"`
}

func (t *Todo) Validate() []string {
	validate := validator.New()
	validate.RegisterValidation("item", validateItem)

	translator := en.New()
	uni := ut.New(translator, translator)

	// // this is usually known or extracted from http 'Accept-Language' header
	// // also see uni.FindTranslator(...)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatal(err)
	}

	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Translate(trans))
		return t
	})

	err := validate.Struct(t)

	errorsArray := []string{}
	for _, e := range err.(validator.ValidationErrors) {
		errText := e.Translate(trans)
		errorsArray = append(errorsArray, errText)
	}

	// return validate.Struct(t)
	return errorsArray
}

func validateItem(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	matches, _ := regexp.MatchString(`^[a-zA-Z ]*$`, str)

	if matches == false {
		return false
	}

	return true
}

func validateUUID(fl validator.FieldLevel) bool {
	// sku is of format abc-absd-dfsdf
	re := regexp.MustCompile(`[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}

	return true
}

type Repository struct {
	DB *gorm.DB
}

func MigrateTodos(db *gorm.DB) error {
	err := db.AutoMigrate(&Todo{})
	return err
}
