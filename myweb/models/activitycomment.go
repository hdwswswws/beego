package models

import (
	"time"
	"fmt"

	_ "github.com/bmizerany/pq"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ActivityComment struct {
	Id           uuid.UUID `json:"id,omitempty"`
	ActivityName string    `json:"activity_name,omitempty"`
	Username     string    `json:"username,omitempty"`
	Comment      string    `json:"comment,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	Report       int       `json:"report"`
	Portrait     string    `json:"portrait" gorm:"-"`
}
func CreateActivityComment(accessToken string, input ActivityComment) (err error) {
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
	var activity Activity
	result := db.Table("activity").Where("activity_name = ?", input.ActivityName).Take(&activity )
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	id := uuid.New()
	input.Id =id
	input.Username = user.Username
	result = db.Table("activity_comment").Create(&input)
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	return nil
}
func SuperAdminFindActivityComment(accessToken string,input ActivityComment,page int, pageSize int) (comments []ActivityComment, total int, err error) {
	_, err = SuperAdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Table("activity_comment").Select([]string{"id", "activity_name", "username", "comment", "created_at","report"}).Where("activity_name LIKE ? AND comment LIKE ?", "%"+input.ActivityName+"%","%"+input.Comment+"%").Limit(pageSize).Offset((page - 1) * pageSize).Order("report desc").Find(&comments)
	db.Table("activity_comment").Where("activity_name LIKE ? AND comment LIKE ?", "%"+input.ActivityName+"%","%"+input.Comment+"%").Count(&total)
	return comments, total, nil
}
func DelActivityComment( id string) (err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	var activityComment ActivityComment
	result := db.Table("activity_comment").Where("id = ? ", id).Delete(&activityComment)
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}
func UserDelActivityComment( username string,id string) (err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	var activityComment ActivityComment
	result := db.Table("activity_comment").Where("id = ? AND username = ?", id,username).Delete(&activityComment)
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}
func UserCountActivityComment(activityname string) ( total int, err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Table("activity_comment").Where("activity_name = ?", activityname).Count(&total)
	return total, nil
}
func UserFindActivityComment(activityname string) ( comments []ActivityComment, err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Raw("SELECT activity_comment.id, activity_comment.activity_name, activity_comment.username,activity_comment.comment, activity_comment.created_at, activity_comment.report,  resume.portrait FROM activity_comment left join resume on activity_comment.username = resume.username where activity_comment.activity_name =  ?", activityname).Order("activity_comment.created_at desc").Find(&comments)
	return comments, nil
}
func UserEditActivityComment(id string) (err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("举报失败")
		return
	}
	var input ActivityComment
	result := db.Table("activity_comment").Where("id = ? ", id).Take(&input)
	if result.Error != nil {
		err = fmt.Errorf("举报失败")
		return
	}
	report := input.Report+1
	result = db.Table("activity_comment").Where("id = ? ", id).Update(&ActivityComment{Report:report})
	if result.Error != nil {
		err = fmt.Errorf("举报失败")
		return
	}
	return nil
}
