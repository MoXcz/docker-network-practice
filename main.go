package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
)

var listenAddr = flag.String("addr", ":4000", "HTTP network address")

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	mux.HandleFunc("/", handler)

	logger.Info("starting server", "addr", *listenAddr)

	err := http.ListenAndServe(*listenAddr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}

func handler(w http.ResponseWriter, r *http.Request) {
	flag.Parse()
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		return
	}

	host, _ := os.Hostname()
	addrs, err := net.LookupHost(host)

	fmt.Println(string(reqDump))

	msg := fmt.Sprintf("Hi from %s (hostname: %s IP: %s)\n", *listenAddr, host, addrs)
	w.Write([]byte(msg))
}
