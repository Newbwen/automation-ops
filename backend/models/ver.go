package models

import "time"

// CaptchaMeta 验证码元数据
type CaptchaMeta struct {
	Value     string    `json:"value"`
	ExpiresAt time.Time `json:"expires_at"`
}
