package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	http "net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	dbName     string
	dbUserName string
	dbPassword string
	dbHost     string
	db         *sql.DB
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		switch pair[0] {
		case "PG_PASSWORD":
			dbPassword = pair[1]
		case "PG_USERNAME":
			dbUserName = pair[1]
		case "PG_DBNAME":
			dbName = pair[1]
		case "DB_HOST":
			dbName = pair[1]
		}
	}

	connStr := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbUserName, dbPassword, dbHost, dbName)
	println(connStr)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error connecting to db ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("db ping failed ", err)
	}

	uptimeTicker := time.NewTicker(5 * time.Second)

	go startServer()

	update()
	for {
		select {
		case <-uptimeTicker.C:
			update()
		}
	}
}

func startServer() {
	http.HandleFunc("/", serveApp)
	http.HandleFunc("/api/creators", getCreators)
	http.ListenAndServe(":8080", nil)
}

func serveApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "test")
}

type Creator struct {
	Name string `json:"name"`
}

func getCreators(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.Query("select patreon_id from creator")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("creators: failed to query rows ", err)
		return
	}

	var creators []Creator

	for rows.Next() {
		var patreonId string
		err := rows.Scan(&patreonId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("creators: error in scan loop ", err)
			return
		}
		creators = append(creators, Creator{patreonId})
	}
	err = rows.Err()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("creators: rows err ", err)
		return
	}

	json, err := json.Marshal(creators)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("creators: err marshal ", err)
		return
	}

	_, err = w.Write(json)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("creators: err write json response ", err)
		return
	}
}

func update() {
	client := http.Client{}
	resp, err := client.Get("https://www.patreon.com/HubertMoszka")
	if err != nil {
		log.Fatal("error fetching page ", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("failed to create doc ", err)
	}

	count := doc.Find("[data-tag=\"CampaignPatronEarningStats-patron-count\"] h2").Text()
	println(count)

	earnings := doc.Find("[data-tag=\"CampaignPatronEarningStats-earnings\"] h2").Text()
	println(earnings)
	earningsInt, err := strconv.ParseInt(normalizeAmerican(removeDollarSign(earnings)), 10, 64)
	println(earningsInt)

	_, err = db.Query("insert into statistics (creator_id, patrons_count, revenues, created_at) values ($1, $2, $3, now())", 1, count, earningsInt)
	if err != nil {
		log.Fatal("failed to insert update to statistics ", err)
	}
}

func normalizeAmerican(str string) string {
	return strings.Replace(str, ",", "", -1)
}

func removeDollarSign(str string) string {
	return strings.Replace(str, "$", "", -1)
}
