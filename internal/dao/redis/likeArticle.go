package redis

import (
	"context"
	"fmt"
)

// 检查用户是否已经点过赞了
func CheckUserLike(articleId int64, userId int64) (bool, error) {
	key := fmt.Sprintf("article:%d:liked_users", articleId)
	// 在 Redis 中检查用户是否已经点赞
	exists, err := ClientRe.SIsMember(context.Background(), key, userId).Result()
	if err != nil {
		return false, err
	}

	return exists, nil
}

// 将用户的点赞信息记录到redis中，
func StoreUserLike(articleId int64, userId int64) error {
	key := fmt.Sprintf("article:%d:liked_users", articleId)

	// 将用户 ID 存储到 Redis 集合中，表示用户已经点赞了该文章
	_, err := ClientRe.SAdd(context.Background(), key, userId).Result()
	return err
}
