package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const DB_PATH = "../forum.db"

func main() {
	db, err := sql.Open("sqlite3", DB_PATH)
	checkErr(err)
	defer db.Close()

	fmt.Println("New")
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS userinfo (
			uid INTEGER PRIMARY KEY AUTOINCREMENT,
			username text,
			department text,
			created timestamp
		);`)
	checkErr(err)
	res, err := stmt.Exec()
	checkErr(err)

	fmt.Println("Inserting")
	stmt, err = db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	checkErr(err)

	res, err = stmt.Exec("astaxie", "software developement", time.Now())
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("id of last inserted row =", id)

	fmt.Println("Updating")
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "row(s) changed")

	fmt.Println("Querying")
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username, department, created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)

		created = "2020-07-15 03:33:27.535094531+05:00"
		parsed, err := time.Parse("2006-01-02 03:04:05-07:00", created)
		checkErr(err)
		fmt.Println(parsed.String())
		fmt.Println("uid | username | department | created")
		fmt.Printf("%3v | %6v | %8v | %6v\n", uid, username, department, created)
	}

	fmt.Println("Deleting")
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "row(s) changed")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
