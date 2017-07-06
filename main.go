package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	port      = 8080
	static, _ = os.Getwd()
)

func init() {
	flag.IntVar(&port, "port", port, "http port")
	flag.StringVar(&static, "dir", static, "static content directory")
}

func setup() {
	flag.Parse()
	if !path.IsAbs(static) {
		cwd, _ := os.Getwd()
		static = path.Join(cwd, static)
	}
}

func hey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "you said: %s\n", r.URL.Path[1:])
}

func localIP(s string) bool {
	return strings.HasPrefix(s, "127.") || strings.Index(s, ":") != -1
}

func myIP() string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !localIP(ipnet.String()) {
			return strings.Split(ipnet.String(), "/")[0]
		}
	}
	return ""
}

func server() {
	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir(static))))
	http.HandleFunc("/", hey)

	fmt.Printf("listening at http://%s:%d\n", myIP(), port)
	httpServer := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(httpServer, nil); err != nil {
		panic(err)
	}
}

func main() {
	setup()
	server()
}
