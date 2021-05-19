package controllers

import (
	"math"
	"myweb/models"
	"os"

	"github.com/astaxie/beego"
	"github.com/google/uuid"
)

type AdminIndexController struct {
	beego.Controller
}
type AdminNoticeController struct {
	beego.Controller
}
type AdminActivityController struct {
	beego.Controller
}
type CreateActivityController struct {
	beego.Controller
}
type FindClubMemberController struct {
	beego.Controller
}
type CreateClubMemberController struct {
	beego.Controller
}
type EditClubMemberController struct {
	beego.Controller
}
type DelClubMemberController struct {
	beego.Controller
}
type CreateClubRoleController struct {
	beego.Controller
}
type DelClubRoleController struct {
	beego.Controller
}
type FindClubRoleController struct {
	beego.Controller
}
type FindResumeController struct {
	beego.Controller
}
type AdminChangePasswordController struct {
	beego.Controller
}
type CreateClubNoticeController struct {
	beego.Controller
}
type DelClubNoticeController struct {
	beego.Controller
}
type EditClubNoticeController struct {
	beego.Controller
}
type FindClubNoticeController struct {
	beego.Controller
}
type CreateClubActivityController struct {
	beego.Controller
}
type DelClubActivityController struct {
	beego.Controller
}
type EditClubActivityController struct {
	beego.Controller
}
type FindClubActivityController struct {
	beego.Controller
}
type FindClubActivityIntroduceController struct {
	beego.Controller
}
type AdminFindClubEnrollController struct {
	beego.Controller
}
type AdminEditClubController struct {
	beego.Controller
}
type AdminEnrollController struct {
	beego.Controller
}
type AdminEnroll1Controller struct {
	beego.Controller
}
type AdminFindClubController struct {
	beego.Controller
}
type AdminFindRecruitEnrollController struct {
	beego.Controller
}
type AdminEditRecruitEnrollController struct {
	beego.Controller
}
type AdminDelRecruitEnrollController struct {
	beego.Controller
}
type AdminFindChangeEnrollController struct {
	beego.Controller
}
type AdminDelChangeEnrollController struct {
	beego.Controller
}
type ActivityEnrollDownloadFileController struct {
	beego.Controller
}
type RecruitEnrollDownloadFileController struct {
	beego.Controller
}
type ChangeEnrollDownloadFileController struct {
	beego.Controller
}


func (c *AdminIndexController) Get() {
	c.TplName = "adminindex.html"
}
func (c *AdminEnrollController) Get() {
	c.TplName = "adminenroll.html"
}
func (c *AdminNoticeController) Get() {
	c.TplName = "adminnotice.html"
}
func (c *AdminActivityController) Get() {
	c.TplName = "adminactivity.html"
}
func (c *AdminEnroll1Controller) Get() {
	c.TplName = "adminenroll1.html"
}
func (c *ActivityEnrollDownloadFileController) Get() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityname := c.GetString("activityname")
	filename, err := models.ActivityEnrollExportExcel(accessToken, activityname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Ctx.Output.Download(filename, filename)
	_ = os.Remove(filename)
}
func (c * RecruitEnrollDownloadFileController) Get() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	filename, err := models.RecruitEnrollExportExcel(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Ctx.Output.Download(filename, filename)
	_ = os.Remove(filename)
}
func (c * ChangeEnrollDownloadFileController) Get() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	filename, err := models.ChangeEnrollExportExcel(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Ctx.Output.Download(filename, filename)
	_ = os.Remove(filename)
}
func (c *AdminEditClubController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	introduce := c.GetString("introduce")
	recruit, err := c.GetInt("recruit")
	change, err := c.GetInt("change")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if recruit != 0 && recruit != 1 && recruit != 2 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "状态无效"}
		c.ServeJSON()
		return
	}
	if change != 0 && change != 1 && change != 2 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "状态无效"}
		c.ServeJSON()
		return
	}
	err = models.AdminEditClub(accessToken, models.Club{Introduce: introduce, Recruit: recruit, Change: change})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改成功"}
	c.ServeJSON()
	return
}
func (c *FindClubMemberController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	name := c.GetString("name")
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
	clubmembers, total, err := models.FindClubMember(accessToken, page, pagesize, name)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	totalpages := int(math.Ceil(float64(total) / float64(pagesize)))
	c.Data["json"] = map[string]interface{}{"total": total, "page": page, "totalpages": totalpages, "data": clubmembers}
	c.ServeJSON()
	return
}
func (c *CreateClubMemberController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	role := c.GetString("role")
	if username == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户不能为空"}
		c.ServeJSON()
		return
	}
	err := models.CreateClubMember(accessToken, username, role)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建成功"}
	c.ServeJSON()
}
func (c *EditClubMemberController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	role := c.GetString("role")
	if username == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户不能为空"}
		c.ServeJSON()
		return
	}
	err := models.EditClubMember(accessToken, models.ClubMember{Username: username, Role: role})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改成功"}
	c.ServeJSON()
}
func (c *DelClubMemberController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	if username == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户不能为空"}
		c.ServeJSON()
		return
	}
	err := models.DelClubMember(accessToken, username)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
}
func (c *CreateClubRoleController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	role := c.GetString("role")
	if role == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "角色不能为空"}
		c.ServeJSON()
		return
	}
	err := models.CreateClubRole(accessToken, role)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建成功"}
	c.ServeJSON()
}
func (c *DelClubRoleController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	role := c.GetString("role")
	if role == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "角色不能为空"}
		c.ServeJSON()
		return
	}
	err := models.DelClubRole(accessToken, role)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
}
func (c *FindClubRoleController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	admin, err := models.AdminCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	clubRoles, err := models.FindClubRole(admin.ClubName)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = clubRoles
	c.ServeJSON()
	return
}
func (c *FindResumeController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	_,err := models.AdminCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	resume, err := models.FindResume(username)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = resume
	c.ServeJSON()
	return
}
func (c *AdminChangePasswordController) Post() {
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
	err := models.AdminChangePassword(accessToken, previousPassword, proposedPassword)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "更改密码成功"}
	c.ServeJSON()
	return
}
func (c *FindClubNoticeController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
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
	clubnotices, total, err := models.FindClubNotice(accessToken, page, pagesize)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	totalpages := int(math.Ceil(float64(total) / float64(pagesize)))
	c.Data["json"] = map[string]interface{}{"total": total, "page": page, "totalpages": totalpages, "data": clubnotices}
	c.ServeJSON()
	return
}
func (c *CreateClubNoticeController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	notice := c.GetString("notice")
	if notice == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "通知不能为空"}
		c.ServeJSON()
		return
	}
	err := models.CreateClubNotice(accessToken, models.ClubNotice{Notice: notice})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建成功"}

	c.ServeJSON()
}
func (c *EditClubNoticeController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	id := c.GetString("id")
	notice := c.GetString("notice")
	if notice == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "通知不能为空"}
		c.ServeJSON()
		return
	}
	err := models.EditClubNotice(accessToken, notice, id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改成功"}
	c.ServeJSON()
}
func (c *DelClubNoticeController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	id := c.GetString("id")
	err := models.DelClubNotice(accessToken, id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
}
func (c *FindClubActivityController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
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
	admin, err := models.AdminCheckAccessToken(accessToken)
	clubactivitys, total, err := models.FindActivity(page, pagesize, models.Activity{ClubName: admin.ClubName})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	totalpages := int(math.Ceil(float64(total) / float64(pagesize)))
	c.Data["json"] = map[string]interface{}{"total": total, "page": page, "totalpages": totalpages, "data": clubactivitys}
	c.ServeJSON()
	return
}
func (c *EditClubActivityController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityname := c.GetString("activityname")
	introduce := c.GetString("introduce")
	admin, err := models.AdminCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	state, err := c.GetInt("state")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if state != 0 && state != 1 && state != 4 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "状态码无效"}
		c.ServeJSON()
		return
	}
	err = models.EditActivity(admin.ClubName, models.Activity{ActivityName: activityname, State: state, Introduce: introduce})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改成功"}
	c.ServeJSON()
	return
}
func (c *CreateClubActivityController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityname := c.GetString("activityname")
	introduce := c.GetString("introduce")
	if activityname == "" || introduce == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "名字或者介绍不能为空"}
		c.ServeJSON()
		return
	}
	f, h, err := c.GetFile("fileName1")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	defer f.Close()
	filename := uuid.New().String()
	c.SaveToFile("fileName1", "static/images/" + filename+h.Filename)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	err = models.CreateActivity(accessToken, models.Activity{ActivityName: activityname, Introduce: introduce,Cover:filename+h.Filename})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建成功"}

	c.ServeJSON()
}
func (c *DelClubActivityController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityname := c.GetString("activityname")
	if activityname == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "名字不能为空"}
		c.ServeJSON()
		return
	}
	err := models.DelActivity(accessToken, models.Activity{ActivityName: activityname})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}

	c.ServeJSON()
}
func (c *FindClubActivityIntroduceController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityName := c.GetString("activityname")
	if activityName == "" {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "活动名不能为空"}
		c.ServeJSON()
		return
	}
	admin, err := models.AdminCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	activity, err := models.FindActivityIntroduce(activityName, admin.ClubName)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "introduce": activity.Introduce}
	c.ServeJSON()
	return
}
func (c *AdminFindClubEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityName := c.GetString("activityname")
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
	clubenrolls, total, err := models.AdminFindActivityEnroll(accessToken, activityName, page, pagesize)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	totalpages := int(math.Ceil(float64(total) / float64(pagesize)))
	c.Data["json"] = map[string]interface{}{"total": total, "page": page, "totalpages": totalpages, "data": clubenrolls}
	c.ServeJSON()
	return
}
func (c *AdminFindClubController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	club, err := models.AdminFindClub(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = club
	c.ServeJSON()
	return
}
func (c *AdminFindRecruitEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	state, err := c.GetInt("state")
	studentNumber := c.GetString("studentnumber")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if state != 0 && state != 1 && state != 2 && state != 3 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "状态无效"}
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
	results, total, err := models.FindRecruitEnroll(accessToken, page, pagesize, state,studentNumber)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	totalpages := int(math.Ceil(float64(total) / float64(pagesize)))
	c.Data["json"] = map[string]interface{}{"total": total, "page": page, "totalpages": totalpages, "data": results}
	c.ServeJSON()
	return
}
func (c *AdminFindChangeEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
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
	results, total, err := models.FindChangeEnroll(accessToken, page, pagesize)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	totalpages := int(math.Ceil(float64(total) / float64(pagesize)))
	c.Data["json"] = map[string]interface{}{"total": total, "page": page, "totalpages": totalpages, "data": results}
	c.ServeJSON()
	return
}
func (c *AdminDelRecruitEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	err := models.DelRecruitEnroll(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
}
func (c *AdminDelChangeEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	err := models.DelChangeEnroll(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
}
func (c *AdminEditRecruitEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	username := c.GetString("username")
	state, err := c.GetInt("state")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	if state != 0 && state != 1 && state != 2 && state != 3 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "状态无效"}
		c.ServeJSON()
		return
	}
	err = models.EditRecruitEnroll(accessToken, username, state)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "更新成功"}
	c.ServeJSON()
}
