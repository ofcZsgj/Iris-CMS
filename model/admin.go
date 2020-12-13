package model

import "time"

/**
 * 定义管理员这个实体的结构体
 * 通过Tag中的xorm限定来制定各个结构体字段的类型
 * 使用json来限定在进行JSON数据序列化时定义的json字段
 */
type Admin struct {
	AdminId    int64     `xorm:"pk autoincr" json:"id"` //tag设置为主键，并且拥有自增属性
	AdminName  string    `xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xorm:"DateTime" json:"create_time"`
	Status     int64     `xorm:"default 0" json:"status"`
	Avatar     string    `xorm:"varchar(255)" json:"avatar"`
	Pwd        string    `xorm:"varchar(255)" json:"pwd"`
	CityName   string    `xorm:"varchar(12)" json:"city_name"`
	CityId     int64     `xorm:"index" json:"city_id"` //索引
}
