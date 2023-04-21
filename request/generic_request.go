package request

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

type GenericRequest struct {
	Session     string
	Method      string
	URL         string
	Host        string
	Headers     http.Header
	QueryParams url.Values
	Body        []byte
	PathParams  map[string]string
	RemoteAddr  string
	UserAgent   string
	Token       string
	r           *http.Request
}

var HTTPXTokenHeader = "Authorization"

// NewGenericHTTPRequestFromHTTPRequest- Returns a new Generic HTTP Request from http.request
func NewGenericHTTPRequestFromHTTPRequest(r *http.Request) (*GenericRequest, error) {
	gr := new(GenericRequest)

	var err error
	gr.Headers = r.Header
	gr.URL = r.URL.String()
	gr.Host = r.Host
	gr.PathParams = mux.Vars(r)
	gr.QueryParams = r.URL.Query()
	gr.Method = r.Method
	gr.RemoteAddr = r.RemoteAddr
	gr.UserAgent = r.UserAgent()
	gr.r = r

	ua := r.Header.Get("User-Agent")
	if ua != "" {
		gr.UserAgent = ua
	}

	tokens := r.Header.Get(HTTPXTokenHeader)
	tokens = strings.TrimPrefix(tokens, "Bearer ")
	gr.Token = tokens

	reader := bufio.NewReader(r.Body)
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	gr.Body = body
	return gr, nil
}
