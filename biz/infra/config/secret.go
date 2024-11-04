package config

import "os"

type SecretKeys struct {
	JWTKey string
}

func GetSecretKey() *SecretKeys {
	jwtKey := os.Getenv("JWT_KEY")
	if len(jwtKey) < 10 {
		panic("环境变量JWT_KEY的长度必须大于10")
	}
	return &SecretKeys{
		JWTKey: jwtKey,
	}
}
