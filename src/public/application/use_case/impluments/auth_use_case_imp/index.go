package auth_use_case_imp

type LoginResponse struct {
	UserId             int    `json:"user_id"`
	Token              string `json:"access_token"`
	AccessTokenExpires int    `json:"access_token_expires"`
	RefreshToken       string `json:"refresh_token"`
	AvatarURL          string `json:"avatar_url"`
}

type ILoginUseCase interface {
	Execute(email, password string) (*LoginResponse, error)
}

type Response struct {
	UserId             int    `json:"user_id"`
	Token              string `json:"access_token"`
	AccessTokenExpires int    `json:"access_token_expires"`
	RefreshToken       string `json:"refreshToken"`
	ImageURL           string `json:"image_url"`
}

type IRefreshTokenUseCase interface {
	Execute(refreshToken string) (*Response, error)
}

type ISignUpUseCase interface {
	Execute(email, password string) error
}
