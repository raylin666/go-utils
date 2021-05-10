package model

import (
	"github.com/raylin666/go-gin-api/pkg/database"
	"gorm.io/gorm"
)

func GetDefaultDB() *gorm.DB {
	return database.GetDB("default")
}

