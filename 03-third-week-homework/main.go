package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
)

// StartBackendServer 启动后台服务
func StartBackendServer(srv *http.Server) error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello world")
	})
	return srv.ListenAndServe()
}

//StartBackendServer 启动性能检测服务
func startDebugServer(srv *http.Server) error {
	http.HandleFunc("/debug", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello debug")
	})
	return srv.ListenAndServe()
}

func main() {
	//定义后台端口
	backendSer := &http.Server{Addr: ":8080"}
	//定义性能检测端口
	debugSer := &http.Server{Addr: ":8081"}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)
	//信号量处理
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel)
	//执行后台服务
	group.Go(func() error {
		return StartBackendServer(backendSer)
	})
	//执行性能检测服务
	group.Go(func() error {
		return startDebugServer(debugSer)
	})
	group.Go(func() error {
		//堵塞等待监听关闭
		<-errCtx.Done()
		err := debugSer.Shutdown(errCtx)
		err = backendSer.Shutdown(errCtx)
		return err
	})
	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-signalChannel:
				cancel()
			}
		}
	})
	err := group.Wait()
	if err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")
}
