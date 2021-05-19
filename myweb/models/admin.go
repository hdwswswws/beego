package models

import (
	"fmt"
	"time"

	_ "github.com/bmizerany/pq"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Id           uuid.UUID `json:"id,omitempty"`
	Username     string    `json:"username,omitempty"`
	PasswordHash string    `json:"password_hash,omitempty"`
	ActiveState  int       `json:"active_state"`
	TokenState   int       `json:"token_state"`
	ClubName     string    `json:"club_name,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

func AdminCheckAccessToken(accessToken string) (admin Admin, err error) {
	result, err := parseJwtToken(accessToken,"admin")
	if err != nil || result["Role"].(string) != "admin" {
		err = fmt.Errorf("用户未登录")
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("用户未登录")
		return
	}
	dbresult := db.Table("admin").Where("username = ? ", result["Username"].(string)).Take(&admin)
	if dbresult.Error != nil {
		err = fmt.Errorf("用户不存在")
		return
	}
	if admin.TokenState == 0 {
		err = fmt.Errorf("用户未登录")
		return
	}
	return
}
func AdminChangePassword(accessToken string, previousPassword string, proposedPassword string) (err error) {
	admin, err := AdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(previousPassword))
	if err != nil {
		err = fmt.Errorf("密码错误")
		return
	}
	if previousPassword == proposedPassword {
		err = fmt.Errorf("两次密码不能一样")
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(proposedPassword), bcrypt.DefaultCost)
	result := db.Table("admin").Where("username = ?", admin.Username).Update(Admin{PasswordHash: string(passwordHash)})
	if result.Error != nil {
		err = fmt.Errorf("修改密码失败")
		return
	}
	return nil
}
