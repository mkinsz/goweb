package main

import (
	"goweb/app/model"
	"goweb/app/route"
)

func init() {

}

func main() {
	model.InitDB()

	r := route.Setup()
	r.Run(":9000")
}
