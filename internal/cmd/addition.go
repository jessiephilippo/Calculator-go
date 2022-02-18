/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newAdditionCmd() *cobra.Command {
	additionCmd := &cobra.Command{
		Use:   "addition",
		Short: "addition your calculation",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("addition called")
		},
	}

	return additionCmd
}
