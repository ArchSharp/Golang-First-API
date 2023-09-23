package Model

import (
	// uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/google/uuid"
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
	Item      string    `json:"item"`
	Owner     *string   `json:"owner,omitempty"`
	Completed bool      `json:"completed"`
}

type AddTodo struct {
	ID        string  `json:"id"`
	Item      string  `json:"item"`
	Owner     *string `json:"owner,omitempty"`
	Completed bool    `json:"completed"`
}

type Repository struct {
	DB *gorm.DB
}

func MigrateTodos(db *gorm.DB) error {
	err := db.AutoMigrate(&Todo{})
	return err
}
