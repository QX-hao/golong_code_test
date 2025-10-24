@echo off
echo 正在启动数据可视化管理平台...
echo.

REM 检查Go环境
go version >nul 2>&1
if errorlevel 1 (
    echo 错误: 未检测到Go环境，请先安装Go语言
    pause
    exit /b 1
)

REM 检查依赖
echo 检查项目依赖...
go mod tidy

REM 创建必要的目录
if not exist uploads mkdir uploads
if not exist static mkdir static
if not exist logs mkdir logs

REM 编译项目
echo 编译项目...
go build -o data-visualization-platform.exe

if errorlevel 1 (
    echo 错误: 编译失败，请检查代码错误
    pause
    exit /b 1
)

echo.
echo 项目编译成功！
echo.
echo 启动服务...
echo 访问地址: http://localhost:9033
echo 管理后台: http://localhost:9033/admin
echo.

REM 启动服务
start data-visualization-platform.exe

echo 服务已启动，按任意键退出...
pause >nul