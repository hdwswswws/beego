package models

import (
	"fmt"
	"os"
	"time"

	_ "github.com/bmizerany/pq"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Club struct {
	Id        uuid.UUID `json:"id,omitempty"`
	ClubName  string    `json:"club_name,omitempty"`
	Introduce string    `json:"introduce,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Recruit   int       `json:"recruit,omitempty"`
	Change    int       `json:"change,omitempty"`
	Logo      string    `json:"logo"`
}

func CreateClub(accessToken string, input Club) (err error) {
	_, err = SuperAdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	var club Club
	result := db.Table("club").Where("club_name = ?", input.ClubName).Take(&club)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("创建失败")
		return
	}
	if result.RowsAffected != 0 {
		err = fmt.Errorf("名字已存在")
		return
	}
	id := uuid.New()
	club = Club{Id: id, ClubName: input.ClubName, Introduce: input.Introduce, Recruit: 1, Change: 1}
	result = db.Table("club").Create(&club)
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	return nil
}
func DelClub(accessToken string, input Club) (err error) {
	_, err = SuperAdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	result := db.Table("admin").Where("club_name = ? ", input.ClubName).Update(map[string]interface{}{"club_name": ""})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	result = db.Table("change_enroll").Where("club_name = ? ", input.ClubName).Delete(ChangeEnroll{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	result = db.Table("recruit_enroll").Where("club_name = ? ", input.ClubName).Delete(&ClubMember{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	result = db.Table("club_member").Where("club_name = ? ", input.ClubName).Delete(&ClubMember{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	result = db.Table("club_notice").Where("club_name = ? ", input.ClubName).Delete(&ClubNotice{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	result = db.Table("club_role").Where("club_name = ? ", input.ClubName).Delete(&ClubRole{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	var club Club
	result = db.Table("club").Where("club_name = ?", input.ClubName).Take(&club)
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	result = db.Table("club").Where("club_name = ? ", input.ClubName).Delete(&Club{})
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	_ = os.Remove("./static/images/" + club.Logo)
	return nil
}
func AdminEditClub(accessToken string, input Club) (err error) {
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
	result := db.Table("club").Where("club_name = ? ", admin.ClubName).Update(Club{Introduce: input.Introduce, Recruit: input.Recruit, Change: input.Change})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	return nil
}
func SAdminEditClub(input Club, clubname string) (err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	result := db.Table("club").Where("club_name = ? ", clubname).Update(input)
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	return nil
}
func EditClubName(accessToken string, clubname string, input Club) (err error) {
	_, err = SuperAdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	var club Club
	result := db.Table("club").Where("club_name = ?", input.ClubName).Take(&club)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("修改失败")
		return
	}
	if result.RowsAffected != 0 {
		err = fmt.Errorf("名字已存在")
		return
	}
	result = db.Table("club").Where("club_name = ? ", clubname).Update(&Club{ClubName: input.ClubName})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	result = db.Table("recruit_enroll").Where("club_name = ? ", clubname).Update(&RecruitEnroll{ClubName: input.ClubName})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	result = db.Table("change_enroll").Where("club_name = ? ", clubname).Update(&ChangeEnroll{ClubName: input.ClubName})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	result = db.Table("club_member").Where("club_name = ? ", clubname).Update(&ClubMember{ClubName: input.ClubName})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	result = db.Table("club_notice").Where("club_name = ? ", clubname).Update(&ClubNotice{ClubName: input.ClubName})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	result = db.Table("club_role").Where("club_name = ? ", clubname).Update(&ClubRole{ClubName: input.ClubName})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	result = db.Table("admin").Where("club_name = ? ", clubname).Update(&Admin{ClubName: input.ClubName})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	return nil
}
func FindClub(accessToken string, page int, pageSize int, clubname string) (clubs []Club, total int, err error) {
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
	db.Table("club").Select([]string{"id", "club_name", "introduce", "created_at"}).Where("club_name LIKE ?", "%"+clubname+"%").Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&clubs)
	db.Table("club").Where("club_name LIKE ?", "%"+clubname+"%").Model(&Club{}).Count(&total)
	return clubs, total, nil
}
func FindAllClub(accessToken string) (clubs []Club, err error) {
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
	db.Table("club").Select([]string{"club_name"}).Find(&clubs)
	return clubs, nil
}
func AdminFindClub(accessToken string) (club Club, err error) {
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
	result := db.Table("club").Where("club_name = ?", admin.ClubName).Find(&club)
	if result.Error != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	return club, nil
}
func UserFindClub(accessToken string) (clubs []Club, total int, err error) {
	_, err = UserCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Table("club").Order("created_at desc").Find(&clubs)
	db.Table("club").Count(&total)
	return clubs, total, nil
}
func UserFindClubRecruit(accessToken string) (clubs []Club, total int, err error) {
	_, err = UserCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Table("club").Where("recruit = ?", 2).Order("created_at desc").Find(&clubs)
	db.Table("club").Where("recruit = ?", 2).Count(&total)
	return clubs, total, nil
}
func UserFindClubChange(accessToken string) (clubs []Club, total int, err error) {
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
	db.Raw("SELECT club.id, club.club_name, club.introduce, club.created_at, club.logo FROM club_member left join club on club_member.club_name = club.club_name where club_member.username = ? and club.change = ?", user.Username, 2).Order("created_at desc").Find(&clubs)
	return clubs, len(clubs), nil
}
func UserClubDetail(accessToken string, clubname string) (club Club, err error) {
	_, err = UserCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Table("club").Where("club_name = ?", clubname).Find(&club)
	return club, nil
}
