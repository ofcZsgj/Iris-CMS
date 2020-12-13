package service

import (
	"Iris-CMS/model"
	"github.com/go-xorm/xorm"
)

/**
 * 管理员服务
 * 标准的开发模式将每个实体的提供的功能以接口标准的形式定义,供控制层进行调用。
 * 将数据提供服务模块设计成接口，这样设计的目的是接口定义和具体的功能编程实现了分离
 * 有助于在不同的实现方案之间进行切换，如数据库由MySQL更换为sqlite只需要修改adminService的实现即可
 */
type AdminService interface {
	//通过管理员 用户名+密码 获取管理员实体 如果查询到返回 管理员实体和true
	//否则返回 nil false
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)
}

func NewAdminService(db *xorm.Engine) AdminService {
	return &adminService{
		engine: db,
	}
}

/**
 * 管理员的服务实现结构体
 */
type adminService struct {
	engine *xorm.Engine
}

/*
 * 通过用户名和密码查询管理员
 */
func (ac *adminService) GetByAdminNameAndPassword(username, password string) (model.Admin, bool) {

	var admin model.Admin

	//通过xorm引擎操作MySQL数据库查询传入的admin_name和pwd是否存在，并将存在的管理员实体返回
	ac.engine.Where(" admin_name = ? and pwd = ? ", username, password).Get(&admin)

	return admin, admin.AdminId != 0

}
