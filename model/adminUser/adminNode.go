package adminUser

type AdminNode struct {
	Id int `xorm:"id"`
	Pid int `xorm:"pid"`
	GroupId int `xorm:"group_id"`
	Name string `xorm:"name"`
	Title string `xorm:"title"`
	Remark string `xorm:"remark"`
	Level int `xorm:"level"`
	Type int `xorm:"type"`
	Sort int `xorm:"sort"`
	Status int `xorm:"status"`
	IsDelete int `xorm:"isdelete"`
}