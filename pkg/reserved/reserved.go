package reserved

import (
	"fmt"
	"strings"

	"github.com/snowboardit/reserved/pkg/data"
)

type Reserved struct {
	languages map[string][]string
}

type Checked struct {
	Reserved map[string][]string
}

var s = data.GetStoreInstance()

// TODO - include language root, so we can filter by programming/database
func (r *Reserved) Load() error {
	// load languages and their words into map
	r.languages = make(map[string][]string)
	for _, l := range s.Data.Programming {
		r.languages[l.Name] = l.Words
	}
	for _, l := range s.Data.Database {
		r.languages[l.Name] = l.Words
	}
	return nil
}

// Check if word(s) are reserved in all languages
// returns a slice of languages the word is reserved in
func (r *Reserved) Check(words ...string) Checked {
	reserved := Checked{Reserved: make(map[string][]string)}

	if r.languages == nil {
		r.Load()
	}

	for _, w := range words {
		for lk, l := range r.languages {
			for _, lw := range l {
				word := strings.ToLower(w)
				languageWord := strings.ToLower(lw)
				if word == languageWord {
					reserved.Reserved[word] = append(reserved.Reserved[word], lk)
				}
			}
		}
	}
	return reserved
}

// TODO
// Check if word(s) are reserved in programming languages
func (r *Reserved) CheckProgramming(words ...string) bool {
	return false
}

// TODO
// Check if word(s) are reserved in database languages
func (r *Reserved) CheckDatabase(words ...string) bool {
	return false
}

// TODO
// Check if word(s) are reserved in stack
func (r *Reserved) CheckStack(words ...string) bool {
	return false
}

// String output when printing checked
func (c Checked) String() string {
	var (
		index  int
		output string
	)

	fmt.Println()
	for w, l := range c.Reserved {
		output += fmt.Sprintf("`%s`", w) + " is reserved in\n"
		for i, lw := range l {
			output += "- " + lw
			if i < len(l)-1 {
				output += "\n"
			}
		}
		if index < len(c.Reserved)-1 {
			output += "\n\n"
		}
		index++
	}
	return output
}
