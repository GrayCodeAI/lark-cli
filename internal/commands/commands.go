package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"lark-cli/internal/client"
)

type Config struct {
	client *client.Client
}

func New(cl *client.Client) *Config {
	return &Config{client: cl}
}

func (c *Config) Whoami() error {
	cfg, err := client.LoadConfig()
	if err != nil {
		return fmt.Errorf("not logged in (run 'lark-cli login')")
	}
	fmt.Println("Server:", cfg.BaseURL)
	fmt.Println("Token:", cfg.Token[:16]+"...")
	return nil
}

func (c *Config) ListWorkspaces() error {
	var workspaces []map[string]any
	if err := c.client.Do("GET", "/v1/workspaces", nil, &workspaces); err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tSlug")
	for _, ws := range workspaces {
		fmt.Fprintf(w, "%s\t%s\t%s\n", ws["id"], ws["name"], ws["slug"])
	}
	w.Flush()
	return nil
}

func (c *Config) ListChannels(workspaceID string) error {
	var channels []map[string]any
	if err := c.client.Do("GET", "/v1/workspaces/"+workspaceID+"/channels", nil, &channels); err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tType\tTopic")
	for _, ch := range channels {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", ch["id"], ch["name"], ch["type"], ch["topic"])
	}
	w.Flush()
	return nil
}

func (c *Config) ListMessages(channelID string, limit int) error {
	path := fmt.Sprintf("/v1/channels/%s/messages", channelID)
	if limit > 0 {
		path += fmt.Sprintf("?limit=%d", limit)
	}
	var messages []map[string]any
	if err := c.client.Do("GET", path, nil, &messages); err != nil {
		return err
	}
	for _, msg := range messages {
		fmt.Printf("[%s] %s: %s\n", msg["created_at"], msg["sender_id"], msg["content"])
	}
	return nil
}

func (c *Config) SendMessage(channelID, content string) error {
	body := map[string]string{"content": content}
	var msg map[string]any
	if err := c.client.Do("POST", "/v1/channels/"+channelID+"/messages", body, &msg); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(msg, "", "  ")
	fmt.Println(string(b))
	return nil
}

func (c *Config) ListTasks(workspaceID, status string) error {
	path := "/v1/workspaces/" + workspaceID + "/tasks"
	if status != "" {
		path += "?status=" + status
	}
	var tasks []map[string]any
	if err := c.client.Do("GET", path, nil, &tasks); err != nil {
		return err
	}
	for _, t := range tasks {
		fmt.Printf("[%s] %s → %s\n", t["status"], t["title"], t["assigned_to"])
	}
	return nil
}
