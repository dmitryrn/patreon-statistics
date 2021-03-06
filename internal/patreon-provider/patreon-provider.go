package patreon_provider

import (
	"database/sql"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func update(db *sql.DB) {
	client := http.Client{}
	resp, err := client.Get("https://www.patreon.com/HubertMoszka")
	if err != nil {
		log.Println("error fetching page ", err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("failed to create doc ", err)
		return
	}

	count := doc.Find("[data-tag=\"CampaignPatronEarningStats-patron-count\"] h2").Text()
	earnings := doc.Find("[data-tag=\"CampaignPatronEarningStats-earnings\"] h2").Text()

	println(count, earnings)
	earningsInt, err := strconv.ParseInt(normalizeAmerican(removeDollarSign(earnings)), 10, 64)

	rows, err := db.Query("insert into statistics (creator_id, patrons_count, revenues, created_at) values ($1, $2, $3, now())", 1, count, earningsInt)
	if err != nil {
		log.Println("failed to insert update to statistics ", err)
		return
	}
	defer rows.Close()
}

func normalizeAmerican(str string) string {
	return strings.Replace(str, ",", "", -1)
}

func removeDollarSign(str string) string {
	return strings.Replace(str, "$", "", -1)
}
