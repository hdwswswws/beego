package models

import (
	"fmt"
	"os"
	"time"

	_ "github.com/bmizerany/pq"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Activity struct {
	Id           uuid.UUID `json:"id,omitempty"`
	ClubName     string    `json:"club_name,omitempty"`
	Introduce    string    `json:"introduce,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	State        int       `json:"state,omitempty"`
	Remarks      string    `json:"remarks"`
	ActivityName string    `json:"activity_name,omitempty"`
	Cover        string    `json:"cover"`
}
type Activitys struct {
	Id           uuid.UUID `json:"id,omitempty"`
	ClubName     string    `json:"club_name,omitempty"`
	Introduce    string    `json:"introduce,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	State        int       `json:"state,omitempty"`
	Remarks      string    `json:"remarks"`
	ActivityName string    `json:"activity_name,omitempty"`
	Cover        string    `json:"cover"`
	Count        int       `json:"count"`
}

func CreateActivity(accessToken string, input Activity) (err error) {
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
	var club Club
	result := db.Table("club").Where("club_name = ?", admin.ClubName).Take(&club)
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	var activity Activity
	result = db.Table("activity").Where("activity_name = ? AND club_name = ?", input.ActivityName, admin.ClubName).Take(&activity)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("创建失败")
		return
	}
	if result.RowsAffected != 0 {
		err = fmt.Errorf("活动名已存在")
		return
	}
	id := uuid.New()
	input.Id = id
	input.ClubName = admin.ClubName
	input.State = 1
	result = db.Table("activity").Create(&input)
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	return nil
}
func EditActivity(clubName string, input Activity) (err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	result := db.Table("activity").Where("activity_name = ? AND club_name = ?", input.ActivityName, clubName).Update(input)
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	return nil
}
func EditActivityState(accessToken string, input Activity) (err error) {
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
	result := db.Table("activity").Where("activity_name = ? ", input.ActivityName).Update(input)
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	return nil
}
func DelActivity(accessToken string, input Activity) (err error) {
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
	var activity Activity
	result := db.Table("activity").Where("activity_name = ? AND club_name = ?", input.ActivityName, admin.ClubName).Take(&activity)
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	var activityComment ActivityComment
	result = db.Table("activity_comment").Where("activity_name = ? ", input.ActivityName).Delete(activityComment)
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	var activityEnroll ActivityEnroll
	result = db.Table("activity_enroll").Where("activity_name = ? ", input.ActivityName).Delete(activityEnroll)
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	result = db.Table("activity").Where("activity_name = ? ", input.ActivityName).Delete(activity)
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	_ = os.Remove("./static/images/" + activity.Cover)
	return nil
}
func FindActivity(page int, pageSize int, input Activity) (activitys []Activity, total int, err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	if input.ClubName == "" && input.State != 0 {
		db.Table("activity").Where("activity_name LIKE ? AND state = ? ", "%"+input.ActivityName+"%", input.State).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&activitys)
		db.Table("activity").Where("activity_name LIKE ? AND state = ?", "%"+"%"+input.ClubName+"%", input.State).Count(&total)
		return activitys, total, nil
	}
	if input.ClubName != "" && input.State == 0 {
		db.Table("activity").Where("activity_name LIKE ? AND club_name = ? ", "%"+input.ActivityName+"%", input.ClubName).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&activitys)
		db.Table("activity").Where("activity_name LIKE ? AND club_name = ? ", "%"+input.ActivityName+"%", input.ClubName).Count(&total)
		return activitys, total, nil
	}
	if input.ClubName == "" && input.State == 0 {
		db.Table("activity").Where("activity_name LIKE ? ", "%"+input.ActivityName+"%").Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&activitys)
		db.Table("activity").Where("activity_name LIKE ?", "%"+input.ActivityName+"%").Count(&total)
		return activitys, total, nil
	}
	db.Table("activity").Where("activity_name  LIKE ? AND state = ? AND club_name = ?", "%"+input.ActivityName+"%", input.State, input.ClubName).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&activitys)
	db.Table("activity").Where("activity_name  LIKE ? AND state = ? AND club_name = ?", "%"+input.ActivityName+"%", input.State, input.ClubName).Count(&total)
	return activitys, total, nil
}
func FindActivityIntroduce(input string, clubname string) (activity Activity, err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	if clubname == "" {
		result := db.Table("activity").Where("activity_name = ? ", input).Take(&activity)
		if result.Error != nil {
			err = fmt.Errorf("获取失败")
			return
		}
		return
	}
	result := db.Table("activity").Where("activity_name = ? AND club_name = ?", input, clubname).Take(&activity)
	if result.Error != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	return
}
func UserFindActivity() (activitys []Activitys, total int, err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Table("activity").Where("state = ?", 3).Order("created_at desc").Find(&activitys)
	for i := 0; i < len(activitys); i++ {
		activitys[i].Count, _ = UserCountActivityComment(activitys[i].ActivityName)
	}
	db.Table("activity").Where("state = ?", 3).Count(&total)
	return activitys, total, nil
}
func UserActivityDetail(activityname string) (activity Activitys, err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("获取失败")
		return
	}
	db.Table("activity").Where("activity_name = ? ", activityname).Find(&activity)
	activity.Count, _ = UserCountActivityComment(activity.ActivityName)
	return activity, nil
}
