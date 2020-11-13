package main

var migrate = []string{
	`
	CREATE TABLE IF NOT EXISTS permission (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name varchar(30),
		name_code varchar(15),
		description text
	);
	`, `
	CREATE TABLE IF NOT EXISTS role (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name varchar(30),
		name_code varchar(15),
		description text
	);
	`, `
	CREATE TABLE IF NOT EXISTS photo (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		upload_dt datetime,
		path text,
		size_mb,
		'format' varchar(10)
	);
	`, `
	CREATE TABLE IF NOT EXISTS category (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name varchar(30),
		name_code varchar(15),
		description text
	);
	`, `
	CREATE TABLE IF NOT EXISTS reaction (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		num_like smallint,
		num_dislike smallint
	);
	`, `
	CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title text,
		create_dt datetime,
		body text,
		photo_id bigint,
		reaction_id bigint,
		FOREIGN KEY(photo_id) REFERENCES photo(id) ON DELETE CASCADE,
		FOREIGN KEY(reaction_id) REFERENCES reaction(id) ON DELETE CASCADE
	);
	`, `
	CREATE TABLE IF NOT EXISTS user (
		id BIGINT PRIMARY KEY,
		username VARCHAR(20) UNIQUE,
		email varchar UNIQUE,
		password varchar,
		date_created datetime DEFAULT CURRENT_TIMESTAMP,
		role_id BIGINT DEFAULT 1,
		permission_id BIGINT DEFAULT 1,
		photo_id BIGINT,
	);
	`, `
	CREATE TABLE IF NOT EXISTS session (
		id text PRIMARY KEY,
		userId text,
		expiry datetime
	);
	`, `
	CREATE TABLE IF NOT EXISTS categories (
		category_id BIGINT PRIMARY KEY,
		post_id BIGINT,
		FOREIGN KEY(category_id) REFERENCES category(id) ON DELETE CASCADE,
		FOREIGN KEY(post_id) REFERENCES post(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS categories_id_index ON categories (category_id);
	`, `
	CREATE TABLE IF NOT EXISTS comment (
		user_id BIGINT PRIMARY KEY,
		post_id BIGINT,
		comment_dt datetime,
		comment text,
		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY(post_id) REFERENCES post(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS comment_id_index ON comment (user_id);
	CREATE INDEX IF NOT EXISTS comment_id_index ON comment (user_id);
	`, `
	CREATE TABLE IF NOT EXISTS user_posts (
		user_id BIGINT PRIMARY KEY,
		post_id BIGINT,
		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY(post_id) REFERENCES post(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS user_posts_id_index ON user_posts (user_id);
	`, `
	CREATE TABLE IF NOT EXISTS user_reactions (
		user_id BIGINT PRIMARY KEY,
		reaction_id BIGINT,
		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY(reaction_id) REFERENCES reaction(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS user_reactions_id_index ON user_reactions (user_id);`,
}
