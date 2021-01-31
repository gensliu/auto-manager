# Auto-Manager
机动车安防后端

# 0. 准备
go 1.15+\
开启go mod, 设置仓库镜像 

MariaDB 10+ (MySQL 5+)

# 1. 检查
    go mod tidy

# 2. 编译
    go build -o auto-manager

# 3. 运行
    ./auto-manager

# Makefile
    make check - 检查第三方依赖并安装
    make build - 编译Go代码，生成可执行文件
    make run   - 直接运行main.go
    make help  - 获取帮助信息

# 接口文档
  编译并执行   
  默认自动化文档地址： http://127.0.0.1:8888/swagger/index.html   
  默认API运行地址： http://127.0.0.1:8888/v1
  