package models

import (
	"Link_shortener_website_S/database"
	"errors"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Link struct {
	ID            int    `json:"id"`
	OriginUrl     string `json:"origin_url"`
	ShortUrl      string `json:"short_url" gorm:"unique"`
	ExpireAt      int64  `json:"expire_at"`
	CreatAt       int64  `json:"creat_at"`
	Clicks        int    `json:"clicks"`
	LastClickTIme int64  `json:"last_click_time"`
}

// var DB = database.DB
func IfLongUrlExist(longUrl string) bool {
	var link Link
	result := database.DB.Where("origin_url = ?", longUrl).First(&link)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}
func CreateLink(originUrl string) (string, error) {
	shortUrl, err := generateShortUrl()
	if err != nil {
		return "", err
	}
	newLink := Link{
		OriginUrl:     originUrl,
		ShortUrl:      shortUrl,
		ExpireAt:      time.Now().Unix() + 24*60*60,
		CreatAt:       time.Now().Unix(),
		Clicks:        0,
		LastClickTIme: 0,
	}
	err = database.DB.Create(&newLink).Error
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}
func generateShortUrl() (string, error) {
	for {
		shortUrl := GetRandomString(6) // 确保这个函数生成的字符串长度和随机性是足够的
		var link Link
		result := database.DB.Where("short_url = ?", shortUrl).First(&link)

		// 处理可能的错误
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// 如果没有找到记录，返回生成的短URL
				return shortUrl, nil
			}
			// 如果有其他错误，返回错误
			return "", result.Error
		}

		// 如果找到了记录，说明短URL已经存在，循环会继续生成新的短URL
		if link.ID == 0 {
			// 如果ID为0，说明没有找到记录，可以使用这个短URL
			return shortUrl, nil
		}

		// 如果ID不为0，说明短URL已存在，循环将继续尝试
	}
	// 注意：理论上这个循环不会退出，除非找到了一个未使用的短URL或者发生了错误
}

func GetOriginUrl(shortUrl string) (Link, error) {
	var link Link
	err := database.DB.Where("short_url = ?", shortUrl).First(&link).Error
	if err != nil {
		return link, err
	}
	return link, nil
}
func GetShortUrl(longUrl string) (Link, error) {
	var link Link
	err := database.DB.Where("origin_url = ?", longUrl).First(&link).Error
	if err != nil {
		return link, err
	}
	return link, nil
}

func GetRandomString(i int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for j := 0; j < i; j++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)

}
func GetAllData() ([]Link, error) {
	var links []Link
	err := database.DB.Find(&links).Error
	if err != nil {
		return links, err
	}
	return links, nil
}
func GetTargetData(shortUrl string) (Link, error) {
	var link Link
	err := database.DB.Where("short_url = ?", shortUrl).First(&link).Error
	if err != nil {
		return link, err
	}
	return link, nil
}
