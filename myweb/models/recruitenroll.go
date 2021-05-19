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

type RecruitEnroll struct {
	Id        uuid.UUID `json:"id,omitempty"`
	ClubName  string    `json:"club_name,omitempty"`
	Username  string    `json:"username,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	State     int       `json:"state,omitempty"`
}
type RecruitEnrollDetails struct {
	Id            uuid.UUID `json:"id,omitempty"`
	clubName      string    `json:"club_name,omitempty"`
	Username      string    `json:"username,omitempty"`
	State         int       `json:"state,omitempty"`
	Name          string    `json:"name,omitempty"`
	Sex           string    `json:"sex,omitempty"`
	StudentNumber string    `json:"student_number,omitempty"`
	Major         string    `json:"major,omitempty"`
	PhoneNumber   string    `json:"phone_number,omitempty"`
	Email         string    `json:"email,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

func DelRecruitEnroll(accessToken string) (err error) {
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
	result := db.Table("recruit_enroll").Where("club_name = ? ", admin.ClubName).Delete(&RecruitEnroll{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}
func RecruitEnrollExportExcel(accessToken string) (filename string, err error) {
	admin, err := AdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	var RecruitEnrollDetails []RecruitEnrollDetails
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Raw("SELECT recruit_enroll.id, recruit_enroll.username, recruit_enroll.state, recruit_enroll.created_at, resume.name, resume.sex, resume.student_number, resume.major, resume.phone_number, resume.email FROM recruit_enroll left join resume on recruit_enroll.username = resume.username where recruit_enroll.club_name = ? AND state = ?", admin.ClubName,2).Order("recruit_enroll.created_at desc").Find(&RecruitEnrollDetails)
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("sheet1")
	title := sheet.AddRow()
	titleRow := title.AddCell()
	titleRow.HMerge = 7 
	titleRow.Value = admin.ClubName+"的纳新面试报名表"
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

	for i := 0; i < len(RecruitEnrollDetails); i++ {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = RecruitEnrollDetails[i].Username
		cell = row.AddCell()
		cell.Value = RecruitEnrollDetails[i].Name
		cell = row.AddCell()
		cell.Value = RecruitEnrollDetails[i].Sex
		cell = row.AddCell()
		cell.Value = RecruitEnrollDetails[i].StudentNumber
		cell = row.AddCell()
		cell.Value = RecruitEnrollDetails[i].Major
		cell = row.AddCell()
		cell.Value = RecruitEnrollDetails[i].PhoneNumber
		cell = row.AddCell()
		cell.Value = RecruitEnrollDetails[i].Email
		cell = row.AddCell()
		cell.Value = RecruitEnrollDetails[i].CreatedAt.Format("2006-01-02 15:04:05")
	}
	filename = "files/" + cast.ToString(time.Now().Unix()) + ".xlsx"
	err = file.Save(filename)
	return filename, err
}
func FindRecruitEnroll(accessToken string, page int, pageSize int, state int,studentNumber string) (recruitenrolls []RecruitEnrollDetails, total int, err error) {
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
	if state == 0{
		db.Raw("SELECT recruit_enroll.id, recruit_enroll.username, recruit_enroll.state, recruit_enroll.created_at, resume.name, resume.sex, resume.student_number, resume.major, resume.phone_number, resume.email FROM recruit_enroll left join resume on recruit_enroll.username = resume.username where recruit_enroll.club_name = ? AND resume.student_number LIKE ?", admin.ClubName,"%"+studentNumber+"%").Limit(pageSize).Offset((page - 1) * pageSize).Order("recruit_enroll.created_at desc").Find(&recruitenrolls)
		db.Table(" recruit_enroll").Where("club_name = ? ", admin.ClubName).Count(&total)
		return
	}
    db.Raw("SELECT recruit_enroll.id, recruit_enroll.username, recruit_enroll.state, recruit_enroll.created_at, resume.name, resume.sex, resume.student_number, resume.major, resume.phone_number, resume.email FROM recruit_enroll left join resume on recruit_enroll.username = resume.username where recruit_enroll.club_name = ? AND state = ? AND resume.student_number LIKE ?", admin.ClubName,state,"%"+studentNumber+"%").Limit(pageSize).Offset((page - 1) * pageSize).Order("recruit_enroll.created_at desc").Find(&recruitenrolls)
	db.Table(" recruit_enroll").Where("club_name = ? AND state = ?", admin.ClubName,state).Count(&total)
	return
}
func EditRecruitEnroll(accessToken string,username string,state int) (err error) {
	admin, err := AdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	result := db.Table("recruit_enroll").Where("club_name = ? AND username = ? ", admin.ClubName,username).Update(&RecruitEnroll{State:state})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	return nil
}
func CreateRecruitEnroll(accessToken string,clubname string) (err error) {
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
	var recruitenroll RecruitEnroll
	result := db.Table("recruit_enroll").Where("username = ? AND club_name = ?", user.Username,clubname).Take(&recruitenroll)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("加入失败")
		return
	}
	if result.RowsAffected != 0 {
		err = fmt.Errorf("已加入")
		return
	}
	db.Table("recruit_enroll").Create(&RecruitEnroll{Id:uuid.New(),ClubName:clubname,Username:user.Username,State:1})
	return nil
}
func UserFindRecruitEnroll(accessToken string) (recruitenroll []RecruitEnroll,err error) {
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
	db.Table("recruit_enroll").Where("username = ? ", user.Username).Find(&recruitenroll)
	return recruitenroll,nil
}
func UserDelRecruitEnroll(accessToken string,clubname string) (err error) {
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
	result := db.Table("recruit_enroll").Where("username = ? AND club_name = ? ", user.Username,clubname).Delete(&RecruitEnroll{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}

