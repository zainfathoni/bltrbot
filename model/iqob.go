package model

import (
	"90/db"
	"time"

	"github.com/jinzhu/gorm"
)

type Iqob struct {
	gorm.Model
	UserId   uint
	State    string
	IqobDate time.Time
	PaidAt   time.Time
}

func (iqob *Iqob) paid() {
	iqob.State = "paid"
	db.MysqlDB().Save(iqob)
}
