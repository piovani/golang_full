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

func (d *Database) Open() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env.HostDB,
		config.Env.PortDB,
		config.Env.UserDB,
		config.Env.PasswordDB,
		config.Env.DatabaseDB,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Database) Migrate() error {
	db, err := d.Open()
	if err != nil {
		return err
	}

	entities := []any{
		entity.Student{},
	}

	err = db.Migrator().DropTable(entities...)
	if err != nil {
		return err
	}

	return db.AutoMigrate(entities...)
}
