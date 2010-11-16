package main

import (
	    "flag"
	    "http"
		"os"
		"fmt"
	    "log"
		"./cvis"
)

var addr = flag.String("addr", ":80", "http service address") // Q=17, R=18

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Exit("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {

	arr := make([]byte,196)
	f, err := os.Open("/dev/urandom", os.O_RDONLY, 0)
	if err != nil {
		os.Exit(1);
	}

	_, err = f.Read(arr)
	if err != nil {
		os.Exit(1);
	}
	b := cvis.Update(arr)
	fmt.Fprintf(w, b.String())
}


