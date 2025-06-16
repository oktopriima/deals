package overtime

import (
	"context"
	"fmt"
	"github.com/oktopriima/deals/helper"
	"github.com/oktopriima/deals/usecase/overtime/dto"
)

const maximumOvertimeDuration = 3

func (o *overtimeUsecase) CreateUsecase(request dto.OvertimeRequest, ctx context.Context) (dto.OvertimeResponse, error) {
	if !helper.CheckOvertimeStart(request.Date) {
		return nil, fmt.Errorf("overtime should be purposed after working hour")
	}

	overtime, err := o.overtimeRepository.FindByUserDate(request.UserId, request.Date, ctx)

	// update the overtime duration
	if err == nil && overtime != nil {
		newDuration := request.Duration + overtime.Duration

		if newDuration > maximumOvertimeDuration {
			return nil, fmt.Errorf("your new duration overtime %d is reach the maximum", newDuration)
		}

		overtime.Duration = newDuration
		if err := o.overtimeRepository.Update(overtime, ctx); err != nil {
			return nil, err
		}

		return dto.NewOvertimeResponse(overtime), nil
	}

	// create the overtime
	if err := o.overtimeRepository.Store(request.ToModel(), ctx); err != nil {
		return nil, err
	}

	return dto.NewOvertimeResponse(request.ToModel()), nil
}
