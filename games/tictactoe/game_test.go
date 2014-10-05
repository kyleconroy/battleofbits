package tictactoe

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyleconroy/battleofbits/games"
)

func TestWinConditions(t *testing.T) {
	wins := []Board{
		Board{
			X, 0, 0,
			X, 0, 0,
			X, 0, 0,
		},
		Board{
			X, X, X,
			0, 0, 0,
			0, 0, 0,
		},
		Board{
			X, O, 0,
			0, X, 0,
			0, O, X,
		},
	}
	for _, win := range wins {
		g := Match{Board: win}
		if !g.Over() {
			t.Errorf("Match should be over %s", win)
		}
	}
}

func TestWebhookHandle(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload Payload
		var move Move
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		err = json.Unmarshal(body, &payload)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		space, err := random(payload.Piece, payload.Board)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		move.Space = space
		blob, err := json.MarshalIndent(move, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(blob)
	}))
	defer ts.Close()

	match := NewMatch(ts.URL, ts.URL)
	err := games.Play(&match)
	if err != nil {
		t.Error(err)
	}
	if !match.Over() {
		t.Errorf("game isn't over")
	}
}
