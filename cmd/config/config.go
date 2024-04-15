package config

import (
	"KubeInsight/iam/model"
	"KubeInsight/internal/common"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func InitSySAdmin() {
	log.Println("Init Sysadmin")
	viper.SetConfigFile("config/iam_sys_admin.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	sysadminUsername := viper.GetString("sysadmin.username")
	sysadminPassword := viper.GetString("sysadmin.password")

	salt, err := bcrypt.GenerateFromPassword([]byte(sysadminPassword), 10)
	user := &model.User{
		Username:  sysadminUsername,
		Password:  string(salt),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err != nil {
		log.Printf("Init SysAdmin: %v \n", err)
	}

	if userError := common.DB.GormClient.Create(user).Error; userError != nil {
		log.Printf("Failed to create sysadmin user: %v \n", userError)
		return
	}

	role := &model.Role{
		Name:        "sysadmin",
		Description: "System Administrator",
	}
	if roleError := common.DB.GormClient.Create(role).Error; roleError != nil {
		log.Printf("Failed to create sysadmin role: %v \n", roleError)
		return
	}
	userRole := &model.UserRole{
		UserID: user.ID,
		RoleID: role.ID,
	}
	if userRoleError := common.DB.GormClient.Create(userRole).Error; userRoleError != nil {
		log.Printf("Failed to associate sysadmin user with role: %v \n", userRoleError)
		return
	}

	log.Println("Sysadmin initialization completed successfully")
}
