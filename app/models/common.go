package models

import (
	"gorm.io/gorm"
)

// 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey;comment:主键ID"`
}

// 创建、更新时间
// type Timestamps struct {
// 	CreateTime  time.Time `json:"create_time" gorm:"not null; comment:开始时间"`
// 	UpdateTime time.Time `json:"update_time" gorm:"not null;  comment:修改时间"`
// }

// 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// status
type Status struct {
	Status int8 `json:"status" gorm:"not null;default:1;comment:状态：0->无效，1->正常"`
}
