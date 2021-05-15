package params

import "time"

type GetTokenAuthReq struct {
	Key    string `json:"key" form:"key" validate:"required" label:"颁布标识 Key"`
	Secret string `json:"secret" form:"secret" validate:"required" label:"颁布标识 Secret"`
	UserID string `json:"user_id" form:"user_id" validate:"required" label:"用户标识 ID"`
	TTL    int    `json:"ttl" form:"ttl"`
}

type GetTokenAuthResp struct {
	Key       string    `json:"key"`
	Secret    string    `json:"secret"`
	UserID    string    `json:"user_id"`
	TTL       int       `json:"ttl"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type VerifyTokenAuthReq struct {
	Key    string `json:"key" form:"key" validate:"required" label:"颁布标识 Key"`
	Secret string `json:"secret" form:"secret" validate:"required" label:"颁布标识 Secret"`
	Token  string `json:"token" form:"token" validate:"required" label:"Token"`
}

type VerifyTokenAuthResp struct {
	Key       string    `json:"key"`
	Secret    string    `json:"secret"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
