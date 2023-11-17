package database

import (
	"fmt"

	"github.com/piovani/go_full/infra/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Open() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo",
		config.Env.HostDB,
		config.Env.UserDB,
		config.Env.PasswordDB,
		config.Env.DatabaseDB,
		config.Env.PortDB,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
