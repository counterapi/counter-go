package counter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/counterapi/api/pkg/models"
)

// Counter is a struct for Counter communication.
type Counter struct {
	BaseURL *url.URL

	httpClient *http.Client
}

// CountOptions is options for counts.
type CountOptions struct {
	Namespace string
	Name      string
	GroupBy   string
	OrderBy   string
}

// NewCounter returns Counter struct.
func NewCounter() Counter {
	baseURL := url.URL{
		Scheme: "https",
		Host:   "api.counterapi.dev",
		Path:   "/v1/",
	}

	return Counter{
		BaseURL:    &baseURL,
		httpClient: http.DefaultClient,
	}
}

// Up counts up.
func (a *Counter) Up(namespace, name string) (models.Counter, error) {
	req, err := a.newRequest(fmt.Sprintf("%s/%s/up", namespace, name), url.Values{})
	if err != nil {
		return models.Counter{}, err
	}

	var counter models.Counter

	_, err = a.do(req, &counter)
	if err != nil {
		return models.Counter{}, err
	}

	return counter, nil
}

// Down counts down.
func (a *Counter) Down(namespace, name string) (models.Counter, error) {
	req, err := a.newRequest(fmt.Sprintf("%s/%s/down", namespace, name), url.Values{})
	if err != nil {
		return models.Counter{}, err
	}

	var counter models.Counter

	_, err = a.do(req, &counter)
	if err != nil {
		return models.Counter{}, err
	}

	return counter, nil
}

// Get fetch counter information.
func (a *Counter) Get(namespace, name string) (models.Counter, error) {
	req, err := a.newRequest(fmt.Sprintf("%s/%s", namespace, name), url.Values{})
	if err != nil {
		return models.Counter{}, err
	}

	var counter models.Counter

	_, err = a.do(req, &counter)
	if err != nil {
		return models.Counter{}, err
	}

	return counter, nil
}

// Counts fetch count list of given counter.
func (a *Counter) Counts(options CountOptions) ([]models.CountGroupResult, error) {
	params := url.Values{}
	params.Add("group_by", options.GroupBy)
	params.Add("order_by", options.OrderBy)

	req, err := a.newRequest(fmt.Sprintf("%s/%s/list", options.Namespace, options.Name), url.Values{})
	if err != nil {
		return []models.CountGroupResult{}, err
	}

	var counts []models.CountGroupResult

	_, err = a.do(req, &counts)
	if err != nil {
		return []models.CountGroupResult{}, err
	}

	return counts, nil
}

// Set sets counter.
func (a *Counter) Set(namespace, name string, count string) (models.Counter, error) {
	params := url.Values{}
	params.Add("count", count)

	req, err := a.newRequest(fmt.Sprintf("%s/%s/set", namespace, name), params)
	if err != nil {
		return models.Counter{}, err
	}

	var counter models.Counter

	_, err = a.do(req, &counter)
	if err != nil {
		return models.Counter{}, err
	}

	return counter, nil
}

// newRequest created new request.
func (a *Counter) newRequest(path string, params url.Values) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := a.BaseURL.ResolveReference(rel)

	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

// do executes request.
func (a *Counter) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)

	return resp, err
}
