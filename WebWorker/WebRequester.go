package WebWorker

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// WebRequester does the mighty requests
type WebRequester struct {
	HTTPClient    *http.Client
	HTTPTransport *http.Transport
}

// NewWebRequester generates a new WebRequester
func NewRequester() *WebRequester {
	wr := &WebRequester{
		HTTPClient:    &http.Client{},
		HTTPTransport: &http.Transport{},
	}

	return wr
}

// PostJSON does a post request with a json Body
func (req *WebRequester) PostJSON(url string, body string) (str string, err error) {
	resp, err := req.HTTPClient.Post(url, "application/json", strings.NewReader(body))

	if nil != err {
		return "", err
	}

	defer resp.Body.Close()
	strOut, _ := ioutil.ReadAll(resp.Body)
	return string(strOut[:]), err
}

// Get does a get request
func (req *WebRequester) Get(url string) (str string, err error) {
	resp, err := req.HTTPClient.Get(url)

	if nil != err {
		return "", err
	}

	defer resp.Body.Close()
	strOut, _ := ioutil.ReadAll(resp.Body)
	return string(strOut[:]), err
}
