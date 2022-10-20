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
	go func() {
		for {
			http.Get("https://questrequest.ru/cron/cron_amo_buy_sell_search_infopart")
			time.Sleep(3 * time.Second)
		}
	}()

	http.HandleFunc("/metrics", metrics)

	log.Println("Starting webserver...")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}

}
