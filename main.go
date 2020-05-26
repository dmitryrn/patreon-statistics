package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	http "net/http"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("https://www.patreon.com/HubertMoszka")
	if err != nil {
		log.Fatal("error fetching page", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("failed to create doc", err)
	}

	count := doc.Find("[data-tag=\"CampaignPatronEarningStats-patron-count\"] h2").Text()
	println(count)

	earnings := doc.Find("[data-tag=\"CampaignPatronEarningStats-earnings\"] h2").Text()
	println(earnings)
}
