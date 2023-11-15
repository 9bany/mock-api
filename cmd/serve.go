package cmd

import (
	"9bany/mapi/api"
	"bufio"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Print the version number of Bit",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		// aware: this code just work with small file
		data := ""
		for scanner.Scan() {
			data += scanner.Text()
		}
		server := api.NewServer([]byte(data))
		server.Run()
	},
}
