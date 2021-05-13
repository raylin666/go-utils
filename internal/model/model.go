package model

var (
	model = new(Model)
)

type Model struct {
	JwtSecretModel *JwtSecretModel
}

func InitModel()  {
	model = &Model{
		JwtSecretModel: NewJwtSecretModel(),
	}
}

func Get() *Model {
	return model
}
