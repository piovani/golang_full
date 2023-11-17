package database

import (
	"fmt"

	"github.com/piovani/go_full/domain/entity"
	"github.com/piovani/go_full/infra/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Open() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env.HostDB,
		config.Env.PortDB,
		config.Env.UserDB,
		config.Env.PasswordDB,
		config.Env.DatabaseDB,
	)

	return gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        dsn,
	}), &gorm.Config{})
}

func (d *Database) Migrate() error {
	db, err := d.Open()
	if err != nil {
		return err
	}

	return db.AutoMigrate(entity.Student{})
}
