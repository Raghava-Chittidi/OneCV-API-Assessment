package util

import (
	"Raghava/OneCV-Assignment/internal/database"
	"Raghava/OneCV-Assignment/internal/models"
)

// Migrates gorm models
func Migrate() error {
	err := database.DB.AutoMigrate(&models.Student{}, &models.Teacher{})
	return err
}