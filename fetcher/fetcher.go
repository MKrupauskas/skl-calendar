package fetcher

import (
	"errors"
	"net/http"
)

func Fetch(url string) (*http.Response, error) {
	// TODO: send some headers that look more like a browser
	resp, err := http.Get(url)
	if err != nil {
			return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("response code not 200")
  }
	return resp, err
}
