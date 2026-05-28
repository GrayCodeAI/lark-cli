package commands

import (
	"fmt"
	"strings"

	"github.com/lark-dev/lark-cli/internal/client"
	"github.com/spf13/cobra"
)

var SendCmd = &cobra.Command{
	Use:   "send [channel] [message]",
	Short: "Send a message to a channel",
	Long:  "Send a message to the specified channel. Message can be multi-word.",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		channelID := args[0]
		message := strings.Join(args[1:], " ")

		cl, err := client.NewFromConfig()
		if err != nil {
			return err
		}

		msg, err := cl.SendMessageToChannel(channelID, message)
		if err != nil {
			return fmt.Errorf("send message: %w", err)
		}

		fmt.Printf("Message sent to %s\n", channelID)
		if id, ok := msg["id"]; ok {
			fmt.Printf("Message ID: %v\n", id)
		}
		return nil
	},
}
