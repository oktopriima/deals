package dto

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/oktopriima/deals/models"
	"time"
)

type runPayrollResponse struct {
	PayrollPeriodId int64  `json:"payroll_period_id"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	IsProcessed     bool   `json:"is_processed"`
}

type RunPayrollResponse interface {
	GetObject() *runPayrollResponse
}

func (r *runPayrollResponse) GetObject() *runPayrollResponse {
	return r
}

func NewRunPayrollResponse(period *models.PayrollPeriod) RunPayrollResponse {
	return &runPayrollResponse{
		PayrollPeriodId: period.ID,
		StartDate:       period.StartDate.Format("2006-01-02"),
		EndDate:         period.EndDate.Format("2006-01-02"),
		IsProcessed:     period.Processed,
	}
}

type detailPayrollResponse struct {
	Id                int64               `json:"id"`
	UserId            int64               `json:"user_id"`
	Username          string              `json:"username"`
	PeriodStart       string              `json:"period_start"`
	PeriodEnd         string              `json:"period_end"`
	BaseSalary        float64             `json:"base_salary"`
	TotalWorkingDays  float64             `json:"total_working_days"`
	TotalAttendance   float64             `json:"total_attendance"`
	TotalHourOvertime float64             `json:"total_hour_overtime"`
	SalaryDeduction   float64             `json:"salary_deduction"`
	OvertimeSalary    float64             `json:"overtime_salary"`
	Reimbursement     float64             `json:"reimbursement"`
	AttendanceList    []attendanceList    `json:"attendance_list"`
	OvertimeList      []overtimeList      `json:"overtime_list"`
	ReimbursementList []reimbursementList `json:"reimbursement_list"`
}

type attendanceList struct {
	Day        string    `json:"day"`
	Date       string    `json:"date"`
	SubmitTime string    `json:"submit_time"`
	Timestamp  time.Time `json:"timestamp"`
}

type overtimeList struct {
	StartHour string    `json:"start_hour"`
	EndHour   string    `json:"end_hour"`
	Duration  int64     `json:"duration"`
	Timestamp time.Time `json:"timestamp"`
}

type reimbursementList struct {
	Timestamp   time.Time `json:"timestamp"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
}

type DetailPayrollResponse interface {
	GetObject() *detailPayrollResponse
}

func (r *detailPayrollResponse) GetObject() *detailPayrollResponse {
	return r
}

func NewDetailPayrollResponse(payslip *models.Payslip, period *models.PayrollPeriod) DetailPayrollResponse {

	var user models.User
	if err := json.Unmarshal([]byte(payslip.Users), &user); err != nil {
		log.Printf("failed read user object. error: %v", err)
		return nil
	}

	var attendances []models.Attendance
	if err := json.Unmarshal([]byte(payslip.ListAttendances), &attendances); err != nil {
		log.Printf("failed read attendances object. error: %v", err)
		return nil
	}

	var overtimes []models.Overtime
	if err := json.Unmarshal([]byte(payslip.ListOvertimes), &overtimes); err != nil {
		log.Printf("failed read overtimes object. error: %v", err)
		return nil
	}

	var reimbursements []models.Reimbursement
	if err := json.Unmarshal([]byte(payslip.ListReimbursements), &reimbursements); err != nil {
		log.Printf("failed read reimbursements object. error: %v", err)
		return nil
	}

	al := make([]attendanceList, 0)
	for _, attendance := range attendances {
		al = append(al, attendanceList{
			Day:        attendance.Timestamp.Weekday().String(),
			Date:       attendance.Timestamp.Format("2006-01-02"),
			SubmitTime: attendance.Timestamp.Format("15:04:05"),
			Timestamp:  attendance.Timestamp,
		})
	}

	ol := make([]overtimeList, 0)
	for _, overtime := range overtimes {
		ol = append(ol, overtimeList{
			StartHour: overtime.Timestamp.Format("15:04:05"),
			EndHour:   overtime.Timestamp.Add(time.Duration(overtime.Duration) * time.Hour).Format("15:04:05"),
			Timestamp: overtime.Timestamp,
			Duration:  overtime.Duration,
		})
	}

	rl := make([]reimbursementList, 0)
	for _, reimbursement := range reimbursements {
		rl = append(rl, reimbursementList{
			Timestamp:   reimbursement.Date,
			Amount:      reimbursement.Amount,
			Description: reimbursement.Description,
		})
	}

	return &detailPayrollResponse{
		Id:                payslip.ID,
		UserId:            payslip.UserID,
		Username:          user.Username,
		PeriodStart:       period.StartDate.Format("2006-01-02"),
		PeriodEnd:         period.EndDate.Format("2006-01-02"),
		BaseSalary:        payslip.BasePayment,
		TotalWorkingDays:  payslip.TotalWorkingDays,
		TotalAttendance:   payslip.TotalAttendance,
		TotalHourOvertime: payslip.TotalHourOvertime,
		SalaryDeduction:   payslip.DeductionAmount,
		OvertimeSalary:    payslip.OvertimePayment,
		Reimbursement:     payslip.Reimbursements,
		AttendanceList:    al,
		OvertimeList:      ol,
		ReimbursementList: rl,
	}
}
