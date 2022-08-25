package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	user     string
	password string
	ip       string
	port     string
	database string
	conexion func(MySQL) *sql.DB
	cerrar   func(*sql.DB)
	ping     func(*sql.DB)
	query    func(*sql.DB, string) *sql.Rows
	exec     func(*sql.DB, string)
}

func mysql_crear(user string, password string, ip string, port string, database string) MySQL {
	var mysql = MySQL{
		user:     user,
		password: password,
		ip:       ip,
		port:     port,
		database: database,
		conexion: mysql_connect,
		cerrar:   mysql_close,
		ping:     mysql_ping,
		query:    mysql_query,
		exec:     mysql_exec,
	}
	return mysql
}

func mysql_connect(mysql MySQL) *sql.DB {
	var url string = mysql.user + ":" + mysql.password + "@tcp(" + mysql.ip + ":" + mysql.port + ")/" + mysql.database
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Pong!")
	}
	return db
}

func mysql_close(db *sql.DB) {
	db.Close()
}

func mysql_ping(db *sql.DB) {
	var err = db.Ping()
	if err != nil {
		fmt.Println("Can't ping to database: " + err.Error())
	}
}

func mysql_query(db *sql.DB, query string) *sql.Rows {
	out, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func mysql_exec(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
