package http

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const HTTP_GET string = "GET"
const HTTP_POST string = "POST"
const HTTP_DELETE string = "DELETE"

const DELETE_204 int = 204 // No Content	Resource has been successfully deleted
const DELETE_404 int = 404 // Not Found	Specified resource does not exist
const DELETE_409 int = 409 // Conflict	Specified version incorrect

type RESTInterface interface {
	Post(data string) (string, error)

	// account_id  string, unique identifier (UUID), required
	Get(account_id string) (string, error)

	// account_id  string, unique identifier (UUID), required
	// version integer, required
	Delete(account_id string, version string) (string, error)
}

type RESTClient struct {
	BaseURL *url.URL
	//UserAgent string

	httpClient *http.Client
}

func (o *RESTClient) Get(account_id string) (string, error) {

	if o.httpClient == nil {
		o.httpClient = &http.Client{}
	}

	q := o.BaseURL.Query() // Get a copy of the query values.
	q.Add("account_id", account_id)
	o.BaseURL.RawQuery = q.Encode()

	log.Printf("%s\n", o.BaseURL.String())

	req, err := http.NewRequest(HTTP_GET, o.BaseURL.String(), nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Date", time.Now().Format(time.RFC822))
	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Content-Type", "application/vnd.api+json")

	res, err := o.httpClient.Do(req)

	if err != nil {
		log.Printf("Network error - %v\n", err)
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ERROR - HTTP Status Code: %s, Body: %v\n", fmt.Sprint(res.StatusCode), err)
		return "", err
	}

	switch {
	case res.StatusCode != 200:
		{
			log.Printf("ERROR - HTTP Status Code: %s, Body: %v\n", fmt.Sprint(res.StatusCode), string(body))
			return "", errors.New(string(body))
		}
	}

	log.Printf("OK - HTTP Status Code: %s, Body: %v\n", fmt.Sprint(res.StatusCode), string(body))
	return string(body), nil
}

func (o *RESTClient) Post(data string) (string, error) {
	o.httpClient = &http.Client{}

	req, err := http.NewRequest(HTTP_POST, o.BaseURL.String(), bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", err
	}

	req.Header.Add("Date", time.Now().Format(time.RFC822))
	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Content-Type", "application/vnd.api+json")

	res, err := o.httpClient.Do(req)

	if err != nil {
		log.Printf("Network error - %v\n", err)
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ERROR - HTTP Status Code: %s, Body: %v\n", fmt.Sprint(res.StatusCode), err)
		return "", err
	}

	switch {
	case res.StatusCode != 201:
		{
			log.Printf("ERROR - HTTP Status Code: %s, Body: %v\n", fmt.Sprint(res.StatusCode), string(body))
			return "", errors.New(string(body))
		}
	}

	log.Printf("OK - HTTP Status Code: %s, Body: %v\n", fmt.Sprint(res.StatusCode), string(body))
	return string(body), nil
}

func (o *RESTClient) Delete(account_id string, version string) (string, error) {
	o.httpClient = &http.Client{}

	o.BaseURL.Path = o.BaseURL.Path + "/" + account_id
	q := o.BaseURL.Query() // Get a copy of the query values.
	q.Add("version", version)
	o.BaseURL.RawQuery = q.Encode()

	log.Printf("%s\n", o.BaseURL.String())

	req, err := http.NewRequest(HTTP_DELETE, o.BaseURL.String(), nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Date", time.Now().Format(time.RFC1123))
	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Content-Type", "application/vnd.api+json")

	res, err := o.httpClient.Do(req)

	if err != nil {
		log.Printf("Network error - %v\n", err)
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ERROR: HTTP Status Code: %s, Body: %v\n", fmt.Sprint(res.StatusCode), err)
		return "", err
	}

	switch {
	case res.StatusCode == DELETE_204:
		{
			log.Printf("DELETE ERROR: No Content - Resource has been successfully deleted - HTTP Status Code: %s, Body: %v\n",
				fmt.Sprint(res.StatusCode), string(body))
			return "", errors.New(string(body))
		}
	case res.StatusCode == DELETE_404:
		{
			log.Printf("DELETE ERROR: Not Found	Specified resource does not exist - HTTP Status Code: %s, Body: %v\n",
				fmt.Sprint(res.StatusCode), string(body))
			return "", errors.New(string(body))
		}
	case res.StatusCode == DELETE_409:
		{
			log.Printf("DELETE ERROR: Conflict	Specified version incorrect - HTTP Status Code: %s, Body: %v\n",
				fmt.Sprint(res.StatusCode), string(body))
			return "", errors.New(string(body))
		}
	}

	log.Printf("DELETE OK - HTTP Status Code: %s, Body: %v\n", fmt.Sprint(res.StatusCode), string(body))
	return string(body), nil
}
