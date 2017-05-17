package main

import (
	"90/db"
	"90/model"
	"fmt"
	"strconv"
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
	value, err := strconv.Atoi(Args[2])
	if err == nil && value > 0 && &Args[1] != nil {
		var report_type int
		if Args[1] == ("iqob") {
			report_type = 1
		} else if Args[1] == "report" {
			report_type = 0
			db.MysqlDB().Model(&CurrentUser).Update("remaining_today", (CurrentUser.RemainingToday - value))
		} else {
			Bot.ReplyToUser("tipe Laporan tidak ditemukan")
			return
		}

		report := model.Report{UserId: CurrentUser.ID, Value: value, Type: report_type}
		db.MysqlDB().Create(&report)
		Bot.ReplyToUser("Laporan berhasil dimasukkan")
	} else {
		Bot.ReplyToUser("Nilai yang anda masukkan salah")
	}
}
