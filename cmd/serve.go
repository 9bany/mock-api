package cmd

import (
	mock "9bany/mapi/server/api"
	"bufio"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve a server that return your data in os.stdin.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		// aware: this code just work with small file
		data := ""
		for scanner.Scan() {
			data += scanner.Text()
		}
		server := mock.NewServer([]byte(data))
		server.Run()
	},
}
