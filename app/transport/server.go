package transport

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func StartServer(db *sql.DB) {
	RegisterRoutes(db)

	http.HandleFunc("/", serveFrontend)

	err := http.ListenAndServe(":8080", nil)
	log.Fatal("failed to start http server ", err)
}

func serveFrontend(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "test")
}
