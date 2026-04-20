package main

import (
	"TestDemo1/server"
	"fmt"
	// "log"
)


func main() {

	config,err := server.InitConfig_v1("config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}