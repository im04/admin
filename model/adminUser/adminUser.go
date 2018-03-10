package adminUser

import (
	. "admin/model/db"
	"admin/utils"
	"errors"
)

type AdminUser struct {
	Id int `json:"id" xorm:"id"`
	Account string `json:"account" xorm:"account"`
	RealName string `json:"realname" xorm:"realname"`
	Password string `json:"password" xorm:"password"`
	LastLoginTime int `xorm:"last_login_time"`
	LastLoginIp string `xorm:"last_login_ip"`
	LoginCount int `xorm:"login_count"`
	Email string `xorm:"email"`
	Mobile string `xorm:"mobile"`
	Remark string `xorm:"remark"`
	Status int `xorm:"status"`
	IsDelete int `xorm:"isdelete"`
	CreateTime int `xorm:"create_time"`
	UpdateTime int `xorm:"update_time"`
	Token string `xorm:"-"`
}

func (u *AdminUser) CheckPassword(password string) error {
	if utils.EncryptPassword(password) == u.Password {
		return nil
	} else {
		return errors.New("密码错误")
	}
}

func (u *AdminUser) Insert() (num int64, err error) {
	return DB.Insert(u)
}

func (u *AdminUser) Update() (num int64, err error) {
	return DB.Update(u)
}


