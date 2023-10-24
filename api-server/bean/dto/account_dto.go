package dto

import "time"

// AccountDTO 定义一个用户实体
type AccountDTO struct {
	ID         string    `gorm:"type:char(36);primary_key"`
	Name       string    `gorm:"type:char(6)"`
	JobNumber  string    `gorm:"type:char(15)"`
	Group      string    `gorm:"type:char(8)"`
	Phone      int       `gorm:"type:char(11)"`
	Email      string    `gorm:"type:char(20)"`
	Sex        string    `gorm:"type:char(2)"`
	ArchGroup  string    `gorm:"type:char(10)"`
	CreateTime time.Time `gorm:"type:datetime"`
	UpdateTime time.Time `gorm:"type:datetime"`
}
