-- 数据可视化管理平台数据库初始化脚本
-- 创建数据库
CREATE DATABASE IF NOT EXISTS `data_visualization` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `data_visualization`;

-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(50) UNIQUE NOT NULL COMMENT '用户名',
    `password` VARCHAR(255) NOT NULL COMMENT '密码（加密）',
    `email` VARCHAR(100) COMMENT '邮箱',
    `phone` VARCHAR(20) COMMENT '手机号',
    `role` VARCHAR(20) DEFAULT 'user' COMMENT '角色：admin/user',
    `status` TINYINT DEFAULT 1 COMMENT '状态：1-正常 0-禁用',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_username` (`username`),
    INDEX `idx_email` (`email`),
    INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 业务数据表
CREATE TABLE IF NOT EXISTS `business_data` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `category` VARCHAR(50) NOT NULL COMMENT '数据分类',
    `value` DECIMAL(15,4) NOT NULL COMMENT '数值',
    `date` DATE NOT NULL COMMENT '日期',
    `description` TEXT COMMENT '描述',
    `tags` VARCHAR(255) COMMENT '标签',
    `created_by` BIGINT COMMENT '创建人',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX `idx_category` (`category`),
    INDEX `idx_date` (`date`),
    INDEX `idx_created_by` (`created_by`),
    INDEX `idx_category_date` (`category`, `date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='业务数据表';

-- 图表配置表
CREATE TABLE IF NOT EXISTS `chart_configs` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL COMMENT '图表名称',
    `type` VARCHAR(20) NOT NULL COMMENT '图表类型：line/bar/pie/radar/scatter',
    `config` JSON COMMENT '图表配置（JSON格式）',
    `data_source` VARCHAR(255) COMMENT '数据源',
    `description` TEXT COMMENT '描述',
    `status` TINYINT DEFAULT 1 COMMENT '状态：1-启用 0-禁用',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_type` (`type`),
    INDEX `idx_status` (`status`),
    INDEX `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='图表配置表';

-- 仪表板配置表
CREATE TABLE IF NOT EXISTS `dashboards` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL COMMENT '仪表板名称',
    `layout` JSON COMMENT '布局配置（JSON格式）',
    `charts` JSON COMMENT '图表ID列表（JSON数组）',
    `description` TEXT COMMENT '描述',
    `is_default` BOOLEAN DEFAULT FALSE COMMENT '是否默认仪表板',
    `created_by` BIGINT COMMENT '创建人',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_created_by` (`created_by`),
    INDEX `idx_is_default` (`is_default`),
    INDEX `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='仪表板配置表';

-- 系统日志表
CREATE TABLE IF NOT EXISTS `system_logs` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT COMMENT '用户ID',
    `action` VARCHAR(50) NOT NULL COMMENT '操作类型',
    `module` VARCHAR(50) NOT NULL COMMENT '模块名称',
    `description` TEXT COMMENT '操作描述',
    `ip` VARCHAR(45) COMMENT 'IP地址',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_action` (`action`),
    INDEX `idx_module` (`module`),
    INDEX `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统日志表';

-- 数据源配置表
CREATE TABLE IF NOT EXISTS `data_sources` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL COMMENT '数据源名称',
    `type` VARCHAR(20) NOT NULL COMMENT '类型：mysql/postgresql/api/file',
    `config` JSON COMMENT '连接配置（JSON格式）',
    `description` TEXT COMMENT '描述',
    `status` TINYINT DEFAULT 1 COMMENT '状态：1-启用 0-禁用',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_type` (`type`),
    INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据源配置表';

-- 插入默认数据
-- 默认管理员用户（密码：admin123，实际使用时应该使用加密密码）
INSERT IGNORE INTO `users` (`username`, `password`, `email`, `role`, `status`) VALUES 
('admin', '$2a$10$r3d6v5t7y9u1i3o5a7c9e', 'admin@example.com', 'admin', 1),
('user1', '$2a$10$r3d6v5t7y9u1i3o5a7c9e', 'user1@example.com', 'user', 1),
('user2', '$2a$10$r3d6v5t7y9u1i3o5a7c9e', 'user2@example.com', 'user', 1);

-- 插入示例业务数据（最近30天的销售数据）
INSERT IGNORE INTO `business_data` (`category`, `value`, `date`, `description`, `tags`, `created_by`) VALUES 
('sales', 15000.50, DATE_SUB(CURDATE(), INTERVAL 1 DAY), '日销售额', 'sales,daily', 1),
('sales', 14800.75, DATE_SUB(CURDATE(), INTERVAL 2 DAY), '日销售额', 'sales,daily', 1),
('sales', 15200.25, DATE_SUB(CURDATE(), INTERVAL 3 DAY), '日销售额', 'sales,daily', 1),
('sales', 14950.80, DATE_SUB(CURDATE(), INTERVAL 4 DAY), '日销售额', 'sales,daily', 1),
('sales', 15120.40, DATE_SUB(CURDATE(), INTERVAL 5 DAY), '日销售额', 'sales,daily', 1),
('sales', 14780.60, DATE_SUB(CURDATE(), INTERVAL 6 DAY), '日销售额', 'sales,daily', 1),
('sales', 15300.90, DATE_SUB(CURDATE(), INTERVAL 7 DAY), '周销售额', 'sales,weekly', 1),
('cost', 8500.25, DATE_SUB(CURDATE(), INTERVAL 1 DAY), '日成本', 'cost,daily', 1),
('cost', 8400.75, DATE_SUB(CURDATE(), INTERVAL 2 DAY), '日成本', 'cost,daily', 1),
('cost', 8600.50, DATE_SUB(CURDATE(), INTERVAL 3 DAY), '日成本', 'cost,daily', 1),
('profit', 6500.25, DATE_SUB(CURDATE(), INTERVAL 1 DAY), '日利润', 'profit,daily', 1),
('profit', 6400.00, DATE_SUB(CURDATE(), INTERVAL 2 DAY), '日利润', 'profit,daily', 1),
('profit', 6600.40, DATE_SUB(CURDATE(), INTERVAL 3 DAY), '日利润', 'profit,daily', 1),
('users', 150.00, DATE_SUB(CURDATE(), INTERVAL 1 DAY), '日新增用户', 'users,daily', 1),
('users', 145.00, DATE_SUB(CURDATE(), INTERVAL 2 DAY), '日新增用户', 'users,daily', 1),
('users', 155.00, DATE_SUB(CURDATE(), INTERVAL 3 DAY), '日新增用户', 'users,daily', 1),
('orders', 320.00, DATE_SUB(CURDATE(), INTERVAL 1 DAY), '日订单量', 'orders,daily', 1),
('orders', 315.00, DATE_SUB(CURDATE(), INTERVAL 2 DAY), '日订单量', 'orders,daily', 1),
('orders', 325.00, DATE_SUB(CURDATE(), INTERVAL 3 DAY), '日订单量', 'orders,daily', 1);

-- 插入默认图表配置
INSERT IGNORE INTO `chart_configs` (`name`, `type`, `config`, `data_source`, `description`, `status`) VALUES 
('销售趋势图', 'line', '{"title": "销售趋势", "color": "#3498db"}', 'SELECT date, SUM(value) as total FROM business_data WHERE category="sales" GROUP BY date', '显示销售数据趋势', 1),
('分类占比图', 'pie', '{"title": "数据分类占比", "colors": ["#e74c3c", "#3498db", "#2ecc71", "#f39c12"]}', 'SELECT category, COUNT(*) as count FROM business_data GROUP BY category', '显示数据分类占比', 1),
('用户增长图', 'bar', '{"title": "用户增长", "color": "#2ecc71"}', 'SELECT date, SUM(value) as total FROM business_data WHERE category="users" GROUP BY date', '显示用户增长情况', 1);

-- 插入默认仪表板
INSERT IGNORE INTO `dashboards` (`name`, `layout`, `charts`, `description`, `is_default`, `created_by`) VALUES 
('默认仪表板', '{"columns": 2, "rows": 2}', '[1, 2, 3]', '系统默认仪表板', TRUE, 1);

-- 插入系统日志示例
INSERT IGNORE INTO `system_logs` (`user_id`, `action`, `module`, `description`, `ip`) VALUES 
(1, 'login', 'auth', '用户登录系统', '127.0.0.1'),
(1, 'create', 'chart', '创建销售趋势图', '127.0.0.1'),
(1, 'view', 'dashboard', '查看仪表板', '127.0.0.1');

-- 创建视图用于数据统计
CREATE OR REPLACE VIEW `daily_stats` AS 
SELECT 
    date,
    SUM(CASE WHEN category = 'sales' THEN value ELSE 0 END) as sales_total,
    SUM(CASE WHEN category = 'cost' THEN value ELSE 0 END) as cost_total,
    SUM(CASE WHEN category = 'profit' THEN value ELSE 0 END) as profit_total,
    SUM(CASE WHEN category = 'users' THEN value ELSE 0 END) as users_total,
    SUM(CASE WHEN category = 'orders' THEN value ELSE 0 END) as orders_total,
    COUNT(*) as record_count
FROM business_data 
GROUP BY date 
ORDER BY date DESC;

-- 创建存储过程用于生成月度报告
DELIMITER //
CREATE PROCEDURE GenerateMonthlyReport(IN report_month DATE)
BEGIN
    SELECT 
        DATE_FORMAT(date, '%Y-%m') as month,
        category,
        SUM(value) as total_value,
        AVG(value) as avg_value,
        COUNT(*) as record_count
    FROM business_data 
    WHERE DATE_FORMAT(date, '%Y-%m') = DATE_FORMAT(report_month, '%Y-%m')
    GROUP BY DATE_FORMAT(date, '%Y-%m'), category
    ORDER BY total_value DESC;
END //
DELIMITER ;

-- 输出初始化完成信息
SELECT '数据库初始化完成！' as message;