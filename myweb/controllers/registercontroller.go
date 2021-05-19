package controllers

import (
	"myweb/models"

	"github.com/astaxie/beego"
)
type RegisterController struct {
	beego.Controller
}
type ConfirmRegisterController struct {
	beego.Controller
}
func (c *RegisterController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名或密码不能为空"}
		c.ServeJSON()
		return
	}
	emailerr := verifyEmail(username)
	if emailerr != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": emailerr.Error()}
		c.ServeJSON()
		return
	}
	pwderr := verifyPwd(password)
	if pwderr != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": pwderr.Error()}
		c.ServeJSON()
		return
	}
	username, err := models.Register(models.User{Username: username, PasswordHash: password})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "username": username}
	c.ServeJSON()
	return
}
func (c *ConfirmRegisterController) Post() {
	username := c.GetString("username")
	code := c.GetString("code")
	if username == "" || code == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名或验证码不能为空"}
		c.ServeJSON()
		return
	}
	emailerr := verifyEmail(username)
	if emailerr != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": emailerr.Error()}
		c.ServeJSON()
		return
	}
	err := models.ConfirmRegister(username, code)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建成功"}
	c.ServeJSON()
	return
}