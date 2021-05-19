package controllers

import (
	"fmt"
	"myweb/models"
	"regexp"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
type ResendConfirmationCodeController struct {
	beego.Controller
}
type ForgotPasswordController struct {
	beego.Controller
}
type ChangePasswordController struct {
	beego.Controller
}
type LogoutController struct {
	beego.Controller
}


func verifyEmail(email string) (err error) {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(email) {
		err = fmt.Errorf("邮箱格式不正确")
		return
	}
	return nil
}
func verifyPwd(password string) (err error) {
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	if password == "" {
		err = fmt.Errorf("密码不能为空")
		return
	}
	if len(password) < 8 {
		err = fmt.Errorf("密码少于8位")
		return
	}
	if b, err := regexp.MatchString(num, password); !b || err != nil {
		err = fmt.Errorf("密码没包含数字")
		return err
	}
	if b, err := regexp.MatchString(A_Z, password); !b || err != nil {
		err = fmt.Errorf("密码没包含大写字母")
		return err
	}
	if b, err := regexp.MatchString(a_z, password); !b || err != nil {
		err = fmt.Errorf("密码没包含小写字母")
		return err
	}
	return nil
}
func (c *LoginController) Get() {
	c.TplName = "login.html"
}
func (c *ChangePasswordController) Get() {
	c.TplName = "changepwd.html"
}
func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	role := c.GetString("role")
	if username == "" || password == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名或密码不能为空"}
		c.ServeJSON()
		return
	}
	if role != "user" && role != "admin" && role != "super_admin" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "没有此角色"}
		c.ServeJSON()
		return
	}
	token, err := models.Login(models.User{Username: username, PasswordHash: password}, role)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	
	c.SetSecureCookie("token","accesstoken",token,0)
	c.Ctx.SetCookie("username",username,259200)
	c.Data["json"] = map[string]interface{}{"code": 0, "accesstoken": token}
	c.ServeJSON()
	return
}
func (c *ResendConfirmationCodeController) Post() {
	username := c.GetString("username")
	if username == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名不能为空"}
		c.ServeJSON()
		return
	}
	emailerr := verifyEmail(username)
	if emailerr != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": emailerr.Error()}
		c.ServeJSON()
		return
	}
	err := models.ResendConfirmationCode(username)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "发送成功"}
	c.ServeJSON()
	return
}
func (c *ForgotPasswordController) Get() {
	c.TplName = "forgot.html"
}
func (c *ForgotPasswordController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	code := c.GetString("code")
	if username == "" || password == "" || code == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "属性不能为空"}
		c.ServeJSON()
		return
	}
	pwderr := verifyPwd(password)
	if pwderr != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": pwderr.Error()}
		c.ServeJSON()
		return
	}
	err := models.ForgotPassword(username, password, code)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "重置成功"}
	c.ServeJSON()
	return
}
func (c *ChangePasswordController) Post() {
	accessToken,_:= c.GetSecureCookie("token","accesstoken")
	previousPassword := c.GetString("previousPassword")
	proposedPassword := c.GetString("proposedPassword")
	if  previousPassword == "" || proposedPassword == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "属性不能为空"}
		c.ServeJSON()
		return
	}
	pwderr := verifyPwd(proposedPassword)
	if pwderr != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": pwderr.Error()}
		c.ServeJSON()
		return
	}
	err := models.ChangePassword(accessToken, previousPassword, proposedPassword)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "更改密码成功"}
	c.ServeJSON()
	return
}
func (c *LogoutController) Post() {
	accessToken,_:= c.GetSecureCookie("token","accesstoken")
	role := c.GetString("role")
	if role != "user" && role != "admin" && role != "super_admin" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "没有此角色"}
		c.ServeJSON()
		return
	}
	err := models.Logout(accessToken,role)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Ctx.SetCookie("accesstoken","",-1,"/")
	c.Ctx.SetCookie("username","",-1,"/")
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "登出成功"}
	c.ServeJSON()
	return
}

