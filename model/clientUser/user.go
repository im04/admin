package clientUser

type User struct {
	Id int `json:"id" xorm:"id int(11)"`
	WxOpenId string `json:"openId" xorm:"wx_open_id varchar(255)"`
	WxNickName string `json:"nickName" xorm:"wx_nick_name varchar(20)"`
	WxCity string `json:"city" xorm:"wx_city varchar(20)"`
	WxProvince string `json:"province" xorm:"wx_province varchar(20)"`
	WxCountry string `json:"country" xorm:"wx_country varchar(20)"`
	WxAvatarPath string `json:"avatarPath" xorm:"wx_avatar_path varchar(1024)"`
	Sex int `json:"sex" xorm:"sex tinyint(1)"`
	CreateTime int `json:"createTime" xorm:"create_time int(11)"`
	UpdateTime int `json:"updateTime" xorm:"update_time int(11)"`
}


