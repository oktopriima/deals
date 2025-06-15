package attendance

import (
	"context"

	"github.com/oktopriima/deals/repository"
	"github.com/oktopriima/deals/usecase/attendance/dto"
)

type attendanceUsecase struct {
	attendanceRepository repository.AttendanceRepository
}

type AttendanceUsecase interface {
	CreateUsecase(req dto.AttendanceRequest, ctx context.Context) (dto.AttendanceResponse, error)
}

func NewAttendanceUsecase(
	attendanceRepository repository.AttendanceRepository,
) AttendanceUsecase {
	return &attendanceUsecase{
		attendanceRepository: attendanceRepository,
	}
}
