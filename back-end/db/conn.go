package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*DB, error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL environment variable not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if os.Getenv("AUTO_MIGRATE") == "true" {
		err = db.AutoMigrate(
			&User{}, &Horse{}, &Rank{}, &JumpHeight{}, &Score{},
		)
		if err != nil {
			return nil, err
		}
	}

	if os.Getenv("SEED_DB") == "true" {
		err = SeedDB(&DB{db})
		if err != nil {
			return nil, err
		}
	}

	return &DB{db}, nil
}

type DB struct {
	*gorm.DB
}
