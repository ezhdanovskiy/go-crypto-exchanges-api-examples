package cry

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func DoRequest(method, refURL string, body io.Reader, headers map[string]string) ([]byte, error) {
	rel, err := url.Parse(refURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, rel.String(), body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	//dumpRequest(req)
	resp, err := http.DefaultClient.Do(req)
	//dumpResponse(resp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func dumpRequest(r *http.Request) {
	if r == nil {
		log.Println("dumpReq ok: <nil>")
		return
	}
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("dumpReq err:", err)
	} else {
		log.Printf("dumpReq ok: %s\n\n", dump)
	}
}

func dumpResponse(r *http.Response) {
	if r == nil {
		log.Println("dumpResponse ok: <nil>")
		return
	}
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Println("dumpResponse err:", err)
	} else {
		log.Printf("dumpResponse ok: %s\n\n", dump)
	}
}
