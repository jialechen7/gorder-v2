package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StructuredLog(l *logrus.Entry) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		elapsed := time.Since(t)
		l.WithFields(logrus.Fields{
			"time_elapsed": fmt.Sprintf("%dms", elapsed.Milliseconds()),
			"request_uri":  c.Request.RequestURI,
			"client_ip":    c.ClientIP(),
			"full_path":    c.FullPath(),
		}).Info("request_out")
	}
}
