package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", req.Method, req.URL, req.Proto)
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", req.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", req.RemoteAddr)

}