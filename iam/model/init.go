package model

import (
	"KubeInsight/internal/common"
	"fmt"
)

func InitModel() error {
	db := common.DB.GormClient
	user := User{}
	userRole := UserRole{}
	role := Role{}
	rolePermission := RolePermission{}
	permission := Permission{}
	if err := db.AutoMigrate(&user, &userRole, &role, &rolePermission, &permission); err != nil {
		return fmt.Errorf("Failed Init IAM Model: %v\n", err)
	}
	return nil
}
