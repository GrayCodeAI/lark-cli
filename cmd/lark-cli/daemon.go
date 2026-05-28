package main

import (
	"fmt"

	"github.com/lark-dev/lark-cli/internal/commands"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Manage local daemons",
	Long:  "List, inspect, and manage local daemon connections and their agents.",
}

var daemonListCmd = &cobra.Command{
	Use:   "list",
	Short: "List connected daemons",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		ws, _ := cmd.Flags().GetString("workspace")
		if ws == "" {
			return fmt.Errorf("--workspace is required")
		}
		return (&commands.DaemonCommands{Client: cl}).ListDaemons(ws)
	},
}

var daemonShowCmd = &cobra.Command{
	Use:   "show [daemon-id]",
	Short: "Show daemon details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		return (&commands.DaemonCommands{Client: cl}).ShowDaemon(args[0])
	},
}

var daemonWakeCmd = &cobra.Command{
	Use:   "wake [daemon-id] [agent-name]",
	Short: "Wake an agent through a daemon",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		channel, _ := cmd.Flags().GetString("channel")
		message, _ := cmd.Flags().GetString("message")
		return (&commands.DaemonCommands{Client: cl}).WakeDaemonAgent(args[0], args[1], channel, message)
	},
}

var daemonRestartCmd = &cobra.Command{
	Use:   "restart [daemon-id] [agent-name]",
	Short: "Restart an agent through a daemon",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := loadClient()
		if err != nil {
			return err
		}
		return (&commands.DaemonCommands{Client: cl}).RestartDaemonAgent(args[0], args[1])
	},
}

func init() {
	daemonListCmd.Flags().String("workspace", "", "Workspace ID")
	daemonWakeCmd.Flags().String("channel", "", "Channel ID")
	daemonWakeCmd.Flags().String("message", "", "Wake message")
}
