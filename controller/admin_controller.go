package controller

import (
	"Iris-CMS/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

/**
 * 管理员控制器
 * 控制器负责来完成我们请求的逻辑流程控制
 */
type AdminController struct {
	//iris框架自动为每个请求都绑定上下文对象
	Ctx iris.Context

	//admin功能实体,通过AdminSerevice实体查询数据库
	Service service.AdminService

	//session对象
	Session *sessions.Session
}

const (
	ADMINTABLENAME = "admin"
	ADMIN          = "adminId"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

/**
 * 管理员登陆功能
 * 接口：/admin/login
 */
func (ac *AdminController) PostLogin(ctx iris.Context) mvc.Result {

	iris.New().Logger().Info(" admin login ")

	//将post提交的json数据使用adminLogin存储
	var adminLogin AdminLogin
	ac.Ctx.ReadJSON(&adminLogin)

	//数据参数检验
	if adminLogin.UserName == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":   "0",
				"succeess": "登陆失败",
				"message":  "用户名或者密码为空，请重新登陆",
			},
		}
	}

	//根据用户名，密码到数据库中查询对应的管理信息
	admin, exist := ac.Service.GetByAdminNameAndPassword(adminLogin.UserName, adminLogin.Password)

	//管理员不存在
	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登陆失败",
				"message": "用户名或密码错误，请重新登陆",
			},
		}
	}

	//管理员存在 设置session
	//userByte := admin.Encoder()
	ac.Session.Set(ADMIN, admin.AdminId)

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  "1",
			"success": "登陆成功",
			"message": "管理员登陆成功",
		},
	}

}
