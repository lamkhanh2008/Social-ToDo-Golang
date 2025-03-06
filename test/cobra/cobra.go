package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var sayHelloCmd = &cobra.Command{
		Use:   "say-hello",
		Short: "Say hello command",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				fmt.Println("hello!")
			} else {
				fmt.Printf("Hello, %s \n", name)
			}
		},
	}
	var rootCmd = &cobra.Command{
		Use:   "cobra-example",
		Short: "cobra CLI example",
		Long:  `Cobra CLI example is a basic example of using the Cobra package in Go.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from cobra")
		},
	}
	sayHelloCmd.Flags().StringP("name", "n", "", "YourName")
	rootCmd.AddCommand(sayHelloCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
