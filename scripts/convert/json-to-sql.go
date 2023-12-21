package main

import (
	"database/sql"
	"strings"

	"github.com/snowboardit/reserved/pkg/data"

	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Data struct {
	// Define your struct fields based on the structure of your JSON data
	// For example:
	Languages []data.Language `json:"languages"`
}

func main() {
	// Read the JSON file
	filePath := "pkg/data/data.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data into a struct
	var jsonData Data
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the Postgres db
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=reserved sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows := getRows(jsonData)

	// Insert the data into the database
	query := "INSERT INTO languages (name, kind, aliases, words) VALUES " + rows
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data ingested successfully!")
}

func getRows(d Data) string {
	returner := ""
	for _, l := range d.Languages {
		returner += fmt.Sprintf("('%s', '%s', '{%s}', '{%s}'),", l.Name, l.Kind, getArray(l.Aliases), getArray(l.Words))
		fmt.Println(fmt.Sprintf("('%s', '%s', '{%s}', '{%s}'),", l.Name, l.Kind, getArray(l.Aliases), getArray(l.Words)))
	}
	returner += strings.TrimSuffix(returner, ",") // remove trailing comma
	return returner
}

func getArray(a []string) string {
	returner := ""
	for i, alias := range a {
		if len(a) == 0 {
			return ""
		}

		returner += alias
		if len(a) != i+1 {
			returner += ","
		}
	}

	return returner
}
