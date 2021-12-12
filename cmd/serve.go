package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevegood/kalmar/pkg/server"
)

const defaultGraphQLPort = "8080"

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts up the GraphQL server",
	Long:  `Starts up the GraphQL server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		port := cmd.Flag("port").Value.String()
		return server.Exec(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().StringP("port", "p", defaultGraphQLPort, "GraphQL server port")
}
