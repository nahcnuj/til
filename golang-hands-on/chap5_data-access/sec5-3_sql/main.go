package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()

	db, err := sql.Open("sqlite3", "./db/data.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for {
		s := input("id")
		if s == "" {
			break
		}

		id, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		row := db.QueryRow("SELECT * FROM mydata WHERE id = ?", id)
		if err != nil {
			panic(err)
		}

		var d Mydata
		if err := row.Scan(&d.ID, &d.Name, &d.Mail, &d.Age); err == sql.ErrNoRows {
			continue
		} else if err != nil {
			panic(err)
		}
		fmt.Println(d.ToString())
	}
}

func input(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt + ": ")
	scanner.Scan()
	return scanner.Text()
}
