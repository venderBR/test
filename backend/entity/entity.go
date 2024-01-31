package entity

import(
	"gorm.io/gorm"
)

type User struct{
	gorm.Model

	Firstname    string   `valid:"required~Firstname is required"`
	Lastname     string   `valid:"required~Lastname is required"`
	StudentID    string   `valid:"required~StudentID is required, matches(^[BMDbmd]\\d{7}$)~StudentID is invalid"`
	Phone        string   `valid:"required~Phone is required , stringlength(10|10)~Phone is invalid"`
	Email        string   `valid:"required~Email is required, email~Email is invalid"`
	LinkedIn  string `valid:"required~LinkedIn is required, url~Url LinkedIn is invalid"`

	// GenderID ทำหน้าที่เป็น FK
	GenderID uint   
	Gender   Gender `gorm:"foreignKey:GenderID"`
}

type Gender struct {
	gorm.Model
	Name string
}