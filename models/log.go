package models

import (
	"Link_shortener_website_S/database"
	"gorm.io/gorm"
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
	// 更新link表中的clicks和last_click_time
	err = database.DB.Model(&Link{}).Where("id = ?", urlId).Update("clicks", gorm.Expr("clicks + ?", 1)).Update("last_click_t_ime", time.Now().Unix()).Error
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
