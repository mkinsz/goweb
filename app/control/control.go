package control

import (
	"goweb/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"content": "1234",
		"title":   "Go Templates !",
	})
}

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tpl", gin.H{})
}

func AboutHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "blank.tpl", gin.H{})
}

func ActionHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tpl", gin.H{})
}

func SimuHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "simu.tpl", gin.H{
		"title": "Graphql",
	})
}

func GrapqlHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tpl", gin.H{})
}

func TemplateHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "template.tpl", gin.H{})
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var todo model.Todo
	c.BindJSON(&todo)
	// 2. 存入数据库
	err := model.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

func GetTodoList(c *gin.Context) {
	// 查询todo这个表里的所有数据
	todoList, err := model.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := model.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = model.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := model.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
