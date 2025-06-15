package authentication

import (
	"context"
	"strconv"

	"github.com/oktopriima/deals/helper"
	jwthandle "github.com/oktopriima/deals/lib/jwtHandle"
	"github.com/oktopriima/deals/usecase/authentication/dto"
)

// LoginUsecase implements AuthenticationUsecase.
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
