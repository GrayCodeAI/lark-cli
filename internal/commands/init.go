package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lark-dev/lark-cli/internal/client"
	"github.com/lark-dev/lark-cli/internal/config"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Lark CLI configuration",
	Long:  "Create .lark/config.yaml with server connection settings.",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		// Server URL
		defaultURL := config.DefaultServerURL
		if cfg, err := config.Load(); err == nil && cfg.ServerURL != "" {
			defaultURL = cfg.ServerURL
		}
		fmt.Printf("Server URL [%s]: ", defaultURL)
		serverURL, _ := reader.ReadString('\n')
		serverURL = strings.TrimSpace(serverURL)
		if serverURL == "" {
			serverURL = defaultURL
		}

		// API Token
		fmt.Print("API Token: ")
		apiToken, _ := reader.ReadString('\n')
		apiToken = strings.TrimSpace(apiToken)

		// Workspace ID (optional)
		fmt.Print("Workspace ID (optional): ")
		workspaceID, _ := reader.ReadString('\n')
		workspaceID = strings.TrimSpace(workspaceID)

		cfg := &config.Config{
			ServerURL:   serverURL,
			APIToken:    apiToken,
			WorkspaceID: workspaceID,
		}

		if err := config.Save(cfg); err != nil {
			return fmt.Errorf("save config: %w", err)
		}

		// Verify connection
		cl := client.New(serverURL, apiToken)
		if err := cl.Do("GET", "/v1/workspaces", nil, nil); err != nil {
			fmt.Printf("Warning: could not verify connection: %v\n", err)
			fmt.Println("Configuration saved. You may need to check your credentials.")
		} else {
			fmt.Println("Configuration saved and connection verified.")
		}

		return nil
	},
}
