package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/kyleconroy/battleofbits/games"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type Tx struct {
	*sql.Tx
}

func Openenv() (*DB, error) {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return nil, fmt.Errorf("missing DATABASE_URL environment variable")
	}
	return Open(url)
}

// Open returns a DB reference for a data source.
func Open(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// Begin starts an returns a new transaction.
func (db *DB) Begin() (*Tx, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{tx}, nil
}

var insertPlayerQuery = `
INSERT INTO bots (game, name, slug, url)
  VALUES ($1, $2, $3, $4)
  RETURNING id
`

func (db *DB) InsertPlayer(game games.Game, name string, u *url.URL) (string, error) {
	var uuid string
	err := db.QueryRow(insertPlayerQuery, string(game), name,
		strings.ToLower(name), u.String()).Scan(&uuid)
	return uuid, err
}

var insertMatchQuery = `
INSERT INTO matches (player_one, player_two)
  VALUES (($1)::uuid, ($2)::uuid)
  RETURNING id::varchar(36)
`

func (db *DB) InsertMatch(one, two string) (string, error) {
	var id string
	err := db.QueryRow(insertMatchQuery, one, two).Scan(&id)
	return id, err
}

var insertMoveQuery = `
INSERT INTO moves (match, state)
  VALUES (($1)::uuid, $2)
  RETURNING id::varchar(36)
`

func (db *DB) InsertMove(match string, state interface{}) (string, error) {
	var id string
	blob, err := json.Marshal(state)
	if err != nil {
		return "", err
	}
	err = db.QueryRow(insertMoveQuery, match, blob).Scan(&id)
	return id, err
}
