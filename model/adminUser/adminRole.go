package adminUser

type AdminRole struct {
	Id int `xorm:"id"`
	Pid int `xorm:"pid"`
	Name string `xorm:"name"`
	Remark string `xorm:"remark"`
	Status string `xorm:"status"`
	IsDelete int `xorm:"isdelete"`
	CreateTime int `xorm:"create_time"`
	UpdateTime int `xorm:"update_time"`
}