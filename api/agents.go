package api

// AgentRegisterRequest is a call to register on the Buildkite Agent API
type AgentRegisterRequest struct {
	Name              string   `json:"name"`
	Hostname          string   `json:"hostname"`
	OS                string   `json:"os"`
	Arch              string   `json:"arch"`
	ScriptEvalEnabled bool     `json:"script_eval_enabled"`
	Priority          string   `json:"priority,omitempty"`
	Version           string   `json:"version"`
	Build             string   `json:"build"`
	Tags              []string `json:"meta_data"`
	PID               int      `json:"pid,omitempty"`
	MachineID         string   `json:"machine_id,omitempty"`
}

// AgentRegisterResponse is the response from the Buildkite Agent API
type AgentRegisterResponse struct {
	UUID              string   `json:"id"`
	Name              string   `json:"name"`
	AccessToken       string   `json:"access_token"`
	Endpoint          string   `json:"endpoint"`
	PingInterval      int      `json:"ping_interval"`
	JobStatusInterval int      `json:"job_status_interval"`
	HeartbeatInterval int      `json:"heartbeat_interval"`
	Tags              []string `json:"meta_data"`
}

// Registers the agent against the Buildkite Agent API. The client for this
// call must be authenticated using an Agent Registration Token
func (c *Client) Register(regReq *AgentRegisterRequest) (*AgentRegisterResponse, *Response, error) {
	req, err := c.newRequest("POST", "register", regReq)
	if err != nil {
		return nil, nil, err
	}

	a := new(AgentRegisterResponse)
	resp, err := c.doRequest(req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}

// Connects the agent to the Buildkite Agent API
func (c *Client) Connect() (*Response, error) {
	req, err := c.newRequest("POST", "connect", nil)
	if err != nil {
		return nil, err
	}

	return c.doRequest(req, nil)
}

// Disconnects the agent to the Buildkite Agent API
func (c *Client) Disconnect() (*Response, error) {
	req, err := c.newRequest("POST", "disconnect", nil)
	if err != nil {
		return nil, err
	}

	return c.doRequest(req, nil)
}
