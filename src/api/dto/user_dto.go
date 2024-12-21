package dto

type TokenDetail struct {
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	AccessTokenExpireTime  int    `json:"access_token_expire_time"`
	RefreshTokenExpireTime int    `json:"refresh_token_expire_time"`
}