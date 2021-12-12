package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:                   "go_kafka",
		Short:                 "go_kafka",
		Version:               "1.0.0",
		DisableFlagsInUseLine: true,
	}

	rootCmd.AddCommand(API)
	rootCmd.AddCommand(ConsumerWorker)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func InitConfig() {

}
