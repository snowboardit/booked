package reserved

import (
	"fmt"
	"strings"

	"github.com/snowboardit/reserved/pkg/data"
)

type Reserved struct {
	languages []data.Language
}

type Checked map[string][]string

// create new reserved instance
func New() *Reserved {
	r := &Reserved{}
	return r
}

// Get all language names loaded
func (r *Reserved) Languages() []string {
	d := data.Get()
	languages := []string{}

	for _, l := range d.Languages {
		row := l.Name
		if len(l.Aliases) > 0 {
			row += fmt.Sprintf(" (%s)", strings.Join(l.Aliases, ", "))
		}
		languages = append(languages, row)
	}

	return languages
}

// Check if word(s) are reserved in all languages
// returns a slice of languages the word is reserved in
func (r *Reserved) Check(words ...string) Checked {
	d := data.Get()
	checked := make(map[string][]string)

	for _, l := range d.Languages {
		for _, w := range words {
			for _, lw := range l.Words {
				if lw == w {
					checked[l.Name] = append(checked[l.Name], w)
				}
			}
		}
	}
	return checked
}

// Check if word(s) are reserved in programming languages
// returns a slice of languages the word is reserved in
func (r *Reserved) CheckProgramming(words ...string) Checked {
	d := data.Get()
	result := make(map[string][]string)

	for _, l := range d.Languages {
		if l.Kind == "programming" {
			for _, w := range words {
				for _, lw := range l.Words {
					if lw == w {
						result[l.Name] = append(result[l.Name], w)
					}
				}
			}
		}
	}
	return result
}

// Check if word(s) are reserved in database languages
// returns a slice of languages the word is reserved in
func (r *Reserved) CheckDatabase(words ...string) Checked {
	d := data.Get()
	result := make(map[string][]string)

	for _, l := range d.Languages {
		if l.Kind == "database" {
			for _, w := range words {
				for _, lw := range l.Words {
					if lw == w {
						result[l.Name] = append(result[l.Name], w)
					}
				}
			}
		}
	}
	return result
}

// String output when printing checked
func (c Checked) String() string {
	var (
		index  int
		output string
	)

	for w, l := range c {
		output += fmt.Sprintf("%s\n", w)
		for i, lw := range l {
			output += "• " + lw
			if i < len(l)-1 {
				output += "\n"
			}
		}
		if index < len(c)-1 {
			output += "\n\n"
		}
		index++
	}

	return output
}
