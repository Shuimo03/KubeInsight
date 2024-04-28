package model

import "time"

type KubeConfig struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"column name"`
	Type      string    `gorm:"column type"`
	Config    string    `gorm:"column config"`
	CreatedAt time.Time `gorm:"column created_at"`
	UpdatedAt time.Time `gorm:"column updated_at"`
}
