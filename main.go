package main

import (
	"ginblog/model"
	"ginblog/router"
)

func main() {
	model.ConnectDb()
	router.CreateRouter()
}