package models

import (
	"Link_shortener_website_S/database"
	"time"
)

type Log struct {
	ID        int    `json:"id"`
	UserIp    string `json:"user_ip"`
	UrlId     int    `json:"url_id"`
	VisitTime int64  `json:"visit_time"`
}

func CreateLog(userIp string, urlId int) error {
	newLog := Log{
		UserIp:    userIp,
		UrlId:     urlId,
		VisitTime: time.Now().Unix(),
	}
	err := database.DB.Create(&newLog).Error
	if err != nil {
		return err
	}
	return nil
}
func GetALlLogs() ([]Log, error) {
	var logs []Log
	err := database.DB.Find(&logs).Error
	if err != nil {
		return logs, err
	}
	return logs, nil
}
