package model

import (
	"go-server/internal/constant"
	"go-server/pkg/database"
	"go-server/pkg/logger"
	"gorm.io/gorm"
	"time"
)

type (
	JwtUsersModel struct {
		Connection *gorm.DB
		Table      string
	}

	JwtUsers struct {
		database.Model
		SecretId  int        `json:"secret_id" gorm:"uniqueIndex:user_secret"`
		UserID    string     `json:"user_id" gorm:"uniqueIndex:user_secret"`
		Token     string     `json:"token"`
		TTL       int        `json:"ttl"`
		ExpiredAt time.Time  `json:"expired_at"`
		RefreshAt *time.Time `json:"refresh_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
)

func NewJwtUsersModel() *JwtUsersModel {
	var connection = database.GetDB(database.DefaultDatabaseConnection)
	return &JwtUsersModel{
		Connection: connection,
		Table:      connection.Config.NamingStrategy.TableName("jwt_users"),
	}
}

// 获取用户数据
func (model *JwtUsersModel) GetSecretUser(user_id string, secret_id int) *JwtUsers {
	var jwt_users *JwtUsers
	model.Connection.Table(model.Table).Where("user_id = ? AND secret_id = ?", user_id, secret_id).First(&jwt_users)
	return jwt_users
}

// 获取 Token 数据
func (model *JwtUsersModel) GetTokenUser(token string, secret_id int) *JwtUsers {
	var jwt_users *JwtUsers
	model.Connection.Table(model.Table).Where("token = ? AND secret_id = ?", token, secret_id).First(&jwt_users)
	return jwt_users
}

// 创建用户数据
func (model *JwtUsersModel) Create(jwtUsers *JwtUsers) uint64 {
	result := model.Connection.Table(model.Table).Create(jwtUsers)
	if result.Error != nil {
		logger.NewWrite(constant.LogSql).WithFields(logger.H{
			"data": jwtUsers,
			"err":  result.Error,
		}.Fields()).Error("create data error")
		return 0
	}

	return jwtUsers.ID
}
