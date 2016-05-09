package WebRequester

import (
	"net/http"
	"strings"
	"io/ioutil"
)

type WebRequester struct {
	HttpClient *http.Client
	HttpTransport *http.Transport
}

// generates a new WebRequester
func New() (*WebRequester) {
	wr := &WebRequester{
		HttpClient: &http.Client{},
		HttpTransport: &http.Transport{},
	}

	return wr
}

// does a post request with a json Body
func (req *WebRequester) PostJson(url string, body string) (str string, err error) {
	resp, err := req.HttpClient.Post(url, "application/json", strings.NewReader(body))

	if nil != err {
		return "", err
	}

	defer resp.Body.Close()
	strOut, _ := ioutil.ReadAll(resp.Body)
	return string(strOut[:]), err
}

// does a get request
func (req *WebRequester) Get(url string) (str string, err error) {
	resp, err := req.HttpClient.Get(url)

	if nil != err {
		return "", err
	}

	defer resp.Body.Close()
	strOut, _ := ioutil.ReadAll(resp.Body)
	return string(strOut[:]), err
}
