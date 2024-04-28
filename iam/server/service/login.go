package service

import (
	"KubeInsight/iam/jwt"
	"KubeInsight/iam/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func (iam *IamService) Login(username, password string) (string, error) {
	// 查询用户名
	var user model.User
	if err := iam.dbHandler.GormClient.Where("username = ?", username).First(&user).Error; err != nil {
		// 其他错误
		return "", fmt.Errorf("error finding user: %v", err)
	}

	token, tokenError := jwt.GenerateToken(user.ID)
	if tokenError != nil {
		return "", fmt.Errorf("Failed GenerateToken: %v", tokenError)
	}

	// 使用 bcrypt 检查密码是否匹配
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {

		// 其他错误
		return "", fmt.Errorf("error comparing password: %v", err)
	}

	// 登录成功
	return token, nil

}
