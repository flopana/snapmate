package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
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

	err := checkIfDatabaseExists()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(conf.DatabasePath), &gorm.Config{})
	if err != nil {
		l.Error("Could not connect to database")
		return nil, err
	}
	return db, nil
}

func checkIfDatabaseExists() error {
	l := logger.NewLogger()
	conf := config.GetConfig()
	if _, err := os.Stat(conf.DatabasePath); os.IsNotExist(err) {
		l.Warn("Database does not exist, creating it")
		_, err := os.Create(conf.DatabasePath)
		if err != nil {
			l.Error("Could not create database")
			return err
		}
	}

	return nil
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

	return &snapshot, nil
}

func GetOldestSnapshots() ([]Snapshot, error) {
	l := logger.NewLogger()
	db, err := getDb()
	if err != nil {
		return nil, err
	}

	var snapshots []Snapshot
	result := db.Order("created_at asc").Find(&snapshots)
	if result.Error != nil {
		l.Error("Could not get oldest snapshots")
		return nil, result.Error
	}
	return snapshots, nil
}

func GetNewestSnapshot() (*Snapshot, error) {
	l := logger.NewLogger()
	db, err := getDb()
	if err != nil {
		return nil, err
	}

	var count int64
	result := db.Model(&Snapshot{}).Count(&count)
	if result.Error != nil {
		l.Error("Could not get count of snapshots")
		return nil, result.Error
	}
	if count == 0 {
		return nil, nil
	}

	var snapshot Snapshot
	result = db.Order("created_at desc").First(&snapshot)
	if result.Error != nil {
		l.Error("Could not get newest snapshot")
		return nil, result.Error
	}
	return &snapshot, nil
}

func DeleteSnapshot(snapshot *Snapshot) error {
	l := logger.NewLogger()
	db, err := getDb()
	if err != nil {
		return err
	}
	result := db.Delete(snapshot)
	if result.Error != nil {
		l.Error("Could not delete snapshot")
		return result.Error
	}
	return nil
}
