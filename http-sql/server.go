package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type User struct {
	id        int
	username  string
	email     string
	password  string
	createdAt time.Time
}

func main() {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testeGolang?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected!!")

	SelectAll()
}

func CreateTableAndSeed() {
	db.Exec(`CREATE TABLE users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );`)
	db.Exec(`INSERT INTO users (username, email, password)
    VALUES
    ('john_doe', 'john@example.com', 'password123'),
    ('jane_smith', 'jane@example.com', 'securepassword'),
    ('alex_jones', 'alex@example.com', 'letmein');`)
}

func SelectAll() {
	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		u := User{}
		if err := rows.Scan(&u.id, &u.username, &u.email, &u.password, &u.createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(u)
	}
}
