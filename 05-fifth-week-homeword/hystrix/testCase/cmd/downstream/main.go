package main

import "go-demo/05-fifth-week-homeword/hystrix/testCase/server"

func main() {
	err := server.NewDownStreamServer(1).Run(":8000")
	if err != nil {
		panic("downStream start fail")
	}
}
