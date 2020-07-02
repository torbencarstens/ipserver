package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			_, _ = fmt.Fprintf(w, "")
		} else {
			ip := net.ParseIP(rawIP)
			_, _ = fmt.Fprintf(w, ip.String())
		}
	}))

	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
