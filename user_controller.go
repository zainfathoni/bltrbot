package main

import (
	"fmt"
	"github.com/ripasfilqadar/bltrbot/db"
	"github.com/ripasfilqadar/bltrbot/model"
	"strconv"
	"time"
)

func (c *Controller) SetTarget() {
	fmt.Println("set target")
	target, err := strconv.Atoi(Args[1])
	if err == nil {
		CurrentUser.SetTarget(target)
		Bot.ReplyToUser("Target Berhasil di atur")
	} else {
		Bot.ReplyToUser("Target harus besar dari 0")
	}
}

func (c *Controller) TodayReport() {
	fmt.Println("today report")
	value, err := strconv.Atoi(Args[1])
	if err == nil && value > 0 {
		report_type := "tilawah"
		remaining_today := CurrentUser.RemainingToday - value
		if remaining_today < 0 {
			remaining_today = 0
		}
		db.MysqlDB().Model(&CurrentUser).Update("remaining_today", remaining_today)

		report := model.Report{UserId: CurrentUser.ID, Value: value, Type: report_type}
		db.MysqlDB().Create(&report)
		Bot.ReplyToUser("Laporan berhasil dimasukkan, sisa tilawah anda adalah " + strconv.Itoa(CurrentUser.RemainingToday) + " halaman")
	} else {
		Bot.ReplyToUser("Nilai yang anda masukkan salah")
	}
}

func (c *Controller) PaidIqob() {
	total, err := strconv.Atoi(Args[1])
	if err == nil && total > 0 {
		iqobs := []model.Iqob{}
		db.MysqlDB().Where("state = ?", "not_paid").Find(&iqobs).Limit(total)
		count := 0
		for _, iqob := range iqobs {
			iqob.PaidAt = time.Now()
			iqob.State = "paid"
			db.MysqlDB().Save(&iqob)
			count++
		}
		Bot.ReplyToUser(strconv.Itoa(count) + " Iqob telah dibayar")
	} else {
		panic(err)
		Bot.ReplyToUser("Total Iqob harus lebih besar dari 0")
	}
}
