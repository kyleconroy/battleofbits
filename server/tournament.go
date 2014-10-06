package server

import (
	"github.com/kyleconroy/battleofbits/games"
	"log"
)

func Play(db *DB, matchID string, match games.Match) error {
	// Create the match
	for {
		// Persist the match
		err := match.Tick()
		if err != nil {
			return err
		}
		id, err := db.InsertMove(matchID, match)
		if err != nil {
			return err
		}
		log.Printf("state change match=%s move=%s\n", matchID, id)
		if match.Over() {
			return nil
		}
	}
}
