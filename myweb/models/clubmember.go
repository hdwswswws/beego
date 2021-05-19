package models

import (
		"fmt"
		"time"
	
		_ "github.com/bmizerany/pq"
		"github.com/google/uuid"
		"github.com/jinzhu/gorm"
)

type ClubMember struct {
	Id        uuid.UUID `json:"id,omitempty"`
	ClubName  string    `json:"club_name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Name      string    `json:"name,omitempty"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
func FindClubMember(accessToken string, page int, pageSize int,name string) (clubmembers []ClubMember, total int, err error) {
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
	db.Table("club_member").Select([]string{"id", "username", "name", "role"}).Where("club_name = ? AND name LIKE ? ", admin.ClubName,"%"+name+"%").Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&clubmembers)
	db.Table("club_member").Where("club_name = ? AND name LIKE ? ", admin.ClubName,"%"+name+"%").Count(&total)
	return clubmembers, total, nil
}
func EditClubMember(accessToken string, input ClubMember) (err error) {
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
	var clubRole ClubRole
	result := db.Table("club_role").Where("club_name = ? AND role = ?", admin.ClubName,input.Role).Take(&clubRole)
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	var clubMember = ClubMember{Role:input.Role}
	result = db.Table("club_member").Where("club_name = ? AND username = ?", admin.ClubName,input.Username).Update(&clubMember)
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	return nil
}
func CreateClubMember(accessToken string, username string,role string) (err error) {
	admin, err := AdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("加入失败")
		return
	}
	var resume Resume
	result := db.Table("resume").Where("username = ?",username ).Take(&resume)
	if result.Error != nil {
		err = fmt.Errorf("加入失败")
		return
	}
	var club ClubMember
	result = db.Table("club_member").Where("club_name = ? AND username = ?", admin.ClubName,username ).Take(&club)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("加入失败")
		return
	}
	if result.RowsAffected != 0 {
		err = fmt.Errorf("用户已加入")
		return
	}
	id := uuid.New()
	clubMember := ClubMember{Id: id, ClubName: admin.ClubName,Username: username,Name:resume.Name,Role:role}
	result = db.Table("club_member").Create(&clubMember)
	if result.Error != nil {
		err = fmt.Errorf("加入失败")
		return
	}
	return nil
}
func DelClubMember(accessToken string, username string) (err error) {
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
	result := db.Table("club_member").Where("club_name = ? AND username = ?", admin.ClubName,username).Delete(ClubMember{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}