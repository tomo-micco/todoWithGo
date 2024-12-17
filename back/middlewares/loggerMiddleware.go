package middlewares

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 処理開始時間
		t := time.Now()
		slog.Info("process start")

		c.Next()

		// 経過時間と送信予定のステータスをログに記載
		elapsed := time.Since(t)
		status := c.Writer.Status()
		slog.Info("process end.", "status", status, "elapsed", elapsed)
	}
}
