package models

import (
	"fmt"
	"time"

	_ "github.com/bmizerany/pq"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type SuperAdmin struct {
	Id           uuid.UUID `json:"id,omitempty"`
	Username     string    `json:"username,omitempty"`
	PasswordHash string    `json:"password_hash,omitempty"`
	TokenState   int       `json:"token_state"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

func SuperAdminCheckAccessToken(accessToken string) (superadmin SuperAdmin, err error) {
	result, err := parseJwtToken(accessToken,"super_admin")
	if err != nil || result["Role"].(string) != "super_admin" {
		err = fmt.Errorf("用户未登录")
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("用户未登录")
		return
	}
	dbresult := db.Table("super_admin").Where("username = ? ", result["Username"].(string)).Take(&superadmin)
	if dbresult.Error != nil {
		err = fmt.Errorf("用户不存在")
		return
	}
	if superadmin.TokenState == 0 {
		err = fmt.Errorf("用户未登录")
		return
	}
	return
}
func FindUser(accessToken string, page int, pageSize int, input User) (users []User, total int, err error) {
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
	if input.ActiveState != 0 {
		db.Select([]string{"id", "username", "active_state", "created_at"}).Where("username LIKE ? AND active_state = ?", "%"+input.Username+"%", input.ActiveState).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&users)
		db.Model(&User{}).Where("username LIKE ? AND active_state = ?", "%"+input.Username+"%", input.ActiveState).Count(&total)
		return users, total, nil
	}
	db.Select([]string{"id", "username", "active_state", "created_at"}).Where("username LIKE ?", "%"+input.Username+"%").Not("active_state = ?", 0).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&users)
	db.Model(&User{}).Where("username LIKE ?", "%"+input.Username+"%").Not("active_state = ?", 0).Count(&total)
	return users, total, nil
}
func FindAdmin(accessToken string, page int, pageSize int, input Admin) (admins []Admin, total int, err error) {
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
	if input.ClubName == "" && input.ActiveState != 0 {
		db.Table("admin").Select([]string{"id", "username", "active_state", "token_state", "created_at", "club_name"}).Where("username LIKE ? AND active_state = ? ", "%"+input.Username+"%", input.ActiveState).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&admins)
		db.Table("admin").Where("username LIKE ? AND active_state = ?", "%"+input.Username+"%", input.ActiveState).Count(&total)
		return admins, total, nil
	}
	if input.ClubName != "" && input.ActiveState == 0 {
		db.Table("admin").Select([]string{"id", "username", "active_state", "token_state", "created_at", "club_name"}).Where("username LIKE ? AND club_name = ? ", "%"+input.Username+"%", input.ClubName).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&admins)
		db.Table("admin").Where("username LIKE ? AND club_name = ?", "%"+input.Username+"%", input.ClubName).Count(&total)
		return admins, total, nil
	}
	if input.ClubName == "" && input.ActiveState == 0 {
		db.Table("admin").Select([]string{"id", "username", "active_state", "token_state", "created_at", "club_name"}).Where("username LIKE ? ", "%"+input.Username+"%").Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&admins)
		db.Table("admin").Where("username LIKE ?", "%"+input.Username+"%").Count(&total)
		return admins, total, nil
	}
	db.Table("admin").Select([]string{"id", "username", "active_state", "token_state", "created_at", "club_name"}).Where("username LIKE ? AND active_state = ? AND club_name = ?", "%"+input.Username+"%", input.ActiveState, input.ClubName).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&admins)
	db.Table("admin").Where("username LIKE ? AND active_state = ? AND club_name = ?", "%"+input.Username+"%", input.ActiveState, input.ClubName).Count(&total)
	return admins, total, nil
}
func CreateUser(accessToken string, input User) (err error) {
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
	var user User
	result := db.Table("users").Where("username = ?", input.Username).Take(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("创建失败")
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.DefaultCost)
	if result.RowsAffected != 0 {
		if user.ActiveState == 0 {
			db.Table("users").Where("username = ?", input.Username).Update(User{Username: input.Username, PasswordHash: string(passwordHash), ActiveState: 1, TokenState: 0})
			return nil
		}
		err = fmt.Errorf("用户已存在")
		return
	}
	id := uuid.New()
	user = User{Id: id, Username: input.Username, PasswordHash: string(passwordHash), ActiveState: 1, TokenState: 0}
	result = db.Table("users").Create(&user)
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	return nil
}
func EditUserPassword(accessToken string, input User) (err error) {
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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.DefaultCost)
	result := db.Table("users").Where("username = ?", input.Username).Update(User{PasswordHash: string(passwordHash)})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	return nil
}
func EditUserState(accessToken string, input User) (err error) {
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
	result := db.Table("users").Where("username = ?", input.Username).Update(map[string]interface{}{"active_state": input.ActiveState})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	return nil
}
func DelUser(accessToken string, input User) (err error) {
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
	var user User
	result := db.Table("users").Where("username = ? ", input.Username).Delete(&user)
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}
func CreateAdmin(accessToken string, input Admin) (err error) {
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
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	var admin Admin
	result = db.Table("admin").Where("username = ?", input.Username).Take(&admin)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("创建失败")
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.DefaultCost)
	if result.RowsAffected != 0 {
		err = fmt.Errorf("用户已存在")
		return
	}
	id := uuid.New()
	admin = Admin{Id: id, Username: input.Username, PasswordHash: string(passwordHash), ActiveState: 1, TokenState: 0, ClubName: input.ClubName}
	result = db.Table("admin").Create(&admin)
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	return nil
}
func DelAdmin(accessToken string, input Admin) (err error) {
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
	var admin Admin
	result := db.Table("admin").Where("username = ? ", input.Username).Delete(&admin)
	if result.Error != nil {
		err = fmt.Errorf("删除失败")
		return
	}
	return nil
}
func EditAdmin(accessToken string, input Admin) (err error) {
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
	if input.PasswordHash != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.DefaultCost)
		result := db.Table("admin").Where("username = ?", input.Username).Update(Admin{PasswordHash: string(passwordHash), ActiveState: input.ActiveState})
		if result.Error != nil {
			err = fmt.Errorf("修改失败")
			return err
		}
		return nil
	}
	result := db.Table("admin").Where("username = ?", input.Username).Update(Admin{ActiveState: input.ActiveState})
	if result.Error != nil {
		err = fmt.Errorf("修改失败")
		return err
	}
	return nil
}
func SuperAdminChangePassword(accessToken string, previousPassword string, proposedPassword string) (err error) {
	var superAdmin SuperAdmin
	superAdmin, err = SuperAdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	err = bcrypt.CompareHashAndPassword([]byte(superAdmin.PasswordHash), []byte(previousPassword))
	if err != nil {
		err = fmt.Errorf("密码错误")
		return
	}
	if previousPassword == proposedPassword {
		err = fmt.Errorf("两次密码不能一样")
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(proposedPassword), bcrypt.DefaultCost)
	result := db.Table("super_admin").Where("username = ?", superAdmin.Username).Update(SuperAdmin{PasswordHash: string(passwordHash)})
	if result.Error != nil {
		err = fmt.Errorf("重置密码失败")
		return
	}
	return nil
}
