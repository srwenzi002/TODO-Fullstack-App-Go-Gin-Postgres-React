package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApiResponse 统一响应格式
type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// UnifiedResponseMiddleware 统一响应中间件
func UnifiedResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建一个 包装函数，用于捕获响应
		c.Next()

		// 捕获响应状态码
		status := c.Writer.Status()
		if status == 0 {
			status = http.StatusOK
		}

		// 准备响应数据
		response := ApiResponse{
			Status:  status,
			Message: http.StatusText(status),
		}

		// 检查是否存在数据
		if len(c.Errors) > 0 {
			// 如果有错误，使用错误信息
			response.Message = c.Errors.String()
		} else {
			// 否则使用上下文中的数据
			data, _ := c.Get("response_data")
			if data != nil {
				response.Data = data
			} else {
				// 如果没有设置数据，尝试获取原始响应
				data, _ := c.Get("response")
				if data != nil {
					response.Data = data
				}
			}
		}

		// 发送统一响应
		c.JSON(status, response)
	}
}