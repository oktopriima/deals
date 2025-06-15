package models

type User struct {
	ID       int64   `gorm:"primaryKey"`
	Username string  `gorm:"uniqueIndex;not null"`
	Password string  `gorm:"not null"`
	IsAdmin  bool    `gorm:"not null;default:false"`
	Salary   float64 `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}
