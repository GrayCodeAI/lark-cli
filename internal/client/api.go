package client

import (
	"fmt"
	"net/url"

	"github.com/lark-dev/lark-cli/internal/config"
)

// NewFromConfig creates a client from the local config.
func NewFromConfig() (*Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	if cfg.APIToken == "" {
		return nil, fmt.Errorf("not authenticated (run 'lark init' to configure)")
	}
	return New(cfg.ServerURL, cfg.APIToken), nil
}

// GetStatus returns server status overview.
func (c *Client) GetStatus() (map[string]any, error) {
	var status map[string]any
	if err := c.Do("GET", "/v1/status", nil, &status); err != nil {
		return nil, err
	}
	return status, nil
}

// ListAllAgents returns all connected agents.
func (c *Client) ListAllAgents() ([]map[string]any, error) {
	var agents []map[string]any
	if err := c.Do("GET", "/v1/agents", nil, &agents); err != nil {
		return nil, err
	}
	return agents, nil
}

// ListAllChannels returns all channels (optionally filtered by workspace).
func (c *Client) ListAllChannels(workspaceID string) ([]map[string]any, error) {
	path := "/v1/channels"
	if workspaceID != "" {
		path = "/v1/workspaces/" + url.PathEscape(workspaceID) + "/channels"
	}
	var channels []map[string]any
	if err := c.Do("GET", path, nil, &channels); err != nil {
		return nil, err
	}
	return channels, nil
}

// SendMessageToChannel sends a message to a channel.
func (c *Client) SendMessageToChannel(channelID, content string) (map[string]any, error) {
	body := map[string]string{"content": content}
	var msg map[string]any
	if err := c.Do("POST", "/v1/channels/"+url.PathEscape(channelID)+"/messages", body, &msg); err != nil {
		return nil, err
	}
	return msg, nil
}
