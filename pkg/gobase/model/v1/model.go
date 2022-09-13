package v1

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint64         `json:"id,omitempty"        gorm:"primary_key;AUTO_INCREMENT;column:id"`
	CreatedAt time.Time      `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdateAt  time.Time      `json:"updateAt,omitempty"  gorm:"column:updateAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deletedAt;index:idx_deletedAt"`
}
