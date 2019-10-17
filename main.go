package main

/*
* returns value of the variable
& returns memory address of variable
*/

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

/*
Defining Constants
*/
const DBHOST = "db"
const DBUSER = "postgres-dev"
const DBPASSWORD = "mysecretpassword"
const DBNAME = "dev"
const DB_MIGRATION = `CREATE TABLE IF NOT EXISTS billboards (
id serial PRIMARY KEY,
author text NOT NULL,
content text NOT NULL,
created_at timestamp with time zone DEFAULT current_timestamp)`

/*
	DB Definitions
*/
const (
	DbHost     = DBHOST
	DbUser     = DBUSER
	DbPassword = DBPASSWORD
	DbName     = DBNAME
	Migration  = DB_MIGRATION
)

/*
	Billboard object
	Added tags to customize JSON key names and include model binding (binds request body to type)
*/
type Board struct {
	Author    string    `json:"author" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

/*
	Global DB Connection
*/
var db *sql.DB

/*
	GetBoards function which returns either a Board array or an error.
*/
func GetBoards() ([]Board, error) {
	const q = `SELECT author, content, created_at FROM billboards ORDER BY created_at DESC LIMIT 100`
	/*
		Query from above returns either a "rows" or an error
	*/
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}

	/*
		make() method creates slices, maps, and channels of the type of the first argument; dynamically sized arrays (slice returns selected elements of an array as a new array object)
		in this case, the second argument (0) specifies that length is 0
	*/
	results := make([]Board, 0)

	for rows.Next() {
		var author string
		var content string
		var createdAt time.Time
		/*
			Scanning data from the returned rows
		*/
		err = rows.Scan(&author, &content, &createdAt)
		if err != nil {
			return nil, err
		}
		/*
			Appending new result
		*/
		results = append(results, Board{author, content, createdAt})
	}

	return results, nil
}

func AddBoard(billboard Board) error {
	const q = `INSERT INTO billboards(author, content, created_at) VALUES ($1, $2, $3)`
	_, err := db.Exec(q, billboard.Author, billboard.Content, billboard.CreatedAt)
	return err
}

func main() {
	var err error
	r := gin.Default()
	//get billboards
	r.GET("/board", func(context *gin.Context) {
		results, err := GetBoards()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error: " + err.Error()})
			return
		}

		context.JSON(http.StatusOK, results)
	})

	r.POST("/board", func(context *gin.Context) {
		var b Board

		if context.Bind(&b) == nil {
			b.CreatedAt = time.Now()
			if err := AddBoard(b); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error: " + err.Error()})
				return
			}

			context.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
		context.JSON(http.StatusUnprocessableEntity, gin.H{"status": "invalid body"})
		return
	})

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName)
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
		log.Println("Failed to connect to db")
	}

	defer db.Close()

	_, err = db.Query(Migration)
	if err != nil {
		log.Println("Failed to run migrations", err.Error())
		return
	}

	log.Println("running...")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
