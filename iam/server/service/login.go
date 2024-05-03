package service

import (
	"KubeInsight/iam/jwt"
	"KubeInsight/iam/model"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func (iam *IamService) Login(username, password string) error {

	var user model.User
	if err := iam.dbHandler.GormClient.Where("username = ?", username).First(&user).Error; err != nil {
		return fmt.Errorf("error finding user: %v", err)
	}

	// 使用 bcrypt 检查密码是否匹配
	if hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); hashErr != nil {
		return fmt.Errorf("error comparing password: %v", hashErr)
	}

	//
	token, err := jwt.GenerateToken(user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("failed Generate to ken: %v", err)
	}
	if tr := iam.cacheHandler.Set(context.TODO(), user.Username, token, 0).Err(); tr != nil {
		return fmt.Errorf("failed Save to token: %v", tr)
	}

	// 登录成功
	return nil
}
