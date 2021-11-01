package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"errors"

	test "github.com/rosfran/interview-accountapi/test_assertions"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func (c Client) UpperCase(word string) (string, error) {
	res, err := http.Get(c.url + "/upper?word=" + word)
	if err != nil {
		return "", errors.Unwrap(err)
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Unwrap(err)
	}

	return string(out), nil
}

func TestClientUpperCase(t *testing.T) {
	expected := "dummy data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()
	c := NewClient(svr.URL)
	res, err := c.UpperCase("Heap32")
	if err != nil {
		t.Errorf("Err should be nil, but got %v.", err)
	}
	// res: expected\r\n
	// due to the http protocol cleanup response
	res = strings.TrimSpace(res)
	if res != expected {
		t.Errorf("RES should be %s, got %s.", expected, res)
	}
}

func ExampleResponseRecorder() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello GoLang!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Output:
	// 200
	// text/html; charset=utf-8
	// <html><body>Hello World!</body></html>
}

func TestPostRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		fmt.Printf("%s", test.AssertEqual(t, "/v1/organisation/accounts", req.URL.String()))
		test.AssertNotEmpty(t, req.Host)
		test.AssertNotEmpty(t, req.Header.Get("Date"))
		test.AssertEqual(t, "application/vnd.api+json", req.Header.Get("Accept"))
		test.AssertEqual(t, "application/vnd.api+json", req.Header.Get("Content-Type"))

		res.WriteHeader(201)
		res.Write([]byte(`my-response`))
	}))
	defer server.Close()

	b, _ := url.Parse(server.URL + "/v1/organisation/accounts")
	restClient := &RESTClient{
		BaseURL:    b,
		httpClient: &http.Client{},
	}

	body, err := restClient.Post("my-response")
	test.AssertNoError(t, err)
	test.AssertEqual(t, "my-response", body)
}

func TestPostRequest2(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(400)
		res.Write([]byte(`Bad Request`))
	}))
	defer server.Close()

	b, _ := url.Parse(server.URL)
	restClient := &RESTClient{
		BaseURL:    b,
		httpClient: &http.Client{},
	}

	_, err := restClient.Post("form3...body")

	test.AssertError(t, err)
}
