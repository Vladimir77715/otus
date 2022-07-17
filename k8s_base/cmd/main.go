package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Vladimir77715/otus/k8s_base/internal/app"
	"github.com/Vladimir77715/otus/k8s_base/internal/version"
	"net"
	"net/http"
	"os/signal"
	"syscall"
)

var (
	port string
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	defer stop()

	flag.StringVar(&port, "port", "8000", "port for server")

	fmt.Printf("version: %s buildTime: %s\n", version.Version, version.BuildTime)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", app.Heath)

	s := http.Server{Addr: net.JoinHostPort("", port), Handler: log(mux)}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			println(err.Error())
			return
		}
	}()

	<-ctx.Done()

	if err := s.Shutdown(ctx); err != nil {
		println(err.Error())
	}

}

func log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("method: %s, url: %s addr: %s\n", request.Method, request.URL, request.RemoteAddr)
		handler.ServeHTTP(writer, request)
	})
}
