package public_api

import (
	"database/sql"
	"net/http"
	creator "patreon-statistics/app/controllers"
)

func tempPassDB(handler func(*sql.DB, http.ResponseWriter, *http.Request), db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(db, w, r)
	}
}

func RegisterRoutes(db *sql.DB) {
	http.HandleFunc("/api/creators", tempPassDB(creator.GetCreators, db))
}
