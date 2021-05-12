package model

const (
	DB_DEFAULT = "default"
)

var (
	model = new(Model)
)

type Model struct {
	JwtSecretModel *JwtSecretModel
}

func InitModel()  {
	model.JwtSecretModel = NewJwtSecretModel()
}

func Get() *Model {
	return model
}

