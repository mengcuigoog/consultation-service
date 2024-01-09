package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func demo() {
	os.Remove("./feedback.db")

	db, err := sql.Open("sqlite3", "./data/db/feedback.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE  if not exists user_feedback (id integer NOT NULL PRIMARY KEY, user_email TEXT, user_feedback TEXT, feedback_time TEXT);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	curTime := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("INSERT INTO user_feedback(user_email, user_feedback, feedback_time) values('564398053@qq.com', '吐槽一下', '%s')", curTime)
	log.Println(sql)
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}
