package model

import "time"

// Follow 关注模块
type Follow struct {
	followerID  int       // 关注人ID
	followingID int       // 被关注人ID
	followTime  time.Time // 关注发起时间
	status      int       // 关注状态（1：关注，默认0）
}

// TableName 修改表名映射
func (Follow) TableName() string {
	return "follow"
}