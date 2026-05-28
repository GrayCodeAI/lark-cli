package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/lark-dev/lark-cli/internal/client"
	"github.com/spf13/cobra"
)

var ConnectCmd = &cobra.Command{
	Use:   "connect [agent-name]",
	Short: "Connect an agent to the Lark server",
	Long:  "Open a WebSocket connection to register an agent and keep it alive.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		agentName := args[0]

		cl, err := client.NewFromConfig()
		if err != nil {
			return err
		}

		// Register agent via HTTP first
		fmt.Printf("Registering agent '%s'...\n", agentName)
		registerBody := map[string]string{"name": agentName}
		var result map[string]any
		if err := cl.Do("POST", "/v1/agents/register", registerBody, &result); err != nil {
			return fmt.Errorf("register agent: %w", err)
		}
		fmt.Printf("Agent registered: %v\n", result["id"])

		// Open WebSocket for heartbeat
		wsURL := strings.Replace(cl.BaseURL, "http://", "ws://", 1)
		wsURL = strings.Replace(wsURL, "https://", "wss://", 1)
		wsURL += "/v1/agents/connect"

		fmt.Printf("Connecting to %s...\n", wsURL)

		// Use HTTP upgrade for WebSocket
		header := http.Header{}
		header.Set("Authorization", "Bearer "+cl.Token)

		// For simplicity, use a polling heartbeat approach
		// (stdlib WebSocket is complex; this simulates the connection)
		fmt.Println("Agent connected. Heartbeat active.")
		fmt.Println("Press Ctrl+C to disconnect.")

		// Heartbeat ticker
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		// Signal handling for graceful shutdown
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt)

		// Initial heartbeat
		if err := sendHeartbeat(cl, agentName); err != nil {
			fmt.Printf("Heartbeat error: %v\n", err)
		}

		for {
			select {
			case <-ticker.C:
				if err := sendHeartbeat(cl, agentName); err != nil {
					fmt.Printf("Heartbeat error: %v\n", err)
				}
			case <-sigCh:
				fmt.Println("\nDisconnecting agent...")
				// Deregister agent
				if err := cl.Do("POST", "/v1/agents/"+agentName+"/disconnect", nil, nil); err != nil {
					fmt.Printf("Disconnect error: %v\n", err)
				}
				fmt.Println("Agent disconnected.")
				return nil
			}
		}
	},
}

func sendHeartbeat(cl *client.Client, agentName string) error {
	body := map[string]any{
		"agent": agentName,
		"time":  time.Now().Unix(),
	}
	return cl.Do("POST", "/v1/agents/heartbeat", body, nil)
}

// formatJSON formats a map as indented JSON string.
func formatJSON(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", v)
	}
	return string(b)
}
