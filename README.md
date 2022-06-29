# Osier

### 目标：优雅、高效、开箱即用的Go框架！
    能满足更多项目基本功能需求，达到开箱即用的效果。
    保证运行效率，每个功能有一定自由度，不需要过多配置。
    代码编写规范，逻辑清晰，上手快，利于二次代码编写。
    绝不增加可有可无功能，防止框架臃肿。


### 目录结构
- app 应用目录
    - controller 控制器
    - model 模型
    - middle 中间件
        - middle.go 全局中间件，所有继承方法都将自动挂载

- boot 启动及全局数据
    - kernel 启动项
    - boot.go 全局所需句柄
    - cache.go 全局缓存所需KEY
    - ctr.go 全局控制器-应用控制器需继承
    - lang.go 全局语言配置
    - mdl.go 全局数据模型-应用模型需继承
    - tool.go 全局公共方法

- config 内部配置
- router 路由器
    - router.go 其他路由文件都需继承该类
    - admin.go 与 api.go 保持方法继承格式后自动加载

- storage 存放和日志
- web 后台管理Vue前端
- config.ini 配置文件


### 其他
- JWT认证 - 因其不可控，不会引入该中间件
- swagger - 计划引入接口文档生成工具


### 注意
    0. 全局将 osier 替换成 你的项目名称
    1. config.ini.example 复制并改名 config.ini
    2. 当前目录执行命令： go mod tidy
    3. go run main.go
