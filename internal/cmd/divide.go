/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newDivideCmd() *cobra.Command {
	newDivideCmd := &cobra.Command{
		Use:   "divide",
		Short: "divide your calculation",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("divide called")
		},
	}

	return newDivideCmd
}
