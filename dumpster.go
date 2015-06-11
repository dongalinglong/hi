package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

var port int

func init() {
	flag.IntVar(&port, "port", 80, "Server port")
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := httputil.DumpRequest(r, true)
	check(err)
	log.Println(string(data))
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)
	hostPort := net.JoinHostPort("", strconv.Itoa(port))
	log.Println("listen on", hostPort)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
