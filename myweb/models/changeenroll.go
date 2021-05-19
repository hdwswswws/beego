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

type ChangeEnroll struct {
	Id        uuid.UUID `json:"id,omitempty"`
	ClubName  string    `json:"club_name,omitempty"`
	Username  string    `json:"username,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Role      string    `json:"role"`
}
type ChangeEnrollDetails struct {
	Id            uuid.UUID `json:"id,omitempty"`
	clubName      string    `json:"club_name,omitempty"`
	Username      string    `json:"username,omitempty"`
	Name          string    `json:"name,omitempty"`
	Sex           string    `json:"sex,omitempty"`
	StudentNumber string    `json:"student_number,omitempty"`
	Major         string    `json:"major,omitempty"`
	Role          string    `json:"role"`
	PhoneNumber   string    `json:"phone_number,omitempty"`
	Email         string    `json:"email,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

func DelChangeEnroll(accessToken string) (err error) {
	admin, err := AdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	result := db.Table("change_enroll").Where("club_name = ? ", admin.ClubName).Delete(&ChangeEnroll{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}
func ChangeEnrollExportExcel(accessToken string) (filename string, err error) {
	admin, err := AdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	var ChangeEnrollDetails []ChangeEnrollDetails
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Raw("SELECT change_enroll.id, change_enroll.username,change_enroll.role, change_enroll.created_at, resume.name, resume.sex, resume.student_number, resume.major, resume.phone_number, resume.email FROM change_enroll left join resume on change_enroll.username = resume.username where change_enroll.club_name = ?", admin.ClubName).Order("change_enroll.created_at desc").Find(&ChangeEnrollDetails)
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("sheet1")
	title := sheet.AddRow()
	titleRow := title.AddCell()
	titleRow.HMerge = 7 
	titleRow.Value = admin.ClubName+"的干部换届报名表"
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

	for i := 0; i < len(ChangeEnrollDetails); i++ {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = ChangeEnrollDetails[i].Username
		cell = row.AddCell()
		cell.Value = ChangeEnrollDetails[i].Name
		cell = row.AddCell()
		cell.Value = ChangeEnrollDetails[i].Sex
		cell = row.AddCell()
		cell.Value = ChangeEnrollDetails[i].StudentNumber
		cell = row.AddCell()
		cell.Value = ChangeEnrollDetails[i].Major
		cell = row.AddCell()
		cell.Value = ChangeEnrollDetails[i].PhoneNumber
		cell = row.AddCell()
		cell.Value = ChangeEnrollDetails[i].Email
		cell = row.AddCell()
		cell.Value = ChangeEnrollDetails[i].CreatedAt.Format("2006-01-02 15:04:05")
	}
	filename = "files/" + cast.ToString(time.Now().Unix()) + ".xlsx"
	err = file.Save(filename)
	return filename, err
}
func FindChangeEnroll(accessToken string, page int, pageSize int) (result []ChangeEnrollDetails, total int, err error) {
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
	db.Raw("SELECT change_enroll.id, change_enroll.username,change_enroll.role, change_enroll.created_at, resume.name, resume.sex, resume.student_number, resume.major, resume.phone_number, resume.email FROM change_enroll left join resume on change_enroll.username = resume.username where change_enroll.club_name = ?", admin.ClubName).Limit(pageSize).Offset((page - 1) * pageSize).Order("change_enroll.created_at desc").Find(&result)
	db.Table("change_enroll").Where("club_name = ? ", admin.ClubName).Count(&total)
	return
}
func CreateChangeEnroll(accessToken string,clubname string,role string) (err error) {
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
	var changeenroll ChangeEnroll
	result := db.Table("change_enroll").Where("username = ? AND club_name = ?", user.Username,clubname).Take(&changeenroll)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("加入失败")
		return
	}
	if result.RowsAffected != 0 {
		err = fmt.Errorf("已加入")
		return
	}
	db.Table("change_enroll").Create(&ChangeEnroll{Id:uuid.New(),ClubName:clubname,Username:user.Username,Role:role})
	return nil
}
func UserFindChangeEnroll(accessToken string) (changeenroll []ChangeEnroll,err error) {
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
	db.Table("change_enroll").Where("username = ? ", user.Username).Find(&changeenroll)
	return changeenroll,nil
}
func UserDelChangeEnroll(accessToken string,clubname string) (err error) {
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
	result := db.Table("change_enroll").Where("username = ? AND club_name =? ", user.Username,clubname).Delete(&ChangeEnroll{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}
