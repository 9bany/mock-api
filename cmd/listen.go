package cmd

import (
	"9bany/mapi/server/listen"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listenCmd)
}

var listenCmd = &cobra.Command{
	Use:   "observe",
	Short: "Serve a server that listening data you send to the API.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server := listen.NewServer()
		server.Run()
	},
}
