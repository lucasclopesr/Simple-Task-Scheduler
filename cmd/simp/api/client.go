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

// Client é um cliente para comunicação com o simp daemon
type Client struct {
	httpClient http.Client
}

// NewClient cria um novo Client para comunicação com um simp daemon
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
		return nil, simperr.NewError().Code(response.StatusCode).Build()
	}
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, err
}

// CreateJob cria um job no simp daemon
func (c *Client) CreateJob(request meta.JobRequest, id string) error {
	body, _ := json.Marshal(request)
	_, err := c.sendRequest("queue/"+id, body, "POST")
	if err != nil {
		return simperr.NewError().Message(err.Error()).NotFound().Build()
	}
	return nil
}

// DeleteExecutingJob sends a request to delete a job in the simp daemon
func (c *Client) DeleteExecutingJob(id string) error {
	_, err := c.sendRequest("job/"+id, nil, "DELETE")
	if err != nil {
		return simperr.NewError().Message(err.Error()).NotFound().Build()
	}
	return nil
}

// DeleteJobFromQueue sends a request to delete a job in the simp daemon
func (c *Client) DeleteJobFromQueue(id string) error {
	_, err := c.sendRequest("queue/"+id, nil, "DELETE")
	if err != nil {
		return &simperr.SimpError{
			Code:    simperr.ErrorNotFound,
			Message: err.Error(),
		}
	}
	return nil
}

// GetExecutingJob gets a job from the simp daemon
func (c *Client) GetExecutingJob(id string) (job meta.Job, err error) {
	resp, err := c.sendRequest("job/"+id, nil, "GET")
	if err != nil {
		return job, simperr.NewError().Message(err.Error()).NotFound().Build()
	}
	json.Unmarshal(resp, &job)
	return
}

// GetJobFromQueue gets a job from the simp daemon queue
func (c *Client) GetJobFromQueue(id string) (job meta.Job, err error) {
	resp, err := c.sendRequest("queue/"+id, nil, "GET")
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
		return nil, simperr.NewError().Message(err.Error()).NotFound().Build()
	}
	json.Unmarshal(resp, &jobs)
	return
}

// GetQueuedJobs pega os jobs enfilerados no simp daemon
func (c *Client) GetQueuedJobs() (jobs []meta.Job, err error) {
	resp, err := c.sendRequest("queued", nil, "GET")
	if err != nil {
		return nil, simperr.NewError().Message(err.Error()).NotFound().Build()
	}
	json.Unmarshal(resp, &jobs)
	return
}

// DeleteQueue deleta os jobs enfilerados no simp daemon
func (c *Client) DeleteQueue() error {
	_, err := c.sendRequest("queued", nil, "DELETE")
	if err != nil {
		return simperr.NewError().Message(err.Error()).NotFound().Build()
	}
	return nil
}

// DeleteExecutingJobs deleta os jobs que estão sendo executados no simp daemon
func (c *Client) DeleteExecutingJobs() error {
	_, err := c.sendRequest("jobs", nil, "DELETE")
	if err != nil {
		return simperr.NewError().Message(err.Error()).NotFound().Build()
	}
	return nil
}
