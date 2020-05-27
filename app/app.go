package main

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
	"log"
	http "net/http"
	DB "patreon-statistics/app/db"
	public_api "patreon-statistics/app/public-api"
	"strconv"
	"strings"
	"time"
)

func main() {
	uptimeTicker := time.NewTicker(5 * time.Second)

	db := DB.InitDbConnection()

	go startServer(db)

	update(db)
	for {
		select {
		case <-uptimeTicker.C:
			update(db)
		}
	}
}

func startServer(db *sql.DB) {
	public_api.RegisterRoutes(db)

	http.HandleFunc("/", serveFrontend)
	http.ListenAndServe(":8080", nil)
}

func serveFrontend(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "test")
}

func update(db *sql.DB) {
	client := http.Client{}
	resp, err := client.Get("https://www.patreon.com/HubertMoszka")
	defer resp.Body.Close()
	if err != nil {
		log.Println("error fetching page ", err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("failed to create doc ", err)
		return
	}

	count := doc.Find("[data-tag=\"CampaignPatronEarningStats-patron-count\"] h2").Text()
	println(count)

	earnings := doc.Find("[data-tag=\"CampaignPatronEarningStats-earnings\"] h2").Text()
	println(earnings)
	earningsInt, err := strconv.ParseInt(normalizeAmerican(removeDollarSign(earnings)), 10, 64)
	println(earningsInt)

	rows, err := db.Query("insert into statistics (creator_id, patrons_count, revenues, created_at) values ($1, $2, $3, now())", 1, count, earningsInt)
	defer rows.Close()
	if err != nil {
		log.Println("failed to insert update to statistics ", err)
		return
	}
}

func normalizeAmerican(str string) string {
	return strings.Replace(str, ",", "", -1)
}

func removeDollarSign(str string) string {
	return strings.Replace(str, "$", "", -1)
}
