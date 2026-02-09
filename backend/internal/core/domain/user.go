package domain

import (
	"time"
)

// User: ข้อมูลเข้าสู่ระบบ
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Character Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 1-to-1 Relation
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

// Character: ข้อมูลเกมของผู้ใช้ (Level, HP, XP)
type Character struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"uniqueIndex"` // Foreign Key
	Level       int     `gorm:"default:1"`
	CurrentXP   int     `gorm:"default:0"`
	MaxXP       int     `gorm:"default:100"` // XP ที่ต้องใช้เพื่อขึ้นเวลถัดไป
	HPCurrent   float64 `gorm:"default:0"`   // เงินที่เหลือ (Budget Left)
	HPMax       float64 `gorm:"default:0"`   // งบประมาณตั้งต้น (Total Budget)
	StreakCount int     `gorm:"default:0"`
	AvatarSkin  string  `gorm:"default:'dragon_egg'"` // skin_id
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
