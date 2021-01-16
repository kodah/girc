package store

import (
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite"

	"github.com/kodah/girc/store/models"
)

var Service *Database

type Database struct {
	store *gorm.DB
	mu    sync.Mutex
}

func (d *Database) AddMessage(name, user, host, command string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.store.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&models.Message{
			Time:    time.Now().UTC(),
			Name:    name,
			User:    user,
			Host:    host,
			Command: command,
		}).Error
	})
}

func New() error {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		return err
	}

	Service = new(Database)
	Service.store = db

	err = db.AutoMigrate(
		&models.Message{},
		&models.Network{},
		&models.Server{})
	if err != nil {
		return err
	}

	return nil
}
