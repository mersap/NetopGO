package routers

import (
	"NetopGO/controllers"
	"NetopGO/models"
	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/user/list", &controllers.UserController{})
	//beego.AutoRouter(&controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "get:Add")
	beego.Router("/user/add", &controllers.UserController{}, "post:Post")
	beego.Router("/user/modify", &controllers.UserController{}, "post:Post")
	beego.Router("/user/del", &controllers.UserController{}, "get:Delete")
	beego.Router("/user/search", &controllers.UserController{}, "get:Search")
	beego.Router("/user/detail", &controllers.UserController{}, "get:Detail")
	beego.Router("/user/bitchDel", &controllers.UserController{}, "post:BitchDelete")
	beego.Router("/user/reset_password", &controllers.UserController{}, "get:ResetPasswd")
	beego.Router("/user/reset_password", &controllers.UserController{}, "post:ResetPasswd")

	beego.Router("/group/list", &controllers.GroupController{})
	beego.Router("/group/add", &controllers.GroupController{}, "get:Add")
	beego.Router("/group/add", &controllers.GroupController{}, "post:Post")
	beego.Router("/group/modify", &controllers.GroupController{}, "post:Post")
	beego.Router("/group/del", &controllers.GroupController{}, "get:Delete")
	beego.Router("/group/bitchDel", &controllers.GroupController{}, "post:BitchDelete")
	beego.Router("/group/search", &controllers.GroupController{}, "get:Search")

	beego.Router("/host/list", &controllers.HostController{})
	beego.Router("/host/add", &controllers.HostController{}, "get:Add")
	beego.Router("/host/add", &controllers.HostController{}, "post:Post")
	beego.Router("/host/modify", &controllers.HostController{}, "post:Post")
	beego.Router("/host/del", &controllers.HostController{}, "get:Delete")
	beego.Router("/host/bitchDel", &controllers.HostController{}, "post:BitchDelete")
	beego.Router("/host/search", &controllers.HostController{}, "get:Search")
	beego.Router("/host/webconsole", &controllers.HostController{}, "get:WebConsole")
	beego.Handler("/console/sshws", websocket.Handler(models.SSHWebSocketHandler))

	beego.Router("/schema/list", &controllers.SchemaController{})
	beego.Router("/schema/add", &controllers.SchemaController{}, "get:Add")
	beego.Router("/schema/add", &controllers.SchemaController{}, "post:Post")
	beego.Router("/schema/modify", &controllers.SchemaController{}, "post:Post")
	beego.Router("/schema/del", &controllers.SchemaController{}, "get:Delete")
	beego.Router("/schema/bitchDel", &controllers.SchemaController{}, "post:BitchDelete")
	beego.Router("/schema/partition", &controllers.SchemaController{}, "get:Partition")

	beego.Router("/db/list", &controllers.DBController{})
	beego.Router("/db/add", &controllers.DBController{}, "get:Add")
	beego.Router("/db/add", &controllers.DBController{}, "post:Post")
	beego.Router("/db/modify", &controllers.DBController{}, "post:Post")
	beego.Router("/db/del", &controllers.DBController{}, "get:Delete")
	beego.Router("/db/bitchDel", &controllers.DBController{}, "post:BitchDelete")
	beego.Router("/db/search", &controllers.DBController{}, "get:Search")
	beego.Router("/db/query", &controllers.DBController{}, "get:Query")
	beego.Router("/db/detail", &controllers.DBController{}, "get:Detail")
	beego.Router("/db/slowlog", &controllers.DBController{}, "get:SlowLog")
	beego.Router("/db/explain", &controllers.DBController{}, "get:Explain")
	beego.Router("/db/sqltext", &controllers.DBController{}, "get:Sqltext")
	beego.Router("/db/query/export", &controllers.DBController{}, "get:Export")

	beego.Router("/record/db/list", &controllers.RecordController{})
	beego.Router("/record/db/add", &controllers.RecordController{}, "get:Add")
	beego.Router("/record/db/add", &controllers.RecordController{}, "post:Post")
	beego.Router("/record/db/del", &controllers.RecordController{}, "get:Delete")
	beego.Router("/record/db/bitchDel", &controllers.RecordController{}, "post:BitchDelete")
	beego.Router("/record/db/detail", &controllers.RecordController{}, "get:Post")
	beego.Router("/record/db/search", &controllers.RecordController{}, "get:Search")

	beego.Router("/record/app/list", &controllers.AppRecordController{})
	beego.Router("/record/app/add", &controllers.AppRecordController{}, "get:Add")
	beego.Router("/record/app/add", &controllers.AppRecordController{}, "post:Post")
	beego.Router("/record/app/del", &controllers.AppRecordController{}, "get:Delete")
	beego.Router("/record/app/bitchDel", &controllers.AppRecordController{}, "post:BitchDelete")
	beego.Router("/record/app/detail", &controllers.AppRecordController{}, "get:Post")
	beego.Router("/record/app/search", &controllers.AppRecordController{}, "get:Search")

	beego.Router("/record/fault/list", &controllers.FaultRecordController{})
	beego.Router("/record/fault/add", &controllers.FaultRecordController{}, "get:Add")
	beego.Router("/record/fault/add", &controllers.FaultRecordController{}, "post:Post")
	beego.Router("/record/fault/del", &controllers.FaultRecordController{}, "get:Delete")
	beego.Router("/record/fault/bitchDel", &controllers.FaultRecordController{}, "post:BitchDelete")
	beego.Router("/record/fault/detail", &controllers.FaultRecordController{}, "get:Post")
	beego.Router("/record/fault/search", &controllers.FaultRecordController{}, "get:Search")

	beego.Router("/audit/list", &controllers.AuditController{})
	beego.Router("/audit/del", &controllers.AuditController{}, "get:Delete")
	beego.Router("/audit/bitchDel", &controllers.AuditController{}, "post:BitchDelete")
	beego.Router("/audit/detail", &controllers.AuditController{}, "get:Detail")
	beego.Router("/audit/search", &controllers.AuditController{}, "get:Search")

	beego.Router("/workorder/app", &controllers.AppWOController{}, "get:AppOrder")
	beego.Router("/workorder/app", &controllers.AppWOController{}, "post:AppOrderPost")
	beego.Router("/workorder/my", &controllers.AppWOController{}, "get:Get")
	beego.Router("/workorder/approve", &controllers.AppWOController{}, "get:Approve")

	beego.Router("/attachment/:all", &controllers.AttachController{})
}
