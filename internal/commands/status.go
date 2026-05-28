package commands

import (
	"fmt"
	"text/tabwriter"
	"os"

	"github.com/lark-dev/lark-cli/internal/client"
	"github.com/lark-dev/lark-cli/internal/config"
	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show Lark CLI and server status",
	Long:  "Display connection status, configured server, and summary of agents/channels.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		// Show config
		fmt.Println("=== Lark CLI Status ===")
		fmt.Printf("Server:   %s\n", cfg.ServerURL)
		tokenDisplay := cfg.APIToken
		if len(tokenDisplay) > 16 {
			tokenDisplay = tokenDisplay[:16] + "..."
		}
		fmt.Printf("Token:    %s\n", tokenDisplay)
		if cfg.WorkspaceID != "" {
			fmt.Printf("Workspace: %s\n", cfg.WorkspaceID)
		}
		fmt.Println()

		// Try to connect and get status
		cl, err := client.NewFromConfig()
		if err != nil {
			fmt.Printf("Connection: Not configured\n")
			return nil
		}

		// Check server health
		var status map[string]any
		if err := cl.Do("GET", "/v1/status", nil, &status); err != nil {
			fmt.Printf("Connection: Failed (%v)\n", err)
			return nil
		}
		fmt.Println("Connection: OK")

		// Show agents summary
		agents, err := cl.ListAllAgents()
		if err == nil && len(agents) > 0 {
			fmt.Printf("\n=== Agents (%d) ===\n", len(agents))
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tNAME\tSTATUS")
			for _, a := range agents {
				name, _ := a["name"].(string)
				id, _ := a["id"].(string)
				status, _ := a["status"].(string)
				if status == "" {
					status = "unknown"
				}
				fmt.Fprintf(w, "%s\t%s\t%s\n", id, name, status)
			}
			w.Flush()
		}

		// Show channels summary
		channels, err := cl.ListAllChannels(cfg.WorkspaceID)
		if err == nil && len(channels) > 0 {
			fmt.Printf("\n=== Channels (%d) ===\n", len(channels))
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tNAME\tTYPE")
			for _, ch := range channels {
				name, _ := ch["name"].(string)
				id, _ := ch["id"].(string)
				chType, _ := ch["type"].(string)
				fmt.Fprintf(w, "%s\t%s\t%s\n", id, name, chType)
			}
			w.Flush()
		}

		return nil
	},
}
