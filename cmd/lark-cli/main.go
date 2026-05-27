package main

import (
	"flag"
	"fmt"
	"os"

	"lark-cli/internal/client"
	"lark-cli/internal/commands"
)

func main() {
	cfg, err := client.LoadConfig()
	if err == nil && cfg != nil {
		// logged in
	}
	cl := client.New("", "")
	_ = cl
	cmds := commands.New(cl)

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	sub := os.Args[1]
	args := os.Args[2:]

	switch sub {
	case "login":
		fs := flag.NewFlagSet("login", flag.ExitOnError)
		host := fs.String("host", "http://127.0.0.1:4001", "Lark server URL")
		token := fs.String("token", "", "JWT or API key")
		fs.Parse(args)
		if *token == "" {
			fmt.Fprintln(os.Stderr, "error: --token is required")
			os.Exit(1)
		}
		if err := cl.Login(*host, *token); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		fmt.Println("logged in")

	case "whoami":
		cfg, err := client.LoadConfig()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: not logged in (run 'lark login')")
			os.Exit(1)
		}
		cl = client.New(cfg.BaseURL, cfg.Token)
		cmds = commands.New(cl)
		if err := cmds.Whoami(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

	case "workspaces":
		cfg, err := mustLoad()
		if err != nil {
			os.Exit(1)
		}
		cl = client.New(cfg.BaseURL, cfg.Token)
		cmds = commands.New(cl)
		if err := cmds.ListWorkspaces(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

	case "channels":
		cfg, err := mustLoad()
		if err != nil {
			os.Exit(1)
		}
		fs := flag.NewFlagSet("channels", flag.ExitOnError)
		ws := fs.String("workspace", "", "Workspace ID")
		fs.Parse(args)
		if *ws == "" {
			fmt.Fprintln(os.Stderr, "error: --workspace is required")
			os.Exit(1)
		}
		cl = client.New(cfg.BaseURL, cfg.Token)
		cmds = commands.New(cl)
		if err := cmds.ListChannels(*ws); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

	case "messages":
		cfg, err := mustLoad()
		if err != nil {
			os.Exit(1)
		}
		fs := flag.NewFlagSet("messages", flag.ExitOnError)
		ch := fs.String("channel", "", "Channel ID")
		limit := fs.Int("limit", 10, "Message limit")
		fs.Parse(args)
		if *ch == "" {
			fmt.Fprintln(os.Stderr, "error: --channel is required")
			os.Exit(1)
		}
		cl = client.New(cfg.BaseURL, cfg.Token)
		cmds = commands.New(cl)
		if err := cmds.ListMessages(*ch, *limit); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

	case "send":
		cfg, err := mustLoad()
		if err != nil {
			os.Exit(1)
		}
		fs := flag.NewFlagSet("send", flag.ExitOnError)
		ch := fs.String("channel", "", "Channel ID")
		content := fs.String("content", "", "Message content")
		fs.Parse(args)
		if *ch == "" || *content == "" {
			fmt.Fprintln(os.Stderr, "error: --channel and --content are required")
			os.Exit(1)
		}
		cl = client.New(cfg.BaseURL, cfg.Token)
		cmds = commands.New(cl)
		if err := cmds.SendMessage(*ch, *content); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

	case "tasks":
		cfg, err := mustLoad()
		if err != nil {
			os.Exit(1)
		}
		fs := flag.NewFlagSet("tasks", flag.ExitOnError)
		ws := fs.String("workspace", "", "Workspace ID")
		status := fs.String("status", "", "Filter by status")
		fs.Parse(args)
		if *ws == "" {
			fmt.Fprintln(os.Stderr, "error: --workspace is required")
			os.Exit(1)
		}
		cl = client.New(cfg.BaseURL, cfg.Token)
		cmds = commands.New(cl)
		if err := cmds.ListTasks(*ws, *status); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", sub)
		printUsage()
		os.Exit(1)
	}
}

func mustLoad() (*client.Config, error) {
	cfg, err := client.LoadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: not logged in (run 'lark login')")
		return nil, err
	}
	return cfg, nil
}

func printUsage() {
	fmt.Println(`Usage: lark-cli <command> [options]

Commands:
  login                     Authenticate with Lark server
    --host <url>              Server URL (default: http://127.0.0.1:4001)
    --token <key>             JWT or API key

  whoami                    Show current auth info

  workspaces                List workspaces

  channels                  List channels
    --workspace <id>          Workspace ID (required)

  messages                  List messages
    --channel <id>            Channel ID (required)
    --limit <n>               Max messages (default: 10)

  send                      Send a message
    --channel <id>            Channel ID (required)
    --content <text>          Message content (required)

  tasks                     List tasks
    --workspace <id>          Workspace ID (required)
    --status <s>              Filter: todo|in_progress|review|done`)
}
