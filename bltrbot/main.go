package main

import (
	"fmt"

	"bufio"

	"os"

	"strings"

	"github.com/ripasfilqadar/bltrbot/bltrbot/db"
	"github.com/ripasfilqadar/bltrbot/bltrbot/model"
)

var Emoji = map[string]string{
	"not_confirm": "👹",
	"smile":       "😇",
	"iqob":        "💀",
	"leave":       "✈",
}

func main() {
	fmt.Println("start")
	initEnv()
	InitRoute()
	InitDB()
	InitTelegram()
	go RunSchedule()
	//	reminderUser()
	//	updateRemaining()
	StartTelegram()
}

func InitDB() {
	db.MysqlDB().AutoMigrate(&model.User{}, &model.Report{}, &model.Iqob{}, &model.Group{})
	db.MysqlDB().Model(&model.User{}).AddIndex("group_id", "user_name", "state")
	db.MysqlDB().Model(&model.Report{}).AddIndex("user_id", "type", "actor_id")
	db.MysqlDB().Model(&model.Iqob{}).AddIndex("user_id", "state")
	db.MysqlDB().Model(&model.Group{}).AddIndex("group_id")
}

func initEnv() {
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		envTemp := strings.Split(scanner.Text(), "=")
		if len(envTemp) == 2 {
			if envTemp[0] == "ADMIN_USERNAME" {
				admins := strings.Fields(envTemp[1])
				for _, admin := range admins {
					os.Setenv(envTemp[0], admin)
				}
			} else {
				os.Setenv(envTemp[0], envTemp[1])
			}
		}
	}
}
