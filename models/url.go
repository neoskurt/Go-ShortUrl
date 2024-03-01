package models

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type URL struct {
    gorm.Model
    LongURL      string     `json:"long_url" gorm:"unique"`
    ShortURL     string     `json:"short_url" gorm:"unique"`
    AccessCount  uint       `json:"access_count"`
    LastAccessed *time.Time `json:"last_accessed"`
    AccessPlace  string     `json:"access_place"`
    Alias        string `gorm:"unique"`
    UserID       uint
}


func (u *URL) GenerateShortURL() {
 rand.Seed(time.Now().UnixNano())
 var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
 b := make([]rune, 6)
 for i := range b {
  b[i] = letterRunes[rand.Intn(len(letterRunes))]
 }
 u.ShortURL = string(b)

}

func CreateURL(url *URL) {
 DB.Create(&url)
}

func GetURLByShortURL(shortURL string) (URL, error) {
    var url URL

    err := DB.Where("alias = ? OR short_url = ?", shortURL, shortURL).First(&url).Error
    if err != nil {
        return url, err
    }
    return url, nil
}

func UpdateURL(url *URL) error {
    result := DB.Save(url)
    return result.Error
}

func CloseDB() {
	DB.Close()
}