package server

import (
	"log"
	"net/url"

	"github.com/kyleconroy/battleofbits/games/fourup"
	"github.com/kyleconroy/migrator"
)

func Migrate() error {
	db, err := Openenv()
	if err != nil {
		return err
	}
	err = migrator.Run(db.DB, "migrations")
	if err != nil {
		return err
	}
	return nil
}

func Battle() error {
	db, err := Openenv()
	if err != nil {
		return err
	}

	alice := "b9aff487-40cc-41b7-b0ca-523ba4bfbc39"
	bob := "f6ec191b-bf30-43f4-9829-9aa43401dafc"

	one, _ := url.Parse("http://127.0.0.1:5000/fourup")
	two, _ := url.Parse("http://127.0.0.1:5000/fourup")

	log.Println(alice, bob)

	match, err := db.InsertMatch(alice, bob)

	if err != nil {
		return err
	}

	m := fourup.NewMatch(one, two)
	return Play(db, match, &m)
}
