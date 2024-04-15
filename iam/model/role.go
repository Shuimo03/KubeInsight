package model

type Role struct {
	ID          uint         `gorm:"primaryKey;autoIncrement"`
	Name        string       // 角色名称
	Description string       // 角色描述
	Users       []User       `gorm:"many2many:user_roles;"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
