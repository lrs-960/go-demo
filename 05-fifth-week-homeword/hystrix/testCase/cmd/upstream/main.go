package main

import (
	"go-demo/05-fifth-week-homeword/hystrix/testCase/server"
	"time"
)

func main() {
	err := server.NewUpStreamServer(
		10,
		50,
		0.8,
		time.Second*5,
	).Run(":9000")
	if err != nil {
		panic("upstream start fail")
	}
}
