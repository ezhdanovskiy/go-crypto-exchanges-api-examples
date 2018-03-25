package cry

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func DoRequest(method, refURL string) ([]byte, error) {
	rel, err := url.Parse(refURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, rel.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
