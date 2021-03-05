# goweb
go web html

## Tree

config：用于存储配置文件
middleware：应用中间件
model：应用数据库模型
route: 路由
control: 逻辑处理
runtime：应用运行时数据

## App Flow

url   --> controller  -->  logic    -->      model
请求  -->   控制器    --> 业务逻辑  -->  模型层的增删改查