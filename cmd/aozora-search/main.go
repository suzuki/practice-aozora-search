package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Entry struct {
	AuthorID string
	Author   string
	TitleID  string
	Title    string
	SiteURL  string
	ZipURL   string
}

func showAuthors(db *sql.DB) error {
	rows, err := db.Query(`SELECT author_id, author FROM authors ORDER BY author_id`)
	if err != nil {
		return err
	}
	defer rows.Close()

	entries := []Entry{}
	for rows.Next() {
		//entry := new(Entry)
		entry := Entry{}
		err := rows.Scan(&entry.AuthorID, &entry.Author)
		if err != nil {
			continue
		}
		entries = append(entries, entry)
	}

	fmt.Println("AuthorID\tAuthor")
	for _, entry := range entries {
		fmt.Printf("%s\t%s\n", entry.AuthorID, entry.Author)
	}

	return nil
}

// func showTitles(db *sql.DB, authorID string) error {

// 	return nil
// }

// func showContent(db *sql.DB, authorID, titleID string) error {

// 	return nil
// }

// func queryContent(db *sql.DB, a string) error {

// 	return nil
// }

func main() {
	var dsn string
	flag.StringVar(&dsn, "d", "database.sqlite", "database")
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch flag.Arg(0) {
	case "authors":
		err = showAuthors(db)

		// case "titles":
		// 	if flag.NArg() != 2 {
		// 		flag.Usage()
		// 		os.Exit(2)
		// 	}
		// 	err = showTitles(db, flag.Arg(1))

		// case "content":
		// 	if flag.NArg() != 3 {
		// 		flag.Usage()
		// 		os.Exit(2)
		// 	}
		// 	err = showContent(db, flag.Arg(1), flag.Arg(2))

		// case "query":
		// 	if flag.NArg() != 2 {
		// 		flag.Usage()
		// 		os.Exit(2)
		// 	}
		// 	err = queryContent(db, flag.Arg(1))
	}

	if err != nil {
		log.Fatal(err)
	}
}
