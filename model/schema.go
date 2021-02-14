package model

var migrate = `
	CREATE TABLE IF NOT EXISTS category (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name varchar(30),
		name_code varchar(15),
		description text
	);
		

	CREATE TABLE IF NOT EXISTS comment (
		id	INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id	BIGINT NOT NULL,
		user_id	BIGINT NOT NULL,
		comment_dt	timestamp DEFAULT CURRENT_TIMESTAMP,
		comment	text NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY(post_id) REFERENCES post(id) ON DELETE CASCADE
	);
	
	
	CREATE TABLE IF NOT EXISTS comment (
		id	INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id	BIGINT NOT NULL,
		user_id	BIGINT NOT NULL,
		comment_dt	timestamp DEFAULT CURRENT_TIMESTAMP,
		comment	text NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY(post_id) REFERENCES post(id) ON DELETE CASCADE
	);
	
	
	CREATE TABLE IF NOT EXISTS photo (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		upload_dt timestamp,
		path text,
		size_mb,
		'format' varchar(10)
	);
	
	
	CREATE TABLE IF NOT EXISTS post (
		id	INTEGER,
		user_id	text DEFAULT '',
		title	text DEFAULT '',
		text	text DEFAULT '',
		created_at	timestamp NOT NULL DEFAULT 'datetime()',
		like_count	integer DEFAULT 0,
		dislike_count	integer DEFAULT 0,
		comment_count	integer DEFAULT 0,
		photo_id	text DEFAULT '',
		PRIMARY KEY(id AUTOINCREMENT),
		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE
	);
	
	
	CREATE TABLE IF NOT EXISTS post (
		id	INTEGER,
		user_id	text DEFAULT '',
		title	text DEFAULT '',
		text	text DEFAULT '',
		created_at	timestamp NOT NULL DEFAULT 'datetime()',
		like_count	integer DEFAULT 0,
		dislike_count	integer DEFAULT 0,
		comment_count	integer DEFAULT 0,
		photo_id	text DEFAULT '',
		PRIMARY KEY(id AUTOINCREMENT),
		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE
	);
	
	
	CREATE TABLE IF NOT EXISTS reaction (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		num_like smallint,
		num_dislike smallint
	);
	
	
	CREATE TABLE IF NOT EXISTS reaction (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		num_like smallint,
		num_dislike smallint
	);
	
	
	CREATE TABLE IF NOT EXISTS session (
		id text PRIMARY KEY,
		userId text,
		expiry datetime
	);
	
	
	CREATE TABLE IF NOT EXISTS user (
		id	BIGINT,
		username	VARCHAR(20) UNIQUE,
		email	varchar UNIQUE,
		password	varchar,
		date_created	timestamp DEFAULT CURRENT_TIMESTAMP,
		role_id	BIGINT DEFAULT 1,
		permission_id	BIGINT DEFAULT 1,
		photo_id	BIGINT
	);
	
	
	CREATE TABLE IF NOT EXISTS user_posts (
		user_id	BIGINT,
		post_id	BIGINT,
		FOREIGN KEY(post_id) REFERENCES post(id) ON DELETE CASCADE
	);
	
	
	CREATE TABLE IF NOT EXISTS user_posts (
		user_id	BIGINT,
		post_id	BIGINT,
		FOREIGN KEY(post_id) REFERENCES post(id) ON DELETE CASCADE
	);
	
	
	CREATE TABLE IF NOT EXISTS userinfo (
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		username text,
		department text,
		created timestamp
	);
	`

var drop = []string{
	`DROP TABLE IF EXISTS permission;`,
	`DROP TABLE IF EXISTS "role";`,
	`DROP TABLE IF EXISTS photo;`,
	`DROP TABLE IF EXISTS category;`,
	`DROP TABLE IF EXISTS reaction;`,
	`DROP TABLE IF EXISTS post;`,
	`DROP TABLE IF EXISTS user;`,
	`DROP TABLE IF EXISTS categories;`,
	`DROP TABLE IF EXISTS comment;`,
	`DROP TABLE IF EXISTS user_posts;`,
	`DROP TABLE IF EXISTS user_reactions;`,
}

var createUsers = []string{
	`INSERT INTO role(name, name_code, description) values("Admin", "admin", "Superuser Administrator");`,
	`INSERT INTO role(name, name_code, description) values("Moderator", "moderator", "Site moderator to control,read,write the posts");`,
	`INSERT INTO role(name, name_code, description) values("User", "user", "Normal user. Can only read and rite posts");`,
	`INSERT INTO permission(name, name_code, description) values("Admin", "admin", "Allowed to: control users and post");`,
	`INSERT INTO permission(name, name_code, description) values("Moderator", "moderator", "Site moderator to control,read,write the posts");`,
	`INSERT INTO permission(name, name_code, description) values("User", "user", "Normal user. Can only read and rite posts");`,
}
