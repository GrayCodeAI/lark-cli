package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	DefaultServerURL = "http://localhost:4001"
	ConfigDir        = ".lark"
	ConfigFile       = "config.yaml"
)

// Config holds the CLI configuration.
type Config struct {
	ServerURL   string `json:"server_url"`
	APIToken    string `json:"api_token"`
	WorkspaceID string `json:"workspace_id"`
}

// ConfigPath returns the path to .lark/config.yaml relative to cwd.
func ConfigPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, ConfigDir, ConfigFile), nil
}

// Load reads the config from .lark/config.yaml.
func Load() (*Config, error) {
	path, err := ConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("not initialized (run 'lark init')")
		}
		return nil, err
	}

	cfg := &Config{
		ServerURL: DefaultServerURL,
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "server_url":
			cfg.ServerURL = value
		case "api_token":
			cfg.APIToken = value
		case "workspace_id":
			cfg.WorkspaceID = value
		}
	}

	return cfg, nil
}

// Save writes the config to .lark/config.yaml.
func Save(cfg *Config) error {
	path, err := ConfigPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	content := fmt.Sprintf(`# Lark CLI Configuration
server_url: %s
api_token: %s
workspace_id: %s
`, cfg.ServerURL, cfg.APIToken, cfg.WorkspaceID)

	return os.WriteFile(path, []byte(content), 0600)
}

// Delete removes the config file.
func Delete() error {
	path, err := ConfigPath()
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// Exists checks if the config file exists.
func Exists() bool {
	path, err := ConfigPath()
	if err != nil {
		return false
	}
	_, err = os.Stat(path)
	return err == nil
}
