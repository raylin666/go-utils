package model

import (
	"github.com/raylin666/go-gin-api/pkg/database"
	"gorm.io/gorm"
)

type (
	JwtSecretModel struct {
		Connection *gorm.DB
		Table      string
	}

	JwtSecret struct {
		database.Model
		Secret string `json:"secret" gorm:"uniqueIndex"`
	}
)

func NewJwtSecretModel() *JwtSecretModel {
	var connection = database.GetDB(DB_DEFAULT)
	return &JwtSecretModel{
		Connection: connection,
		Table:      connection.Config.NamingStrategy.TableName("jwt_secret"),
	}
}

// 判断 Secret 是否存在
func (model *JwtSecretModel) ExistSecret(secret string) bool {
	var jwt_secret JwtSecret
	model.Connection.Table(model.Table).Where("secret = ?", secret).Select("id").First(&jwt_secret)

	if jwt_secret.ID > 0 {
		return true
	}

	return false
}
