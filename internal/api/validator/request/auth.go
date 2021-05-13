package request

type GetTokenAuthValidate struct {
	Secret string `validate:"required" label:"颁布标识 Secret"`
}