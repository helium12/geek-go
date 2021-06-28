package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"
)

//接收系统取消信号，关闭http服务进程
func main() {
	//create cancel context
	ctxWithCancel, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	//use cancel context
	g, ctx := errgroup.WithContext(ctxWithCancel)

	//manage http server with err group
	//init http server with mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Qcon! N+2")
	})
	s := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	//use err group to start http server
	g.Go(func() error {
		return s.ListenAndServe()
	})
	//use err group with cancel context to stop http server
	g.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("stopping http server")
			s.Shutdown(ctx)
			fmt.Println("stopped http server")
			return errors.New("task done")
		}
		return nil
	})

	//deal with signal with err group
	g.Go(func() error {
		c := make(chan os.Signal, 1)
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
