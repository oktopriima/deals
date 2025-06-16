package overtime

import (
	"context"
	"github.com/oktopriima/deals/repository"
	"github.com/oktopriima/deals/usecase/overtime/dto"
)

type overtimeUsecase struct {
	overtimeRepository repository.OvertimeRepository
}

type OvertimeUsecase interface {
	CreateUsecase(request dto.OvertimeRequest, ctx context.Context) (dto.OvertimeResponse, error)
}

func NewOvertimeUsecase(
	overtimeRepository repository.OvertimeRepository,
) OvertimeUsecase {
	return &overtimeUsecase{
		overtimeRepository: overtimeRepository,
	}
}
