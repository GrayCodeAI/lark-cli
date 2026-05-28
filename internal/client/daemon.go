package client

// Daemon represents a registered local daemon.
type Daemon struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Status     string       `json:"status"`
	AgentCount int          `json:"agent_count"`
	StartedAt  string       `json:"started_at"`
	Agents     []DaemonAgent `json:"agents"`
}

// DaemonAgent represents an agent registered through a daemon.
type DaemonAgent struct {
	Name    string `json:"name"`
	AgentID string `json:"agent_id"`
}

// ListDaemons returns all connected daemons for a workspace.
func (c *Client) ListDaemons(workspaceID string) ([]Daemon, error) {
	var daemons []Daemon
	err := c.Do("GET", "/v1/workspaces/"+workspaceID+"/daemons", nil, &daemons)
	return daemons, err
}

// GetDaemon returns details of a specific daemon.
func (c *Client) GetDaemon(daemonID string) (*Daemon, error) {
	var daemon Daemon
	err := c.Do("GET", "/v1/daemons/"+daemonID, nil, &daemon)
	return &daemon, err
}

// WakeDaemonAgent wakes a specific agent through a daemon.
func (c *Client) WakeDaemonAgent(daemonID, agentName, channelID, message string) error {
	return c.Do("POST", "/v1/daemons/"+daemonID+"/agents/"+agentName+"/wake", map[string]string{
		"channel_id": channelID,
		"message":    message,
	}, nil)
}

// RestartDaemonAgent restarts an agent through a daemon.
func (c *Client) RestartDaemonAgent(daemonID, agentName string) error {
	return c.Do("POST", "/v1/daemons/"+daemonID+"/agents/"+agentName+"/restart", nil, nil)
}
