package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"
)

type InitServer struct {
	addr string
	url  string
	msg  string
}

//接收系统取消信号，关闭http服务进程
func main() {
	//create cancel context
	ctxWithCancel, cancel := context.WithTimeout(context.Background(), time.Second*600)
	defer cancel()

	//use cancel context
	g, _ := errgroup.WithContext(ctxWithCancel)

	//use err group to start http servers
	https := []*InitServer{
		&InitServer{"127.0.0.1:8001", "/h", "Hello, Qcon! N+1"},
		&InitServer{"127.0.0.1:8002", "/h", "Hello, Qcon! N+2"},
	}
	servers := initServers(https)
	for _, s := range servers {
		fmt.Println(".........", s.Handler)

		g.Go(func() error {
			return s.ListenAndServe()
		})

		//g.Go(func() error {
		//	select {
		//	case <-ctx.Done():
		//		fmt.Println("stopping http server: ", s.Addr)
		//		s.Shutdown(ctx)
		//		fmt.Println("stopped http server: ", s.Addr)
		//		//return errors.New("task done")
		//	}
		//	return nil
		//})
	}

	//use err group to stop http servers
	//for _,st := range servers{
	//	g.Go(func() error {
	//		select {
	//		case <-ctx.Done():
	//			fmt.Println("stopping http server: ", st.Addr)
	//			st.Shutdown(ctx)
	//			fmt.Println("stopped http server: ", st.Addr)
	//			//return errors.New("task done")
	//		}
	//		return nil
	//	})
	//}

	////manage http server with err group
	////init http server with mux
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, Qcon! N+2")
	//})
	//s := http.Server{
	//	Addr:    "0.0.0.0:8080",
	//	Handler: mux,
	//}
	//
	////use err group to start http server
	//g.Go(func() error {
	//	return s.ListenAndServe()
	//})
	////use err group with cancel context to stop http server
	//g.Go(func() error {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("stopping http server")
	//		s.Shutdown(ctx)
	//		fmt.Println("stopped http server")
	//		return errors.New("task done")
	//	}
	//	return nil
	//})

	//deal with signal with err group
	g.Go(func() error {
		c := make(chan os.Signal)
		defer close(c)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM,
			syscall.SIGKILL)
		select {
		case <-c:
			fmt.Println("got cancel signal, prepare to close the http server.")
			cancel()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("task done, exit.")
	}
}

func initServers(svrs []*InitServer) (servers []*http.Server) {
	//servers := []*InitServer{
	//	&InitServer{"127.0.0.1:8080","/","Hello, Qcon! N+1"},
	//	&InitServer{"127.0.0.1:8081","/","Hello, Qcon! N+2"},
	//}
	//var servers []*http.Server

	for _, svr := range svrs {
		mux := http.NewServeMux()
		//mux.HandleFunc(svr.url, func(w http.ResponseWriter, r *http.Request) {
		//	fmt.Fprintf(w, svr.msg)
		//})

		fmt.Println("init svr: ", svr.url)
		mux.HandleFunc(svr.url, handler)

		servers = append(servers, &http.Server{Addr: svr.addr, Handler: mux})
	}
	return servers
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "3 欢迎访问 www.ydook.com !")
}
