package main 

import (
	    "http"
		"fmt"
	    "log"
)

var mychan chan []byte

func StartWebServer() (chan []byte) {
	mychan = make(chan []byte)
	//fmt.Fprintf(os.Stderr, "Make: %#v\n", mychan)

	http.Handle("/", http.HandlerFunc(QR))

	go func() {
		err := http.ListenAndServe(":80", nil)
		if err != nil {
			log.Exit("ListenAndServe:", err)
		}
	}()

	return mychan
}

func QR(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(os.Stderr, "Receive: %#v\n", mychan)

	arr := <-mychan;
	//fmt.Fprintf(os.Stdout, "%v\n", arr)
	b := Update(arr)
	fmt.Fprintf(w, b.String())
}


