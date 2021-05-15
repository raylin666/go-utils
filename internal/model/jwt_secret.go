package model

import (
	"go-server/pkg/database"
	"gorm.io/gorm"
	"time"
)

type (
	JwtSecretModel struct {
		Connection *gorm.DB
		Table      string
	}

	JwtSecret struct {
		database.Model
		App       string    `json:"app"`
		Key       string    `json:"key" gorm:"uniqueIndex:key_secret"`
		Secret    string    `json:"secret" gorm:"uniqueIndex:key_secret"`
		ExpiredAt time.Time `json:"expired_at"`
	}
)

func NewJwtSecretModel() *JwtSecretModel {
	var connection = database.GetDB(database.DefaultDatabaseConnection)
	return &JwtSecretModel{
		Connection: connection,
		Table:      connection.Config.NamingStrategy.TableName("jwt_secret"),
	}
}

// 判断 Key Secret 是否存在
func (model *JwtSecretModel) ExistKeySecret(key string, secret string) bool {
	var jwt_secret JwtSecret
	model.Connection.Table(model.Table).Where("`key` = ? AND secret = ?", key, secret).Select("id").First(&jwt_secret)

	if jwt_secret.ID > 0 {
		return true
	}

	return false
}

// 获取 Key Secret 数据
func (model *JwtSecretModel) GetKeySecretFirst(key string, secret string) *JwtSecret {
	var jwt_secret *JwtSecret
	model.Connection.Table(model.Table).Where("`key` = ? AND secret = ?", key, secret).First(&jwt_secret)
	return jwt_secret
}
