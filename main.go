package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/signal"
	"syscall"
	"time"
)

func start(srv *http.Server) {
	env := GetEnv()

	innerUrl := "http://" + env.InnerHost
	url, err := url.Parse(innerUrl)
	if err != nil {
		log.Fatalf("[WRP] Invalid inner url: %s\n", innerUrl)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	log.Printf("[WRP] Proxy URL: %s\n", innerUrl)

	srv.Addr = env.ServrAddr()
	srv.Handler = proxy
	log.Printf("[WRP] Serve on %s\n", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("[WRP] Serve is end: %v\n", err)
	} else {
		log.Println("[WRP] Serve is end gracefully")
	}
}

func stop(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("[WRP] Server stopped: %v\n", err)
	} else {
		log.Println("[WRP] Server stopped")
	}
}

func main() {
	srv := &http.Server{}
	go start(srv)

	// Wait for interrupt signal to gracefully shutdown the server
	ctx, reset := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	log.Println("[WRP] Press Ctrl+C to shutdown server...")
	<-ctx.Done()

	log.Println("[WRP] Shutdown server...")
	reset()
	stop(srv)
}
