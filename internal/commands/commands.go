package commands

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"

	"github.com/lark-dev/lark-cli/internal/client"
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
	display := cfg.Token
	if len(display) > 16 {
		display = display[:16] + "..."
	}
	fmt.Println("Token:", display)
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
	if err := c.client.Do("GET", "/v1/workspaces/"+url.PathEscape(workspaceID)+"/channels", nil, &channels); err != nil {
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
	path := fmt.Sprintf("/v1/channels/%s/messages", url.PathEscape(channelID))
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
	if err := c.client.Do("POST", "/v1/channels/"+url.PathEscape(channelID)+"/messages", body, &msg); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(msg, "", "  ")
	fmt.Println(string(b))
	return nil
}

func (c *Config) SearchMessages(query, channelID string) error {
	path := "/v1/search/messages?q=" + url.QueryEscape(query)
	if channelID != "" {
		path += "&channel_id=" + url.QueryEscape(channelID)
	}
	var messages []map[string]any
	if err := c.client.Do("GET", path, nil, &messages); err != nil {
		return err
	}
	if len(messages) == 0 {
		fmt.Println("no results")
		return nil
	}
	for _, msg := range messages {
		fmt.Printf("[%s] %s: %s\n", msg["created_at"], msg["sender_id"], msg["content"])
	}
	return nil
}

func (c *Config) ListTasks(workspaceID, status string) error {
	path := "/v1/workspaces/" + url.PathEscape(workspaceID) + "/tasks"
	if status != "" {
		path += "?status=" + url.QueryEscape(status)
	}
	var tasks []map[string]any
	if err := c.client.Do("GET", path, nil, &tasks); err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTitle\tStatus\tAssigned")
	for _, t := range tasks {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", t["id"], t["title"], t["status"], t["assigned_to"])
	}
	w.Flush()
	return nil
}

func (c *Config) ListAgents() error {
	var agents []map[string]any
	if err := c.client.Do("GET", "/v1/admin/agents", nil, &agents); err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tWorkspace\tOnline")
	for _, a := range agents {
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\n", a["id"], a["name"], a["workspace"], a["online"])
	}
	w.Flush()
	return nil
}

func (c *Config) ListFiles(workspaceID string) error {
	var files []map[string]any
	if err := c.client.Do("GET", "/v1/workspaces/"+url.PathEscape(workspaceID)+"/files", nil, &files); err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tSize\tType")
	for _, f := range files {
		fmt.Fprintf(w, "%s\t%s\t%v\t%s\n", f["id"], f["filename"], f["size"], f["content_type"])
	}
	w.Flush()
	return nil
}

func (c *Config) ListPins(channelID string) error {
	var pins []map[string]any
	if err := c.client.Do("GET", "/v1/channels/"+url.PathEscape(channelID)+"/pins", nil, &pins); err != nil {
		return err
	}
	for _, p := range pins {
		fmt.Printf("[%s] pinned message %s\n", p["created_at"], p["message_id"])
	}
	return nil
}

func (c *Config) ListApprovals(workspaceID string) error {
	var approvals []map[string]any
	if err := c.client.Do("GET", "/v1/workspaces/"+url.PathEscape(workspaceID)+"/approvals", nil, &approvals); err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTitle\tStatus\tRequester")
	for _, a := range approvals {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", a["id"], a["title"], a["status"], a["requester_id"])
	}
	w.Flush()
	return nil
}

func (c *Config) ListDMs() error {
	var dms []map[string]any
	if err := c.client.Do("GET", "/v1/users/me/dms", nil, &dms); err != nil {
		return err
	}
	for _, dm := range dms {
		fmt.Printf("%s (%s)\n", dm["name"], dm["id"])
	}
	return nil
}

func (c *Config) ShowUnread() error {
	var unread map[string]any
	if err := c.client.Do("GET", "/v1/users/me/unread", nil, &unread); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(unread, "", "  ")
	fmt.Println(string(b))
	return nil
}

func (c *Config) EditMessage(messageID, content string) error {
	body := map[string]string{"content": content}
	var msg map[string]any
	if err := c.client.Do("PATCH", "/v1/messages/"+url.PathEscape(messageID), body, &msg); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(msg, "", "  ")
	fmt.Println(string(b))
	return nil
}

func (c *Config) GetThread(messageID string) error {
	var thread []map[string]any
	if err := c.client.Do("GET", "/v1/messages/"+url.PathEscape(messageID)+"/thread", nil, &thread); err != nil {
		return err
	}
	for _, msg := range thread {
		fmt.Printf("[%s] %s: %s\n", msg["created_at"], msg["sender_id"], msg["content"])
	}
	return nil
}

func (c *Config) ReplyThread(messageID, channelID, content string) error {
	body := map[string]string{"content": content, "thread_id": messageID}
	var msg map[string]any
	if err := c.client.Do("POST", "/v1/channels/"+url.PathEscape(channelID)+"/messages", body, &msg); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(msg, "", "  ")
	fmt.Println(string(b))
	return nil
}

func (c *Config) ListNotifications(unreadOnly bool) error {
	path := "/v1/notifications"
	if unreadOnly {
		path += "?unread=true"
	}
	var notifications []map[string]any
	if err := c.client.Do("GET", path, nil, &notifications); err != nil {
		return err
	}
	if len(notifications) == 0 {
		fmt.Println("No notifications")
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tType\tTitle\tRead")
	for _, n := range notifications {
		read := "no"
		if n["is_read"].(bool) {
			read = "yes"
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", n["id"], n["type"], n["title"], read)
	}
	w.Flush()
	return nil
}

func (c *Config) MarkNotificationRead(id string) error {
	return c.client.Do("PATCH", "/v1/notifications/"+url.PathEscape(id)+"/read", nil, nil)
}

func (c *Config) MarkAllNotificationsRead() error {
	return c.client.Do("POST", "/v1/notifications/mark-all-read", nil, nil)
}

func (c *Config) ListIntegrations() error {
	var integrations []map[string]any
	if err := c.client.Do("GET", "/v1/integrations", nil, &integrations); err != nil {
		return err
	}
	if len(integrations) == 0 {
		fmt.Println("No integrations available")
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tType\tDescription")
	for _, i := range integrations {
		desc := ""
		if d, ok := i["description"].(string); ok {
			if len(d) > 40 {
				d = d[:37] + "..."
			}
			desc = d
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", i["id"], i["name"], i["type"], desc)
	}
	w.Flush()
	return nil
}

func (c *Config) ListWorkspaceIntegrations(workspaceID string) error {
	var integrations []map[string]any
	if err := c.client.Do("GET", "/v1/workspaces/"+url.PathEscape(workspaceID)+"/integrations", nil, &integrations); err != nil {
		return err
	}
	if len(integrations) == 0 {
		fmt.Println("No integrations installed")
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tIntegration ID\tEnabled")
	for _, i := range integrations {
		enabled := "no"
		if e, ok := i["enabled"].(bool); ok && e {
			enabled = "yes"
		}
		fmt.Fprintf(w, "%s\t%s\t%s\n", i["id"], i["integration_id"], enabled)
	}
	w.Flush()
	return nil
}

func (c *Config) ListCalls(limit int) error {
	path := "/v1/calls"
	if limit > 0 {
		path += fmt.Sprintf("?limit=%d", limit)
	}
	var calls []map[string]any
	if err := c.client.Do("GET", path, nil, &calls); err != nil {
		return err
	}
	if len(calls) == 0 {
		fmt.Println("No calls")
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tType\tStatus\tCaller\tCallee")
	for _, call := range calls {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", call["id"], call["type"], call["status"], call["caller_id"], call["callee_id"])
	}
	w.Flush()
	return nil
}

func (c *Config) ListWorkflows(workspaceID string) error {
	var workflows []map[string]any
	if err := c.client.Do("GET", "/v1/workspaces/"+url.PathEscape(workspaceID)+"/workflows", nil, &workflows); err != nil {
		return err
	}
	if len(workflows) == 0 {
		fmt.Println("No workflows")
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tTrigger\tEnabled")
	for _, wf := range workflows {
		enabled := "no"
		if e, ok := wf["enabled"].(bool); ok && e {
			enabled = "yes"
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", wf["id"], wf["name"], wf["trigger_type"], enabled)
	}
	w.Flush()
	return nil
}

func (c *Config) TriggerWorkflow(workflowID string) error {
	var run map[string]any
	if err := c.client.Do("POST", "/v1/workflows/"+url.PathEscape(workflowID)+"/trigger", nil, &run); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(run, "", "  ")
	fmt.Println(string(b))
	return nil
}

func (c *Config) GetBilling(workspaceID string) error {
	var billing map[string]any
	if err := c.client.Do("GET", "/v1/workspaces/"+url.PathEscape(workspaceID)+"/billing", nil, &billing); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(billing, "", "  ")
	fmt.Println(string(b))
	return nil
}

func (c *Config) GetUsage(workspaceID string) error {
	var usage map[string]any
	if err := c.client.Do("GET", "/v1/workspaces/"+url.PathEscape(workspaceID)+"/usage", nil, &usage); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(usage, "", "  ")
	fmt.Println(string(b))
	return nil
}
