package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"data-visualization-platform/handlers"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	// API路由组
	api := r.Group("/api")
	{
		// 数据可视化API
		api.GET("/charts", handlers.GetCharts)
		api.GET("/charts/:id", handlers.GetChart)
		api.POST("/charts", handlers.CreateChart)
		api.PUT("/charts/:id", handlers.UpdateChart)
		api.DELETE("/charts/:id", handlers.DeleteChart)

		// 业务数据API
		api.GET("/data", handlers.GetBusinessData)
		api.POST("/data", handlers.CreateBusinessData)
		api.PUT("/data/:id", handlers.UpdateBusinessData)
		api.DELETE("/data/:id", handlers.DeleteBusinessData)

		// 仪表板API
		api.GET("/dashboards", handlers.GetDashboards)
		api.GET("/dashboards/:id", handlers.GetDashboard)
		api.POST("/dashboards", handlers.CreateDashboard)
		api.PUT("/dashboards/:id", handlers.UpdateDashboard)
		api.DELETE("/dashboards/:id", handlers.DeleteDashboard)

		// 数据统计API
		api.GET("/stats/summary", handlers.GetSummaryStats)
		api.GET("/stats/trends", handlers.GetTrendStats)
		api.GET("/stats/categories", handlers.GetCategoryStats)

		// 文件上传API
		api.POST("/upload", handlers.UploadFile)
		api.POST("/upload/csv", handlers.UploadCSV)
	}

	// 页面路由
	r.GET("/", handlers.IndexPage)
	r.GET("/dashboard", handlers.DashboardPage)
	r.GET("/charts", handlers.ChartsPage)
	r.GET("/data", handlers.DataPage)
	r.GET("/settings", handlers.SettingsPage)

	// 静态文件服务
	r.Static("/static", "./static")
	r.Static("/uploads", "./uploads")

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "data-visualization-platform",
			"version": "1.0.0",
		})
	})
}