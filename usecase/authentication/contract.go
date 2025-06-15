package authentication

import (
	"context"

	jwthandle "github.com/oktopriima/deals/lib/jwtHandle"
	"github.com/oktopriima/deals/repository"
	"github.com/oktopriima/deals/usecase/authentication/dto"
)

type authenticationUsecase struct {
	jwtHandle      jwthandle.AccessToken
	userRepository repository.UserRepository
}

type AuthenticationUsecase interface {
	LoginUsecase(req dto.AuthenticationRequest, ctx context.Context) (dto.AuthenticationResponse, error)
}

func NewAuthenticationUsecase(
	jwtHandle jwthandle.AccessToken,
	userRepository repository.UserRepository,
) AuthenticationUsecase {
	return &authenticationUsecase{
		jwtHandle:      jwtHandle,
		userRepository: userRepository,
	}
}
