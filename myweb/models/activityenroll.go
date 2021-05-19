package models

import (
	"fmt"
	"time"

	_ "github.com/bmizerany/pq"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"github.com/tealeg/xlsx"
)

type ActivityEnroll struct {
	Id           uuid.UUID `json:"id,omitempty"`
	ActivityName string    `json:"activity_name,omitempty"`
	Username     string    `json:"username,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
type ActivityEnrollDetails struct {
	Id            uuid.UUID `json:"id,omitempty"`
	ActivityName  string    `json:"activity_name,omitempty"`
	Username      string    `json:"username,omitempty"`
	Name          string    `json:"name,omitempty"`
	Sex           string    `json:"sex,omitempty"`
	StudentNumber string    `json:"student_number,omitempty"`
	Major         string    `json:"major,omitempty"`
	PhoneNumber   string    `json:"phone_number,omitempty"`
	Email         string    `json:"email,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

func ActivityEnrollExportExcel(accessToken string, activityName string) (filename string, err error) {
	_, err = AdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	var ActivityEnrollDetails []ActivityEnrollDetails
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Raw("SELECT activity_enroll.id,activity_enroll.activity_name, activity_enroll.username,activity_enroll.created_at,resume.name,resume.sex,resume.student_number,resume.major,resume.phone_number,resume.email FROM  activity_enroll left join resume on activity_enroll.username = resume.username where activity_enroll.activity_name = ?", activityName).Order("activity_enroll.created_at desc").Find(&ActivityEnrollDetails)
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("sheet1")
	title := sheet.AddRow()
	titleRow := title.AddCell()
	titleRow.HMerge = 7 
	titleRow.Value = "活动（" + activityName + "）的报名表"
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "用户名"
	cell = row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "性别"
	cell = row.AddCell()
	cell.Value = "学号"
	cell = row.AddCell()
	cell.Value = "专业"
	cell = row.AddCell()
	cell.Value = "电话"
	cell = row.AddCell()
	cell.Value = "电子邮件"
	cell = row.AddCell()
	cell.Value = "报名时间"

	for i := 0; i < len(ActivityEnrollDetails); i++ {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = ActivityEnrollDetails[i].Username
		cell = row.AddCell()
		cell.Value = ActivityEnrollDetails[i].Name
		cell = row.AddCell()
		cell.Value = ActivityEnrollDetails[i].Sex
		cell = row.AddCell()
		cell.Value = ActivityEnrollDetails[i].StudentNumber
		cell = row.AddCell()
		cell.Value = ActivityEnrollDetails[i].Major
		cell = row.AddCell()
		cell.Value = ActivityEnrollDetails[i].PhoneNumber
		cell = row.AddCell()
		cell.Value = ActivityEnrollDetails[i].Email
		cell = row.AddCell()
		cell.Value = ActivityEnrollDetails[i].CreatedAt.Format("2006-01-02 15:04:05")
	}
	filename = "files/" + cast.ToString(time.Now().Unix()) + ".xlsx"
	err = file.Save(filename)
	return filename, err
}
func CreateActivityEnroll(accessToken string, input ActivityEnroll) (err error) {
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
	result := db.Table("activity").Where("activity_name = ?", input.ActivityName).Take(&activity)
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	var activityEnroll ActivityEnroll
	result = db.Table("activity_enroll").Where("activity_name = ? AND username = ?", input.ActivityName, user.Username).Take(&activityEnroll)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("创建失败")
		return
	}
	if result.RowsAffected != 0 {
		err = fmt.Errorf("活动已加入")
		return
	}
	id := uuid.New()
	input.Id = id
	input.Username = user.Username
	result = db.Table("activity_enroll").Create(&input)
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	return nil
}
func DelActivityEnroll(accessToken string, activityName string) (err error) {
	user, err := UserCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	var activityenroll ActivityEnroll
	result := db.Table("activity_enroll").Where("username = ? AND activity_name ", user.Username, activityName).Delete(&activityenroll)
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}
func AdminFindActivityEnroll(accessToken string, activityName string, page int, pageSize int) (ActivityEnrollDetails []ActivityEnrollDetails, total int, err error) {
	admin, err := AdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	var activity Activity
	result := db.Table("activity").Where("activity_name = ? AND club_name = ?", activityName, admin.ClubName).Take(&activity)
	if result.Error != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Raw("SELECT activity_enroll.id,activity_enroll.activity_name, activity_enroll.username,activity_enroll.created_at,resume.name,resume.sex,resume.student_number,resume.major,resume.phone_number,resume.email FROM  activity_enroll left join resume on activity_enroll.username = resume.username where activity_enroll.activity_name = ?", activityName).Limit(pageSize).Offset((page - 1) * pageSize).Order("activity_enroll.created_at desc").Find(&ActivityEnrollDetails)
	db.Table("activity_enroll").Where("activity_name = ? ", activityName).Count(&total)
	return ActivityEnrollDetails, total, nil
}
func UserFindActivityEnroll(accessToken string) (activityenroll []ActivityEnroll,err error) {
	user, err := UserCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Table("activity_enroll").Where("username = ? ", user.Username).Find(&activityenroll)
	return activityenroll,nil
}
func UserDelActivityEnroll(accessToken string,activityname string) (err error) {
	user, err := UserCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	result := db.Table("activity_enroll").Where("username = ? AND activity_name =?", user.Username,activityname).Delete(&ActivityEnroll{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}

