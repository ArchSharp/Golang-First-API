package Database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewDBPool(config *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error in connecting to PostgreSQL-1")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error in connecting to PostgreSQL-2")
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL")
	return db, nil
}

func Connection(configs *Config) gin.HandlerFunc {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		configs.Host, configs.Port, configs.User, configs.Password, configs.DBName, configs.SSLMode,
	)
	log.Println("Connection: 0", connStr)
	config, err := pgxpool.ParseConfig(connStr)
	log.Println("Connection: 1", config.ConnString())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection: 2")
	//pool, err := pgxpool.ConnectConfig(context.Background(), config)
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection: 3")
	return func(c *gin.Context) {
		c.Set("pool", pool)
		c.Next()
	}
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

// func main() {
// 	db, err := NewDBPool()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Now you can use 'db' in your repositories or other parts of your application.
// }
