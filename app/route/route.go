package route

import (
	"goweb/app/control"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	/**
	  所有的接口都要由路由来进行管理。
	      Gin的路由支持GET,POST,PUT,DELETE,PATCH,HEAD,OPTIONS等请求
	      同时还有一个Any函数，可以同时支持以上的所有请求。

	  创建路由(router)并引入默认中间件
	      router := gin.Default()
	      在源码中,首先是New一个engine,紧接着通过Use方法传入了Logger()和Recovery()这两个中间件。
	      其中 Logger 是对日志进行记录，而 Recovery 是对有 painc时, 进行500的错误处理。

	  创建路由(router)无中间件
	      router := gin.New()
	*/
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.Static("/assets", "./assets")

	//定义GET方法
	r.GET("/", control.IndexHandler)
	r.GET("/login", control.LoginHandler)
	r.GET("/about", control.AboutHandler)
	r.GET("/simu", control.SimuHandler)
	r.GET("/graphql", control.GrapqlHandler)
	r.GET("/template", control.TemplateHandler)

	// 待办事项
	apiv1 := r.Group("v1")
	{
		// 添加
		apiv1.POST("/todo", control.CreateTodo)
		// 查看所有的待办事项
		apiv1.GET("/todo", control.GetTodoList)
		// 修改某一个待办事项
		apiv1.PUT("/todo/:id", control.UpdateATodo)
		// 删除某一个待办事项
		apiv1.DELETE("/todo/:id", control.DeleteATodo)
	}

	return r
}
