package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Handle struct {
	*sql.DB
	connection string
}

func NewDB() *Handle {
	var connStr = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, ipAddr, port, dbName)

	db, err := sql.Open("mysql", connStr)
	check(err)

	err = db.Ping()
	check(err)

	var res = &Handle{
		DB:         db,
		connection: connStr,
	}
	defer res.SetupTables()

	return res
}

func (db *Handle) SetupTables() {
	tx, err := db.Begin()
	defer tx.Commit()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("rollback...")
			fmt.Println(r)
			tx.Rollback()
		}
	}()
	check(err)
	stmt, err := tx.Prepare(`CREATE TABLE IF NOT EXISTS roster(
		session VARCHAR(200) NOT NULL,
		PRIMARY KEY (session)
	);`)
	check(err)
	_, err = stmt.Exec()
	check(err)

	stmt, err = tx.Prepare(`CREATE TABLE IF NOT EXISTS student(
		first VARCHAR(200) NOT NULL,
		last  VARCHAR(200) NOT NULL,
		grade INT(11) NOT NULL,
		session VARCHAR(200),
		PRIMARY KEY (first, last)
	);`)
	check(err)
	_, err = stmt.Exec()
	check(err)

	for roster, classList := range Roster {
		stmt, err := tx.Prepare(`REPLACE INTO roster(session) VALUES (?);`)
		check(err)
		_, err = stmt.Exec(roster)
		check(err)
		for _, student := range classList {
			stmt, err := tx.Prepare(`REPLACE INTO student
			(first, last, grade, session)
			VALUES
			(?, ?, ?, ?);`)
			check(err)
			_, err = stmt.Exec(student.first, student.last, student.grade, roster)
			check(err)
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var ipAddr = "209.97.153.160"

// var username = "root"
// var password = "RpR[+6ko?bngGy&P&Lh7CxFFDnBFyfdaLmh*CaMNkQeVKftPrD2RCX2UuUo9hDA3"
var username = "robbie"
var password = "NrP2EKk2nHGAQX2xecyXLhQvC3v2mN2oBPr4LfPCPgQAT9qch67rZLDqDnaQPiqR"
var port = "3306"
var dbName = "cchs"
