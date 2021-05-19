package controllers

import (
	"math"
	"myweb/models"

	"github.com/astaxie/beego"
	"github.com/google/uuid"
)

type FindUserController struct {
	beego.Controller
}
type FindAdminController struct {
	beego.Controller
}
type CreateUserController struct {
	beego.Controller
}
type EditUserPasswordController struct {
	beego.Controller
}
type EditUserStateController struct {
	beego.Controller
}
type DelUserController struct {
	beego.Controller
}
type FindClubController struct {
	beego.Controller
}
type CreateClubController struct {
	beego.Controller
}
type EditClubNameController struct {
	beego.Controller
}
type DelClubController struct {
	beego.Controller
}
type CreateAdminController struct {
	beego.Controller
}
type EditAdminController struct {
	beego.Controller
}
type DelAdminController struct {
	beego.Controller
}
type SuperAdminIndexController struct {
	beego.Controller
}
type SuperAdminClubController struct {
	beego.Controller
}
type SuperAdminController struct {
	beego.Controller
}
type FindAllClubController struct {
	beego.Controller
}
type SuperAdminActivityController struct {
	beego.Controller
}
type FindActivityController struct {
	beego.Controller
}
type EditActivityController struct {
	beego.Controller
}
type FindActivityIntroduceController struct {
	beego.Controller
}
type SuperAdminCommentController struct {
	beego.Controller
}
type FindCommentController struct {
	beego.Controller
}
type DelCommentController struct {
	beego.Controller
}
type SuperAdminChangePasswordController struct {
	beego.Controller
}
type ClubLoGoController struct {
	beego.Controller
}

func (c *SuperAdminActivityController) Get() {
	c.TplName = "sadminactivity.html"
}
func (c *SuperAdminIndexController) Get() {
	c.TplName = "sadminindex.html"
}
func (c *SuperAdminClubController) Get() {
	c.TplName = "sadminclub.html"
}
func (c *SuperAdminController) Get() {
	c.TplName = "sadmin.html"
}
func (c *SuperAdminCommentController) Get() {
	c.TplName = "sadmincomment.html"
}
func (c *FindUserController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	page, err := c.GetInt("page")
	username := c.GetString("username")
	state, err := c.GetInt("state")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	pagesize, err := c.GetInt("pagesize")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if page <= 0 || pagesize <= 0 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "属性不能小于等于0"}
		c.ServeJSON()
		return
	}
	if state != 0 && state != 1 && state != 2 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "无效状态码"}
		c.ServeJSON()
		return
	}
	users, total, err := models.FindUser(accessToken, page, pagesize, models.User{Username: username, ActiveState: state})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	totalpages := int(math.Ceil(float64(total) / float64(pagesize)))
	c.Data["json"] = map[string]interface{}{"code": 0, "total": total, "page": page, "totalpages": totalpages, "data": users}
	c.ServeJSON()
	return
}
func (c *FindAdminController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	page, err := c.GetInt("page")
	state, err := c.GetInt("state")
	username := c.GetString("username")
	clubname := c.GetString("clubname")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	pagesize, err := c.GetInt("pagesize")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if page <= 0 || pagesize <= 0 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "属性不能小于等于0"}
		c.ServeJSON()
		return
	}
	if state != 0 && state != 1 && state != 2 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "无效状态码"}
		c.ServeJSON()
		return
	}
	admins, total, err := models.FindAdmin(accessToken, page, pagesize, models.Admin{Username: username, ActiveState: state, ClubName: clubname})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	totalpages := int(math.Ceil(float64(total) / float64(pagesize)))
	c.Data["json"] = map[string]interface{}{"total": total, "page": page, "totalpages": totalpages, "data": admins}
	c.ServeJSON()
	return
}
func (c *CreateUserController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名或密码不能为空"}
		c.ServeJSON()
		return
	}
	pwderr := verifyPwd(password)
	if pwderr != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": pwderr.Error()}
		c.ServeJSON()
		return
	}
	err := models.CreateUser(accessToken, models.User{Username: username, PasswordHash: password})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建成功"}
	c.ServeJSON()
}
func (c *EditUserPasswordController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名或密码不能为空"}
		c.ServeJSON()
		return
	}
	pwderr := verifyPwd(password)
	if pwderr != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": pwderr.Error()}
		c.ServeJSON()
		return
	}
	err := models.EditUserPassword(accessToken, models.User{Username: username, PasswordHash: password})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改成功"}
	c.ServeJSON()
	return
}
func (c *EditUserStateController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	state, err := c.GetInt("state")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if username == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名不能为空"}
		c.ServeJSON()
		return
	}
	if state != 1 && state != 2 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "无效状态码"}
		c.ServeJSON()
		return
	}
	err = models.EditUserState(accessToken, models.User{Username: username, ActiveState: state})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改成功"}
	c.ServeJSON()
	return
}
func (c *DelUserController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	if username == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名不能为空"}
		c.ServeJSON()
		return
	}
	err := models.DelUser(accessToken, models.User{Username: username})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
	return
}
func (c *CreateClubController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubname := c.GetString("clubname")
	introduce := c.GetString("introduce")
	if clubname == "" || introduce == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "组织名或简介不能为空"}
		c.ServeJSON()
		return
	}
	err := models.CreateClub(accessToken, models.Club{ClubName: clubname, Introduce: introduce})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建成功"}
	c.ServeJSON()
}
func (c *FindClubController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubname := c.GetString("clubname")
	page, err := c.GetInt("page")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	pagesize, err := c.GetInt("pagesize")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if page <= 0 || pagesize <= 0 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "属性不能小于等于0"}
		c.ServeJSON()
		return
	}
	clubs, total, err := models.FindClub(accessToken, page, pagesize, clubname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	totalpages := int(math.Ceil(float64(total) / float64(pagesize)))
	c.Data["json"] = map[string]interface{}{"total": total, "page": page, "totalpages": totalpages, "data": clubs}
	c.ServeJSON()
	return
}
func (c *EditClubNameController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubname := c.GetString("clubname")
	newClubname := c.GetString("newclubname")
	if clubname == "" || newClubname == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "组织名不能为空"}
		c.ServeJSON()
		return
	}
	err := models.EditClubName(accessToken, clubname, models.Club{ClubName: newClubname})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改成功"}
	c.ServeJSON()
	return
}
func (c *DelClubController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubname := c.GetString("clubname")
	if clubname == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "组织名不能为空"}
		c.ServeJSON()
		return
	}
	err := models.DelClub(accessToken, models.Club{ClubName: clubname})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
	return
}
func (c *CreateAdminController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	password := c.GetString("password")
	club := c.GetString("club")
	if username == "" || password == "" || club == "" {
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
	err := models.CreateAdmin(accessToken, models.Admin{Username: username, PasswordHash: password, ClubName: club})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建成功"}
	c.ServeJSON()
}
func (c *EditAdminController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	password := c.GetString("password")
	state, err := c.GetInt("state")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if username == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名不能为空"}
		c.ServeJSON()
		return
	}
	if password != "" {
		pwderr := verifyPwd(password)
		if pwderr != nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": pwderr.Error()}
			c.ServeJSON()
			return
		}
	}
	if state != 1 && state != 2 && state != 0 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "无效状态码"}
		c.ServeJSON()
		return
	}
	err = models.EditAdmin(accessToken, models.Admin{Username: username, PasswordHash: password, ActiveState: state})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改成功"}
	c.ServeJSON()
	return
}
func (c *DelAdminController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	if username == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名不能为空"}
		c.ServeJSON()
		return
	}
	err := models.DelAdmin(accessToken, models.Admin{Username: username})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
	return
}
func (c *FindAllClubController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubs, err := models.FindAllClub(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = clubs
	c.ServeJSON()
	return
}
func (c *FindActivityController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityName := c.GetString("activityname")
	clubName := c.GetString("clubname")
	state, err := c.GetInt("state")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	page, err := c.GetInt("page")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	pagesize, err := c.GetInt("pagesize")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if page <= 0 || pagesize <= 0 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "属性不能小于等于0"}
		c.ServeJSON()
		return
	}
	if state != 0 && state != 1 && state != 2 && state != 3 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "无效状态码"}
		c.ServeJSON()
		return
	}
	_, err = models.SuperAdminCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	activitys, total, err := models.FindActivity(page, pagesize, models.Activity{ActivityName: activityName, ClubName: clubName, State: state})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	totalpages := int(math.Ceil(float64(total) / float64(pagesize)))
	c.Data["json"] = map[string]interface{}{"total": total, "page": page, "totalpages": totalpages, "data": activitys}
	c.ServeJSON()
	return
}
func (c *EditActivityController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityName := c.GetString("activityname")
	state, err := c.GetInt("state")
	remarks := c.GetString("remarks")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if activityName == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "活动名不能为空"}
		c.ServeJSON()
		return
	}
	if state != 1 && state != 2 && state != 3 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "无效状态码"}
		c.ServeJSON()
		return
	}
	err = models.EditActivityState(accessToken, models.Activity{ActivityName: activityName, State: state, Remarks: remarks})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改成功"}
	c.ServeJSON()
	return
}
func (c *FindActivityIntroduceController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityName := c.GetString("activityname")
	if activityName == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "活动名不能为空"}
		c.ServeJSON()
		return
	}
	_, err := models.SuperAdminCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	activity, err := models.FindActivityIntroduce(activityName, "")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "introduce": activity.Introduce}
	c.ServeJSON()
	return
}
func (c *FindCommentController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityName := c.GetString("activityname")
	page, err := c.GetInt("page")
	pagesize, err := c.GetInt("pagesize")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if page <= 0 || pagesize <= 0 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "属性不能小于等于0"}
		c.ServeJSON()
		return
	}
	comment := c.GetString("comment")
	comments, total, err := models.SuperAdminFindActivityComment(accessToken, models.ActivityComment{ActivityName: activityName, Comment: comment}, page, pagesize)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	totalpages := int(math.Ceil(float64(total) / float64(pagesize)))
	c.Data["json"] = map[string]interface{}{"total": total, "page": page, "totalpages": totalpages, "data": comments}
	c.ServeJSON()
	return
}
func (c *DelCommentController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	id := c.GetString("id")
	if id == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "id不能为空"}
		c.ServeJSON()
		return
	}
	_, err := models.SuperAdminCheckAccessToken(accessToken)
	if err != nil {
		return
	}
	err = models.DelActivityComment(id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
	return
}
func (c *SuperAdminChangePasswordController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	previousPassword := c.GetString("previousPassword")
	proposedPassword := c.GetString("proposedPassword")
	if previousPassword == "" || proposedPassword == "" {
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
	err := models.SuperAdminChangePassword(accessToken, previousPassword, proposedPassword)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "更改密码成功"}
	c.ServeJSON()
	return
}
func (c *ClubLoGoController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	_,err := models.SuperAdminCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	clubname := c.GetString("clubname")
	f, h, err := c.GetFile("fileName1")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	defer f.Close()
	filename := uuid.New().String()
	c.SaveToFile("fileName1", "static/images/" + filename+h.Filename)
	err = models.SAdminEditClub(models.Club{Logo:filename+h.Filename},clubname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "上传成功"}
		c.ServeJSON()
		return
}