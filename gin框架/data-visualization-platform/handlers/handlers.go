package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"data-visualization-platform/models"
)

// GetCharts 获取图表列表
func GetCharts(c *gin.Context) {
	sqlDB, err := db.GetConnection("default")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败"})
		return
	}

	rows, err := sqlDB.Query("SELECT * FROM chart_configs WHERE status = 1 ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	defer rows.Close()

	var charts []models.ChartConfig
	for rows.Next() {
		var chart models.ChartConfig
		err := rows.Scan(&chart.ID, &chart.Name, &chart.Type, &chart.Config, &chart.DataSource,
			&chart.Description, &chart.Status, &chart.CreatedAt, &chart.UpdatedAt)
		if err != nil {
			continue
		}
		charts = append(charts, chart)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    charts,
	})
}

// GetChart 获取单个图表
func GetChart(c *gin.Context) {
	id := c.Param("id")
	sqlDB, err := db.GetConnection("default")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败"})
		return
	}

	var chart models.ChartConfig
	err = sqlDB.QueryRow("SELECT * FROM chart_configs WHERE id = ?", id).Scan(
		&chart.ID, &chart.Name, &chart.Type, &chart.Config, &chart.DataSource,
		&chart.Description, &chart.Status, &chart.CreatedAt, &chart.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "图表不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    chart,
	})
}

// CreateChart 创建图表
func CreateChart(c *gin.Context) {
	var chart models.ChartConfig
	if err := c.ShouldBindJSON(&chart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	sqlDB, err := db.GetConnection("default")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败"})
		return
	}

	result, err := sqlDB.Exec("INSERT INTO chart_configs (name, type, config, data_source, description, status) VALUES (?, ?, ?, ?, ?, ?)",
		chart.Name, chart.Type, chart.Config, chart.DataSource, chart.Description, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}

	id, _ := result.LastInsertId()
	chart.ID = id

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    chart,
	})
}

// GetBusinessData 获取业务数据
func GetBusinessData(c *gin.Context) {
	category := c.Query("category")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	sqlDB, err := db.GetConnection("default")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败"})
		return
	}

	query := "SELECT * FROM business_data WHERE 1=1"
	args := []interface{}{}

	if category != "" {
		query += " AND category = ?"
		args = append(args, category)
	}

	if startDate != "" {
		query += " AND date >= ?"
		args = append(args, startDate)
	}

	if endDate != "" {
		query += " AND date <= ?"
		args = append(args, endDate)
	}

	query += " ORDER BY date DESC"

	rows, err := sqlDB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	defer rows.Close()

	var data []models.BusinessData
	for rows.Next() {
		var item models.BusinessData
		err := rows.Scan(&item.ID, &item.Category, &item.Value, &item.Date, &item.Description,
			&item.Tags, &item.CreatedBy, &item.CreatedAt)
		if err != nil {
			continue
		}
		data = append(data, item)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

// CreateBusinessData 创建业务数据
func CreateBusinessData(c *gin.Context) {
	var data models.BusinessData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	sqlDB, err := db.GetConnection("default")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败"})
		return
	}

	result, err := sqlDB.Exec("INSERT INTO business_data (category, value, date, description, tags, created_by) VALUES (?, ?, ?, ?, ?, ?)",
		data.Category, data.Value, data.Date, data.Description, data.Tags, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}

	id, _ := result.LastInsertId()
	data.ID = id

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    data,
	})
}

// GetSummaryStats 获取统计摘要
func GetSummaryStats(c *gin.Context) {
	sqlDB, err := db.GetConnection("default")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败"})
		return
	}

	// 总数据量
	var totalCount int
	err = sqlDB.QueryRow("SELECT COUNT(*) FROM business_data").Scan(&totalCount)
	if err != nil {
		totalCount = 0
	}

	// 总价值
	var totalValue float64
	err = sqlDB.QueryRow("SELECT SUM(value) FROM business_data").Scan(&totalValue)
	if err != nil {
		totalValue = 0
	}

	// 今日数据
	var todayCount int
	today := time.Now().Format("2006-01-02")
	err = sqlDB.QueryRow("SELECT COUNT(*) FROM business_data WHERE date = ?", today).Scan(&todayCount)
	if err != nil {
		todayCount = 0
	}

	// 分类数量
	var categoryCount int
	err = sqlDB.QueryRow("SELECT COUNT(DISTINCT category) FROM business_data").Scan(&categoryCount)
	if err != nil {
		categoryCount = 0
	}

	stats := gin.H{
		"total_count":    totalCount,
		"total_value":    totalValue,
		"today_count":    todayCount,
		"category_count": categoryCount,
		"avg_value":      totalValue / float64(totalCount),
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    stats,
	})
}

// GetTrendStats 获取趋势统计
func GetTrendStats(c *gin.Context) {
	period := c.DefaultQuery("period", "7") // 默认7天
	days, _ := strconv.Atoi(period)

	sqlDB, err := db.GetConnection("default")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败"})
		return
	}

	query := `
		SELECT 
			DATE(date) as day,
			COUNT(*) as count,
			SUM(value) as total_value,
			AVG(value) as avg_value
		FROM business_data 
		WHERE date >= DATE_SUB(CURDATE(), INTERVAL ? DAY)
		GROUP BY DATE(date)
		ORDER BY day
	`

	rows, err := sqlDB.Query(query, days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	defer rows.Close()

	type TrendData struct {
		Date       string  `json:"date"`
		Count      int     `json:"count"`
		TotalValue float64 `json:"total_value"`
		AvgValue   float64 `json:"avg_value"`
	}

	var trends []TrendData
	for rows.Next() {
		var trend TrendData
		err := rows.Scan(&trend.Date, &trend.Count, &trend.TotalValue, &trend.AvgValue)
		if err != nil {
			continue
		}
		trends = append(trends, trend)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    trends,
	})
}

// UploadFile 文件上传
func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件上传失败"})
		return
	}
	defer file.Close()

	// 创建上传目录
	uploadDir := "./uploads"
	os.MkdirAll(uploadDir, 0755)

	// 生成唯一文件名
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)
	filepath := filepath.Join(uploadDir, filename)

	// 保存文件
	out, err := os.Create(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "上传成功",
		"data": gin.H{
			"filename": filename,
			"path":     filepath,
			"size":     header.Size,
		},
	})
}

// UploadCSV CSV文件上传并解析
func UploadCSV(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件上传失败"})
		return
	}
	defer file.Close()

	// 检查文件类型
	if filepath.Ext(header.Filename) != ".csv" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "只支持CSV文件"})
		return
	}

	// 解析CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CSV解析失败"})
		return
	}

	if len(records) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CSV文件内容为空"})
		return
	}

	// 处理数据（这里简化处理，实际应该根据具体格式解析）
	headers := records[0]
	data := records[1:]

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "CSV解析成功",
		"data": gin.H{
			"headers": headers,
			"rows":    data,
			"count":   len(data),
		},
	})
}

// 页面处理函数
func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "数据可视化管理平台",
		"page":  "dashboard",
	})
}

func DashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title": "仪表板 - 数据可视化管理平台",
		"page":  "dashboard",
	})
}

func ChartsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "charts.html", gin.H{
		"title": "图表管理 - 数据可视化管理平台",
		"page":  "charts",
	})
}

func DataPage(c *gin.Context) {
	c.HTML(http.StatusOK, "data.html", gin.H{
		"title": "数据管理 - 数据可视化管理平台",
		"page":  "data",
	})
}

func SettingsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.html", gin.H{
		"title": "系统设置 - 数据可视化管理平台",
		"page":  "settings",
	})
}

// 其他处理函数（简化实现）
func UpdateChart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

func DeleteChart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func UpdateBusinessData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

func DeleteBusinessData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func GetDashboards(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
}

func GetDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": map[string]interface{}{}})
}

func CreateDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

func UpdateDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

func DeleteDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func GetCategoryStats(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": map[string]interface{}{}})
}