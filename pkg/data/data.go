package data

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

type LanguageKind string

const (
	ProgrammingType LanguageKind = "programming"
	DatabaseType    LanguageKind = "database"
)

type Language struct {
	Name    string   `json:"name"`
	Kind    string   `json:"kind"`
	Aliases []string `json:"aliases"`
	Words   []string `json:"words"`
}

type Store struct {
	Languages []Language `json:"languages"`
}

// embed data file in binary
//
//go:embed data.json
var bytes []byte

// where we store the embedded data
var data Store

func init() {
	if err := json.Unmarshal(bytes, &data); err != nil {
		fmt.Printf("Error loading data:\n%e", err)
	}
}

func Get() *Store {
	// unmarshal json data into Data struct
	return &data
}
