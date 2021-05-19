package models

import (
	"fmt"
	"time"

	_ "github.com/bmizerany/pq"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Resume struct {
	Id               uuid.UUID `json:"id"`
	Username         string    `json:"username"`
	CreatedAt        time.Time `json:"created_at"`
	Name             string    `json:"name"`
	Sex              string    `json:"sex"`
	StudentNumber    string    `json:"student_number"`
	Major            string    `json:"major"`
	PoliticalOutlook string    `json:"political_outlook"`
	PhoneNumber      string    `json:"phone_number"`
	Email            string    `json:"email"`
	Introduction     string    `json:"introduction"`
	Birthday         string    `json:"birthday"`
	Portrait         string    `json:"portrait"`
}

func CreateResume(accessToken string, input Resume) (err error) {
	user, err := UserCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	id := uuid.New()
	input.Id = id
	input.Username = user.Username
	resume := input
	var resume1 Resume
	result := db.Table("resume").Where("username = ? ", user.Username).Take(&resume1)
	if result.Error != nil {
		result = db.Table("resume").Create(&resume)
		return
	}
	result = db.Table("resume").Where("username = ? ", user.Username).Update(&resume)
	return nil
}
func EditResume(accessToken string, input Resume) (err error) {
	user, err := UserCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("修改失败")
		return
	}

	result := db.Table("resume").Where("username = ?", user.Username).Update(input)
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return err
	}
	return nil
}
func FindResume(username string) (resume Resume, err error) {
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	result := db.Table("resume").Where("username = ? ", username).Find(&resume)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("获取失败")
		return
	}

	return
}
