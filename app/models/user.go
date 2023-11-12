package models

import "strconv"

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"not null;index;comment:用户手机号"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	// Timestamps
	LastTime string `json:"last_time" gorm:"comment:上一次登录时间"`
    Ip       string `json:"ip" gorm:"comment:IP"`
    Status   string `json:"status" gorm:"default:'1';comment:用户状态"`
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
