package main

import (
	"net/http"
	"time"
)

func main() {

	for {
		http.Get("https://questrequest.ru/cron/cron_amo_buy_sell_search_infopart")
		time.Sleep(3 * time.Second)
	}

}
