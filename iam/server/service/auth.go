package service

import (
	"KubeInsight/iam/jwt"
	"context"
	"fmt"
)

func (iam *IamService) Auth(userName string) error {
	token, err := iam.cacheHandler.Get(context.TODO(), userName).Result()
	if err != nil {
		return fmt.Errorf("failed get Token: %v", err)
	}
	if tokenError := jwt.ParseToken(token); tokenError != nil {
		return fmt.Errorf("failed parse token: %v", tokenError)
	}
	return nil
}
