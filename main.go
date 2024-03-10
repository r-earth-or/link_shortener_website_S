package main

import (
	"Link_shortener_website_S/database"
	"Link_shortener_website_S/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (err error) {
	DB, err = database.Connect()
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&models.Link{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&models.Log{})
	if err != nil {
		return err
	}
	return nil
}
func main() {

	err := InitDB()
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err := database.CloseDB()
		if err != nil {
			fmt.Println(err)
		}
	}()
	r := gin.Default()
	r.POST("/api/links", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		longUrl := c.Query("longUrl")
		if models.IfLongUrlExist(longUrl) {
			shortUrl, err := models.GetShortUrl(longUrl)
			if err != nil {
				c.JSON(500, gin.H{
					"message": err.Error(),
				})
			}
			c.JSON(
				200, gin.H{
					"shortUrl": shortUrl.ShortUrl,
				})
		}
		shortUrl, err := models.CreateLink(longUrl)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(
			200, gin.H{
				"shortUrl": shortUrl,
			})
	})
	r.GET("/api/logs", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		queryTarget := c.Query("target")
		if queryTarget == "all" {
			logs, err := models.GetAllData()
			if err != nil {
				c.JSON(500, gin.H{
					"message": err.Error(),
				})
			}
			c.JSON(200, logs)
		} else if models.IfLongUrlExist(queryTarget) {
			shortUrl, err := models.GetShortUrl(queryTarget)
			if err != nil {
				c.JSON(500, gin.H{
					"message": err.Error(),
				})
			}
			logs, err := models.GetTargetData(shortUrl.ShortUrl)
			if err != nil {
				c.JSON(500, gin.H{
					"message": err.Error(),
				})
			}
			c.JSON(200, logs)
		} else {
			logs, err := models.GetTargetData(queryTarget)
			if err != nil {
				c.JSON(500, gin.H{
					"message": err.Error(),
				})
			}
			c.JSON(200, logs)
		}
	})
	r.GET("/:shortUrl", func(c *gin.Context) {
		shortUrl := c.Param("shortUrl")
		link, err := models.GetOriginUrl(shortUrl)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			err = models.CreateLog(c.ClientIP(), link.ID)
			if err != nil {
				c.JSON(500, gin.H{
					"message": err.Error(),
				})
			}
			c.Redirect(301, link.OriginUrl)
		}
	})
	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}

}
