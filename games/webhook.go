package games

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func FetchMove(url *url.URL, payload, move interface{}) error {
	blob, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}
	resp, err := http.Post(url.String(), "application/json", bytes.NewReader(blob))
	if err != nil {
		return err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("hook %s returned status %d", url, resp.StatusCode)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &move)
}
