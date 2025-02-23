package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/steveyiyo/url-shortener/internal/tools"
)

// Create file and table
func CreateTable() {
	// Create Table
	db, err := sql.Open("sqlite3", "./data.db")
	tools.ErrCheck(err)
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS urlinfo (ShortID string, Link string, Expireat string);")
	tools.ErrCheck(err)
	stmt.Exec()
	tools.ErrCheck(err)
	db.Close()
}

// Add data to DB
func AddData(ShortID string, Link string, ExpireAt int64) {
	db, err := sql.Open("sqlite3", "./data.db")
	tools.ErrCheck(err)
	stmt, err := db.Prepare("INSERT INTO urlinfo(ShortID, Link, Expireat) values(?,?,?)")
	tools.ErrCheck(err)
	res, err := stmt.Exec(ShortID, Link, ExpireAt)
	tools.ErrCheck(err)
	res.LastInsertId()
	db.Close()
}

// Get data from DB
func QueryData(ID string) (bool, string) {
	db, err := sql.Open("sqlite3", "./data.db")
	tools.ErrCheck(err)
	now := time.Now().Unix()
	rows, err := db.Query("SELECT * FROM urlinfo WHERE ShortID = ?", ID)
	tools.ErrCheck(err)
	status := false
	URL := ""
	for rows.Next() {
		var Link string
		var Expireat int64
		var ShortLink string
		err = rows.Scan(&ShortLink, &Link, &Expireat)
		tools.ErrCheck(err)
		if Link != "" {
			URL = Link
			if Expireat > now {
				status = true
			} else {
				status = false
			}
		} else {
			URL = ""
			status = false
		}
	}
	return status, URL
}
