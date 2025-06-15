package dto

type authenticationResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type AuthenticationResponse interface {
	GetObject() *authenticationResponse
}

func (r *authenticationResponse) GetObject() *authenticationResponse {
	return r
}

func NewAuthenticationResponse(accessToken, refreshToken string, expiresIn int64) AuthenticationResponse {
	return &authenticationResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}
}
