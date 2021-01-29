package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/simperr"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/transport"
)

// Client is a client for communicating with the simp daemon
type Client struct {
	httpClient http.Client
}

// NewClient creates a new client for communicating with the simp daemon
func NewClient() ClientInterface {
	return &Client{
		httpClient: transport.NewUnixSocketClient(),
	}
}

func (c *Client) sendRequest(route string, body []byte, method string) ([]byte, error) {
	req, err := http.NewRequest(method, "http://unix.com/"+route, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, &simperr.SimpError{
			Code:    response.StatusCode,
			Message: response.Status,
		}
	}
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, err
}

// CreateJob sends a request to create a new job in the simp daemon
func (c *Client) CreateJob(request meta.JobRequest, id string) error {
	body, _ := json.Marshal(request)
	_, err := c.sendRequest("job/"+id, body, "POST")
	if err != nil {
		return &simperr.SimpError{
			Code:    simperr.ErrorNotFound,
			Message: err.Error(),
		}
	}
	return nil
}

// DeleteJob sends a request to delete a job in the simp daemon
func (c *Client) DeleteJob(id string) error {
	_, err := c.sendRequest("job/"+id, nil, "DELETE")
	if err != nil {
		return &simperr.SimpError{
			Code:    simperr.ErrorNotFound,
			Message: err.Error(),
		}
	}
	return nil
}

// GetJob gets a job from the simp daemon
func (c *Client) GetJob(id string) (job meta.Job, err error) {
	resp, err := c.sendRequest("job/"+id, nil, "GET")
	if err != nil {
		return meta.Job{}, &simperr.SimpError{
			Code:    simperr.ErrorNotFound,
			Message: err.Error(),
		}
	}
	json.Unmarshal(resp, &job)
	return
}

// GetExecutingJobs gets the current executing jobs from the simp daemon
func (c *Client) GetExecutingJobs() (jobs []meta.Job, err error) {
	resp, err := c.sendRequest("jobs", nil, "GET")
	if err != nil {
		return nil, &simperr.SimpError{
			Code:    simperr.ErrorNotFound,
			Message: err.Error(),
		}
	}
	json.Unmarshal(resp, &jobs)
	return
}

// GetQueuedJobs gets the queued jobs from the simp daemon
func (c *Client) GetQueuedJobs() (jobs []meta.Job, err error) {
	resp, err := c.sendRequest("queue", nil, "GET")
	if err != nil {
		return nil, &simperr.SimpError{
			Code:    simperr.ErrorNotFound,
			Message: err.Error(),
		}
	}
	json.Unmarshal(resp, &jobs)
	return
}

// DeleteQueue deletes the queued jobs from the simp daemon
func (c *Client) DeleteQueue() error {
	_, err := c.sendRequest("queue", nil, "DELETE")
	if err != nil {
		return &simperr.SimpError{
			Code:    simperr.ErrorNotFound,
			Message: err.Error(),
		}
	}
	return nil
}

// DeleteExecutingJobs deletes the current executing jobs from the simp daemon
func (c *Client) DeleteExecutingJobs() error {
	_, err := c.sendRequest("jobs", nil, "DELETE")
	if err != nil {
		return &simperr.SimpError{
			Code:    simperr.ErrorNotFound,
			Message: err.Error(),
		}
	}
	return nil
}
