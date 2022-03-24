package postgresql

import (
	"database/sql"
	"fmt"
	"os"
)

func InitPostgreSQL() Handler {

	database, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		`CREATE TABLE IF NOT EXISTS users (
			index		SERIAL	PRIMARY	KEY,
			sessionid	VARCHAR,
			id			VARCHAR,
			password	VARCHAR,
			name		VARCHAR,
			phone		VARCHAR,
			email		VARCHAR,
			createdAt	TIMESTAMP
			);`)

	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		`CREATE TABLE IF NOT EXISTS containers (
			index		SERIAL	PRIMARY	KEY,
			name		VARCHAR,
			id			VARCHAR
			);`)

	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		`CREATE TABLE IF NOT EXISTS files (
			name		VARCHAR,
			size		INT,
			ftype		VARCHAR,
			author		VARCHAR,
			container	VARCHAR,
			downloadurl	VARCHAR,
			createdAt	TIMESTAMP
			);`)

	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		`CREATE TABLE IF NOT EXISTS overview (
			index		SERIAL	PRIMARY	KEY,
			name		VARCHAR,
			size		INT,
			count		INT
			);`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected PostgreSQL")
	return &connectDB{db: database}
}
