package adminUser

type AdminGroup struct {
	Id int `xorm:"id"`
	Name string `xorm:"name"`
	Icon string `xorm:"icon"`
	Sort int `xorm:"sort"`
	Status int `xorm:"status"`
	Remark string `xorm:"remark"`
	IsDelete int `xorm:"isdelete"`
	CreateTime string `xorm:"create_time"`
	UpdateTime string `xorm:"update_time"`
}
