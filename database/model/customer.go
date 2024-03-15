package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	Email      string    `gorm:"type:varchar(100)" json:"email"`
	Title      string    `gorm:"type:varchar(100)" json:"title"`
	Content    string    `gorm:"type:text" json:"content"`
	MailingID  int       `gorm:"type:integer" json:"mailing_id"`
	InsertTime time.Time `gorm:"type:timestamp" json:"insert_time"`
}
