package mpv

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"
)

// Response received from mpv. Can be an event or a user requested response.
type Response struct {
	Err       string      `json:"error"`
	Data      interface{} `json:"data"` // May contain float64, bool or string
	Event     string      `json:"event"`
	RequestID int         `json:"request_id"`
	Name      string      `json:"name"`
}

// request sent to mpv. Includes request_id for mapping the response.
type request struct {
	Command   []interface{}  `json:"command"`
	RequestID int            `json:"request_id"`
	Response  chan *Response `json:"-"`
}

// Client is a more comfortable higher level interface
// to LLClient. It can use any LLClient implementation.
type Client struct {
	socket          string
	timeout         time.Duration
	requests        chan *request
	mu              sync.Mutex
	reqMap          map[int]*request // Maps RequestIDs to Requests for response association
	observeFloatCB  map[string]func(value float64)
	observeStringCB map[string]func(value string)
}

// NewClient creates a new highlevel client based on a lowlevel client.
func NewClient(socket string) (*Client, error) {
	conn, err := net.Dial("unix", socket)

	if err != nil {
		return nil, err
	}

	client := &Client{
		socket:   socket,
		timeout:  2 * time.Second,
		requests: make(chan *request),
		reqMap:   make(map[int]*request),
	}

	go client.receiveLoop(conn)
	go client.sendLoop(conn)

	return client, nil
}

func (c *Client) sendLoop(conn io.Writer) {
	for req := range c.requests {
		b, err := json.Marshal(req)
		if err != nil {
			continue
		}
		c.mu.Lock()
		c.reqMap[req.RequestID] = req
		c.mu.Unlock()
		b = append(b, '\n')
		_, _ = conn.Write(b)
	}
}

func (c *Client) receiveLoop(conn io.Reader) {
	rd := bufio.NewReader(conn)
	for {
		data, err := rd.ReadBytes('\n')
		if err != nil {
			continue
		}
		var response Response
		err = json.Unmarshal(data, &response)
		if err != nil {
			continue
		}
		c.dispatch(&response)
	}
}

// dispatch dispatches responses to the corresponding request
func (c *Client) dispatch(resp *Response) {
	if resp.Event == "" { // No Event
		c.mu.Lock()
		defer c.mu.Unlock()
		if request, ok := c.reqMap[resp.RequestID]; ok { // Lookup requestID in request map
			delete(c.reqMap, resp.RequestID)
			request.Response <- resp
			return
		}
	} else if resp.Event == "property-change" {
		//resp.

		// TODO: Implement Event support
	}
}

func (c *Client) ObserveProperty() {
	//c.Exec()
}

var (
	ErrTimeoutSend = errors.New("timeout while sending command")
	ErrTimeoutRecv = errors.New("timeout while receiving response")
)

func (c *Client) Exec(command ...interface{}) (*Response, error) {
	req := &request{
		Command:   command,
		RequestID: rand.Intn(10000),
		Response:  make(chan *Response, 1),
	}

	select {
	case c.requests <- req:
	case <-time.After(c.timeout):
		return nil, ErrTimeoutSend
	}

	select {
	case res, ok := <-req.Response:
		if !ok {
			panic("Response channel closed")
		}
		return res, nil
	case <-time.After(c.timeout):
		return nil, ErrTimeoutRecv
	}
}
