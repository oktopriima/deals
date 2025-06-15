package authentication

import (
	"context"
	"strconv"

	"github.com/oktopriima/deals/helper"
	jwthandle "github.com/oktopriima/deals/lib/jwtHandle"
	"github.com/oktopriima/deals/usecase/authentication/dto"
)

// LoginUsecase handles the authentication process for a user login.
// It takes an AuthenticationRequest and a context, then attempts to find the user by username.
// If the user is found and the password is correct, it generates a JWT token and returns an AuthenticationResponse.
// Returns an error if the user is not found, the password is incorrect, or token generation fails.
func (a *authenticationUsecase) LoginUsecase(req dto.AuthenticationRequest, ctx context.Context) (dto.AuthenticationResponse, error) {
	user, err := a.userRepository.FindByUsername(req.Username, ctx)
	if err != nil {
		return nil, err
	}

	if !helper.CheckPassword(req.Password, user.Password) {
		return nil, helper.ErrInvalidCredentials
	}

	token, err := a.jwtHandle.GenerateToken(jwthandle.Params{
		ID:  strconv.Itoa(int(user.ID)),
		Obj: user,
	})
	if err != nil {
		return nil, err
	}

	return dto.NewAuthenticationResponse(
		token.GetStringToken(),
		token.GetStringRefreshToken(),
		token.GetTimeExpiredAt().Unix(),
	), nil
}
