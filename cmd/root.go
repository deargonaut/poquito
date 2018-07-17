package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "poquito",
	Short: "Poquite will orchestrate your Istio Microservice Mesh",
	Long: `Easily deploy and manage microservices in Istio Mesh
				running on Kubernetes`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
