package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	r.GET("/api/word", GetWordHandler)

	r.Run()
}

func GetWordHandler(c *gin.Context) {
	connStr := "postgresql://postgres:mysecretpassword@0.0.0.0/golangman_dictionary?sslmode=disable"
	var x []string
	var res string

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT word FROM words order by random() limit 1")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, "An error occured")
	}

	for rows.Next() {
		rows.Scan(&res)
		x = append(x, res)
	}

	c.JSON(http.StatusOK, x)
}
