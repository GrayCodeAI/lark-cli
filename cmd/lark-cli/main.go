package main

import (
	"fmt"
	"os"

	"github.com/lark-dev/lark-cli/internal/client"
	"github.com/lark-dev/lark-cli/internal/commands"
	"github.com/spf13/cobra"
)

var version = "dev"

var rootCmd = &cobra.Command{
	Use:     "lark-cli",
	Short:   "CLI client for the Lark messaging platform",
	Long:    "Command-line interface for interacting with Lark agent-native messaging platform.",
	Version: version,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// loadClient loads config and returns a configured client, or exits if not logged in.
func loadClient() (*client.Client, error) {
	cfg, err := client.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("not logged in (run 'lark-cli login')")
	}
	return client.New(cfg.BaseURL, cfg.Token), nil
}

func init() {
	// Auth
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(whoamiCmd)

	// Resources
	rootCmd.AddCommand(workspacesCmd)
	rootCmd.AddCommand(channelsCmd)
	rootCmd.AddCommand(messagesCmd)
	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(tasksCmd)
	rootCmd.AddCommand(agentsCmd)
	rootCmd.AddCommand(filesCmd)
	rootCmd.AddCommand(pinsCmd)
	rootCmd.AddCommand(approvalsCmd)
	rootCmd.AddCommand(dmCmd)
	rootCmd.AddCommand(unreadCmd)
	rootCmd.AddCommand(editMessageCmd)
	rootCmd.AddCommand(threadCmd)
	rootCmd.AddCommand(replyCmd)
	rootCmd.AddCommand(notificationsCmd)
	rootCmd.AddCommand(markReadCmd)
	rootCmd.AddCommand(markAllReadCmd)
	rootCmd.AddCommand(integrationsCmd)
	rootCmd.AddCommand(wsIntegrationsCmd)
	rootCmd.AddCommand(callsCmd)
	rootCmd.AddCommand(workflowsCmd)
	rootCmd.AddCommand(triggerWorkflowCmd)
	rootCmd.AddCommand(billingCmd)
	rootCmd.AddCommand(usageCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with Lark server",
	RunE: func(cmd *cobra.Command, args []string) error {
		host, _ := cmd.Flags().GetString("host")
		token, _ := cmd.Flags().GetString("token")
		if token == "" {
			return fmt.Errorf("--token is required")
		}
		cl := client.New(host, token)
		if err := cl.Login(host, token); err != nil {
			return err
		}
		fmt.Println("logged in to", host)
		return nil
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove saved credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := client.DeleteConfig(); err != nil {
			return err
		}
		fmt.Println("logged out")
		return nil
	},
}

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show current auth info",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		return commands.New(cl).Whoami()
	},
}

var workspacesCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "List workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		return commands.New(cl).ListWorkspaces()
	},
}

var channelsCmd = &cobra.Command{
	Use:   "channels",
	Short: "List channels in a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ws, _ := cmd.Flags().GetString("workspace")
		if ws == "" {
			return fmt.Errorf("--workspace is required")
		}
		return commands.New(cl).ListChannels(ws)
	},
}

var messagesCmd = &cobra.Command{
	Use:   "messages",
	Short: "List messages in a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ch, _ := cmd.Flags().GetString("channel")
		if ch == "" {
			return fmt.Errorf("--channel is required")
		}
		limit, _ := cmd.Flags().GetInt("limit")
		return commands.New(cl).ListMessages(ch, limit)
	},
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ch, _ := cmd.Flags().GetString("channel")
		content, _ := cmd.Flags().GetString("content")
		if ch == "" || content == "" {
			return fmt.Errorf("--channel and --content are required")
		}
		return commands.New(cl).SendMessage(ch, content)
	},
}

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search messages",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ch, _ := cmd.Flags().GetString("channel")
		return commands.New(cl).SearchMessages(args[0], ch)
	},
}

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List tasks in a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ws, _ := cmd.Flags().GetString("workspace")
		if ws == "" {
			return fmt.Errorf("--workspace is required")
		}
		status, _ := cmd.Flags().GetString("status")
		return commands.New(cl).ListTasks(ws, status)
	},
}

var agentsCmd = &cobra.Command{
	Use:   "agents",
	Short: "List agents in a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		return commands.New(cl).ListAgents()
	},
}

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "List files in a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ws, _ := cmd.Flags().GetString("workspace")
		if ws == "" {
			return fmt.Errorf("--workspace is required")
		}
		return commands.New(cl).ListFiles(ws)
	},
}

var pinsCmd = &cobra.Command{
	Use:   "pins",
	Short: "List pinned messages in a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ch, _ := cmd.Flags().GetString("channel")
		if ch == "" {
			return fmt.Errorf("--channel is required")
		}
		return commands.New(cl).ListPins(ch)
	},
}

var approvalsCmd = &cobra.Command{
	Use:   "approvals",
	Short: "List approval requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ws, _ := cmd.Flags().GetString("workspace")
		if ws == "" {
			return fmt.Errorf("--workspace is required")
		}
		return commands.New(cl).ListApprovals(ws)
	},
}

var dmCmd = &cobra.Command{
	Use:   "dm",
	Short: "List direct message conversations",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		return commands.New(cl).ListDMs()
	},
}

var unreadCmd = &cobra.Command{
	Use:   "unread",
	Short: "Show unread message count",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		return commands.New(cl).ShowUnread()
	},
}

var editMessageCmd = &cobra.Command{
	Use:   "edit-message",
	Short: "Edit a message",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		id, _ := cmd.Flags().GetString("id")
		content, _ := cmd.Flags().GetString("content")
		if id == "" || content == "" {
			return fmt.Errorf("--id and --content are required")
		}
		return commands.New(cl).EditMessage(id, content)
	},
}

var threadCmd = &cobra.Command{
	Use:   "thread",
	Short: "View a message thread",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		id, _ := cmd.Flags().GetString("message")
		if id == "" {
			return fmt.Errorf("--message is required")
		}
		return commands.New(cl).GetThread(id)
	},
}

var replyCmd = &cobra.Command{
	Use:   "reply",
	Short: "Reply to a thread",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		msgID, _ := cmd.Flags().GetString("message")
		channelID, _ := cmd.Flags().GetString("channel")
		content, _ := cmd.Flags().GetString("content")
		if msgID == "" || channelID == "" || content == "" {
			return fmt.Errorf("--message, --channel, and --content are required")
		}
		return commands.New(cl).ReplyThread(msgID, channelID, content)
	},
}

var notificationsCmd = &cobra.Command{
	Use:   "notifications",
	Short: "List notifications",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		unread, _ := cmd.Flags().GetBool("unread")
		return commands.New(cl).ListNotifications(unread)
	},
}

var markReadCmd = &cobra.Command{
	Use:   "mark-read [id]",
	Short: "Mark a notification as read",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		return commands.New(cl).MarkNotificationRead(args[0])
	},
}

var markAllReadCmd = &cobra.Command{
	Use:   "mark-all-read",
	Short: "Mark all notifications as read",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		return commands.New(cl).MarkAllNotificationsRead()
	},
}

var integrationsCmd = &cobra.Command{
	Use:   "integrations",
	Short: "List available integrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		return commands.New(cl).ListIntegrations()
	},
}

var wsIntegrationsCmd = &cobra.Command{
	Use:   "ws-integrations",
	Short: "List installed integrations in a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ws, _ := cmd.Flags().GetString("workspace")
		if ws == "" {
			return fmt.Errorf("--workspace is required")
		}
		return commands.New(cl).ListWorkspaceIntegrations(ws)
	},
}

var callsCmd = &cobra.Command{
	Use:   "calls",
	Short: "List recent calls",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		limit, _ := cmd.Flags().GetInt("limit")
		return commands.New(cl).ListCalls(limit)
	},
}

var workflowsCmd = &cobra.Command{
	Use:   "workflows",
	Short: "List workflows in a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ws, _ := cmd.Flags().GetString("workspace")
		if ws == "" {
			return fmt.Errorf("--workspace is required")
		}
		return commands.New(cl).ListWorkflows(ws)
	},
}

var triggerWorkflowCmd = &cobra.Command{
	Use:   "trigger-workflow [id]",
	Short: "Trigger a workflow manually",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		return commands.New(cl).TriggerWorkflow(args[0])
	},
}

var billingCmd = &cobra.Command{
	Use:   "billing",
	Short: "Show billing status for a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ws, _ := cmd.Flags().GetString("workspace")
		if ws == "" {
			return fmt.Errorf("--workspace is required")
		}
		return commands.New(cl).GetBilling(ws)
	},
}

var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Show usage metrics for a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ws, _ := cmd.Flags().GetString("workspace")
		if ws == "" {
			return fmt.Errorf("--workspace is required")
		}
		return commands.New(cl).GetUsage(ws)
	},
}

func init() {
	loginCmd.Flags().String("host", "http://127.0.0.1:4001", "Lark server URL")
	loginCmd.Flags().String("token", "", "JWT or API key")

	channelsCmd.Flags().String("workspace", "", "Workspace ID")
	messagesCmd.Flags().String("channel", "", "Channel ID")
	messagesCmd.Flags().Int("limit", 10, "Max messages")
	sendCmd.Flags().String("channel", "", "Channel ID")
	sendCmd.Flags().String("content", "", "Message content")
	searchCmd.Flags().String("channel", "", "Channel ID (optional, search all if empty)")
	tasksCmd.Flags().String("workspace", "", "Workspace ID")
	tasksCmd.Flags().String("status", "", "Filter: todo|in_progress|review|done")
	filesCmd.Flags().String("workspace", "", "Workspace ID")
	pinsCmd.Flags().String("channel", "", "Channel ID")
	approvalsCmd.Flags().String("workspace", "", "Workspace ID")
	editMessageCmd.Flags().String("id", "", "Message ID")
	editMessageCmd.Flags().String("content", "", "New content")
	threadCmd.Flags().String("message", "", "Parent message ID")
	replyCmd.Flags().String("message", "", "Parent message ID")
	replyCmd.Flags().String("channel", "", "Channel ID")
	replyCmd.Flags().String("content", "", "Reply content")

	notificationsCmd.Flags().Bool("unread", false, "Only show unread notifications")

	wsIntegrationsCmd.Flags().String("workspace", "", "Workspace ID")

	callsCmd.Flags().Int("limit", 20, "Max results")

	workflowsCmd.Flags().String("workspace", "", "Workspace ID")

	billingCmd.Flags().String("workspace", "", "Workspace ID")
	usageCmd.Flags().String("workspace", "", "Workspace ID")
}
