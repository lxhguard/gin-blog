package main

import (
	"ginblog/config"
    "fmt"
)

func main() {
	fmt.Println(config.AppMode)
	fmt.Println(config.DbHost)
}