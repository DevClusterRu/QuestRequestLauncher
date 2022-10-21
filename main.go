package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func metrics(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "alive 1\n")
}

func main() {

	conf, _ := newConfig("local.yml")

	for _, v := range conf.LINKS_COLLECTION {
		fmt.Println("Starting ", v.Uri, " per ", v.Interval)
		go func() {
			for {
				http.Get(v.Uri)
				time.Sleep(time.Duration(v.Interval) * time.Millisecond)
			}
		}()
	}

	http.HandleFunc("/metrics", metrics)

	log.Println("Starting webserver...")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}

}
