package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := Connect()
	if err != nil {
		fmt.Printf("%v", err)
	}

	stmt, err := db.Prepare(
		`insert into role(name, name_code, description) values($1, $2, $3); SELECT last_insert_rowid()`)
	checkErr(err)

	res, err := stmt.Exec("User", "user", "Casual user")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("id of last inserted row =", id)

	// fmt.Println("Querying")
	// rows, err := db.Query("SELECT * FROM user")
	// checkErr(err)

	// for rows.Next() {
	// 	var uid int
	// 	var username, department, created string
	// 	err = rows.Scan(&id, &username, &password, &)
	// 	checkErr(err)
	// 	fmt.Println("uid | username | department | created")
	// 	fmt.Printf("%3v | %6v | %8v | %6v\n", uid, username, department, created)
	// }
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Connect connects to a store.
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	// for _, q := range migrate {
	// 	_, err := db.Exec(q)
	// 	if err != nil {
	// 		break
	// 	}
	// }
	if err != nil {
		return nil, err
	}
	return db, nil
}

// var migrate = []string{
// 	`
// 	CREATE TABLE IF NOT EXISTS permission (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name varchar(30),
// 		name_code varchar(15),
// 		desciption text
// 	);
// 	`, `
// 	CREATE TABLE IF NOT EXISTS role (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name varchar(30),
// 		name_code varchar(15),
// 		desciption text
// 	);
// 	`, `
// 	CREATE TABLE IF NOT EXISTS photo (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		upload_dt timestamp,
// 		path text,
// 		size_mb,
// 		'format' varchar(10)
// 	);
// 	`, `
// 	CREATE TABLE IF NOT EXISTS category (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name varchar(30),
// 		name_code varchar(15),
// 		desciption text
// 	);
// 	`, `
// 	CREATE TABLE IF NOT EXISTS reaction (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		num_like smallint,
// 		num_dislike smallint
// 	);
// 	`, `
// 	CREATE TABLE IF NOT EXISTS post (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		title text,
// 		create_dt timestamp,
// 		body text,
// 		photo_id bigint,
// 		reaction_id bigint,
// 		FOREIGN KEY(photo_id) REFERENCES photo(id) ON DELETE CASCADE,
// 		FOREIGN KEY(reaction_id) REFERENCES reaction(id) ON DELETE CASCADE
// 	);
// 	`, `
// 	CREATE TABLE IF NOT EXISTS user (
// 		id BIGINT PRIMARY KEY,
// 		username VARCHAR(20) UNIQUE,
// 		email varchar UNIQUE,
// 		password varchar,
// 		date_created timestamp DEFAULT CURRENT_TIMESTAMP,
// 		role_id BIGINT DEFAULT 1,
// 		permission_id BIGINT DEFAULT 1,
// 		photo_id BIGINT,
// 	);
// 	`, `
// 	CREATE TABLE IF NOT EXISTS session (
// 		id text PRIMARY KEY,
// 		userId text,
// 		expiry time
// 	);
// 	`, `
// 	CREATE TABLE IF NOT EXISTS categories (
// 		category_id BIGINT PRIMARY KEY,
// 		post_id BIGINT,
// 		FOREIGN KEY(category_id) REFERENCES category(id) ON DELETE CASCADE,
// 		FOREIGN KEY(post_id) REFERENCES post(id) ON DELETE CASCADE
// 	);
// 	CREATE INDEX IF NOT EXISTS categories_id_index ON categories (category_id);
// 	`, `
// 	CREATE TABLE IF NOT EXISTS comment (
// 		user_id BIGINT PRIMARY KEY,
// 		post_id BIGINT,
// 		comment_dt timestamp,
// 		comment text,
// 		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE,
// 		FOREIGN KEY(post_id) REFERENCES post(id) ON DELETE CASCADE
// 	);
// 	CREATE INDEX IF NOT EXISTS comment_id_index ON comment (user_id);
// 	CREATE INDEX IF NOT EXISTS comment_id_index ON comment (user_id);
// 	`, `
// 	CREATE TABLE IF NOT EXISTS user_posts (
// 		user_id BIGINT PRIMARY KEY,
// 		post_id BIGINT,
// 		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE,
// 		FOREIGN KEY(post_id) REFERENCES post(id) ON DELETE CASCADE
// 	);
// 	CREATE INDEX IF NOT EXISTS user_posts_id_index ON user_posts (user_id);
// 	`, `
// 	CREATE TABLE IF NOT EXISTS user_reactions (
// 		user_id BIGINT PRIMARY KEY,
// 		reaction_id BIGINT,
// 		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE,
// 		FOREIGN KEY(reaction_id) REFERENCES reaction(id) ON DELETE CASCADE
// 	);
// 	CREATE INDEX IF NOT EXISTS user_reactions_id_index ON user_reactions (user_id);`,
// }
