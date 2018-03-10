package adminUser

type AdminNodeLoad struct {
	Id int `xorm:"id"`
	Title string `xorm:"title"`
	Name string `xorm:"name"`
	Status string `xorm:"status"`
}