package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"backend/db"
	"backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := db.New()
	if err != nil {
		panic(err)
	}

	fmt.Println("Server started running...")
	http.ListenAndServe("0.0.0.0:8080", routes.Routes(db))
}
