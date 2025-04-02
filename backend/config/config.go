package config

// Config 结构体用于集中管理应用配置
type Config struct {
	JWTSecret string // JWT 签名密钥（生产环境应通过环境变量注入，不要硬编码）
}

// AppConfig 全局配置实例（演示用，实际建议使用环境变量加载）
var AppConfig = &Config{
	JWTSecret: "2+BqMqctKfKVpP64J96ZprLQrw1iMihQhNGbf+5Urlc=", // 至少32位复杂字符串，建议生产环境使用 openssl rand -base64 32 生成
}
