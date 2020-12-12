package main

import (
	"Iris-CMS/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {

	//构建APP
	app := newApp()

	//应用APP设置
	configation(app)

	//从配置文件读取配置参数
	conf := config.InitConfig()
	addr := ":" + conf.Port

	//监听
	app.Run(
		iris.Addr(addr),//在端口8080上运行
		iris.WithoutServerError(iris.ErrServerClosed),//无服务错误提示
		iris.WithOptimizations,//对json数据序列化更快的配置
		)

}

//构建APP
func newApp() *iris.Application {

	app := iris.New()

	//设定应用图标
	app.Favicon("./static/favicon.ico")

	//设置日志级别，开发阶段为debug
	app.Logger().SetLevel("debug")

	//注册静态资源
	//将URL请求的目录文件映射到当前项目中的指定目录
	app.HandleDir("/static", "./static")
	app.HandleDir("/manage/static", "./static")

	//注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	//请求localhost:8080页面时返回index.html
	app.Get("/", func(ctx context.Context) {
		ctx.View("index.html")
	})

	return app

}

//项目设置
func configation(app *iris.Application) {

	//配置字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//错误配置
	//页面未找到
	app.OnErrorCode(iris.StatusNotFound, func(ctx context.Context) {
		ctx.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg": "404 not found",
			"data": iris.Map{},
		})
	})

	//服务端出错
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx context.Context) {
		ctx.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg": "server error",
			"data": iris.Map{},
		})
	})

}