package payroll

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/oktopriima/deals/helper"
	"github.com/oktopriima/deals/models"
	"github.com/oktopriima/deals/usecase/payroll/dto"
	"log"
)

func (p *payrollUsecase) RunUsecase(ctx context.Context, req dto.RunPayrollRequest) (dto.RunPayrollResponse, error) {
	payrollPeriod, err := p.payrollPeriodRepo.Find(ctx, req.PayrollPeriodId)
	if err != nil {
		return nil, err
	}

	if payrollPeriod.Processed {
		return nil, errors.New("payroll period is already processed")
	}

	users, err := p.userRepo.ListEmployees(ctx)
	if err != nil {
		return nil, err
	}

	totalWorkingDays := helper.WorkingDaysCount(payrollPeriod.StartDate, payrollPeriod.EndDate)

	for _, user := range users {
		attendances, err := p.attendanceRepo.ListAttendanceByUserID(ctx, user.ID, payrollPeriod.StartDate, payrollPeriod.EndDate)
		if err != nil {
			log.Printf("error found while fetching attendance for user: %d\n", user.ID)
			continue
		}

		if len(attendances) == 0 {
			continue
		}

		salary := user.Salary

		dailySalary := salary / float64(totalWorkingDays)
		deduction := dailySalary * (float64(totalWorkingDays) - float64(len(attendances)))

		// overtime section
		hourlySalary := dailySalary / float64(8)

		overtimes, err := p.overtimeRepo.ListOvertimeByUserId(user.ID, payrollPeriod.StartDate, payrollPeriod.EndDate, ctx)
		if err != nil {
			log.Printf("error found while fetching overtime for user: %d\n", user.ID)
		}

		var totalOvertimeHour int64
		for _, overtime := range overtimes {
			totalOvertimeHour += overtime.Duration
		}

		overtimeSalary := (hourlySalary * 2) * float64(totalOvertimeHour)

		// reimbursement
		reimbursements, err := p.reimbursementRepo.ListByUserID(user.ID, payrollPeriod.StartDate, payrollPeriod.EndDate, ctx)
		if err != nil {
			log.Printf("error found while fetching reimburse for user: %d\n", user.ID)
		}

		var totalReimbursement float64
		for _, reimbursement := range reimbursements {
			totalReimbursement += reimbursement.Amount
		}
		totalPayment := (salary - deduction) + totalReimbursement + overtimeSalary

		attendancesByte, _ := json.Marshal(attendances)
		overtimeByte, _ := json.Marshal(overtimes)
		reimbursementByte, _ := json.Marshal(reimbursements)
		// store payslip
		if err := p.payslipRepo.Store(ctx, &models.Payslip{
			UserID:             user.ID,
			PayrollPeriodId:    payrollPeriod.ID,
			BasePayment:        salary,
			DeductionAmount:    deduction,
			OvertimePayment:    overtimeSalary,
			Reimbursements:     totalReimbursement,
			TotalPayment:       totalPayment,
			ListAttendances:    string(attendancesByte),
			ListOvertimes:      string(overtimeByte),
			ListReimbursements: string(reimbursementByte),
		}); err != nil {
			log.Printf("error found while storing payslip for user: %d\n", user.ID)
			continue
		}
	}

	// update the payroll period
	payrollPeriod.Processed = true
	if err := p.payrollPeriodRepo.Update(ctx, payrollPeriod); err != nil {
		return nil, err
	}

	return dto.NewRunPayrollResponse(payrollPeriod), nil
}
