package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/lark-dev/lark-cli/internal/client"
	"github.com/spf13/cobra"
)

var AgentsListCmd = &cobra.Command{
	Use:   "agents",
	Short: "List connected agents",
	Long:  "Display all agents registered with the Lark server.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := client.NewFromConfig()
		if err != nil {
			return err
		}

		agents, err := cl.ListAllAgents()
		if err != nil {
			return fmt.Errorf("list agents: %w", err)
		}

		if len(agents) == 0 {
			fmt.Println("No agents connected.")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tSTATUS\tWORKSPACE")
		for _, a := range agents {
			id, _ := a["id"].(string)
			name, _ := a["name"].(string)
			status, _ := a["status"].(string)
			if status == "" {
				status = "online"
			}
			workspace, _ := a["workspace"].(string)
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", id, name, status, workspace)
		}
		w.Flush()
		return nil
	},
}
