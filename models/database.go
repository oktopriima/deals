package models

import "gorm.io/gorm"

type DB struct {
	*gorm.DB
}

func (db *DB) AutoMigrate() error {
	return db.DB.AutoMigrate(
		&User{},
		&Attendance{},
		&Overtime{},
		&Reimbursement{},
		&PayrollPeriod{},
		&Payslip{},
		&AuditLog{},
	)
}
