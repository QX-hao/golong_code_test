package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
)

// 用户表结构
type User struct {
	ID        int64     `json:"id" db:"id" primaryKey:"true"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	Role      string    `json:"role" db:"role"`
	Status    int       `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// 业务数据表结构
type BusinessData struct {
	ID          int64     `json:"id" db:"id" primaryKey:"true"`
	Category    string    `json:"category" db:"category"`
	Value       float64   `json:"value" db:"value"`
	Date        time.Time `json:"date" db:"date"`
	Description string    `json:"description" db:"description"`
	Tags        string    `json:"tags" db:"tags"`
	CreatedBy   int64     `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// 图表配置表结构
type ChartConfig struct {
	ID          int64     `json:"id" db:"id" primaryKey:"true"`
	Name        string    `json:"name" db:"name"`
	Type        string    `json:"type" db:"type"` // line, bar, pie, etc.
	Config      string    `json:"config" db:"config"` // JSON配置
	DataSource  string    `json:"data_source" db:"data_source"`
	Description string    `json:"description" db:"description"`
	Status      int       `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// 仪表板配置表结构
type Dashboard struct {
	ID          int64     `json:"id" db:"id" primaryKey:"true"`
	Name        string    `json:"name" db:"name"`
	Layout      string    `json:"layout" db:"layout"` // JSON布局配置
	Charts      string    `json:"charts" db:"charts"` // 图表ID列表
	Description string    `json:"description" db:"description"`
	IsDefault   bool      `json:"is_default" db:"is_default"`
	CreatedBy   int64     `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// 系统日志表结构
type SystemLog struct {
	ID          int64     `json:"id" db:"id" primaryKey:"true"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Action      string    `json:"action" db:"action"`
	Module      string    `json:"module" db:"module"`
	Description string    `json:"description" db:"description"`
	IP          string    `json:"ip" db:"ip"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// 初始化数据库表
func InitModels() error {
	// 获取数据库连接
	sqlDB, err := db.GetConnection("default")
	if err != nil {
		return fmt.Errorf("获取数据库连接失败: %v", err)
	}

	// 创建用户表
	err = createUserTable(sqlDB)
	if err != nil {
		return err
	}

	// 创建业务数据表
	err = createBusinessDataTable(sqlDB)
	if err != nil {
		return err
	}

	// 创建图表配置表
	err = createChartConfigTable(sqlDB)
	if err != nil {
		return err
	}

	// 创建仪表板表
	err = createDashboardTable(sqlDB)
	if err != nil {
		return err
	}

	// 创建系统日志表
	err = createSystemLogTable(sqlDB)
	if err != nil {
		return err
	}

	// 插入默认数据
	err = insertDefaultData(sqlDB)
	if err != nil {
		log.Printf("插入默认数据失败: %v", err)
	}

	log.Println("数据库表初始化完成")
	return nil
}

// 创建用户表
func createUserTable(sqlDB *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(100),
		phone VARCHAR(20),
		role VARCHAR(20) DEFAULT 'user',
		status TINYINT DEFAULT 1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		INDEX idx_username (username),
		INDEX idx_email (email)
	)
	`
	_, err := sqlDB.Exec(query)
	return err
}

// 创建业务数据表
func createBusinessDataTable(sqlDB *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS business_data (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		category VARCHAR(50) NOT NULL,
		value DECIMAL(15,4) NOT NULL,
		date DATE NOT NULL,
		description TEXT,
		tags VARCHAR(255),
		created_by BIGINT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		INDEX idx_category (category),
		INDEX idx_date (date),
		INDEX idx_created_by (created_by)
	)
	`
	_, err := sqlDB.Exec(query)
	return err
}

// 创建图表配置表
func createChartConfigTable(sqlDB *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS chart_configs (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		type VARCHAR(20) NOT NULL,
		config JSON,
		data_source VARCHAR(255),
		description TEXT,
		status TINYINT DEFAULT 1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		INDEX idx_type (type),
		INDEX idx_status (status)
	)
	`
	_, err := sqlDB.Exec(query)
	return err
}

// 创建仪表板表
func createDashboardTable(sqlDB *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS dashboards (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		layout JSON,
		charts JSON,
		description TEXT,
		is_default BOOLEAN DEFAULT FALSE,
		created_by BIGINT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		INDEX idx_created_by (created_by),
		INDEX idx_is_default (is_default)
	)
	`
	_, err := sqlDB.Exec(query)
	return err
}

// 创建系统日志表
func createSystemLogTable(sqlDB *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS system_logs (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		user_id BIGINT,
		action VARCHAR(50) NOT NULL,
		module VARCHAR(50) NOT NULL,
		description TEXT,
		ip VARCHAR(45),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		INDEX idx_user_id (user_id),
		INDEX idx_action (action),
		INDEX idx_module (module),
		INDEX idx_created_at (created_at)
	)
	`
	_, err := sqlDB.Exec(query)
	return err
}

// 插入默认数据
func insertDefaultData(sqlDB *sql.DB) error {
	// 插入默认管理员用户（密码：admin123，实际使用时应该加密）
	query := `INSERT IGNORE INTO users (username, password, email, role, status) VALUES (?, ?, ?, ?, ?)`
	_, err := sqlDB.Exec(query, "admin", "$2a$10$r3d6v5t7y9u1i3o5a7c9e", "admin@example.com", "admin", 1)
	if err != nil {
		return err
	}

	// 插入示例业务数据
	query = `INSERT IGNORE INTO business_data (category, value, date, description, tags, created_by) VALUES (?, ?, ?, ?, ?, ?)`
	today := time.Now()
	for i := 0; i < 30; i++ {
		date := today.AddDate(0, 0, -i)
		value := 1000.0 + float64(i)*50 + float64(i%10)*20
		_, err = sqlDB.Exec(query, "sales", value, date.Format("2006-01-02"), "销售数据", "sales,daily", 1)
		if err != nil {
			return err
		}
	}

	return nil
}