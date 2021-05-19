package routers

import (
	"myweb/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)
var FilterUser = func(ctx *context.Context) {
	accessToken,err := ctx.GetSecureCookie("token","accesstoken")
	if accessToken == "" || !err{
		ctx.Redirect(302, "/")
}
	}
func init() {
	beego.InsertFilter("/api/*", beego.BeforeRouter,FilterUser )
	beego.InsertFilter("/superadmin/*", beego.BeforeRouter,FilterUser )
	beego.InsertFilter("/admin/*", beego.BeforeRouter,FilterUser )
	beego.InsertFilter("/user/*", beego.BeforeRouter,FilterUser )
	beego.Router("/", &controllers.LoginController{},"get:Get")
	beego.Router("/login",&controllers.LoginController{},"post:Post")
	beego.Router("/register",&controllers.RegisterController{},"post:Post")
	beego.Router("/confirmregister",&controllers.ConfirmRegisterController{},"post:Post")
	beego.Router("/resendconfirmationcode",&controllers.ResendConfirmationCodeController{},"post:Post")
	beego.Router("/forgotpassword",&controllers.ForgotPasswordController{},"post:Post")
	beego.Router("/forgotpassword",&controllers.ForgotPasswordController{},"get:Get")
	beego.Router("/api/changepassword",&controllers.ChangePasswordController{},"post:Post")
	beego.Router("/api/logout",&controllers.LogoutController{},"post:Post")
	beego.Router("/api/finduser",&controllers.FindUserController{},"post:Post")
	beego.Router("/api/findadmin",&controllers.FindAdminController{},"post:Post")
	beego.Router("/api/createuser",&controllers.CreateUserController{},"post:Post")
	beego.Router("/api/edituserpassword",&controllers.EditUserPasswordController{},"post:Post")
	beego.Router("/api/edituserstate",&controllers.EditUserStateController{},"post:Post")
	beego.Router("/api/deluser",&controllers.DelUserController{},"post:Post")
	beego.Router("/api/createclub",&controllers.CreateClubController{},"post:Post")
	beego.Router("/api/findclub",&controllers.FindClubController{},"post:Post")
	beego.Router("/api/findallclub",&controllers.FindAllClubController{},"post:Post")
	beego.Router("/api/editclubname",&controllers.EditClubNameController{},"post:Post")
	beego.Router("/api/delclub",&controllers.DelClubController{},"post:Post")
	beego.Router("/api/createadmin",&controllers.CreateAdminController{},"post:Post")
	beego.Router("/api/editadmin",&controllers.EditAdminController{},"post:Post")
	beego.Router("/api/deladmin",&controllers.DelAdminController{},"post:Post")
	beego.Router("/api/createresume",&controllers.CreateResumeController{},"post:Post")
	beego.Router("/api/editresume",&controllers.EditResumeController{},"post:Post")
	beego.Router("/superadmin/index",&controllers.SuperAdminIndexController{},"get:Get")
	beego.Router("/superadmin/club",&controllers.SuperAdminClubController{},"get:Get")
	beego.Router("/superadmin/admin",&controllers.SuperAdminController{},"get:Get")
	beego.Router("/superadmin/activity",&controllers.SuperAdminActivityController{},"get:Get")
	beego.Router("/api/findactivity",&controllers.FindActivityController{},"post:Post")
	beego.Router("/api/editactivity",&controllers.EditActivityController{},"post:Post")
	beego.Router("/api/findactivitydetals",&controllers.FindActivityIntroduceController{},"post:Post")
	beego.Router("/superadmin/comment",&controllers.SuperAdminCommentController{},"get:Get")
	beego.Router("/api/sadminfindcomment",&controllers.FindCommentController{},"post:Post")
	beego.Router("/api/sadmindelcomment",&controllers.DelCommentController{},"post:Post")
	beego.Router("/api/sadminchangepassword",&controllers.SuperAdminChangePasswordController{},"post:Post")
	beego.Router("/admin/index",&controllers.AdminIndexController{},"get:Get")
	beego.Router("/admin/notice",&controllers.AdminNoticeController{},"get:Get")
	beego.Router("/admin/activity",&controllers.AdminActivityController{},"get:Get")
	beego.Router("/admin/enroll",&controllers.AdminEnrollController{},"get:Get")
	beego.Router("/admin/enroll1",&controllers.AdminEnroll1Controller{},"get:Get")
	beego.Router("/api/findclubmember",&controllers.FindClubMemberController{},"post:Post")
	beego.Router("/api/createclubmember",&controllers.CreateClubMemberController{},"post:Post")
	beego.Router("/api/editclubmember",&controllers.EditClubMemberController{},"post:Post")
	beego.Router("/api/delclubmember",&controllers.DelClubMemberController{},"post:Post")
	beego.Router("/api/findclubrole",&controllers.FindClubRoleController{},"post:Post")
	beego.Router("/api/delclubrole",&controllers.DelClubRoleController{},"post:Post")
	beego.Router("/api/createclubrole",&controllers.CreateClubRoleController{},"post:Post")
	beego.Router("/api/findclubrole",&controllers.FindClubRoleController{},"post:Post")
	beego.Router("/api/findresume",&controllers.FindResumeController{},"post:Post")
	beego.Router("/api/adminchangepassword",&controllers.AdminChangePasswordController{},"post:Post")
	beego.Router("/api/delclubnotice",&controllers.DelClubNoticeController{},"post:Post")
	beego.Router("/api/createclubnotice",&controllers.CreateClubNoticeController{},"post:Post")
	beego.Router("/api/findclubnotice",&controllers.FindClubNoticeController{},"post:Post")
	beego.Router("/api/editclubnotice",&controllers.EditClubNoticeController{},"post:Post")
	beego.Router("/api/delclubactivity",&controllers.DelClubActivityController{},"post:Post")
	beego.Router("/api/createclubactivity",&controllers.CreateClubActivityController{},"post:Post")
	beego.Router("/api/findclubactivity",&controllers.FindClubActivityController{},"post:Post")
	beego.Router("/api/editclubactivity",&controllers.EditClubActivityController{},"post:Post")
	beego.Router("/api/findclubactivityintroduce",&controllers.FindClubActivityIntroduceController{},"post:Post")
	beego.Router("/api/adminfindclubenroll",&controllers.AdminFindClubEnrollController{},"post:Post")
	beego.Router("/api/adminfindclub",&controllers.AdminFindClubController{},"post:Post")
	beego.Router("/api/adminfindrecruitenroll",&controllers.AdminFindRecruitEnrollController{},"post:Post")
	beego.Router("/api/adminfindchangeenroll",&controllers.AdminFindChangeEnrollController{},"post:Post")
	beego.Router("/api/admineditrecruitenroll",&controllers.AdminEditRecruitEnrollController{},"post:Post")
	beego.Router("/api/admindelrecruitenroll",&controllers.AdminDelRecruitEnrollController{},"post:Post")
	beego.Router("/api/admindelchangeenroll",&controllers.AdminDelChangeEnrollController{},"post:Post")
	beego.Router("/api/admineditclub",&controllers.AdminEditClubController{},"post:Post")
	beego.Router("/api/activityenrolldownloadfile",&controllers.ActivityEnrollDownloadFileController{},"get:Get")
	beego.Router("/api/recruitenrolldownloadfile",&controllers.RecruitEnrollDownloadFileController{},"get:Get")
	beego.Router("/api/changeenrolldownloadfile",&controllers.ChangeEnrollDownloadFileController{},"get:Get")
	beego.Router("/user/index",&controllers.UserIndexController{},"get:Get")
	beego.Router("/user/activity",&controllers.UserActivityController{},"get:Get")
	beego.Router("/user/detail",&controllers.UserActivityDetailController{},"get:Get")
	beego.Router("/user/clubdetail",&controllers.UserClubDetailController{},"get:Get")
	beego.Router("/user/resume",&controllers.UserResumeController{},"get:Get")
	beego.Router("/api/userclubdetail",&controllers.UserClubDetailController{},"post:Post")
	beego.Router("/api/useractivitydetail",&controllers.UserActivityDetailController{},"post:Post")
	beego.Router("/api/userfindactivity",&controllers.UserFindActivityController{},"post:Post")
	beego.Router("/api/userfindclub",&controllers.UserFindClubController{},"post:Post")
	beego.Router("/api/upclublogo",&controllers.ClubLoGoController{},"post:Post")
	beego.Router("/api/userfindactivitycomment",&controllers.UserFindActivityCommentController{},"post:Post")
	beego.Router("/api/usercreateactivitycomment",&controllers.UserCreateActivityCommentController{},"post:Post")
	beego.Router("/api/userdelactivitycomment",&controllers.UserDelActivityCommentController{},"post:Post")
	beego.Router("/api/usereditactivitycomment",&controllers.UserEditActivityCommentController{},"post:Post")
	beego.Router("/api/usercreateactivityenroll",&controllers.UserCreateActivityEnrollController{},"post:Post")
	beego.Router("/api/userfindresume",&controllers.UserFindResumeController{},"post:Post")
	beego.Router("/api/userfindnotices",&controllers.UserFindnoticesController{},"post:Post")
	beego.Router("/api/userresumeportrait",&controllers.ResumePortraitController{},"post:Post")
	beego.Router("/api/usercreateresume",&controllers.UserCreateResumeController{},"get:Get")
	beego.Router("/user/usereditresume",&controllers.UserEditResumeController{},"get:Get")
	beego.Router("/user/myclub",&controllers.MyClubController{},"get:Get")
	beego.Router("/user/changepwd",&controllers.ChangePasswordController{},"get:Get")
	beego.Router("/user/enroll",&controllers.UserEnrollController{},"get:Get")
	beego.Router("/user/clubenroll",&controllers.UserClubEnrollController{},"get:Get")
	beego.Router("/api/userfindclubrecruit",&controllers.UserFindClubRecruitController{},"post:Post")
	beego.Router("/api/userfindclubchange",&controllers.UserFindClubChangeController{},"post:Post")
	beego.Router("/api/createrecruit",&controllers.UserCreateClubRecruitController{},"post:Post")
	beego.Router("/api/createchange",&controllers.UserCreateClubChangeController{},"post:Post")
	beego.Router("/api/userfindrole",&controllers.UserFindRoleController{},"post:Post")
	beego.Router("/api/userfindactivityenroll",&controllers.UserFindActivityEnrollController{},"post:Post")
	beego.Router("/api/userfindrecruitenroll",&controllers.UserFindRecruitEnrollController{},"post:Post")
	beego.Router("/api/userfindchangeenroll",&controllers.UserFindChangeEnrollController{},"post:Post")
	beego.Router("/api/userdelactivityenroll",&controllers.UserDelActivityEnrollController{},"post:Post")
	beego.Router("/api/userdelrecruitenroll",&controllers.UserDelRecruitEnrollController{},"post:Post")
	beego.Router("/api/userdelchangeenroll",&controllers.UserDelChangeEnrollController{},"post:Post")
}