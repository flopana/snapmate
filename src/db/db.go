package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"snapmate/config"
	"snapmate/logger"
)

func Migrate() error {
	l := logger.NewLogger()
	db, err := getDb()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Snapshot{})
	if err != nil {
		l.Error("Could not migrate database")
		return err
	}

	return nil
}

func getDb() (*gorm.DB, error) {
	l := logger.NewLogger()
	conf := config.GetConfig()
	db, err := gorm.Open(sqlite.Open(conf.DatabasePath), &gorm.Config{})
	if err != nil {
		l.Error("Could not connect to database")
		return nil, err
	}
	return db, nil
}

func CreateSnapshot(name string, comment string) (*Snapshot, error) {
	l := logger.NewLogger()
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	snapshot := Snapshot{Name: name, Comment: comment}
	result := db.Create(&snapshot)
	if result.Error != nil {
		l.Error("Could not create snapshot in database")
		return nil, result.Error
	}

	return GetSnapshotById(snapshot.ID)
}

func GetSnapshotById(id uint) (*Snapshot, error) {
	l := logger.NewLogger()
	db, err := getDb()
	if err != nil {
		return nil, err
	}

	var snapshot Snapshot
	result := db.First(&snapshot, id)
	if result.Error != nil {
		l.Error("Could not get snapshot")
		return nil, result.Error
	}
	return &snapshot, nil
}
