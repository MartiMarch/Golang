package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	user     string
	password string
	ip       string
	port     string
	database string
	conexion func(MySQL) *sql.DB
}

func mysql_crear(user string, password string, ip string, port string, database string) MySQL {
	var mysql = MySQL{
		user:     user,
		password: password,
		ip:       ip,
		port:     port,
		database: database,
	}
	return mysql
}

func mysql_conexion(mysql MySQL) *sql.DB {
	var url string = mysql.user + ":" + mysql.password + "@tcp(" + mysql.ip + ":" + mysql.port + ")/" + mysql.database
	fmt.Println(url)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}
	return db
}
