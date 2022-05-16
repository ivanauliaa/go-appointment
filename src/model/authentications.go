package model

type Authentication struct {
	Token string `gorm:"not null" json:"token"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenPayload struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	RefreshToken string `json:"refreshToken"`
}

type RequestHeader struct {
	Authorization string
}

type AccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}
