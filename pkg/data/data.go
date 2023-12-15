package data

import (
	_ "embed"
	"encoding/json"
	"sync"
)

type Language struct {
	Name    string   `json:"name"`
	Aliases []string `json:"aliases"`
	Words   []string `json:"words"`
}

type Languages struct {
	Programming []Language `json:"programming"`
	Database    []Language `json:"database"`
}

type Store struct {
	Data Languages
}

//go:embed data.json
var bytes []byte

var (
	storeInstance *Store
	once          sync.Once
)

func GetStoreInstance() *Store {
	// ensure singleton
	once.Do(func() {
		storeInstance = &Store{}
		err := storeInstance.Load()
		if err != nil {
			panic(err)
		}
	})
	return storeInstance
}

func (s *Store) Load() error {
	// unmarshal json data into Data struct
	err := json.Unmarshal(bytes, &s.Data)
	if err != nil {
		return err
	}
	return nil
}
