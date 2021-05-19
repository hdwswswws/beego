package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/bmizerany/pq"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

const (
	dbConfig               = "host=localhost user=postgres dbname=test sslmode=disable password=123456"
	timeLayout             = "2006-01-02 15:04:05"
	userJwtSecretKey       = "userlogin"
	adminJwtSecretKey      = "adminlogin"
	superAdminJwtSecretKey = "superlogin"
	jwtExpireSecond        = 259200
)

type User struct {
	Id                   uuid.UUID `json:"id,omitempty"`
	Username             string    `json:"username,omitempty"`
	PasswordHash         string    `json:"password_hash,omitempty"`
	ActiveState          int       `json:"active_state"`
	TokenState           int       `json:"token_state"`
	ConfirmationCodeHash string    `json:"confirmation_code_hash,omitempty"`
	CodeTime             string    `json:"code_time,omitempty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
}

func UserCheckAccessToken(accessToken string) (user User, err error) {
	result, err := parseJwtToken(accessToken,"user")
	if err != nil || result["Role"].(string) != "user" {
		err = fmt.Errorf("用户未登录")
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("用户未登录")
		return
	}
	dbresult := db.Table("users").Where("username = ? ", result["Username"].(string)).Take(&user)
	if dbresult.Error != nil {
		err = fmt.Errorf("用户不存在")
		return
	}
	if user.TokenState == 0 {
		err = fmt.Errorf("用户未登录")
		return
	}
	return
}
func createAccessToken(name string, role string) (token string, err error) {
	if role == "user" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"Username": name,
			"Role":     "user",
			"exp":      time.Now().Add(time.Second * time.Duration(jwtExpireSecond)).Unix(),
		})
		return token.SignedString([]byte(userJwtSecretKey))
	}
	if role == "admin" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"Username": name,
			"Role":     "admin",
			"exp":      time.Now().Add(time.Second * time.Duration(jwtExpireSecond)).Unix(),
		})
		return token.SignedString([]byte(adminJwtSecretKey))
	}
	if role == "super_admin" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"Username": name,
			"Role":     "super_admin",
			"exp":      time.Now().Add(time.Second * time.Duration(jwtExpireSecond)).Unix(),
		})
		return token.SignedString([]byte(superAdminJwtSecretKey))
	}
	err = fmt.Errorf("unknown role")
	return
}
func parseJwtToken(token string, role string) (jwt.MapClaims, error) {
	var key jwt.Keyfunc
	if role == "user" {
		key = func(token *jwt.Token) (interface{}, error) {
			return []byte(userJwtSecretKey), nil
		}
	}
	if role == "admin" {
		key = func(token *jwt.Token) (interface{}, error) {
			return []byte(adminJwtSecretKey), nil
		}
	}
	if role == "super_admin" {
		key = func(token *jwt.Token) (interface{}, error) {
			return []byte(superAdminJwtSecretKey), nil
		}
	}
	result, error := jwt.Parse(token, key)
	if error != nil {
		return nil, error
	}
	finToken := result.Claims.(jwt.MapClaims)
	return finToken, nil
}
func nowTime() string {
	timeUnix := time.Now().Unix()
	formatTimeStr := time.Unix(timeUnix, 0).Format(timeLayout)
	return formatTimeStr
}

func timeSub(input string) (err error) {
	local, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, input, local)
	TimeNow := time.Now()
	left := TimeNow.Sub(theTime)
	if left.Seconds() > 300 {
		err = fmt.Errorf("验证码超时，请重新获取")
		return
	}
	return
}
func verificationCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}
func SendMail(stuEmail string, subject string, body string) (err error) {
	mailTo := []string{stuEmail}
	port, _ := strconv.Atoi("465")

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress("hd0728@qq.com", "验证码"))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.qq.com", port, "hd0728@qq.com", "adzrbbjhhzrfbbff")
	err = d.DialAndSend(m)
	if err != nil {
		return
	}
	return
}
func Register(input User) (username string, err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("注册失败")
		return
	}
	code := verificationCode()
	var user User
	result := db.Table("users").Where("username = ?", input.Username).Take(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("注册失败")
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.DefaultCost)
	codeHash, _ := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
	if result.RowsAffected != 0 {
		if user.ActiveState == 0 {
			db.Table("users").Where("username = ?", input.Username).Update(User{PasswordHash: string(passwordHash), ConfirmationCodeHash: string(codeHash), CodeTime: nowTime()})
			if SendMail(input.Username, "邮箱验证码", code) != nil {
				err = fmt.Errorf("验证码发送失败")
				return
			}
			username = input.Username
			return
		}
		err = fmt.Errorf("用户已存在")
		return
	}
	id := uuid.New()
	user = User{Id: id, Username: input.Username, PasswordHash: string(passwordHash), ActiveState: 0, TokenState: 0, ConfirmationCodeHash: string(codeHash), CodeTime: nowTime()}
	result = db.Table("users").Create(&user)
	if result.Error != nil {
		err = fmt.Errorf("注册失败")
		return
	}
	if SendMail(input.Username, "邮箱验证码", code) != nil {
		err = fmt.Errorf("验证码发送失败")
		return
	}
	username = input.Username
	return
}
func ConfirmRegister(username string, code string) (err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	var user User
	result := db.Table("users").Where("username = ?", username).Take(&user)
	if result.Error != nil {
		err = fmt.Errorf("验证码不存在")
		return
	}
	err = timeSub(user.CodeTime)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.ConfirmationCodeHash), []byte(code))
	if err != nil {
		err = fmt.Errorf("验证码错误")
		return
	}
	result = db.Table("users").Where("username = ?", username).Update(User{ActiveState: 1})
	if result.Error != nil {
		err = fmt.Errorf("创建失败")
		return
	}
	return
}
func Login(input User, role string) (token string, err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("登录失败")
		return
	}
	if role == "user" {
		var user User
		result := db.Table("users").Where("username = ?", input.Username).Take(&user)
		if result.Error == gorm.ErrRecordNotFound {
			err = fmt.Errorf("用户不存在")
			return
		}
		if result.Error != nil {
			err = fmt.Errorf("登录失败")
			return
		}
		if user.ActiveState == 0 {
			err = fmt.Errorf("用户不存在")
			return
		}
		if user.ActiveState == 2 {
			err = fmt.Errorf("用户未激活")
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.PasswordHash))
		if err != nil {
			err = fmt.Errorf("密码错误")
			return
		}
		result = db.Table("users").Where("username = ?", input.Username).Update(User{TokenState: 1})
		token, _ = createAccessToken(input.Username, "user")
		return
	}
	if role == "admin" {
		var admin Admin
		result := db.Table("admin").Where("username = ?", input.Username).Take(&admin)
		if result.Error == gorm.ErrRecordNotFound {
			err = fmt.Errorf("用户不存在")
			return
		}
		if result.Error != nil {
			err = fmt.Errorf("登录失败")
			return
		}
		if admin.ActiveState == 2 {
			err = fmt.Errorf("用户未激活")
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(input.PasswordHash))
		if err != nil {
			err = fmt.Errorf("密码错误")
			return
		}
		result = db.Table("admin").Where("username = ?", input.Username).Update(Admin{TokenState: 1})
		token, _ = createAccessToken(input.Username, "admin")
		return
	}
	if role == "super_admin" {
		var superadmin SuperAdmin
		result := db.Table("super_admin").Where("username = ?", input.Username).Take(&superadmin)
		if result.Error == gorm.ErrRecordNotFound {
			err = fmt.Errorf("用户不存在")
			return
		}
		if result.Error != nil {
			err = fmt.Errorf("登录失败")
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(superadmin.PasswordHash), []byte(input.PasswordHash))
		if err != nil {
			err = fmt.Errorf("密码错误")
			return
		}
		result = db.Table("super_admin").Where("username = ?", input.Username).Update(SuperAdmin{TokenState: 1})
		token, _ = createAccessToken(input.Username, "super_admin")
		return
	}
	err = fmt.Errorf("未知角色")
	return
}
func ResendConfirmationCode(username string) (err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("验证码发送失败")
		return
	}
	var user User
	result := db.Table("users").Where("username = ?", username).Take(&user)
	if result.Error != nil {
		err = fmt.Errorf("用户不存在")
		return
	}
	code := verificationCode()
	codeHash, _ := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
	db.Table("users").Where("username = ?", username).Update(User{ConfirmationCodeHash: string(codeHash), CodeTime: nowTime()})
	if SendMail(username, "邮箱验证码", code) != nil {
		err = fmt.Errorf("验证码发送失败")
		return
	}
	return nil
}
func ForgotPassword(username string, password string, code string) (err error) {
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("重置密码失败")
		return
	}
	var user User
	result := db.Table("users").Where("username = ?", username).Take(&user)
	if result.Error != nil {
		err = fmt.Errorf("用户不存在")
		return
	}
	err = timeSub(user.CodeTime)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.ConfirmationCodeHash), []byte(code))
	if err != nil {
		err = fmt.Errorf("验证码错误")
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	result = db.Table("users").Where("username = ?", username).Update(User{PasswordHash: string(passwordHash)})
	if result.Error != nil {
		err = fmt.Errorf("重置密码失败")
		return
	}
	return
}
func ChangePassword(accessToken string, previousPassword string, proposedPassword string) (err error) {
	var user User
	user, err = UserCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(previousPassword))
	if err != nil {
		err = fmt.Errorf("密码错误")
		return
	}
	if previousPassword == proposedPassword {
		err = fmt.Errorf("两次密码不能一样")
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(proposedPassword), bcrypt.DefaultCost)
	result := db.Table("users").Where("username = ?", user.Username).Update(User{PasswordHash: string(passwordHash)})
	if result.Error != nil {
		err = fmt.Errorf("重置密码失败")
		return
	}
	return nil
}
func Logout(accessToken string, role string) (err error) {
	if role == "user" {
		var user User
		user, err = UserCheckAccessToken(accessToken)
		if err != nil {
			return err
		}
		db, err := gorm.Open("postgres", dbConfig)
		defer db.Close()
		if err != nil {
			err = fmt.Errorf("登出失败")
			return err
		}
		result := db.Table("users").Where("username = ?", user.Username).Update(map[string]interface{}{"token_state": 0})
		if result.Error != nil {
			err = fmt.Errorf("登出失败")
			return err
		}
		return nil
	}
	if role == "admin" {
		var admin Admin
		admin, err = AdminCheckAccessToken(accessToken)
		if err != nil {
			return err
		}
		db, err := gorm.Open("postgres", dbConfig)
		defer db.Close()
		if err != nil {
			err = fmt.Errorf("登出失败")
			return err
		}
		result := db.Table("admin").Where("username = ?", admin.Username).Update(map[string]interface{}{"token_state": 0})
		if result.Error != nil {
			err = fmt.Errorf("登出失败")
			return err
		}
		return nil
	}
	var superadmin SuperAdmin
	superadmin, err = SuperAdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	db, err := gorm.Open("postgres", dbConfig)
	defer db.Close()
	if err != nil {
		err = fmt.Errorf("登出失败")
		return
	}
	result := db.Table("super_admin").Where("username = ?", superadmin.Username).Update(map[string]interface{}{"token_state": 0})
	if result.Error != nil {
		err = fmt.Errorf("登出失败")
		return
	}
	return nil
}