/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newMultiplyCmd() *cobra.Command {
	multiplyCmd := &cobra.Command{
		Use:   "multiply",
		Short: "multiply your calculation",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add called")
		},
	}

	return multiplyCmd
}
