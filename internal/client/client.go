package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Client struct {
	BaseURL    string
	Token      string
	httpClient *http.Client
}

type Config struct {
	BaseURL string `json:"base_url"`
	Token   string `json:"token"`
}

func configPath() string {
	dir, _ := os.UserConfigDir()
	return filepath.Join(dir, "lark-cli", "config.json")
}

func LoadConfig() (*Config, error) {
	b, err := os.ReadFile(configPath())
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func SaveConfig(cfg *Config) error {
	dir := filepath.Dir(configPath())
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath(), b, 0600)
}

func New(baseURL, token string) *Client {
	return &Client{
		BaseURL:    baseURL,
		Token:      token,
		httpClient: &http.Client{},
	}
}

func (c *Client) Do(method, path string, body, out any) error {
	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, c.BaseURL+path, reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%d: %s", resp.StatusCode, string(respBody))
	}

	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}
	return nil
}

func (c *Client) Login(host, token string) error {
	c.BaseURL = host
	c.Token = token

	var workspaces []map[string]any
	err := c.Do("GET", "/v1/workspaces", nil, &workspaces)
	if err != nil {
		return fmt.Errorf("auth failed: %w", err)
	}
	return SaveConfig(&Config{BaseURL: host, Token: token})
}
