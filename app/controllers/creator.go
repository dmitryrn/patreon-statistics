package creator

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"patreon-statistics/app/domain/creator"
)

func GetCreators(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.Query("select id, patreon_id, created_at from creator")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("creators: failed to query rows ", err)
		return
	}

	var creators []creator.Creator

	for rows.Next() {
		crtr := creator.Creator{}
		err := rows.Scan(crtr.Id, crtr.PatreonId, crtr.CreatedAt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("creators: error in scan loop ", err)
			return
		}
		creators = append(creators, crtr)
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
