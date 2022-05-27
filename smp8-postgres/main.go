package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Account struct {
	Id        int
	Owner     string
	Balance   int
	Currency  string
	CreatedAt time.Time
}

func init() {
}

func main() {
	var Db *sql.DB

	// Db, err := sql.Open("postgres", "host=192.168.1.10 port=15432 user=to4ya password=postgres1048 dbname=testDB sslmode=disable")
	Db, err := sql.Open("mysql", "host=192.168.1.25 port=33060 user=go_test password=password dbname=go_database sslmode=disable")
	// Db, err := sql.Open("postgres", "postgres://to4ya:postgres1048@192.168.1.10:15432/testDB")
	if err != nil {
		log.Fatal(err)
	}
	/*
		// $1, $2 はプレースホルダー（引数に指定する実際の値の順番が$1, $2, $3...になるので順番に気をつける）
		sql := "INSERT INTO accounts (id, owner, balance, currency, created_at) VALUES ($1, $2, $3, $4, $5)"
		pstatement, err := Db.Prepare(sql)
		if err != nil {
			log.Fatal("1:>>", err)
		}
		defer pstatement.Close()

		createdat := time.Now()
		t := createdat.Format("2006/01/02 03:04:05")
		_, err = pstatement.Exec(6, "Toshi", 9086, "Hira", t)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("書き込み完了")
	*/
	sql := "SELECT id, owner, balance, currency, created_at FROM accounts WHERE id=$1;"
	pstatement, err := Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()

	queryID := 4
	var account Account

	err = pstatement.QueryRow(queryID).Scan(&account.Id,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Id, account.Owner, account.Balance, account.Currency, account.CreatedAt)
}
