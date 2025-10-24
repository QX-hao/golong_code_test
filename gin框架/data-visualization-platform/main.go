package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/gin-gonic/gin"

	"data-visualization-platform/models"
	"data-visualization-platform/routes"
)

func main() {
	// 初始化Gin路由
	r := gin.Default()

	// 初始化Go-Admin引擎
	e := engine.Default()

	// 加载配置
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       "127.0.0.1",
				Port:       "3306",
				User:       "root",
				Pwd:        "password",
				Name:       "data_visualization",
				MaxIdleCon: 30,
				MaxOpenCon: 100,
				Driver:     config.DriverMysql,
			},
		},
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.CN,
		Theme:       "adminlte",
		Title:       "数据可视化管理平台",
		Logo:        template.HTML(`<b>数据可视化</b>平台`),
		MiniLogo:    template.HTML(`<b>DV</b>`),
		IndexUrl:    "/admin",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}

	// 设置模板主题
	template.AddComp(chartjs.NewChart())

	// 初始化Admin插件
	adminPlugin := admin.NewAdmin(&table.Generators{
		// 这里可以添加自定义的表格生成器
	})

	// 添加插件到引擎
	err := e.AddConfig(cfg).
		AddPlugins(adminPlugin).
		Use(r)

	if err != nil {
		log.Fatal(err)
	}

	// 注册自定义路由
	routes.RegisterRoutes(r)

	// 初始化数据库表
	err = models.InitModels()
	if err != nil {
		log.Fatal("初始化数据库表失败:", err)
	}

	fmt.Println("数据可视化管理平台启动成功!")
	fmt.Println("访问地址: http://localhost:9033/admin")
	fmt.Println("API地址: http://localhost:9033/api")

	// 启动服务
	err = r.Run(":9033")
	if err != nil {
		log.Fatal(err)
	}
}

// 初始化函数
func init() {
	// 创建必要的目录
	os.MkdirAll("./uploads", 0755)
	os.MkdirAll("./logs", 0755)
	os.MkdirAll("./temp", 0755)
}