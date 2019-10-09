package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

const (
	DbHost = "db"
	DbUser = "postgres-dev"
	DbPassword = "password1"
	DbName = "dev"
	Migration = `CREATE TABLE IF NOT EXISTS bulletin (
		id serial PRIMARY KEY,
		author text NOT NULL,
		content text NOT NULL,
		created_at timestamp with time zone DEFAULT current_timestamp
	)`
)

type Bulletin struct {
	Author string `json:"author" binding: "required"`
	Content string `json:content binding: "required"`
	CreatedAt time.Time `json:created_at"`
}

var db *sql.DB

func getBulletins() ([]Bulletin, error) {
	return nil, nil
}

func addBulletin(bulletin Bulletin) error {
	return nil
}

func main() {
	var err error

	var r = gin.Default()
	r.GET("/board", func(context *gin.Context) {
		results, err := getBulletins()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error: " + err.Error()})
			return
		}
	})
}
