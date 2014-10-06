package fourup

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyleconroy/battleofbits/games"
)

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

		// For testing, have red and black pick a predetermined column
		if payload.Piece == Red {
			move.Column = 1
		} else {
			move.Column = 2
		}

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
