package model

import (
	"fmt"
	"github.com/raylin666/go-gin-api/pkg/database"
)

type (
	JwtSecret struct {
		database.Model
		Secret string `json:"secret" gorm:"uniqueIndex"`
	}
)

func (model JwtSecret) ExistSecret(secret string) bool {
	r := GetDefaultDB().Model(model).Where("secret = ?", secret).Select("id").Scan(model)

	fmt.Println(r)

	return true
}
