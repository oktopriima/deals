package attendance

import (
	"context"

	"github.com/oktopriima/deals/helper"
	"github.com/oktopriima/deals/usecase/attendance/dto"
)

func (a *attendanceUsecase) CreateUsecase(req dto.AttendanceRequest, ctx context.Context) (dto.AttendanceResponse, error) {
	if helper.IsWeekend(req.Timestamp) {
		return nil, helper.ErrWeekendNotAllowed
	}

	if req.Timestamp.IsZero() {
		return nil, helper.ErrInvalidTimestamp
	}

	// return success when already exist
	if a.attendanceRepository.CheckAlreadyExists(req.UserId, req.Timestamp.Format("2006-01-02"), ctx) {
		return dto.NewAttendanceResponse(req.ToModel()), nil
	}

	if err := a.attendanceRepository.Store(req.ToModel(), ctx); err != nil {
		return nil, err
	}

	return dto.NewAttendanceResponse(req.ToModel()), nil
}
