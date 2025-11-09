package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"todolist/config"
	"todolist/pkg/jwtUtil"
	"todolist/pkg/logUtil"
)

// JWT 认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization 头部获取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "请求未携带token，无权限访问，应在Authorization头部添加token",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Token格式不正确, 应为 Bearer [token]",
			})
			c.Abort()
			return
		}

		// parts[1] 是获取到的 tokenString
		userId, err := jwtUtil.ParseToken(parts[1], config.Get().Jwt.SecretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "无效的Token",
			})
			logUtil.Logger.Info("无效的Token: %v", zap.Error(err))
			return
		}

		logUtil.Logger.Info("用户已通过JWT认证, 用户ID: " + userId)
		idUint, err := strconv.ParseUint(userId, 10, 64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "无效的Token",
			})
			return
		}
		c.Set("user_id", uint(idUint))

		c.Next()
	}
}
