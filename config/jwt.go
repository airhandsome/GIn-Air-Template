package config

// JWTConfig 包含了JWT的配置信息
type JWTConfig struct {
	SecretKey string
	Issuer    string
	ExpiresAt int64
}

// JWT变量存储了JWT配置
var JWT JWTConfig

func init() {
	JWT = JWTConfig{
		SecretKey: "your_secret_key", // 使用你的密钥
		Issuer:    "your_issuer",
		ExpiresAt: 3600, // 令牌有效期，例如1小时
	}
}
