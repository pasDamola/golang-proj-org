package main

import (
	"database/sql"
	"net/http"

	"github.com/PacktPublishing/Go-Recipes-for-Developers/src/chp1/webform/pkg/commentdb"
	_ "modernc.org/sqlite"
)

func main() {
	if db, err := sql.Open("sqlite", "webform.db"); err != nil {
		panic(err)
	}

	commentdb.InitDB(db)

	server := http.Server{
		Addr:    ":8181",
		Handler: http.FileServer(http.Dir("web/static")),
	}

	server.ListenAndServe()
}
