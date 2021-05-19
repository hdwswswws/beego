package controllers

import (
	"myweb/models"

	"github.com/astaxie/beego"
	"github.com/google/uuid"
)
type CreateResumeController struct {
	beego.Controller
}
type EditResumeController struct {
	beego.Controller
}
type UserIndexController struct {
	beego.Controller
}
type UserFindActivityController struct {
	beego.Controller
}
type UserActivityController struct {
	beego.Controller
}
type UserFindClubController struct {
	beego.Controller
}
type UserFindResumeController struct {
	beego.Controller
}
type UserCreateResumeController struct {
	beego.Controller
}
type UserActivityDetailController struct {
	beego.Controller
}
type UserClubDetailController struct {
	beego.Controller
}
type MyClubController struct {
	beego.Controller
}
type UserClubEnrollController struct {
	beego.Controller
}
type UserResumeController struct {
	beego.Controller
}
type UserEditResumeController struct {
	beego.Controller
}
type UserFindActivityCommentController struct {
	beego.Controller
}
type UserCreateActivityCommentController struct {
	beego.Controller
}
type UserDelActivityCommentController struct {
	beego.Controller
}
type UserEditActivityCommentController struct {
	beego.Controller
}
type UserCreateActivityEnrollController struct {
	beego.Controller
}
type UserFindnoticesController struct {
	beego.Controller
}
type UserFindClubRecruitController struct {
	beego.Controller
}
type UserFindClubChangeController struct {
	beego.Controller
}
type UserCreateClubChangeController struct {
	beego.Controller
}
type UserCreateClubRecruitController struct {
	beego.Controller
}
type UserFindRoleController struct {
	beego.Controller
}
type UserEnrollController struct {
	beego.Controller
}
type UserFindRecruitEnrollController struct {
	beego.Controller
}
type UserFindActivityEnrollController struct {
	beego.Controller
}
type UserFindChangeEnrollController struct {
	beego.Controller
}
type UserDelRecruitEnrollController struct {
	beego.Controller
}
type UserDelChangeEnrollController struct {
	beego.Controller
}
type UserDelActivityEnrollController struct {
	beego.Controller
}
type ResumePortraitController struct {
	beego.Controller
}
func (c *UserActivityDetailController) Get() {
	activityname := c.GetString("activityname")
	c.Ctx.SetCookie("activityname",activityname,0)
	c.TplName = "useractivitydetail.html"
}
func (c *UserClubDetailController) Get() {
	clubname := c.GetString("clubname")
	c.Ctx.SetCookie("clubname",clubname,0)
	c.TplName = "userclubdetail.html"
}
func (c *UserEnrollController) Get() {
	c.TplName = "userenroll.html"
}
func (c *UserIndexController) Get() {
	c.TplName = "userindex.html"
}
func (c *UserActivityController) Get() {
	c.TplName = "useractivity.html"
}
func (c *MyClubController) Get() {
	c.TplName = "myclub.html"
}
func (c *UserClubEnrollController) Get() {
	c.TplName = "clubenroll.html"
}
func (c *UserResumeController) Get() {
	c.TplName = "userresume.html"
}
func (c *UserEditResumeController) Get() {
	c.TplName = "usereditresume.html"
}
func (c *CreateResumeController) Post() {
	accessToken,_:= c.GetSecureCookie("token","accesstoken")
	var resume models.Resume
	resume.Name = c.GetString("name")
	resume.Sex = c.GetString("sex")
	resume.Birthday = c.GetString("birthday")
	resume.PhoneNumber = c.GetString("phone_number ")
	resume.PoliticalOutlook = c.GetString("political_outlook ")
	resume.StudentNumber = c.GetString("student_number ")
	resume.Major = c.GetString("major ")
	resume.Introduction = c.GetString("introduction")
	err := models.CreateResume(accessToken, resume)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建成功"}
	c.ServeJSON()
}
func (c *EditResumeController) Post() {
	accessToken,_:= c.GetSecureCookie("token","accesstoken")
	var resume models.Resume
	resume.Name = c.GetString("name")
	resume.Sex = c.GetString("sex")
	resume.Birthday = c.GetString("birthday")
	resume.PhoneNumber = c.GetString("phone_number ")
	resume.PoliticalOutlook = c.GetString("political_outlook ")
	resume.StudentNumber = c.GetString("student_number ")
	resume.Major = c.GetString("major ")
	resume.Introduction = c.GetString("introduction")
	err := models.EditResume(accessToken, resume)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建成功"}
	c.ServeJSON()
}
func (c *ResumePortraitController) Post() {
	accessToken,_:= c.GetSecureCookie("token","accesstoken")
	f, h, err := c.GetFile("fileName1")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	defer f.Close()
	filename := uuid.New().String()
	c.SaveToFile("fileName1", "static/images/" + filename+h.Filename)
	var resume models.Resume
	resume.Portrait = filename+h.Filename
	err = models.EditResume(accessToken, resume)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建成功"}
	c.ServeJSON()

}
func (c *UserFindActivityController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	_, err := models.UserCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	activitys, total, err := models.UserFindActivity()
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"total": total, "data": activitys}
	c.ServeJSON()
	return
}
func (c *UserFindClubController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubs, total, err := models.UserFindClub(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"total": total, "data": clubs}
	c.ServeJSON()
	return
}
func (c * UserActivityDetailController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityname := c.GetString("activityname")
	_, err := models.UserCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	activity, err := models.UserActivityDetail(activityname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = activity
	c.ServeJSON()
	return
}
func (c * UserFindActivityCommentController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityname := c.GetString("activityname")
	_, err := models.UserCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	activitycomment, err := models.UserFindActivityComment(activityname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = activitycomment
	c.ServeJSON()
	return
}
func (c * UserCreateActivityCommentController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityname := c.GetString("activityname")
	comment := c.GetString("comment")
	err := models.CreateActivityComment(accessToken,models.ActivityComment{ActivityName:activityname,Comment:comment})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "评论成功"}
	c.ServeJSON()
	return
}
func (c * UserDelActivityCommentController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	id := c.GetString("id")
	user, err := models.UserCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	err = models.UserDelActivityComment(user.Username,id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
	return
}
func (c * UserEditActivityCommentController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	id := c.GetString("id")
	_, err := models.UserCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	err = models.UserEditActivityComment(id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "举报成功"}
	c.ServeJSON()
	return
}
func (c * UserClubDetailController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubname := c.GetString("clubname")
	clubs, err := models.UserClubDetail(accessToken,clubname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = clubs
	c.ServeJSON()
	return
}
func (c * UserCreateActivityEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityname := c.GetString("activityname")
	err := models.CreateActivityEnroll(accessToken,models.ActivityEnroll{ActivityName:activityname})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "加入成功"}
	c.ServeJSON()
	return
}
func (c * UserFindResumeController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	user, err := models.UserCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	resume,err := models.FindResume(user.Username)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = resume
	c.ServeJSON()
	return
}
func (c * UserCreateResumeController) Get() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	name := c.GetString("name")
	studentname := c.GetString("studentname")
	major := c.GetString("major")
	sex := c.GetString("sex")
	phone_number := c.GetString("phone_number")
	email := c.GetString("email")
	birthday := c.GetString("birthday")
	political_outlook := c.GetString("political_outlook")
	introduce := c.GetString("introduce")
	err := models.CreateResume(accessToken,models.Resume{Name:name,StudentNumber:studentname,Major:major,Sex:sex,PhoneNumber:phone_number,Email:email,Birthday:birthday,PoliticalOutlook:political_outlook,Introduction:introduce})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.TplName = "userresume.html"
	return
}
func (c * UserFindnoticesController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	user, err := models.UserCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	clubs,notices,err := models.UserFindClubNotice(user.Username)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] =  map[string]interface{}{"club": clubs, "notices": notices}
	c.ServeJSON()
	return
}
func (c *UserFindClubRecruitController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubs, total, err := models.UserFindClubRecruit(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"total": total, "data": clubs}
	c.ServeJSON()
	return
}
func (c *UserFindClubChangeController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubs, total, err := models.UserFindClubChange(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"total": total, "data": clubs}
	c.ServeJSON()
	return
}
func (c *UserCreateClubChangeController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubname := c.GetString("clubname")
	role := c.GetString("role")
	err := models.CreateChangeEnroll(accessToken,clubname,role)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "加入成功"}
	c.ServeJSON()
	return
}
func (c *UserCreateClubRecruitController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubname := c.GetString("clubname")
	err := models.CreateRecruitEnroll(accessToken,clubname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "加入成功"}
	c.ServeJSON()
	return
}
func (c *UserFindRoleController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubname := c.GetString("clubname")
	_, err := models.UserCheckAccessToken(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	role,err := models.FindClubRole(clubname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = role
	c.ServeJSON()
	return
}
func (c *UserFindRecruitEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	enroll,err := models.UserFindRecruitEnroll(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = enroll
	c.ServeJSON()
	return
}
func (c *UserFindChangeEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	enroll,err := models.UserFindChangeEnroll(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = enroll
	c.ServeJSON()
	return
}
func (c *UserFindActivityEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	enroll,err := models.UserFindActivityEnroll(accessToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = enroll
	c.ServeJSON()
	return
}
func (c *UserDelActivityEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	activityname := c.GetString("activityname")
	err := models.UserDelActivityEnroll(accessToken,activityname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
	return
}
func (c *UserDelRecruitEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubname := c.GetString("clubname")
	err := models.UserDelRecruitEnroll(accessToken,clubname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
	return
}
func (c *UserDelChangeEnrollController) Post() {
	accessToken, _ := c.GetSecureCookie("token", "accesstoken")
	clubname := c.GetString("clubname")
	err := models.UserDelChangeEnroll(accessToken,clubname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除成功"}
	c.ServeJSON()
	return
}







