package helper

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gojektech/heimdall/v6/httpclient"
)

// APICall returns
type APICall struct {
	URL       string            `json:"url"`
	Method    string            `json:"method"`
	FormParam string            `json:"form_param"`
	Header    map[string]string `json:"header"`
}

// URLHttpResponse return
type URLHttpResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Body       string      `json:"body"`
	Header     http.Header `json:"header"`
}

// Call call to third party endpoint
func (ac *APICall) Call() (result URLHttpResponse, err error) {
	client := httpclient.NewClient()

	// Create an http.Request instance
	req, _ := http.NewRequest(ac.Method, ac.URL, bytes.NewBuffer([]byte(ac.FormParam)))

	req.Header.Add("Content-Type", "application/json")
	for index, value := range ac.Header {
		req.Header.Add(index, value)
	}

	// Call the `Do` method, which has a similar interface to the `http.Do` method
	res, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	result.Status = res.Status
	result.StatusCode = res.StatusCode
	result.Body = string(body)
	result.Header = res.Header

	return
}
