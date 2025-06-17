package payroll

import (
	"context"
	"github.com/oktopriima/deals/repository"
	"github.com/oktopriima/deals/usecase/payroll/dto"
)

type payrollUsecase struct {
	userRepo          repository.UserRepository
	attendanceRepo    repository.AttendanceRepository
	overtimeRepo      repository.OvertimeRepository
	reimbursementRepo repository.ReimbursementRepository
	payrollPeriodRepo repository.PayrollPeriodRepository
	payslipRepo       repository.PayslipRepository
}

type PayrollUsecase interface {
	RunUsecase(ctx context.Context, req dto.RunPayrollRequest) (dto.RunPayrollResponse, error)
	FindByUserUsecase(ctx context.Context, periodId int64) (dto.DetailPayrollResponse, error)
}

func NewPayrollUsecase(userRepo repository.UserRepository,
	attendanceRepo repository.AttendanceRepository,
	overtimeRepo repository.OvertimeRepository,
	reimbursementRepo repository.ReimbursementRepository,
	payrollPeriodRepo repository.PayrollPeriodRepository,
	payslipRepo repository.PayslipRepository,
) PayrollUsecase {
	return &payrollUsecase{
		userRepo:          userRepo,
		attendanceRepo:    attendanceRepo,
		overtimeRepo:      overtimeRepo,
		reimbursementRepo: reimbursementRepo,
		payrollPeriodRepo: payrollPeriodRepo,
		payslipRepo:       payslipRepo,
	}
}
