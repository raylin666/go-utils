package model

var (
	model = new(Model)
)

type Model struct {
	JwtSecretModel *JwtSecretModel
	JwtUsersModel  *JwtUsersModel
}

func InitModel()  {
	model = &Model{
		JwtSecretModel: NewJwtSecretModel(),
		JwtUsersModel:  NewJwtUsersModel(),
	}
}

func Get() *Model {
	return model
}
