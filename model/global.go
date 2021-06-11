package model

import (
	"time"
)

type GdoModel struct {
	ID         uint           `json:"id" gorm:"AUTO_INCREMENT; primary_key"` // 主键ID
	CreatedAt  time.Time      // 创建时间
	UpdatedAt  time.Time      // 更新时间
}