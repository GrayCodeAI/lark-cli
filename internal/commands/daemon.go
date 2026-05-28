package commands

import (
	"fmt"
	"text/tabwriter"
	"os"

	"github.com/lark-dev/lark-cli/internal/client"
)

// DaemonCommands handles daemon management CLI commands.
type DaemonCommands struct {
	Client *client.Client
}

// ListDaemons lists registered daemons.
func (d *DaemonCommands) ListDaemons(workspaceID string) error {
	daemons, err := d.Client.ListDaemons(workspaceID)
	if err != nil {
		return fmt.Errorf("list daemons: %w", err)
	}
	if len(daemons) == 0 {
		fmt.Println("no daemons connected")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tAGENTS\tSTATUS")
	for _, d := range daemons {
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\n", d.ID, d.Name, d.AgentCount, d.Status)
	}
	w.Flush()
	return nil
}

// ShowDaemon shows details of a specific daemon.
func (d *DaemonCommands) ShowDaemon(daemonID string) error {
	daemon, err := d.Client.GetDaemon(daemonID)
	if err != nil {
		return fmt.Errorf("get daemon: %w", err)
	}

	fmt.Printf("Daemon: %s\n", daemon.Name)
	fmt.Printf("  ID:      %s\n", daemon.ID)
	fmt.Printf("  Status:  %s\n", daemon.Status)
	fmt.Printf("  Agents:  %d\n", daemon.AgentCount)
	fmt.Printf("  Started: %s\n", daemon.StartedAt)

	if len(daemon.Agents) > 0 {
		fmt.Println("\n  Registered Agents:")
		for _, a := range daemon.Agents {
			fmt.Printf("    - %s (%s)\n", a.Name, a.AgentID)
		}
	}

	return nil
}

// WakeDaemonAgent wakes a specific agent through a daemon.
func (d *DaemonCommands) WakeDaemonAgent(daemonID, agentName, channelID, message string) error {
	return d.Client.WakeDaemonAgent(daemonID, agentName, channelID, message)
}

// RestartDaemonAgent restarts an agent through a daemon.
func (d *DaemonCommands) RestartDaemonAgent(daemonID, agentName string) error {
	return d.Client.RestartDaemonAgent(daemonID, agentName)
}
