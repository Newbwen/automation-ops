// middlewares/jwt.go
package utils

import (
	"github.com/Newbwen/automation-ops/backend/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

// JWTAuth JWT 认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization 头获取令牌（格式：Bearer <token>）
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "缺少认证令牌"})
			return
		}

		// 提取令牌部分（去除 Bearer 前缀）
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // 未找到 Bearer 前缀
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "令牌格式错误"})
			return
		}

		// 解析并验证令牌
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(config.AppConfig.JWTSecret), nil
		})

		// 处理验证错误
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效或过期的令牌"})
			return
		}

		// 提取声明信息并存入上下文
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("id", claims["id"]) // 用户ID将在后续处理中使用
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的令牌声明"})
			return
		}

		c.Next() // 继续处理请求
	}
}
