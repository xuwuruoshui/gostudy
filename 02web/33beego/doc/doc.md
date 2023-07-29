# Beego使用教程

## 快速开始

```sh
# 安装bee工具
go get github.com/beego/bee/v2
# 创建项目
bee new your-project
# 下载依赖
go mod tidy
# 运行
bee run
```

## 路由

### 路由风格

1. 函数式风格

```go
beego.Get("/demo", GetHello)
func GetHello(ctx *context.Context) {
ctx.WriteString("HelloWorld1")
}
```

2. 控制器风格

```go
// 新版
beego.CtrlGet("/test", (*controllers.MainController).GetHelloWorld)
type MainController struct {
beego.Controller
}
func (c *MainController) GetHelloWorld() {
c.Ctx.WriteString("新版路由")
}

// 旧版
beego.Router("/hello", &controllers.MainController{}, "get:GetHello")

type MainController struct {
beego.Controller
}
func (c *MainController) GetHello() {
c.Ctx.WriteString("Hello World!!!")
}
```

### 命名空间

```go
// 类型于gin的路由组
ns := beego.NewNamespace("/v1",
beego.NSGet("/a", controllers.GetA),
beego.NSPost("/b", controllers.POSTA))
beego.AddNamespace(ns)
```

### 过滤器

```go
beego.InsertFilter("/v1/*", beego.BeforeRouter, func(ctx *context.Context) {
ctx.WriteString("禁止访问")
})
```

## 模板语法
1. 普通数据获取
```go
func (c *MainController) Get() {
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
```
```html
<a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
```
2. 遍历
```go
pages := []struct {
Num int
}{{10}, {20}, {30}}
c.Data["Numbers"] = pages

c.TplName = "index.html"
```
```html
{{range $index,$value := .Numbers}}
{{$index}}={{$value.Num}} of {{$.Website}}<br>
{{end}}
```
3. 引入html
```html
{{template "footer.html" .}}
```

## ORM
### 连接数据库
```shell
go get -u github.com/beego/beego/v2/hello/orm
go get -u github.com/go-sql-driver/mysql
```
```config
appname = 01test
httpport = 8080
runmode = dev

[dev]
defaultdb = root:root@(192.168.0.110)/test?charset=utf8mb4
```
```go
func main() {
    db, _ := beego.AppConfig.String("defaultdb")
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", db)
    beego.Run()
}
```
### 相关操作
app.conf
```config
# 能接受到json
copyrequestbod-y = true
```

```go

```