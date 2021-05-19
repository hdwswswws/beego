package models

import (
	"fmt"
	"time"

	_ "github.com/bmizerany/pq"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ClubNotice struct {
	Id        uuid.UUID `json:"id,omitempty"`
	ClubName  string    `json:"club_name,omitempty"`
	Notice    string    `json:"notice,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
func FindClubNotice(accessToken string, page int, pageSize int) (clubNotices []ClubNotice, total int, err error) {
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
	db.Table("club_notice").Select([]string{"id", "notice", "created_at"}).Where("club_name = ?", admin.ClubName,).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&clubNotices)
	db.Table("club_notice").Where("club_name = ? ", admin.ClubName).Count(&total)
	return clubNotices, total, nil
}
func EditClubNotice(accessToken string,notice string,id string) (err error) {
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
	var clubNotice = ClubNotice{Notice:notice}
	result := db.Table("club_notice").Where("club_name = ? AND id = ?", admin.ClubName,id).Update(&clubNotice)
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	return nil
}
func DelClubNotice(accessToken string, id string) (err error) {
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
	result := db.Table("club_notice").Where("club_name = ? AND id = ?", admin.ClubName,id).Delete(ClubNotice{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}
func CreateClubNotice(accessToken string, input ClubNotice) (err error) {
	admin, err := AdminCheckAccessToken(accessToken)
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
	input.Id=id
	input.ClubName=admin.ClubName
	result := db.Table("club_notice").Create(&input)
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	return nil
}
func UserFindClubNotice(username string) (clubs []ClubMember,clubNotices []ClubNotice, err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Table("club_member").Select([]string{"club_name,role"}).Where("username = ?",username).Find(&clubs)
	db.Raw("SELECT club_member.username,club_notice.id, club_notice.club_name,club_notice.notice, club_notice.created_at FROM club_notice left join club_member on club_notice.club_name = club_member.club_name where club_member.username = ?", username).Order("club_notice.created_at desc").Find(&clubNotices)
	return clubs,clubNotices, nil
}