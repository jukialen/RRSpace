package database

import (
	"log"

	"rrspace/backend/internal/config"
	"rrspace/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DBDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	// AutoMigrate w dev; w prod użyj migracji
	if cfg.AppEnv == "development" {
		if err := db.AutoMigrate(
			&models.Studio{},
			&models.Resource{},
			&models.Schedule{},
			&models.ScheduleException{},
			&models.Booking{},
			&models.Review{},
			&models.Photo{},
		); err != nil {
			log.Fatalf("migration failed: %v", err)
		}
	}

	return db, nil
}