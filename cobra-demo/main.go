package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// kubectl get pod -A
// program argument: get pod -A
func main() {
	var rootCmd = &cobra.Command{
		Use:   "kubectl",
		Short: "simple example for kubectl",
		Long:  `simple example for kubectl get pod -A`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(os.Stdout, args)
		},
	}

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "get pod info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(os.Stdout, args)
		},
	}

	var isAll = false
	var podCmd = &cobra.Command{
		Use:   "pod",
		Short: "get pod info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(os.Stdout, "result: ", isAll)
		},
	}
	podCmd.Flags().BoolVarP(&isAll, "all", "A", false, "it will get all")

	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(podCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
