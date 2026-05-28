package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/lark-dev/lark-cli/internal/client"
	"github.com/lark-dev/lark-cli/internal/config"
	"github.com/spf13/cobra"
)

var ChannelsListCmd = &cobra.Command{
	Use:   "channels",
	Short: "List channels",
	Long:  "Display channels in the workspace. Uses workspace from config if not specified.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := client.NewFromConfig()
		if err != nil {
			return err
		}

		workspaceID, _ := cmd.Flags().GetString("workspace")
		if workspaceID == "" {
			cfg, err := config.Load()
			if err == nil && cfg.WorkspaceID != "" {
				workspaceID = cfg.WorkspaceID
			}
		}

		channels, err := cl.ListAllChannels(workspaceID)
		if err != nil {
			return fmt.Errorf("list channels: %w", err)
		}

		if len(channels) == 0 {
			fmt.Println("No channels found.")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tTYPE\tTOPIC")
		for _, ch := range channels {
			id, _ := ch["id"].(string)
			name, _ := ch["name"].(string)
			chType, _ := ch["type"].(string)
			topic, _ := ch["topic"].(string)
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", id, name, chType, topic)
		}
		w.Flush()
		return nil
	},
}
