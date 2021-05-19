package models

import (
	"fmt"
	"time"

	_ "github.com/bmizerany/pq"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ClubRole struct {
	Id        uuid.UUID `json:"id,omitempty"`
	ClubName  string    `json:"club_name,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
func CreateClubRole(accessToken string,role string) (err error) {
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
	var clubRole ClubRole
	result := db.Table("club_role").Where("club_name = ? AND role = ?", admin.ClubName,role ).Take(&clubRole)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("创建失败")
		return
	}
	if result.RowsAffected != 0 {
		err = fmt.Errorf("角色已存在")
		return
	}
	id := uuid.New()
	clubRole1 := ClubRole{Id: id, ClubName: admin.ClubName,Role:role}
	result = db.Table("club_role").Create(&clubRole1)
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	return nil
}
func DelClubRole(accessToken string, role string) (err error) {
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
	result := db.Table("club_member").Where("club_name = ? AND role = ?", admin.ClubName,role).Update(map[string]interface{}{"role": ""})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	result = db.Table("club_role").Where("club_name = ? AND role = ?", admin.ClubName,role).Delete(ClubMember{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}
func FindClubRole(clubname string) (clubroles []ClubRole, err error) {
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Table("club_role").Where("club_name = ? ", clubname).Find(&clubroles)
	return clubroles, nil
}
